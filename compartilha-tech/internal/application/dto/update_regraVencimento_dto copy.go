package dto

type UpdateRegra struct {
	Descricao   string `json:"descricao"`
	TipoRegra int  `json:"tipoRegra"`
	TipoPeriodo int `json:"tipoPeriodo"`
	DiaInicio int	`json:"diaInicio"`
	DiaVcto int `json:"diaVcto"`
	PreGerarProxMes string `json:"preGerarProxMes"`
	DiaReferenciaTelefonia int `json:"diaRefenciaTelefonia"`
	Inativo string `json:"inativo"`
}