package tools

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	_ "modernc.org/sqlite"
)

func defaultDocsDBPath() string {
	// Look next to the binary first, then $HOME/kentik-docs.db
	exe, err := os.Executable()
	if err == nil {
		candidate := filepath.Join(filepath.Dir(exe), "docs", "kentik-docs.db")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		// Also check sibling docs/ relative to repo root (dev mode)
		candidate2 := filepath.Join(filepath.Dir(exe), "..", "docs", "kentik-docs.db")
		if _, err := os.Stat(candidate2); err == nil {
			return candidate2
		}
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "kentik-docs.db")
}

func registerDocsTools(s *server.MCPServer) {
	searchDocs := mcp.NewTool("kentik_search_docs",
		mcp.WithDescription("Search the local Kentik documentation knowledge base. "+
			"Returns relevant excerpts from Kentik KB articles matching your query. "+
			"Useful for looking up API details, feature explanations, configuration options, and troubleshooting guidance."),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query. E.g. 'BGP monitoring API', 'synthetics test configuration', 'alerting thresholds'."),
		),
		mcp.WithNumber("limit",
			mcp.Description("Number of results to return. Default: 5"),
		),
		mcp.WithString("db_path",
			mcp.Description("Path to the kentik-docs.db file. Auto-detected if not specified."),
		),
	)
	s.AddTool(searchDocs, makeSearchDocsHandler())
}

func makeSearchDocsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := request.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		limit := 5.0
		if l, err := request.RequireFloat("limit"); err == nil && l > 0 {
			limit = l
		}

		dbPath, _ := request.RequireString("db_path")
		if dbPath == "" {
			dbPath = defaultDocsDBPath()
		}

		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			return mcp.NewToolResultError(fmt.Sprintf(
				"Docs database not found at %s. Run scripts/ingest_docs.py to build it.", dbPath,
			)), nil
		}

		db, err := sql.Open("sqlite", dbPath+"?mode=ro")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Open DB: %v", err)), nil
		}
		defer db.Close()

		// FTS5 full-text search with snippet highlighting
		rows, err := db.QueryContext(ctx, `
			SELECT
				title,
				source_url,
				snippet(docs_fts, 4, '**', '**', '...', 30)
			FROM docs_fts
			WHERE docs_fts MATCH ?
			ORDER BY rank
			LIMIT ?
		`, query, int(limit))
		if err != nil {
			// FTS5 syntax errors are common with raw user input — retry with quoted phrase
			quoted := `"` + strings.ReplaceAll(query, `"`, `""`) + `"`
			rows, err = db.QueryContext(ctx, `
				SELECT title, source_url, snippet(docs_fts, 4, '**', '**', '...', 30)
				FROM docs_fts WHERE docs_fts MATCH ?
				ORDER BY rank LIMIT ?
			`, quoted, int(limit))
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Search failed: %v", err)), nil
			}
		}
		defer rows.Close()

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("## Kentik Docs: \"%s\"\n\n", query))

		count := 0
		for rows.Next() {
			var title, sourceURL, snippet string
			if err := rows.Scan(&title, &sourceURL, &snippet); err != nil {
				continue
			}
			count++
			sb.WriteString(fmt.Sprintf("### %d. %s\n", count, title))
			sb.WriteString(fmt.Sprintf("📄 %s\n\n", sourceURL))
			sb.WriteString(snippet)
			sb.WriteString("\n\n---\n\n")
		}

		if count == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("No results found for \"%s\".", query)), nil
		}

		sb.WriteString(fmt.Sprintf("*%d results from local Kentik docs*\n", count))
		return mcp.NewToolResultText(sb.String()), nil
	}
}
