package multi

import (
	"context"
	"sync"
)

// Spawner is responsible for spawning various instances of the waffle services.
type Spawner struct {
	mu sync.Mutex
}

func NewSpawner() *Spawner {
	return &Spawner{}
}

func (s *Spawner) Run(ctx context.Context) error {

}
