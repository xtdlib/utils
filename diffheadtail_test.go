package utils

import "testing"

func TestDiffHeadTail(t *testing.T) {
	err := DiffFileHeadTail("a", "a")
	t.Log(err)
}

func TestDiffHeadTail2(t *testing.T) {
	err := DiffFileHeadTail("a", "b")
	t.Log(err)
}
