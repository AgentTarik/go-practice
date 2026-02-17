# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

A Go practice project (module: `practice`, Go 1.25.0) for algorithm/coding exercises. Currently implements a 2D grid island detection algorithm.

## Commands

```bash
go run main.go      # Run the program
go build            # Build binary
go fmt ./...        # Format code
go vet ./...        # Static analysis
go test ./...       # Run tests (none yet)
```

## Architecture

Single-package project (`package main`) with all code in `main.go`. No external dependencies.

Current implementation — island detection in a 2D grid (`[][]int`, 1 = land, 0 = water):

- `IslandDetector(ocean [][]int) int` — counts distinct islands by scanning row-by-row and checking connectivity
- `IsPartOfIsland(ocean [][]int, i, j int) bool` — checks if `(i, j)` is adjacent to a previously seen land cell (above or left only)
- `ShowMap(ocean [][]int)` — prints the grid using `#` (land) and `~` (water)

The connectivity check is directional (up and left only), which means the island count reflects a specific scan-order definition of island membership rather than full flood-fill connectivity.
