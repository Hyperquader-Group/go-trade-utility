package plot

import (
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
