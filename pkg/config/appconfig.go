package config

import "html/template"

type Appconfig struct {
	EnableCache bool
	Templates   map[string]*template.Template
}
