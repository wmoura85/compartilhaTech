CREATE TABLE person(
    id varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    age integer NULL,
    active boolean NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS mk_regras_vencimento (
	codvenc serial NOT NULL,
	descricao varchar(150) NULL,
	tipo varchar(1) NULL,
	tipo_regra integer NULL,
	tipo_periodo integer NULL,
	dia_inicio integer NULL,
	dia_vcto integer NULL,
	pre_gerar_prox_mes varchar(1) NULL,
	inativo varchar(1) NULL,
	dia_referencia_telefonia integer NULL
      );