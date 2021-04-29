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
        5.668343999999995,
        45.111511000000014,
        5.852471999999996,
        45.26800200000002,
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

## Testing Issue

A few of the tests are currently failing with floating point differences around the 6th decimal point, a roughly 11 centimeter error (at the equator). This needs to be diagnosed.
