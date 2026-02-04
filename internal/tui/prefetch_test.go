package tui

import (
	"os"
	"testing"

	"github.com/steipete/gifgrep/internal/testutil"
)

func TestPrefetchGIFToTempRespectsSize(t *testing.T) {
	data := testutil.MakeTestGIF()
	rt := &testutil.FakeTransport{GIFData: data}
	testutil.WithTransport(t, rt, func() {
		dir := t.TempDir()
		path, err := prefetchGIFToTemp("https://example.test/full.gif", dir, int64(len(data)+1))
		if err != nil {
			t.Fatalf("prefetch failed: %v", err)
		}
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("missing temp file: %v", err)
		}

		if _, err := prefetchGIFToTemp("https://example.test/full.gif", dir, int64(len(data)-1)); err == nil {
			t.Fatalf("expected size cap error")
		}
	})
}
