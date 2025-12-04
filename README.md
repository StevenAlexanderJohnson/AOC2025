# Advent of Code 2025

This repository contains my personal solutions for Advent of Code 2025.

## Overview

Each day's solution lives in its own directory (`day1/`, `day2/`, ...). A small per-day `main.go` bootstraps the program and calls the shared `utils.Runner`.

Please note: the maintainer requests that puzzle text or questions are not reposted in this repository. This repo only contains my implementations and helper code.

## File layout

- `dayN/` — per-day code and inputs
  - `main.go` — per-day entry that calls `utils.Runner`
  - `part1.go`, `part2.go` — solutions for each part
  - `shared.go` — per-day helpers and parsing
  - `input.txt`, `test_input.txt` — input files
- `utils/` — shared runner and helpers
  - `runner.go` — `utils.Runner` that reads flags, parses input, times execution, and calls parts
  - `flags.go`, `input.go` — helpers used by the runner
- `go.mod` — Go module file

## How to run

From the repository root you can run any day's solution with `go run` and the usual flags.

### Flags

All days use the same flags (handled in `utils/flags.go`). To see the flags available for a given day run:

```bash
go run ./day1 -h
```

Examples:

Run day 1, part 1:

```bash
go run ./day1 -input_path=day1/input.txt -part=1
```

Run day 2, part 2:

```bash
go run ./day2 -input_path=day2/input.txt -part=2
```

Notes:
- The runner expects `-input_path` pointing to the input file and `-part` set to `1` or `2` (see `utils/flags.go`).
- Each day's `main.go` passes that day's `part1`, `part2`, and `parseInput` functions to `utils.Runner`.

## Testing and benchmarks

Testing and benchmarking are optional for this repository — it's primarily a place for solving the puzzles and experimenting. If a puzzle is tricky or important, a small test or benchmark may be included in the corresponding day's package.
