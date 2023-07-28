package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/color"
	"math"

	"github.com/gravestench/bitstream"
)

const (
	fileSignature = "Woo!\x01"
)

const (
	numHeaderBytes          = 12
	bytesPerGlyph           = 14
	signatureBytesCount     = 5
	unknownHeaderBytesCount = 7
	unknown1BytesCount      = 1
	unknown2BytesCount      = 3
	unknown3BytesCount      = 4
)

// Font represents a displayable font
type Font struct {
	table  []byte
	Glyphs map[rune]*Glyph
	color  color.Color
}

// Load loads a new font from byte slice
func Load(data []byte) (*Font, error) {
	r := bitstream.NewReader(bytes.NewReader(data))

	signature, err := r.Next(signatureBytesCount).Bytes().AsBytes()
	if err != nil {
		return nil, err
	}

	if string(signature) != fileSignature {
		return nil, fmt.Errorf("invalid font table format")
	}

	font := &Font{
		table: data,
		color: color.White,
	}

	if _, err = r.Next(unknownHeaderBytesCount).Bytes().AsBytes(); err != nil {
		return nil, fmt.Errorf("skipping header bytes: %v", err)
	}

	if err = font.initGlyphs(r); err != nil {
		return nil, fmt.Errorf("initializing glyphs: %v", err)
	}

	return font, nil
}

// SetColor sets the fonts color
func (f *Font) SetColor(c color.Color) {
	f.color = c
}

// GetTextMetrics returns the dimensions of the Font element in pixels
func (f *Font) GetTextMetrics(text string) (width, height int) {
	var lineWidth, lineHeight int

	for _, c := range text {
		if c == '\n' {
			width = maxInt(width, lineWidth)
			height += lineHeight
			lineWidth = 0
			lineHeight = 0
		} else if glyph, ok := f.Glyphs[c]; ok {
			lineWidth += glyph.Width()
			lineHeight = maxInt(lineHeight, glyph.Height())
		}
	}

	width = maxInt(width, lineWidth)
	height += lineHeight

	return width, height
}

func (f *Font) initGlyphs(sr *bitstream.Reader) error {
	glyphs := make(map[rune]*Glyph)

	// for i := numHeaderBytes; i < len(f.table); i += bytesPerGlyph {
	for i := numHeaderBytes; true; i += bytesPerGlyph {
		code, err := sr.Next(2).Bytes().AsUInt16()
		if err != nil {
			break
		}

		// byte of 0
		if _, err = sr.Next(unknown1BytesCount).Bytes().AsByte(); err != nil {
			return err
		}

		width, err := sr.Next(1).Bytes().AsByte()
		if err != nil {
			return err
		}

		height, err := sr.Next(1).Bytes().AsByte()
		if err != nil {
			return err
		}

		// 1, 0, 0
		if _, err = sr.Next(unknown2BytesCount).Bytes().AsByte(); err != nil {
			return err
		}

		frame, err := sr.Next(2).Bytes().AsUInt16()
		if err != nil {
			return err
		}

		// 1, 0, 0, character code repeated, and further 0.
		if _, err = sr.Next(unknown3BytesCount).Bytes().AsByte(); err != nil {
			return err
		}

		glyph := newGlyph(int(frame), int(width), int(height))

		glyphs[rune(code)] = glyph
	}

	f.Glyphs = glyphs

	return nil
}

// Marshal encodes font back into byte slice
// Marshal returns the binary representation of the Font data structure.
func (f *Font) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	pad := func(length int) []byte {
		return make([]byte, length)
	}

	// Write the signature
	if err := binary.Write(buf, binary.LittleEndian, []byte(fileSignature)); err != nil {
		return nil, err
	}

	// Write the unknown header bytes
	if err := binary.Write(buf, binary.LittleEndian, pad(unknownHeaderBytesCount)); err != nil {
		return nil, err
	}

	// Write the glyphs
	for code, glyph := range f.Glyphs {
		// Write the character code as a uint16
		if err := binary.Write(buf, binary.LittleEndian, uint16(code)); err != nil {
			return nil, err
		}

		// Write unknown1BytesCount number of 0 bytes
		if err := binary.Write(buf, binary.LittleEndian, pad(unknown1BytesCount)); err != nil {
			return nil, err
		}

		// Write the width and height as bytes
		if err := binary.Write(buf, binary.LittleEndian, byte(glyph.Width())); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.LittleEndian, byte(glyph.Height())); err != nil {
			return nil, err
		}

		// Write unknown2BytesCount number of 0 bytes
		if err := binary.Write(buf, binary.LittleEndian, pad(unknown2BytesCount)); err != nil {
			return nil, err
		}

		// Write the frame as a uint16
		if err := binary.Write(buf, binary.LittleEndian, uint16(glyph.FrameIndex())); err != nil {
			return nil, err
		}

		// Write unknown3BytesCount number of 0 bytes
		if err := binary.Write(buf, binary.LittleEndian, pad(unknown3BytesCount)); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
