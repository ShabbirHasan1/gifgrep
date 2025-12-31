# Changelog

## 0.1.0 - 2025-12-30

### Added
- **CLI search mode**: `gifgrep <query...>` prints `<title>\\t<url>` (with `--json` and `--number` options).
- **Subcommands**: `search`, `tui`, `still`, `sheet`.
- **Providers**: `--source` (`auto`, `tenor`, `giphy`) with `TENOR_API_KEY` and `GIPHY_API_KEY` support.
- **Output polish**: color on rich TTYs, `--no-color`, `--quiet`, `--verbose`.
- **Interactive TUI**: raw-mode terminal UI with query/browse states, arrow-key navigation, status line, and key hints.
- **Kitty graphics preview**: inline GIF rendering with automatic cleanup; aspect-ratio-aware sizing with `GIFGREP_CELL_ASPECT`.
- **Animation handling**: Kitty animation stream when supported; software re-render fallback (Ghostty auto-detect or `GIFGREP_SOFTWARE_ANIM=1`).
- **Responsive layout**: list + preview side-by-side on wide terminals, stacked preview on narrow terminals.
- **Preview caching**: in-memory cache keyed by preview URL for fast browsing.
- **Stills extraction**: `still` subcommand for a single frame; `sheet` subcommand for a sampled PNG grid; output to file or stdout.
- **Giphy attribution**: inline logo display when Giphy is the active source.

### Developer Experience
- **Formatter + lint**: gofumpt and golangci-lint, with Makefile/justfile targets.
- **Benchmarks**: synthetic and fixture-backed decode benchmarks.
- **Fixtures**: small, licensed GIF corpus with documented sources (`docs/gif-sources.md`).
- **Visual checks**: Ghostty-web screenshot harness (`pnpm snap`).
- **Docs site**: GitHub Pages content in `docs/`.
