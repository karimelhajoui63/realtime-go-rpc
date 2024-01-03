package outbound

import "rpc-server/internal/core/domain/enum"

type ColorRepository interface {
	Get() (enum.Color, error)
	Update(enum.Color)
}
