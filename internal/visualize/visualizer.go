package visualize

import "waffle/internal/request"

type Visualizer struct {
	requestWrappersChan chan request.Wrapper
}

func NewVisualizer() *Visualizer {
	return &Visualizer{
		requestWrappersChan: make(chan request.Wrapper),
	}
}

func (v *Visualizer) VisualizeRequestWrapper(rw request.Wrapper) {
	v.requestWrappersChan <- rw
}
