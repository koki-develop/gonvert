package json

type JSON struct {
	indent int
	minify bool
}

type JSONConfig struct {
	Indent int
	Minify bool
}

func New(cfg *JSONConfig) *JSON {
	return &JSON{
		indent: cfg.Indent,
		minify: cfg.Minify,
	}
}
