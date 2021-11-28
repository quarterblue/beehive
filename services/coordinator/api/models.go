package api

import (
	"database/sql"
)

type Models struct {
	Nodes NodeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Nodes: NodeModel{DB: db},
	}
}
