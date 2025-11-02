package string_compression

import (
	"strings"
	"testing"
)

func TestLocalizer_Compress(t *testing.T) {
	l := NewLocalizer()
	tests := []struct {
		long     string
		expected string
	}{
		{
			"internationalization",
			"i18n",
		},
		{
			"kubernetes",
			"k8s",
		},
	}

	for _, tt := range tests {
		res := l.Compress(tt.long)
		if strings.Compare(res, tt.expected) != 0 {
			t.Errorf("expected %s, got %s\n", tt.expected, res)
		}
	}
}

func TestLocalizer_Decompress(t *testing.T) {
	l := NewLocalizer()
	tests := []struct {
		long  string
		short string
	}{
		{
			"internationalisation",
			"i18n",
		},
		{
			"kubernetes",
			"k8s",
		},
	}

	for _, tt := range tests {
		l.Compress(tt.long)
	}

	for _, tt := range tests {
		res := l.Decompress(tt.short)
		if strings.Compare(res, tt.long) != 0 {
			t.Errorf("expected %s, got %s\n", tt.short, res)
		}
	}
}
