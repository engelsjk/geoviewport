# geoviewport

A Go port of [mapbox/geo-viewport](https://github.com/mapbox/geo-viewport).

## Installation

```bash
go get github.com/engelsjk/geoviewport
```

## Example

```Go
center, zoom := geoviewport.Viewport(
    []float64{
        5.668343999999995, 45.111511000000014,
        5.852471999999996, 45.26800200000002,
    },
    []float64{640, 480},
    0, 0,
    0,
    false,
)

// yields
center: [5.76040800000003 45.18981028279085]
zoom: 11
```

```Go
bounds := geoviewport.Bounds(
    []float64{-75.03, 35.25},
    14,
    []float64{600, 400},
    256,
)

// yields
bounds: [-75.05574920654297 35.2359802066683 -75.00425079345703 35.26401736929553]
```

## Testing Issue

A few of the tests are currently failing with floating point differences around the 6th decimal point, a roughly 11 centimeter error (at the equator). This needs to be diagnosed.
