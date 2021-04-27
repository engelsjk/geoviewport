package geoviewport

import (
	"math"
	"testing"
)

type viewport struct {
	Center []float64
	Zoom   float64
}

var decDegreesFloatTolerance = 8.0

func precisionRound(number, precision float64) float64 {
	factor := math.Pow(10, precision)
	return math.Round(number*factor) / factor
}

var sampleBounds = []float64{
	5.668343999999995,
	45.111511000000014,
	5.852471999999996,
	45.26800200000002,
}

var expectedCenter = []float64{
	5.760407969355583,
	45.189810341718136,
}

func isApproximatelyEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-10
}

func areViewportsApproximatelyEqual(v1, v2 viewport) bool {
	return isApproximatelyEqual(v1.Center[0], v2.Center[0]) &&
		isApproximatelyEqual(v1.Center[1], v2.Center[1]) &&
		isApproximatelyEqual(v1.Zoom, v2.Zoom)
}

func TestViewport(t *testing.T) {

	var center []float64
	var zoom float64
	var got, want viewport

	center, zoom = Viewport(sampleBounds, []float64{640, 480}, 0, 0, 0, false)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: expectedCenter, Zoom: 11}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#1: got %v; want %v", got, want)
	}

	center, zoom = Viewport(sampleBounds, []float64{64, 48}, 0, 0, 0, false)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: expectedCenter, Zoom: 8}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#2: got %v; want %v", got, want)
	}

	center, zoom = Viewport(sampleBounds, []float64{10, 10}, 0, 0, 0, false)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: expectedCenter, Zoom: 5}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#3: got %v; want %v", got, want)
	}
}

func TestViewportSouthernHemisphere(t *testing.T) {

	var center []float64
	var zoom float64
	var got, want viewport

	center, zoom = Viewport([]float64{10, -20, 20, -10}, []float64{500, 250}, 0, 0, 0, false)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: []float64{14.999999776482582, -15.058651551491899}, Zoom: 5}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#1: got %v; want %v", got, want)
	}

	center, zoom = Viewport([]float64{-10, -60, 10, -30}, []float64{500, 250}, 0, 0, 0, false)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: []float64{0, -47.05859720188612}, Zoom: 2}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#2: got %v; want %v", got, want)
	}
}

func TestBounds512pxTiles(t *testing.T) {

	var got, want float64

	bounds := Bounds([]float64{-77.036556, 38.897708}, 17, []float64{1080, 350}, 512)
	xMin := bounds[0]
	yMin := bounds[1]
	xMax := bounds[2]
	yMax := bounds[3]

	got = precisionRound(xMin, decDegreesFloatTolerance)
	want = -77.03945339
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(yMin, decDegreesFloatTolerance)
	want = 38.89697827
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(xMax, decDegreesFloatTolerance)
	want = -77.03365982
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(yMax, decDegreesFloatTolerance)
	want = 38.89843951
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestBoundsForFloatZooms(t *testing.T) {

	var got, want float64

	zoom := 16.52
	bounds := Bounds([]float64{-77.036556, 38.897708}, zoom, []float64{1080, 350}, 512)
	xMin := bounds[0]
	yMin := bounds[1]
	xMax := bounds[2]
	yMax := bounds[3]

	got = precisionRound(xMin, decDegreesFloatTolerance)
	want = -77.04059627
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(yMin, decDegreesFloatTolerance)
	want = 38.89668897
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(xMax, decDegreesFloatTolerance)
	want = -77.03251573
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got = precisionRound(yMax, decDegreesFloatTolerance)
	want = 38.89872702
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestViewportForFloatZooms(t *testing.T) {

	var center []float64
	var zoom float64
	var got, want viewport

	center, zoom = Viewport(sampleBounds, []float64{10, 10}, 0, 0, 256, true)
	got = viewport{Center: center, Zoom: zoom}
	want = viewport{Center: expectedCenter, Zoom: 5.984828902182182}
	if !areViewportsApproximatelyEqual(got, want) {
		t.Errorf("#2: got %v; want %v", got, want)
	}
}
