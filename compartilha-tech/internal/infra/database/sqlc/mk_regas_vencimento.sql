-- name: CreateRegra :exec
INSERT INTO mk_regras_vencimento (
    tipo_regra, descricao, tipo_periodo, dia_vcto, dia_inicio, dia_referencia_telefonia, pre_gerar_prox_mes, inativo
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetRegraById :one
SELECT * FROM mk_regras_vencimento
WHERE codvenc = $1;

-- name: GetRegra :many
SELECT * FROM mk_regras_vencimento;

-- name: UpdateRegra :one
UPDATE mk_regras_vencimento
SET tipo_regra = $2, descricao = $3, tipo_periodo = $4, dia_vcto = $5, dia_inicio = $6,
    dia_referencia_telefonia = $7, pre_gerar_prox_mes = $8, inativo = $9
WHERE codvenc = $1
RETURNING *;

-- name: DeleteRegra :exec
DELETE FROM mk_regras_vencimento
WHERE codvenc = $1;