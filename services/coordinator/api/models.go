package api

import (
	"database/sql"
)

type Models struct {
	Nodes NodeModel
	Jobs  JobModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Nodes: NodeModel{DB: db},
		Jobs:  JobModel{DB: db},
	}
}
