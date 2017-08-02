package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestMinAvg(t *testing.T) {
	tests := []struct {
		in  string
		out uint64
	}{
		{"input01", 9},
		{"input02", 6},
		{"input04", 80289690037},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			path := filepath.Join("testdata", tc.in+".txt")
			r, close := setupTest(path)
			defer close()

			got := MinAvgTime(r)
			if got != tc.out {
				t.Errorf("got %v, want %d", got, tc.out)
			}
		})
	}

}

func setupTest(path string) (io.Reader, func()) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f, func() { f.Close() }
}
