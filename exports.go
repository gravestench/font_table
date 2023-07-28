package font_table

import (
	"github.com/gravestench/font_table/pkg"
)

type (
	Font  = pkg.Font
	Glyph = pkg.Glyph
)

func Load(data []byte) (*Font, error) {
	return pkg.Load(data)
}
