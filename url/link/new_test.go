package link

import (
	"testing"
)

func TestNewLinks(t *testing.T) {
	testLinks := []string{"1", "2", "3", "4", "5"}
	lastID := "1"

	result, err := NewLinks(testLinks, lastID, OrderDesc)
	t.Log(result, err)
}
