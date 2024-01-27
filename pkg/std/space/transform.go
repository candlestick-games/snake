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

func Scale(geoM ebiten.GeoM, origin Vec2F, scale Vec2F) ebiten.GeoM {
	geoM.Translate(-origin.X, -origin.Y)
	geoM.Scale(scale.X, scale.Y)
	geoM.Translate(origin.X, origin.Y)
	return geoM
}
