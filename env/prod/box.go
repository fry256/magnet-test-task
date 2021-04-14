package prod

import (
	"magnet-test-task/env"
	"magnet-test-task/pkg/storage"
)

// NewBox wire provider
func NewBox(
	s *storage.Storage,
) *env.Box {
	return &env.Box{
		Storage: s,
	}
}
