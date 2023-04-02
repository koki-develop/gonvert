package yaml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type YAML struct {
	indent int
}

type YAMLConfig struct {
	Indent int
}

func New(cfg *YAMLConfig) *YAML {
	return &YAML{
		indent: cfg.Indent,
	}
}

func (conv *YAML) FromJSON(r io.Reader) ([]byte, error) {
	var root yaml.Node

	dec := json.NewDecoder(r)
	if err := conv.parseJSON(dec, &root); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	enc := yaml.NewEncoder(buf)
	enc.SetIndent(conv.indent)
	if err := enc.Encode(&root); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (conv *YAML) parseJSON(dec *json.Decoder, node *yaml.Node) error {
	token, err := dec.Token()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	switch token := token.(type) {
	case json.Delim:
		switch token {
		case '{':
			node.Kind = yaml.MappingNode
			for dec.More() {
				var keyNode yaml.Node
				var valueNode yaml.Node
				if err := conv.parseJSON(dec, &keyNode); err != nil {
					return err
				}
				if err := conv.parseJSON(dec, &valueNode); err != nil {
					return err
				}
				node.Content = append(node.Content, &keyNode, &valueNode)
			}
			if _, err = dec.Token(); err != nil {
				return err
			}
		case '[':
			node.Kind = yaml.SequenceNode
			for dec.More() {
				var childNode yaml.Node
				if err := conv.parseJSON(dec, &childNode); err != nil {
					return err
				}
				node.Content = append(node.Content, &childNode)
			}
			if _, err := dec.Token(); err != nil {
				return err
			}
		}
	default:
		node.Kind = yaml.ScalarNode
		node.Tag = ""
		node.Value = fmt.Sprintf("%v", token)
	}

	return nil
}
