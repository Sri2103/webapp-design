package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Inproduction  bool
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
