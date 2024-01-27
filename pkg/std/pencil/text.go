package pencil

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func TextCentered(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignCenter, AlignCenter, x, y, clr)
}

func TextTopLeft(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignTop, AlignLeft, x, y, clr)
}

func TextTopRight(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignTop, AlignRight, x, y, clr)
}

func TextBottomLeft(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignBottom, AlignLeft, x, y, clr)
}

func TextBottomRight(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignBottom, AlignRight, x, y, clr)
}

func TextBottomCenter(dst *ebiten.Image, fontFace font.Face, s string, x, y float64, clr color.Color) {
	TextAligned(dst, fontFace, s, AlignBottom, AlignCenter, x, y, clr)
}

type Alignment uint

const (
	_ Alignment = iota
	AlignStart
	AlignEnd
	AlignCenter

	AlignLeft   = AlignStart
	AlignRight  = AlignEnd
	AlignTop    = AlignStart
	AlignBottom = AlignEnd
)

func TextAligned(
	dst *ebiten.Image, fontFace font.Face, s string, vertical, horizontal Alignment, x, y float64, clr color.Color,
) {
	lines := countLines(s)
	if lines == 0 {
		return
	}

	var longestLine string
	if lines == 1 {
		longestLine = s
	} else {
		ls, le := longestLineIndexes(s, lines)
		longestLine = s[ls:le]
	}

	bounds, _ := font.BoundString(fontFace, longestLine)

	minX := float64(bounds.Min.X.Round())
	minY := float64(bounds.Min.Y.Round())

	maxX := float64((bounds.Max.X - bounds.Min.X).Round())
	maxY := float64((bounds.Max.Y-bounds.Min.Y).Round())*float64(lines) + float64(lines-1)*float64(fontFace.Metrics().XHeight.Round())

	var dstX, dstY float64

	switch horizontal {
	case AlignStart:
		dstX = x - minX
	case AlignCenter:
		dstX = x - minX - maxX/2
	case AlignEnd:
		dstX = x - minX - maxX
	default:
		panic(fmt.Errorf("unsupported horizontal alignment: %d", horizontal))
	}

	switch vertical {
	case AlignStart:
		dstY = y - minY
	case AlignCenter:
		dstY = y - minY - maxY/2
	case AlignEnd:
		dstY = y - minY - maxY
	default:
		panic(fmt.Errorf("unsupported vertical alignment: %d", vertical))
	}

	textAt(dst, 1, fontFace, s, dstX, dstY, clr)
}

func textAt(dst *ebiten.Image, deviceScale float64, font font.Face, s string, x, y float64, clr color.Color) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(clr)
	op.GeoM.Scale(deviceScale, deviceScale)
	op.GeoM.Translate(x*deviceScale, y*deviceScale)
	text.DrawWithOptions(dst, s, font, op)
}

func countLines(s string) int {
	if s == "" {
		return 0
	}
	return 1 + strings.Count(s, "\n")
}

func longestLineIndexes(s string, lines int) (start, end int) {
	longestLength := 0
	longestStart := 0
	longestEnd := 0

	startIndex := 0
	endIndex := strings.IndexByte(s, '\n')

	if lines <= 0 {
		return -1, -1
	}

	for i := 0; i < lines && endIndex != -1; i++ {
		lineLength := endIndex - startIndex

		if lineLength > longestLength {
			longestLength = lineLength
			longestStart = startIndex
			longestEnd = endIndex
		}

		startIndex = endIndex + 1
		endIndex = strings.IndexByte(s[startIndex:], '\n')

		if endIndex != -1 {
			endIndex += startIndex
		}
	}

	lastLineLength := len(s) - startIndex
	if lastLineLength > longestLength {
		longestStart = startIndex
		longestEnd = len(s)
	}

	return longestStart, longestEnd
}
