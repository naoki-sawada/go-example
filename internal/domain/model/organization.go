package model

type Organization struct {
	ID   string `db:"id" json:"id" validate:"required,uuid4"`
	Name string `db:"name" json:"name" validate:"required,max=100"`
}
