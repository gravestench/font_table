// Package d2fontglyph represents a single font glyph
package pkg

func newGlyph(frame, width, height int) *Glyph {
	result := &Glyph{
		frame:  frame,
		width:  width,
		height: height,
	}

	return result
}

// Glyph represents a single font glyph
type Glyph struct {
	frame  int
	width  int
	height int
}

// SetSize sets glyph's size to w, h
func (fg *Glyph) SetSize(w, h int) {
	fg.width, fg.height = w, h
}

// Size returns glyph's size
func (fg *Glyph) Size() (w, h int) {
	return fg.width, fg.height
}

// Width returns font width
func (fg *Glyph) Width() int {
	return fg.width
}

// Height returns glyph's height
func (fg *Glyph) Height() int {
	return fg.height
}

// SetFrameIndex sets frame index to idx
func (fg *Glyph) SetFrameIndex(idx int) {
	fg.frame = idx
}

// FrameIndex returns glyph's frame
func (fg *Glyph) FrameIndex() int {
	return fg.frame
}

// Unknown1 returns unknowns bytes
func (fg *Glyph) Unknown1() []byte {
	return []byte{0}
}

// Unknown2 returns unknowns bytes
func (fg *Glyph) Unknown2() []byte {
	return []byte{1, 0, 0}
}

// Unknown3 returns unknowns bytes
func (fg *Glyph) Unknown3() []byte {
	return []byte{0, 0, 0, 0}
}
