package perpetua

import (
	"image/color"
	"strings"

	internalcolor "github.com/perpetuatheme/go/internal/color"
)

//go:generate go run generate.go

type Palette interface {
	Name() string
	Red() Color
	Orange() Color
	Yellow() Color
	Lime() Color
	Green() Color
	Turquoise() Color
	Cyan() Color
	Cerulean() Color
	Blue() Color
	Violet() Color
	Lavender() Color
	Pink() Color
	RedBack() Color
	OrangeBack() Color
	YellowBack() Color
	LimeBack() Color
	GreenBack() Color
	TurquoiseBack() Color
	CyanBack() Color
	CeruleanBack() Color
	BlueBack() Color
	VioletBack() Color
	LavenderBack() Color
	PinkBack() Color
	Base0() Color
	Base1() Color
	Base2() Color
	Base3() Color
	Base4() Color
	Base5() Color
	Over0() Color
	Over1() Color
	Over2() Color
	Text0() Color
	Text1() Color
	Text2() Color
}

func FromAmbience(ambience string) Palette {
	for _, p := range []Palette{Light, Dark} {
		if strings.EqualFold(p.Name(), ambience) {
			return p
		}
	}
	return nil
}

type Color struct {
	Name  string
	Index uint

	hex   string
	rgb   color.RGBA
	okhsl internalcolor.Okhsl
}

var _ color.Color = Color{}

func (c Color) Hex() string {
	return c.hex
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return c.rgb.RGBA()
}

func (c Color) Okhsl() (h, s, l float64) {
	return c.okhsl.HSL()
}
