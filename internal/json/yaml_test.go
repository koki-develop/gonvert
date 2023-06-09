package json

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromYAML(t *testing.T) {
	tests := []struct {
		yaml    string
		want    string
		wantErr bool
	}{
		{
			yaml: `string: hello
int: 10
float: 10.11
boolean: true
null: null
array:
  - "hello"
  - 10
  - true
  - null
struct:
  string: hello
  int: 10
  float: 10.11
  boolean: true
  null: null
  array:
    - "hello"
    - 10
    - true
    - null
`,
			want: `{
  "string": "hello",
  "int": 10,
  "float": 10.11,
  "boolean": true,
  "null": null,
  "array": [
    "hello",
    10,
    true,
    null
  ],
  "struct": {
    "string": "hello",
    "int": 10,
    "float": 10.11,
    "boolean": true,
    "null": null,
    "array": [
      "hello",
      10,
      true,
      null
    ]
  }
}`,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			conv := New(&JSONConfig{Indent: 2, Minify: false})
			got, err := conv.FromYAML(strings.NewReader(tt.yaml))

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, string(got))
				assert.NoError(t, err)
			}
		})
	}
}
