package usecase

import (
	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/inbound"
	"rpc-server/internal/port/outbound"
)

type PainterUseCase struct {
	broadcasterRepo outbound.BroadcasterRepository
	colorRepo       outbound.ColorRepository
}

func NewPainterUseCase(broadcasterRepo outbound.BroadcasterRepository, colorRepo outbound.ColorRepository) inbound.PainterUseCase {
	return &PainterUseCase{
		broadcasterRepo: broadcasterRepo,
		colorRepo:       colorRepo,
	}
}

func (p *PainterUseCase) ChangeColor(color enum.Color) error {
	p.colorRepo.Update(color)
	return p.broadcasterRepo.Publish(color)
}

func (p *PainterUseCase) GetColor() (enum.Color, error) {
	return p.colorRepo.Get()
}

func (p *PainterUseCase) GetColorStream() (<-chan enum.Color, error) {
	return p.broadcasterRepo.Subscribe()
}
