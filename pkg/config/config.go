// Package config used to save whole app configuration
package config

import "text/template"

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
