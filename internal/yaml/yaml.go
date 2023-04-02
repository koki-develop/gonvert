package yaml

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
