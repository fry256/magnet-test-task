package app

import "magnet-test-task/env"

type Service struct {
	Box *env.Box
}

// New инициализирует сервис
func New(box *env.Box) *Service {
	return &Service{box}
}

func (srv *Service) OnShutdown() {
	// do smth on shutdown...
}
