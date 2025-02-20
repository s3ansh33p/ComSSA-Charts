# ComSSA-Charts

Will override the radar rendering:

`git clone https://github.com/vicanso/go-charts`

Edit line go-charts/radar_chart.go:224
```diff
+ FillColor:   color.WithAlpha(120),
- FillColor:   color.WithAlpha(20),
```