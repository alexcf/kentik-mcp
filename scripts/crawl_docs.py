#!/usr/bin/env python3
"""Crawl kb.kentik.com/docs and save each page as markdown."""
# (unchanged — see ingest_docs.py for the DB ingestion script)

import os
import re
import sys
import time
from collections import deque
from urllib.parse import urljoin, urlparse

import requests
from bs4 import BeautifulSoup
from markdownify import markdownify

BASE_URL = "https://kb.kentik.com"
START_URL = "https://kb.kentik.com/docs"
OUT_DIR = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "docs")
DELAY = 0.3  # seconds between requests

session = requests.Session()
session.headers["User-Agent"] = "kentik-mcp-docs-crawler/1.0"


def slug_from_url(url: str) -> str:
    path = urlparse(url).path.strip("/")
    path = re.sub(r"[^a-zA-Z0-9/_-]", "-", path)
    return path or "index"


def is_docs_url(url: str) -> bool:
    p = urlparse(url)
    return p.netloc in ("kb.kentik.com", "") and p.path.startswith("/docs")


def extract_content(soup: BeautifulSoup) -> tuple[str, str]:
    """Return (title, markdown_content)."""
    title = ""
    if soup.title:
        title = soup.title.string or ""
    # Try article/main content first, fall back to body
    content_el = (
        soup.find("article")
        or soup.find("main")
        or soup.find(class_=re.compile(r"content|article|prose|markdown", re.I))
        or soup.body
    )
    if not content_el:
        return title, ""
    # Strip nav, sidebar, footer, scripts
    for tag in content_el.find_all(["nav", "footer", "script", "style", "aside"]):
        tag.decompose()
    md = markdownify(str(content_el), heading_style="ATX", strip=["a"]).strip()
    # Collapse 3+ blank lines to 2
    md = re.sub(r"\n{3,}", "\n\n", md)
    return title, md


def crawl():
    os.makedirs(OUT_DIR, exist_ok=True)
    visited = set()
    queue = deque([START_URL])
    index = []  # (title, slug, url)

    print(f"Crawling {START_URL} → {OUT_DIR}/")

    while queue:
        url = queue.popleft()
        if url in visited:
            continue
        visited.add(url)

        try:
            resp = session.get(url, timeout=15)
            if resp.status_code != 200:
                print(f"  SKIP {resp.status_code} {url}")
                continue
        except Exception as e:
            print(f"  ERROR {url}: {e}")
            continue

        soup = BeautifulSoup(resp.text, "html.parser")
        title, md = extract_content(soup)

        slug = slug_from_url(url)
        filepath = os.path.join(OUT_DIR, slug + ".md")

        if os.path.exists(filepath):
            # Already scraped — still enqueue links by reading saved source URL
            print(f"  ~ skip (exists) {slug}")
            index.append((slug, slug, url))
        elif md:
            os.makedirs(os.path.dirname(filepath), exist_ok=True)
            with open(filepath, "w", encoding="utf-8") as f:
                f.write(f"# {title}\n\nSource: {url}\n\n---\n\n{md}\n")
            index.append((title, slug, url))
            print(f"  ✓ {slug} ({len(md)} chars)")

        # Enqueue new docs links
        for a in soup.find_all("a", href=True):
            href = a["href"].split("#")[0]  # drop anchors
            full = urljoin(url, href)
            if is_docs_url(full) and full not in visited:
                queue.append(full)

        time.sleep(DELAY)

    # Write index
    index_path = os.path.join(OUT_DIR, "README.md")
    with open(index_path, "w", encoding="utf-8") as f:
        f.write("# Kentik Documentation Index\n\n")
        f.write(f"Crawled {len(index)} pages from {BASE_URL}\n\n")
        for title, slug, url in sorted(index, key=lambda x: x[1]):
            f.write(f"- [{title or slug}]({slug}.md) — [{url}]({url})\n")

    print(f"\nDone. {len(index)} pages saved to {OUT_DIR}/")
    print(f"Index: {index_path}")


if __name__ == "__main__":
    crawl()
