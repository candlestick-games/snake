package loader

import (
	"io"
	"io/fs"

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFile(fs fs.FS, filename string) fs.File {
	file, err := fs.Open(filename)
	if err != nil {
		log.Fatal("Open file", "filename", filename, "error", err)
	}
	return file
}

func LoadData(fs fs.FS, filename string) []byte {
	file := LoadFile(fs, filename)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Read data", "filename", filename, "error", err)
	}
	return data
}

func LoadImage(fs fs.FS, filename string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(fs, filename)
	if err != nil {
		log.Fatal("Load image", "filename", filename, "error", err)
	}
	return img
}

func LoadFont(fs fs.FS, filename string) *opentype.Font {
	openTypeFont, err := opentype.Parse(LoadData(fs, filename))
	if err != nil {
		log.Fatal("Parse font", "filename", filename, "error", err)
	}
	return openTypeFont
}

func NewFontFace(openTypeFont *opentype.Font, size float64) font.Face {
	fontFace, err := opentype.NewFace(openTypeFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Create font face", "error", err)
	}
	return fontFace
}
