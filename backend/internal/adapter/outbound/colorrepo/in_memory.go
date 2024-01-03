package colorrepo

import (
	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/outbound"
)

type colorRepository struct {
	// TODO: move colorStored here?
}

var colorStored enum.Color = enum.Red

func NewInMemoryColorRepository() (outbound.ColorRepository, error) {
	return &colorRepository{}, nil

}

func (c *colorRepository) Get() (enum.Color, error) {
	return colorStored, nil
}

func (c *colorRepository) Update(color enum.Color) {
	colorStored = color
}
