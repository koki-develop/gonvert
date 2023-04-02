package yaml

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAML_FromJSON(t *testing.T) {
	tests := []struct {
		json    string
		want    string
		wantErr bool
	}{
		{
			json: `{
  "string": "hello",
  "int": 10,
  "float": 10.11,
  "boolean": true,
  "array": [
    "hello",
    10,
    true
  ],
  "struct": {
    "string": "hello",
    "int": 10,
    "float": 10.11,
    "boolean": true,
    "array": [
      "hello",
      10,
      true
    ]
  }
}`,
			want: `string: hello
int: 10
float: 10.11
boolean: true
array:
  - hello
  - 10
  - true
struct:
  string: hello
  int: 10
  float: 10.11
  boolean: true
  array:
    - hello
    - 10
    - true
`,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			conv := &YAML{indent: 2}

			got, err := conv.FromJSON(strings.NewReader(tt.json))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, string(got))
			}
		})
	}
}
