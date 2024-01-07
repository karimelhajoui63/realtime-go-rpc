package colorrepo

import (
	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/outbound"
)

type InMemoryColorRepository struct {
	color enum.Color
}

func NewInMemoryColorRepository() (outbound.ColorRepository, error) {
	return &InMemoryColorRepository{}, nil
}

func (imcr *InMemoryColorRepository) Get() (enum.Color, error) {
	return imcr.color, nil
}

func (imcr *InMemoryColorRepository) Update(color enum.Color) {
	imcr.color = color
}
