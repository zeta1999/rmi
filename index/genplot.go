package index

import (
	"image/color"
	"math"
	"rmi/linear"

	"gonum.org/v1/gonum/floats"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

/*
Genplot takes an index and plots its keys, CDF, its approximation, and writes a plot.png file in assets/folder
*/
func Genplot(index *LearnedIndex, indexedCol []float64, plotfilepath string) {
	linearRegFn := func(x float64) float64 { return index.M.Predict(x)*float64(index.Len) - 1 }
	x, _ := linear.Cdf(indexedCol)
	idxFromCDF := func(i float64) float64 { return stat.CDF(i, stat.Empirical, x, nil)*float64(index.Len) - 1 }

	p, _ := plot.New()
	p.Title.Text = "Learned Index RMI"
	p.X.Label.Text = "Age"
	p.Y.Label.Text = "Index"

	courbeKeys := plotter.XYs{}
	courbePreds := plotter.XYs{}
	for i, k := range x {
		courbeKeys = append(courbeKeys, plotter.XY{X: k, Y: float64(i)})
		pred := math.Round(index.M.Predict(k)*float64(index.Len) - 1)
		courbePreds = append(courbePreds, plotter.XY{X: k, Y: pred})
	}
	approxFn := plotter.NewFunction(linearRegFn)
	approxFn.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	approxFn.Width = vg.Points(2)
	approxFn.Color = color.RGBA{G: 255, A: 255}

	cdfFn := plotter.NewFunction(idxFromCDF)
	cdfFn.Width = vg.Points(1)
	cdfFn.Color = color.RGBA{A: 255, B: 255}

	plotutil.AddLinePoints(p, "Keys", courbeKeys)
	p.Add(approxFn)
	p.Legend.Add("Approx (lr)", approxFn)
	p.Add(cdfFn)
	p.Legend.Add("CDF", cdfFn)
	p.X.Min = 0
	p.X.Max = floats.Max(index.ST.keys)
	p.Y.Min = 0
	p.Y.Max = float64(index.Len)
	p.Save(4*vg.Inch, 4*vg.Inch, plotfilepath)
}
