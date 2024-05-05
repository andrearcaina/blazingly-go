package models

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
)

type BaseHandler struct {
	DB *sql.DB
}

type Handler interface {
	InitRoutes() chi.Router
}
