package usecase

import (
	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/inbound"
	"rpc-server/internal/port/outbound"
)

type painterUseCase struct {
	broadcasterRepo outbound.BroadcasterRepository
	colorRepo       outbound.ColorRepository
}

func NewPainterUseCase(broadcasterRepo outbound.BroadcasterRepository, colorRepo outbound.ColorRepository) inbound.PainterUseCase {
	return &painterUseCase{
		broadcasterRepo: broadcasterRepo,
		colorRepo:       colorRepo,
	}
}

func (p *painterUseCase) ChangeColor(color enum.Color) error {
	p.colorRepo.Update(color)
	return p.broadcasterRepo.Publish(color)
}

func (p *painterUseCase) GetColor() (enum.Color, error) {
	return p.colorRepo.Get()
}

func (p *painterUseCase) GetColorStream() (<-chan enum.Color, error) {
	return p.broadcasterRepo.Subscribe()
}
