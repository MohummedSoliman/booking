// Package config used to save whole app configuration
package config

import (
	"log"
	"text/template"

	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
