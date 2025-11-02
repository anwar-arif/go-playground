package string_compression

import (
	"strconv"
	"sync"
)

type Localizer struct {
	cache map[string]string
	mu    sync.RWMutex
}

func NewLocalizer() *Localizer {
	return &Localizer{
		cache: make(map[string]string),
	}
}

func (l *Localizer) Compress(word string) string {
	n := len(word)
	if n <= 2 {
		return word
	}

	compressed := word
	compressed = string(word[0]) + strconv.Itoa(n-2) + string(word[n-1])

	l.mu.Lock()
	l.cache[compressed] = word
	l.mu.Unlock()

	return compressed
}

func (l *Localizer) Decompress(short string) string {
	if len(short) <= 2 {
		return short
	}
	return l.cache[short]
}
