package geoviewport

import (
	"math"

	sm "github.com/engelsjk/sphericalmercator"
)

type SMCache map[int]sm.SphericalMercator

var (
	smCache = SMCache{}
)

func fetchMerc(tileSize int) sm.SphericalMercator {
	if tileSize == 0 {
		tileSize = 256
	}

	if _, ok := smCache[tileSize]; !ok {
		smCache[tileSize] = sm.New(&sm.Options{Size: tileSize})
	}

	return smCache[tileSize]
}

func getAdjusted(base float64, ratios []float64, allowFloat bool) float64 {
	adjusted := math.Min(
		base-(math.Log(ratios[0])/math.Log(2)),
		base-(math.Log(ratios[1])/math.Log(2)),
	)
	if allowFloat {
		return adjusted
	}
	return math.Floor(adjusted)
}

type VP struct {
	Center []float64
	Zoom   float64
}

func Viewport(bounds, dimensions []float64, minZoom, maxZoom float64, tileSize int, allowFloat bool) VP {
	if minZoom == 0 {
		minZoom = 0
	}
	if maxZoom == 0 {
		maxZoom = 20
	}
	merc := fetchMerc(tileSize)
	base := maxZoom
	bl := merc.Px([]float64{bounds[0], bounds[1]}, base)
	tr := merc.Px([]float64{bounds[2], bounds[3]}, base)
	width := tr[0] - bl[0]
	height := bl[1] - tr[1]
	centerPixelX := bl[0] + (width / 2.0)
	centerPixelY := tr[1] + (height / 2.0)
	ratios := []float64{width / dimensions[0], height / dimensions[1]}
	adjusted := getAdjusted(base, ratios, allowFloat)

	center := merc.LL([]float64{centerPixelX, centerPixelY}, base)
	zoom := math.Max(minZoom, math.Min(maxZoom, adjusted))

	return VP{Center: center, Zoom: zoom}
}

func Bounds(viewport []float64, zoom float64, dimensions []float64, tileSize int) []float64 {
	merc := fetchMerc(tileSize)
	px := merc.Px(viewport, zoom)
	tl := merc.LL(
		[]float64{
			px[0] - (dimensions[0] / 2.0),
			px[1] - (dimensions[1] / 2.0),
		}, zoom)
	br := merc.LL(
		[]float64{
			px[0] - (dimensions[0] / 2.0),
			px[1] - (dimensions[1] / 2.0),
		}, zoom)
	return []float64{tl[0], br[1], br[0], tl[1]}
}
