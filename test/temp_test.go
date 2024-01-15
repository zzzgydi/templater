package tests

import (
	"testing"

	"github.com/zzzgydi/templater"
)

func TestTemplaterFunc(t *testing.T) {
	tests := []struct {
		name     string
		template string
		args     map[string]interface{}
		expected string
	}{
		{
			name:     "Basic Variable Replacement",
			template: "Hello, ${name}! Your score is ${score}.",
			args:     map[string]interface{}{"name": "Alice", "score": 95},
			expected: "Hello, Alice! Your score is 95.",
		},
		{
			name:     "Formatted Variable Replacement",
			template: "Price: ${price|%.2f}",
			args:     map[string]interface{}{"price": 123.456},
			expected: "Price: 123.46",
		},
		{
			name:     "Escaped Percent Sign",
			template: "Discount: 10%",
			args:     map[string]interface{}{},
			expected: "Discount: 10%",
		},
		{
			name:     "Multiple Variables",
			template: "${greeting}, ${name}! You are ${age} years old.",
			args:     map[string]interface{}{"greeting": "Hello", "name": "Bob", "age": 30},
			expected: "Hello, Bob! You are 30 years old.",
		},
		{
			name:     "Nonexistent Variable",
			template: "Welcome, ${name}!",
			args:     map[string]interface{}{"username": "Charlie"},
			expected: "Welcome, <nil>!",
		},
		{
			name:     "Empty Template and Args",
			template: "",
			args:     map[string]interface{}{},
			expected: "",
		},
		{
			name:     "Invalid Formatting String",
			template: "Your score: ${score|%x}",
			args:     map[string]interface{}{"score": "eighty"},
			expected: "Your score: 656967687479",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			templater := templater.NewTemplater(test.template)
			output := templater.Parse(test.args)

			if output != test.expected {
				t.Errorf("Test '%s' failed: expected '%s', got '%s'", test.name, test.expected, output)
			}
		})
	}
}
