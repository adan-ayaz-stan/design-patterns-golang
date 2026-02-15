package facade

import "fmt"

type Amplifier struct{}

func (a *Amplifier) On() {
	fmt.Println("Amp on")
}

func (a *Amplifier) SetVolume(v int) {
	fmt.Println("Volume set to", v)
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector on")
}

type Lights struct{}

func (l *Lights) Dim() {
	fmt.Println("Lights dimmed")
}

type HomeTheaterFacade struct {
	amp       *Amplifier
	projector *Projector
	lights    *Lights
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("ready to watch a movie...")
	h.lights.Dim()
	h.projector.On()
	h.amp.On()
	h.amp.SetVolume(10)
}

func main() {
	theater := &HomeTheaterFacade{&Amplifier{}, &Projector{}, &Lights{}}
	theater.WatchMovie()
}
