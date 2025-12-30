package gifdecode

import (
	"bytes"
	"image/png"
	"testing"
)

type fixtureCase struct {
	name             string
	file             string
	minFrames        int
	wantTransparency bool
}

func TestDecodeFixtures(t *testing.T) {
	cases := []fixtureCase{
		{name: "animexample2", file: "animexample2.gif", minFrames: 2, wantTransparency: false},
		{name: "youtube-loading-3", file: "youtube-loading-3.gif", minFrames: 2, wantTransparency: true},
		{name: "knowledge-human-pink", file: "knowledge-human-pink.gif", minFrames: 2, wantTransparency: true},
		{name: "animation-loading1", file: "animation-loading1.gif", minFrames: 2, wantTransparency: true},
		{name: "walk-cycle", file: "walk-cycle.gif", minFrames: 2, wantTransparency: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := readFixture(t, tc.file)
			frames, err := Decode(data, DefaultOptions())
			if err != nil {
				t.Fatalf("decode failed: %v", err)
			}
			if frames.Width <= 0 || frames.Height <= 0 {
				t.Fatalf("invalid size %dx%d", frames.Width, frames.Height)
			}
			if len(frames.Frames) < tc.minFrames {
				t.Fatalf("expected at least %d frames, got %d", tc.minFrames, len(frames.Frames))
			}
			if frames.Frames[0].Delay <= 0 {
				t.Fatalf("expected positive delay")
			}
			img, err := png.Decode(bytes.NewReader(frames.Frames[0].PNG))
			if err != nil {
				t.Fatalf("png decode failed: %v", err)
			}
			b := img.Bounds()
			if b.Dx() != frames.Width || b.Dy() != frames.Height {
				t.Fatalf("expected png %dx%d, got %dx%d", frames.Width, frames.Height, b.Dx(), b.Dy())
			}
			if tc.wantTransparency && !hasTransparentPixel(frames) {
				t.Fatalf("expected transparency")
			}
		})
	}
}

func hasTransparentPixel(frames *Frames) bool {
	for _, frame := range frames.Frames {
		img, err := png.Decode(bytes.NewReader(frame.PNG))
		if err != nil {
			continue
		}
		b := img.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				_, _, _, a := img.At(x, y).RGBA()
				if a < 0xffff {
					return true
				}
			}
		}
	}
	return false
}
