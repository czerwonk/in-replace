package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		name         string
		regex        *regexp.Regexp
		content      string
		replacement  string
		group        uint8
		expected     string
		wantsReplace bool
	}{
		{
			name:         "match and replace",
			regex:        regexp.MustCompile("b*"),
			content:      "aaabbbccc",
			replacement:  "x",
			expected:     "aaaxccc",
			wantsReplace: true,
		},
		{
			name:         "no match",
			regex:        regexp.MustCompile("b*"),
			content:      "aaaccc",
			replacement:  "x",
			expected:     "aaaccc",
			wantsReplace: false,
		},
		{
			name:         "in group replacement (middle)",
			regex:        regexp.MustCompile("0(b*)0"),
			content:      "aaa0bbb0ccc",
			replacement:  "x",
			expected:     "aaa0x0ccc",
			group:        1,
			wantsReplace: true,
		},
		{
			name:         "in group replacement (start)",
			regex:        regexp.MustCompile("(b*)00"),
			content:      "aaabbb00ccc",
			replacement:  "x",
			expected:     "aaax00ccc",
			group:        1,
			wantsReplace: true,
		},
		{
			name:         "in group replacement (end)",
			regex:        regexp.MustCompile("00(b*)"),
			content:      "aaa00bbbccc",
			replacement:  "x",
			expected:     "aaa00xccc",
			group:        1,
			wantsReplace: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, replaced := replace(test.content, test.regex, &Replacement{
				Replacement: test.replacement,
				Group:       test.group,
			})

			assert.Equal(t, test.wantsReplace, replaced)
			assert.Equal(t, test.expected, res)
		})
	}
}
