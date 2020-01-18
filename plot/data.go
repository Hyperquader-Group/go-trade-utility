package plot

import (
	"fmt"

	pp "gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Outputer interface {
	Save(output string) error
}

type Line struct {
	Title string
	// Value
	Length         int
	xLabel, yLabel string
	Value          []float64
}

func NewLine(title, xLabel, yLabel string, values []float64) *Line {
	return &Line{
		Title:  title,
		Length: len(values),
		xLabel: xLabel,
		yLabel: yLabel,
		Value:  values,
	}
}

func (p *Line) Save(output string) error {
	if p.Length < 1 {
		return fmt.Errorf("struct has not length")
	}

	plt, err := pp.New()
	if err != nil {
		return err
	}

	plt.Title.Text = p.Title
	plt.Y.Label.Text = "data"
	if p.yLabel != "" {
		plt.Y.Label.Text = p.yLabel
	}
	plt.X.Label.Text = "time"
	if p.xLabel != "" {
		plt.X.Label.Text = p.xLabel
	}

	points := make(plotter.XYs, p.Length)
	for i := range points {
		points[i].X = float64(i)
		points[i].Y = p.Value[i]
	}

	if err = plotutil.AddLines(plt, points); err != nil {
		return err
	}

	if err := plt.Save(4*vg.Inch, 3*vg.Inch, output); err != nil {
		return err
	}

	return nil
}

type Scatter struct {
	Title string
	// Value
	Length         int
	xLabel, yLabel string
	xValue, yValue []float64
}

func NewScatter(title, xLabel, yLabel string, xValues, yValues []float64) *Scatter {
	if len(xValues) != len(yValues) {
		return &Scatter{}
	}

	return &Scatter{
		Title:  title,
		Length: len(xValues),
		xLabel: xLabel,
		yLabel: yLabel,
		xValue: xValues,
		yValue: yValues,
	}
}

func (p *Scatter) Save(output string) error {
	if p.Length < 1 {
		return fmt.Errorf("struct has not length")
	}

	plt, err := pp.New()
	if err != nil {
		return err
	}

	plt.Title.Text = p.Title
	plt.Y.Label.Text = "data_y"
	if p.yLabel != "" {
		plt.Y.Label.Text = p.yLabel
	}
	plt.X.Label.Text = "data_x"
	if p.xLabel != "" {
		plt.X.Label.Text = p.xLabel
	}

	points := make(plotter.XYs, p.Length)
	for i := range points {
		points[i].X = p.xValue[i]
		points[i].Y = p.yValue[i]
	}

	if err = plotutil.AddScatters(plt, points); err != nil {
		return err
	}

	if err := plt.Save(4*vg.Inch, 4*vg.Inch, output); err != nil {
		return err
	}

	return nil
}
