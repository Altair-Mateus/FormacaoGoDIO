package models

type Cliente struct {
	Id       int             `json:"id"`
	Nome     string          `json:"nome"`
	Idade    int             `json:"idade"`
	Endereco EnderecoCliente `json:"endereco"`
}

type EnderecoCliente struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
}
