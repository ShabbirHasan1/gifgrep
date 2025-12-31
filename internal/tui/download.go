package tui

import (
	"bufio"

	"github.com/steipete/gifgrep/internal/download"
	"github.com/steipete/gifgrep/internal/reveal"
)

func downloadSelected(state *appState, out *bufio.Writer) {
	if state.selected < 0 || state.selected >= len(state.results) {
		state.status = "No selection"
		state.renderDirty = true
		return
	}
	item := state.results[state.selected]
	if item.URL == "" {
		state.status = "No URL"
		state.renderDirty = true
		return
	}
	state.status = "Downloading..."
	state.renderDirty = true
	render(state, out, state.lastRows, state.lastCols)
	_ = out.Flush()

	filePath, err := download.ToDownloads(item)
	if err != nil {
		state.status = "Download error: " + err.Error()
		state.renderDirty = true
		return
	}
	state.lastSavedPath = filePath
	if state.opts.Reveal {
		if err := reveal.Reveal(filePath); err != nil {
			state.status = "Saved " + filePath + " (reveal failed)"
			state.renderDirty = true
			return
		}
		state.status = "Saved " + filePath + " (revealed)"
		state.renderDirty = true
		return
	}
	state.status = "Saved " + filePath
	state.renderDirty = true
}
