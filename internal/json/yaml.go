package json

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"gopkg.in/yaml.v3"
)

func FromYAML(r io.Reader) ([]byte, error) {
	var root yaml.Node
	if err := yaml.NewDecoder(r).Decode(&root); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := encodeJSON(buf, &root); err != nil {
		return nil, err
	}

	j := new(bytes.Buffer)
	if err := json.Indent(j, buf.Bytes(), "", "  "); err != nil {
		return nil, err
	}
	return j.Bytes(), nil
}

func encodeJSON(w *bytes.Buffer, yamlNode *yaml.Node) error {
	switch yamlNode.Kind {
	case yaml.DocumentNode:
		return encodeJSON(w, yamlNode.Content[0])

	case yaml.MappingNode:
		w.WriteString("{")
		for i := 0; i < len(yamlNode.Content); i += 2 {
			if i > 0 {
				w.WriteString(",")
			}

			key := yamlNode.Content[i]
			val := yamlNode.Content[i+1]
			keyBytes, err := json.Marshal(key.Value)
			if err != nil {
				return err
			}
			w.Write(keyBytes)
			w.WriteString(":")
			if err := encodeJSON(w, val); err != nil {
				return err
			}
		}
		w.WriteString("}")

	case yaml.SequenceNode:
		w.WriteString("[")
		for i, child := range yamlNode.Content {
			if i > 0 {
				w.WriteString(",")
			}
			if err := encodeJSON(w, child); err != nil {
				return err
			}
		}
		w.WriteString("]")

	case yaml.ScalarNode:
		var v interface{}
		var err error
		switch yamlNode.Tag {
		case "!!str":
			v = yamlNode.Value
		case "!!int":
			v, err = strconv.Atoi(yamlNode.Value)
		case "!!float":
			v, err = strconv.ParseFloat(yamlNode.Value, 64)
		case "!!bool":
			v, err = strconv.ParseBool(yamlNode.Value)
		default:
			v = yamlNode.Value
		}
		if err != nil {
			return err
		}
		valBytes, err := json.Marshal(v)
		if err != nil {
			return err
		}
		w.Write(valBytes)
	}

	return nil
}
