package main

import (
  "bytes"
  "flag"
  "fmt"
  "image"
  "image/png"
  "os"
  "strconv"
  "strings"

  "github.com/vicanso/go-charts/v2"
)

func parseValues(vals string) ([]float64, error) {
  strVals := strings.Split(vals, ",")
  if len(strVals) != 6 {
    return nil, fmt.Errorf("exactly 6 values are required, but got %d", len(strVals))
  }
  floatVals := make([]float64, len(strVals))
  for i, strVal := range strVals {
    val, err := strconv.ParseFloat(strVal, 64)
    if err != nil {
      return nil, err
    }
    floatVals[i] = val
  }
  return floatVals, nil
}


func main() {
  vals := flag.String("vals", "", "Comma-separated list of values")
  flag.Parse()

  if *vals == "" {
    fmt.Println("Please provide values using --vals")
    return
  }

  values, err := parseValues(*vals)
  if err != nil {
    fmt.Println("Error parsing values:", err)
    return
  }

  data := [][]float64{values}

  charts.AddTheme("custom", charts.ThemeOption{
    IsDarkMode: true,
    AxisStrokeColor: charts.Color{ R: 0, G: 0, B: 0, A: 0 },
    AxisSplitLineColor: charts.Color{ R: 72, G: 71, B: 83, A: 255 },
    BackgroundColor: charts.Color{ R: 0, G: 0, B: 0, A: 0 },
    TextColor: charts.Color{ R: 238, G: 238, B: 238, A: 255 },
    SeriesColors: []charts.Color{
      {R: 212, G: 110, B: 37, A: 255},
    },
  })

  p, err := charts.RadarRender(data,
    charts.WidthOptionFunc(800),
    charts.HeightOptionFunc(800),
    charts.ThemeOptionFunc("custom"),
    charts.RadarIndicatorOptionFunc([]string{
      "", "", "", "", "", "",
    }, []float64{
        10, 10, 10, 10, 10, 10,
      }),
    )

  if err != nil {
    panic(err)
  }

  buf, err := p.Bytes()
  if err != nil {
    panic(err)
  }

  img, err := png.Decode(bytes.NewReader(buf))
  if err != nil {
    panic(err)
  }

  bounds := img.Bounds()
  minX, minY, maxX, maxY := bounds.Max.X, bounds.Max.Y, bounds.Min.X, bounds.Min.Y

  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      _, _, _, a := img.At(x, y).RGBA()
      if a != 0 {
        if x < minX {
          minX = x
        }
        if y < minY {
          minY = y
        }
        if x > maxX {
          maxX = x
        }
        if y > maxY {
          maxY = y
        }
      }
    }
  }

  croppedImg := img.(interface {
    SubImage(r image.Rectangle) image.Image
  }).SubImage(image.Rect(minX, minY, maxX+1, maxY+1))

  var croppedBuf bytes.Buffer
  err = png.Encode(&croppedBuf, croppedImg)
  if err != nil {
    panic(err)
  }

  err = os.WriteFile("out.png", croppedBuf.Bytes(), 0600)
  if err != nil {
    panic(err)
  }
}
