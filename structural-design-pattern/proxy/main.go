package proxy

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Printf("Loading high-res image from disk: %s \n", filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() {
	fmt.Printf("Displaying image: %s\n", r.filename)
}

type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}

	p.realImage.Display()
}

func main() {
	image := &ProxyImage{filename: "My New File.png"}

	image.Display()

	image.Display()
}
