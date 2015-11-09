package pair

import (
	"errors"
	"github.com/lazywei/go-opencv/opencv"
)

type Camera interface {
	GrabFrame() bool
	RetrieveFrame(int) *opencv.IplImage
	Release()
}

type Cascade interface {
	DetectObjects(*opencv.IplImage) []*opencv.Rect
	Release()
}

type Checker struct {
	Camera  Camera
	Cascade Cascade
}

func (c *Checker) Present() (bool, error) {
	defer c.Camera.Release()
	defer c.Cascade.Release()

	ok := c.Camera.GrabFrame()
	if !ok {
		return false, errors.New("Unable to grab frame from the camera")
	}
	i := c.Camera.RetrieveFrame(1)
	if i == nil {
		return false, errors.New("Unable to retrieve image from the camera")
	}
	f := c.Cascade.DetectObjects(i)
	return len(f) >= 2, nil
}
