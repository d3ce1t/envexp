package exporter

import (
	"encoding/json"
	"fmt"
	"os"
)

// WebExporter exports environment to web
type WebExporter struct {
	output *os.File
	pretty bool
}

// NewWebTransformer creates a WebExporter object
func NewWebTransformer(output *os.File, pretty bool) *WebExporter {
	return &WebExporter{output, pretty}
}

// Export to web
func (o *WebExporter) Export(env Environment) error {

	var outputTemplate string
	var b []byte
	var err error

	if o.pretty {
		outputTemplate = "(function(window) {\n%s})(window);\n"
		b, err = json.MarshalIndent(env, "", "\t")
	} else {
		outputTemplate = "(function(window){%s})(window);"
		b, err = json.Marshal(env)
	}

	if err != nil {
		return err
	}

	data := fmt.Sprintf("window._env=%s", string(b))
	fmt.Fprintf(o.output, outputTemplate, data)

	return nil
}
