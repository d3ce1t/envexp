package exporter

import (
	"os"
	"strings"
)

// Environment as key:value map
type Environment map[string]string

// NewEnvironment creates a new environment
func NewEnvironment() Environment {
	env := make(Environment)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		env.Set(pair[0], pair[1])
	}
	return env
}

// Set a new environment variable
func (env Environment) Set(key string, value string) {
	env[key] = value
}

// Unset an existing environment variable
func (env Environment) Unset(key string) {
	delete(env, key)
}

// Clone creates a copy of a map
func (env Environment) Clone() Environment {
	copy := make(Environment)
	for k, v := range env {
		copy.Set(k, v)
	}
	return copy
}

// EnvExporter exports an Environment
type EnvExporter interface {
	Export(env Environment) error
}

// NewExporter creates a EnvExporter object of the given style
func NewExporter(style string, output *os.File, pretty bool) EnvExporter {
	if style == "web" {
		return NewWebTransformer(output, pretty)
	} else if style == "json" {
		return NewJSONTransformer(output, pretty)
	}
	return nil
}