package test

import (
	"github.com/stirante/go-simple-eval/eval"
	"testing"
)

func TestMain(m *testing.M) {
	eval.Init()
	m.Run()
}
