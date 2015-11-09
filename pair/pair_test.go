package pair_test

import (
	"github.com/acrmp/pairshaped/pair"
	"github.com/lazywei/go-opencv/opencv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeCamera struct {
	grab     bool
	retrieve bool
	reqFrame int
	released bool
}

func (c *fakeCamera) GrabFrame() bool {
	return c.grab
}

func (c *fakeCamera) RetrieveFrame(i int) *opencv.IplImage {
	c.reqFrame = i
	if !c.retrieve {
		return nil
	}
	return &opencv.IplImage{}
}

func (c *fakeCamera) Release() {
	c.released = true
}

type fakeCascade struct {
	faces    int
	released bool
}

func (c *fakeCascade) DetectObjects(*opencv.IplImage) []*opencv.Rect {
	return make([]*opencv.Rect, c.faces)
}

func (c *fakeCascade) Release() {
	c.released = true
}

var _ = Describe("Pair", func() {
	Describe("Image capture", func() {
		var c *fakeCamera
		var cs *fakeCascade

		BeforeEach(func() {
			cs = &fakeCascade{}
		})
		Context("when the camera is unavailable", func() {
			It("errors", func() {
				c = &fakeCamera{grab: false}
				p := &pair.Checker{Camera: c, Cascade: cs}
				t, err := p.Present()
				Expect(t).To(BeFalse())
				Expect(err).To(MatchError("Unable to grab frame from the camera"))
			})
		})
		Context("when an image is not captured", func() {
			It("errors", func() {
				c = &fakeCamera{grab: true, retrieve: false}
				p := &pair.Checker{Camera: c, Cascade: cs}
				t, err := p.Present()
				Expect(t).To(BeFalse())
				Expect(err).To(MatchError("Unable to retrieve image from the camera"))
			})
		})
		Context("when an image is captured", func() {
			It("retrieves the first frame", func() {
				c = &fakeCamera{grab: true, retrieve: true}
				p := &pair.Checker{Camera: c, Cascade: cs}
				p.Present()
				Expect(c.reqFrame).To(Equal(1))
			})
		})
		AfterEach(func() {
			Expect(c.released).To(BeTrue())
		})
	})
	Describe("Face detection", func() {
		var cs *fakeCascade
		var p *pair.Checker

		BeforeEach(func() {
			cs = &fakeCascade{}
			p = &pair.Checker{
				Camera:  &fakeCamera{grab: true, retrieve: true},
				Cascade: cs,
			}
		})
		Context("when there are no faces", func() {
			It("is false", func() {
				cs.faces = 0
				t, err := p.Present()
				Expect(t).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("when there is a lonely developer", func() {
			It("is false", func() {
				cs.faces = 1
				t, err := p.Present()
				Expect(t).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("when there is a pair rockin it", func() {
			It("is true", func() {
				cs.faces = 2
				t, err := p.Present()
				Expect(t).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("when there is a mob of hipsters", func() {
			It("is true", func() {
				cs.faces = 8
				t, err := p.Present()
				Expect(t).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			})
		})
		AfterEach(func() {
			Expect(cs.released).To(BeTrue())
		})
	})
})
