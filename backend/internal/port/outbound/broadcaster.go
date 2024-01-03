package outbound

import "rpc-server/internal/core/domain/enum"

type BroadcasterRepository interface {
	Publish(enum.Color) error
	Subscribe() (<-chan enum.Color, error)
}
