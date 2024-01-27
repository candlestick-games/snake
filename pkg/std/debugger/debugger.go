package debugger

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go.uber.org/atomic"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/candlestick-games/snake/pkg/std/loader"
	"github.com/candlestick-games/snake/pkg/std/pencil"
)

var (
	debug = atomic.NewBool(false)

	//go:embed JetBrainsMonoNL-Regular.ttf
	debugFontData []byte

	debugFontFace font.Face

	screenInfo  = false
	cursorCoord = false
)

func init() {
	openTypeFont, err := opentype.Parse(debugFontData)
	if err != nil {
		log.Fatal("Parse debug font", "error", err)
	}

	debugFontFace = loader.NewFontFace(openTypeFont, 16)
}

func Enabled() bool {
	return debug.Load()
}

func Enable() {
	debug.Store(true)
}

func Disable() {
	debug.Store(false)
}

func Toggle() {
	debug.Toggle()
}

func Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		Toggle()
	}

	if !Enabled() {
		return
	}

	if !ebiten.IsKeyPressed(ebiten.KeyF4) {
		return
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		screenInfo = !screenInfo
	case inpututil.IsKeyJustPressed(ebiten.KeyC):
		cursorCoord = !cursorCoord
	}
}

func Draw(screen *ebiten.Image) {
	if !Enabled() {
		return
	}

	screenSize := screen.Bounds().Max

	debugText := strings.Builder{}
	debugText.WriteString(fmt.Sprintf("FPS: %.2f\nTPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))

	if screenInfo {
		debugText.WriteString(fmt.Sprintf("\nScreen: %dx%d", screenSize.X, screenSize.Y))
		debugText.WriteString(fmt.Sprintf("\nDevice scale: %.2f", ebiten.DeviceScaleFactor()))
	}

	const dx = 1
	pencil.TextTopLeft(
		screen, debugFontFace,
		debugText.String(),
		8+dx, 8+dx,
		colornames.Gray,
	)

	pencil.TextTopLeft(
		screen, debugFontFace,
		debugText.String(),
		8, 8,
		colornames.White,
	)

	if cursorCoord {
		cursorX, cursorY := ebiten.CursorPosition()
		x, y := float64(cursorX), float64(cursorY)

		const crossSize = 8
		const crossThinness = 2
		pencil.Line(
			screen,
			x-crossSize, y,
			x+crossSize, y,
			crossThinness,
			colornames.White,
		)
		pencil.Line(
			screen,
			x, y-crossSize,
			x, y+crossSize,
			crossThinness,
			colornames.White,
		)

		pencil.TextTopLeft(
			screen, debugFontFace,
			fmt.Sprintf("%d %d", cursorX, cursorY),
			x+8, y+16,
			colornames.White,
		)
	}
}
