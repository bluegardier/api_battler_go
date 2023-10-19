package models

type People struct {
	ID int64 `json:"id"`
	Apelido configPeople `json:"apelido"`
	Nome configPeople `json:"nome"`
	Nascimento string `json:"nascimento"`
	Stack []configPeople `json:"stack"`
}

type configPeople struct {
	Value string
	MinLen int16
	MaxLen int16
	IsSet bool
	IsUnique bool

}