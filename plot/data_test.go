package plot

import (
	"github.com/go-gota/gota/dataframe"
	"os"
	"path/filepath"
	"testing"

	"fmt"
)

func TestSaveForLine(t *testing.T) {
	data := NewLine("test", "timestamp", "y", []float64{10, 20, 20, 10})

	dir, _ := os.Getwd()
	filename := filepath.Join(dir, "img", "line.png")
	err := data.Save(filename)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("filename: %s\n", filename)
}

func TestSaveForScatter(t *testing.T) {
	data := NewScatter("test", "x", "y", []float64{10, 20}, []float64{20, 10})

	dir, _ := os.Getwd()
	filename := filepath.Join(dir, "img", "scatter.png")
	err := data.Save(filename)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("filename: %s\n", filename)
}

func TestSubPlot(t *testing.T) {
	f, _ := os.Open("/Users/numb/Google ドライブ/02_Work/project_bitcoin/価格変動無い出来高/volume - volume.csv.csv")
	df := dataframe.ReadCSV(f)
	defer f.Close()

	fmt.Printf("%+v\n", df)

	s := NewSeaborn(nil)
	s.Data = df
	if err := s.SubPlot(8, 8, "./thisplot.png"); err != nil {
		t.Fatal(err)
	}
}
