# gifgrep — Grep the GIF. Stick the landing.

Two modes, one tool.
- Scriptable CLI: URLs/JSON for pipes
- TUI: arrow-key browse + inline preview (Kitty graphics)
- Stills: extract a single frame or a PNG sheet of key frames

Website: `https://gifgrep.com`

## Install
- Homebrew: `brew install steipete/tap/gifgrep`
- Go: `go install github.com/steipete/gifgrep/cmd/gifgrep@latest`

## Quickstart
```bash
gifgrep cats --max 5
gifgrep cats --json | jq '.[] | .url'
gifgrep tui "office handshake"

gifgrep still ./clip.gif --at 1.5s -o still.png
gifgrep sheet ./clip.gif --frames 9 --cols 3 -o sheet.png
```

## Sheet
Single PNG grid of sampled frames. Use `--frames` to pick how many frames, `--cols` to control the grid.

## Providers
Select via `--source` (search + TUI):
- `auto` (default) - picks giphy when `GIPHY_API_KEY` is set, else tenor
- `tenor` - uses public demo key if `TENOR_API_KEY` unset
- `giphy` - requires `GIPHY_API_KEY`

## CLI
```text
gifgrep [global flags] <query...>
gifgrep search [flags] <query...>
gifgrep tui [flags] [<query...>]
gifgrep still <gif> --at <time> [-o <file>|-]
gifgrep sheet <gif> [--frames <N>] [--cols <N>] [--padding <px>] [-o <file>|-]
```

## JSON output
`--json` prints an array with: `id`, `title`, `url`, `preview_url`, `tags`, `width`, `height`.

## Environment
- `TENOR_API_KEY` (optional)
- `GIPHY_API_KEY` (required for `--source giphy`)
- `GIFGREP_SOFTWARE_ANIM=1` (force software animation)
- `GIFGREP_CELL_ASPECT=0.5` (tweak preview cell geometry)

## Test fixtures licensing
See `docs/gif-sources.md`.

## Development
```bash
go test ./...
go run ./cmd/gifgrep --help
```
Ghostty web snapshot:
```bash
pnpm install
pnpm snap
```

## GitHub Pages
Landing page lives in `docs/` (GitHub Pages -> `main` -> `/docs`).
