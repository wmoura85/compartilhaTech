package services

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"compartilhatech/internal/domain/entities"
	"compartilhatech/internal/infra/database/sqlc/queries"
	"context"
	"database/sql"
)

type RegraService struct {
	db               *sql.DB
	RegraRepository []entities.MkRegrasVencimento
}

func NewRegraService(db *sql.DB) service_interface.RegraService {
	return &RegraService {
		db:               db,
		RegraRepository: []entities.MkRegrasVencimento{},
	}
}

func (s *RegraService) InsertRegra(data dto.CreateRegra) (*entities.MkRegrasVencimento, error) {
	/*r := entities.NewRegra()
	r.TipoRegra = data.TipoRegra
	r.Descricao = data.Descricao
	r.DiaVcto = data.DiaVcto
	r.DiaInicio = data.DiaInicio
	r.TipoPeriodo = data.TipoPeriodo

	if data.TipoRegra != 0 {
		r.TipoRegra = *&data.TipoRegra
	}*/

	//s.PersonRepository = append(s.PersonRepository, *p)

	dbConn := queries.New(s.db)

	err := dbConn.CreateRegra(context.Background(), queries.CreateRegraParams{
		TipoRegra:              sql.NullInt32{Int32: int32(data.TipoRegra), Valid: true},
		Descricao:              sql.NullString{String: string(data.Descricao), Valid: true},
		TipoPeriodo:            sql.NullInt32{Int32: int32(data.TipoPeriodo), Valid: true},
		DiaVcto:                sql.NullInt32{Int32: int32(data.DiaVcto), Valid: true},
		DiaInicio:              sql.NullInt32{Int32: int32(data.DiaInicio), Valid: true},
		DiaReferenciaTelefonia: sql.NullInt32{Int32: int32(data.DiaReferenciaTelefonia), Valid: true},
		PreGerarProxMes:        sql.NullString{String: string(data.PreGerarProxMes), Valid: true},
		Inativo:                sql.NullString{String: string(data.Inativo), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return nil, err

}

func (s *RegraService) ListRegra() ([]entities.MkRegrasVencimento, error) {
	dbConn := queries.New(s.db)

	r, err := dbConn.GetRegra(context.Background())
	if err != nil {
		return nil, err
	}

	regras := []entities.MkRegrasVencimento{}
	for _, v := range r {
		regras = append(regras, entities.MkRegrasVencimento{
			Codvenc: int(v.Codvenc),
			Descricao: v.Descricao.String,
			TipoRegra: int(v.TipoRegra.Int32),
			TipoPeriodo: int(v.TipoPeriodo.Int32),
			DiaInicio: int(v.DiaInicio.Int32),
			DiaVcto: int(v.DiaVcto.Int32),
			DiaReferenciaTelefonia: int(v.DiaReferenciaTelefonia.Int32),
			PreGerarProxMes: v.PreGerarProxMes.String,
			Inativo: v.Inativo.String,
		})
		

	}
		
	return regras, nil
}


func (s *RegraService) GetByIdRegra(Codvenc int32) (*entities.MkRegrasVencimento, error) {

	Regra := new(entities.MkRegrasVencimento)

	for _, p := range s.RegraRepository {
		if p.Codvenc == int(Codvenc) {
			Regra = &p
			return Regra, nil
		}
	}

	return nil, nil

	
}

func (s *RegraService) UpdateRegra(Codvenc int32, data dto.UpdateRegra) (*entities.MkRegrasVencimento, error) {
	dbConn := queries.New(s.db)

	p, err := dbConn.UpdateRegra(context.Background(), queries.UpdateRegraParams{Codvenc:                Codvenc, TipoRegra:              sql.NullInt32{Int32: int32(data.TipoRegra), Valid: true}, Descricao:              sql.NullString{String: string(data.Descricao),Valid: true}, TipoPeriodo:            sql.NullInt32{Int32: int32(data.TipoPeriodo), Valid: true}, DiaVcto:                sql.NullInt32{Int32: int32(data.DiaVcto), Valid: true}, DiaInicio:              sql.NullInt32{Int32: int32(data.DiaInicio), Valid: true}, DiaReferenciaTelefonia: sql.NullInt32{Int32: int32(data.DiaReferenciaTelefonia,), Valid: true}, PreGerarProxMes:        sql.NullString{String: string(data.PreGerarProxMes), Valid: true}, Inativo:                sql.NullString{String: string(data.Inativo), Valid: true}})
	if err!= nil {
		return nil, err
	}

	regras := new(entities.MkRegrasVencimento)

	regras.Codvenc = int(p.Codvenc)
	regras.Descricao = p.Descricao.String
	regras.DiaInicio = int(p.DiaInicio.Int32)
	regras.DiaReferenciaTelefonia = int(p.DiaReferenciaTelefonia.Int32)
	regras.TipoRegra = int(p.TipoRegra.Int32)
	regras.TipoPeriodo = int(p.TipoPeriodo.Int32)
	regras.PreGerarProxMes = p.PreGerarProxMes.String
	regras.DiaVcto = int(p.DiaVcto.Int32)
	regras.Inativo = p.Inativo.String
	/*for i, r := range s.RegraRepository {
		if r.Codvenc == int(Codvenc) {
			if data.Descricao != "" {
				r.Descricao = *data.Descricao
			}
			if data.TipoRegra != 0 {
				r.TipoRegra = *&data.TipoRegra
			}
			if data.TipoPeriodo != 0 {
				r.TipoPeriodo = *&data.TipoPeriodo
			}
			if data.DiaInicio != 0 {
				r.DiaInicio = *&data.DiaInicio
			}
			if data.DiaVcto != 0 {
				r.DiaVcto = *&data.DiaVcto
			}

			s.RegraRepository[i] = r

			return &r, nil
		}*/
	return regras, nil


}

func (s *RegraService) DeleteRegra(Codvenc int32) error {
	dbConn := queries.New(s.db)
	err := dbConn.DeleteRegra(context.Background(), Codvenc)

	return err

}
