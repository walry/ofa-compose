package tests

import (
	"ofa/services/logic"
	"testing"
)

func TestHasIdy(t *testing.T){
	idy := &logic.Industry{}
	idy.Has("02","林业")
}

