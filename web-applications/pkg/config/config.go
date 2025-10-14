package config

import (
	"html/template"
	"log"
)

// Avoid circular imports by only importing from the std inputs, not packages
// holds thr application config.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
