package inbound

import (
	"rpc-server/internal/core/domain/enum"
)

type PainterUseCase interface {
	ChangeColor(enum.Color) error
	GetColor() (enum.Color, error)
	GetColorStream() (<-chan enum.Color, error)
}
