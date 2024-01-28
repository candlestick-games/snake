package space

import "github.com/hajimehoshi/ebiten/v2"

func ImgResizeTo(geoM ebiten.GeoM, img *ebiten.Image, size Vec2F) ebiten.GeoM {
	imgSize := ImageSize(img).ToF()
	ratio := size.Div(imgSize)
	return Scale(geoM, Vec2F{}, ratio)
}

func Translate(geoM ebiten.GeoM, pos Vec2F) ebiten.GeoM {
	geoM.Translate(pos.X, pos.Y)
	return geoM
}

func Rotate(geoM ebiten.GeoM, origin Vec2F, angle float64) ebiten.GeoM {
	geoM.Translate(-origin.X, -origin.Y)
	geoM.Rotate(angle)
	geoM.Translate(origin.X, origin.Y)
	return geoM
}

func Scale(geoM ebiten.GeoM, origin Vec2F, scale Vec2F) ebiten.GeoM {
	geoM.Translate(-origin.X, -origin.Y)
	geoM.Scale(scale.X, scale.Y)
	geoM.Translate(origin.X, origin.Y)
	return geoM
}

func Flip(geoM ebiten.GeoM, origin Vec2F, horizontal, vertical bool) ebiten.GeoM {
	s := Vec2F{
		X: 1,
		Y: 1,
	}
	if horizontal {
		s.X = -1
	}
	if vertical {
		s.Y = -1
	}
	return Scale(geoM, origin, s)
}
