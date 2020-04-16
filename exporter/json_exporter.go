package exporter

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSONExporter exports environment to json
type JSONExporter struct {
	output *os.File
	pretty bool
}

// NewJSONTransformer creates a JSONExporter object
func NewJSONTransformer(output *os.File, pretty bool) *JSONExporter {
	return &JSONExporter{output, pretty}
}

// Export to JSON
func (o *JSONExporter) Export(env Environment) error {
	var b []byte
	var err error

	if o.pretty {
		b, err = json.MarshalIndent(env, "", "\t")
	} else {
		b, err = json.Marshal(env)
	}

	if err != nil {
		return err
	}

	fmt.Fprintln(o.output, string(b))

	return nil
}
