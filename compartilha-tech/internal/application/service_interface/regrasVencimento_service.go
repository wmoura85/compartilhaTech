package service_interface

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/domain/entities"
)

type RegraService interface {
	InsertRegra(data dto.CreateRegra) (*entities.MkRegrasVencimento, error)
	ListRegra() ([]entities.MkRegrasVencimento, error)
	GetByIdRegra(Codvenc int32) (*entities.MkRegrasVencimento, error)
	UpdateRegra(Codvenc int32, data dto.UpdateRegra) (*entities.MkRegrasVencimento, error)
	DeleteRegra(Codvenc int32) error
}
