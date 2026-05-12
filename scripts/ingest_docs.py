#!/usr/bin/env python3
"""Ingest scraped Kentik docs markdown files into a SQLite FTS5 database."""

import os
import re
import sqlite3

DOCS_DIR = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "docs")
DB_PATH = os.path.join(DOCS_DIR, "kentik-docs.db")


def extract_metadata(content: str) -> tuple:
    title, source = "", ""
    for line in content.splitlines()[:5]:
        if line.startswith("# "):
            title = line[2:].strip()
        elif line.startswith("Source: "):
            source = line[8:].strip()
    return title, source


def clean_markdown(content: str) -> str:
    lines = content.splitlines()
    body = "\n".join(lines[5:]) if len(lines) > 5 else content
    body = re.sub(r"#{1,6}\s+", "", body)
    body = re.sub(r"\*{1,2}([^*]+)\*{1,2}", r"\1", body)
    body = re.sub(r"`{1,3}[^`]*`{1,3}", " ", body)
    body = re.sub(r"\[([^\]]+)\]\([^\)]+\)", r"\1", body)
    body = re.sub(r"\n{3,}", "\n\n", body)
    return body.strip()


def chunk_content(content: str, chunk_size: int = 800) -> list:
    paragraphs = re.split(r"\n\n+", content)
    chunks, current, current_len = [], [], 0
    for para in paragraphs:
        para = para.strip()
        if not para:
            continue
        if current_len + len(para) > chunk_size and current:
            chunks.append("\n\n".join(current))
            current = [current[-1], para]
            current_len = sum(len(p) for p in current)
        else:
            current.append(para)
            current_len += len(para)
    if current:
        chunks.append("\n\n".join(current))
    return chunks or [content[:chunk_size]]


def ingest():
    if os.path.exists(DB_PATH):
        os.remove(DB_PATH)

    conn = sqlite3.connect(DB_PATH)
    conn.executescript("""
        CREATE VIRTUAL TABLE docs_fts USING fts5(
            slug, title, source_url, chunk_id UNINDEXED, body,
            tokenize = 'porter ascii'
        );
        CREATE TABLE doc_meta (
            slug TEXT PRIMARY KEY, title TEXT, source_url TEXT, chunk_count INTEGER
        );
    """)

    # Walk all subdirectories
    md_files = []
    for root, _, files in os.walk(DOCS_DIR):
        for f in files:
            if f.endswith(".md") and f != "README.md" and "kentik-docs.db" not in f:
                md_files.append(os.path.join(root, f))
    md_files.sort()

    for filepath in md_files:
        slug = os.path.relpath(filepath, DOCS_DIR)[:-3]  # relative path minus .md
        with open(filepath, encoding="utf-8") as f:
            raw = f.read()

        title, source_url = extract_metadata(raw)
        body = clean_markdown(raw)
        chunks = chunk_content(body)

        for i, chunk in enumerate(chunks):
            conn.execute(
                "INSERT INTO docs_fts VALUES (?,?,?,?,?)",
                (slug, title, source_url, i, chunk),
            )
        conn.execute(
            "INSERT INTO doc_meta VALUES (?,?,?,?)",
            (slug, title, source_url, len(chunks)),
        )

    conn.commit()
    count = conn.execute("SELECT COUNT(*) FROM docs_fts").fetchone()[0]
    docs = conn.execute("SELECT COUNT(*) FROM doc_meta").fetchone()[0]
    print(f"Ingested {docs} docs -> {count} chunks into {DB_PATH}")
    conn.close()


if __name__ == "__main__":
    ingest()
