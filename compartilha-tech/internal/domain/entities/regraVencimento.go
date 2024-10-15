package entities

type MkRegrasVencimento struct {
	Codvenc                int
	Descricao              string
	TipoRegra              int
	TipoPeriodo            int
	DiaInicio              int
	DiaVcto                int
	PreGerarProxMes        string
	Inativo                string
	DiaReferenciaTelefonia int
}

/*func NewRegra() * MkRegrasVencimento {
	return &MkRegrasVencimento{
		Codvenc:                0,
		Descricao:              "",
		Tipo:                   0,
		TipoRegra:              0,
		TipoPeriodo:            0,
		DiaInicio:              0,
		DiaVcto:                0,
		PreGerarProxMes:        "N",
		Inativo:                "N",
		DiaReferenciaTelefonia: 0,
	}
}*/