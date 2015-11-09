package main

import (
	"github.com/acrmp/pairshaped/pair"
	"github.com/lazywei/go-opencv/opencv"
	"os"
	"path"
	"path/filepath"
)

func main() {
	hd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	cp := path.Join(hd, "haarcascade_frontalface_alt.xml")

	if _, err := os.Stat(cp); err != nil {
		panic(err)
	}

	cm := opencv.NewCameraCapture(0)
	cs := opencv.LoadHaarClassifierCascade(cp)

	p := &pair.Checker{Camera: cm, Cascade: cs}
	t, err := p.Present()

	if err != nil {
		panic(err)
	}

	if !t {
		os.Stderr.WriteString("Your pair must be playing ping pong\n")
		os.Exit(1)
	}
}
