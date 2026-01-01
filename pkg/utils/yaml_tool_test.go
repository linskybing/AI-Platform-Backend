package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitYAMLDocuments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "single document",
			input: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: config1
`,
			expected: []string{
				`apiVersion: v1
kind: ConfigMap
metadata:
  name: config1`,
			},
		},
		{
			name: "multiple documents with --- separator",
			input: `
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: config1
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: config2
`,
			expected: []string{
				`apiVersion: v1
kind: ConfigMap
metadata:
  name: config1`,
				`apiVersion: v1
kind: ConfigMap
metadata:
  name: config2`,
			},
		},
		{
			name: "trailing separator",
			input: `
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: config1
---
`,
			expected: []string{
				`apiVersion: v1
kind: ConfigMap
metadata:
  name: config1`,
			},
		},
		{
			name:     "empty input",
			input:    ``,
			expected: []string{},
		},
		{
			name:     "only separators",
			input:    `---`,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SplitYAMLDocuments(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("SplitYAMLDocuments() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestReplacePlaceholders(t *testing.T) {
	template := "apiVersion: v1\nkind: Pod\nmetadata:\n  name: {{name}}\nspec:\n  image: {{image}}"
	values := map[string]string{
		"name":  "my-pod",
		"image": "nginx:latest",
	}

	expected := "apiVersion: v1\nkind: Pod\nmetadata:\n  name: my-pod\nspec:\n  image: nginx:latest"
	result := ReplacePlaceholders(template, values)

	assert.Equal(t, expected, result)
}

func TestReplacePlaceholdersInJSON(t *testing.T) {
	jsonStr := `{
		"metadata": {
			"name": "{{name}}"
		},
		"spec": {
			"containers": [{
				"image": "{{image}}"
			}]
		}
	}`

	values := map[string]string{
		"name":  "my-app",
		"image": "nginx:1.25",
	}

	result, err := ReplacePlaceholdersInJSON(jsonStr, values)
	assert.NoError(t, err)
	assert.Contains(t, result, `"name": "my-app"`)
	assert.Contains(t, result, `"image": "nginx:1.25"`)
}

func TestYAMLToJSON(t *testing.T) {
	yamlContent := `
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: nginx
      image: nginx:1.25
`

	expectedContains := `"kind": "Pod"`
	jsonStr, err := YAMLToJSON(yamlContent)
	assert.NoError(t, err)
	assert.Contains(t, jsonStr, expectedContains)
}
