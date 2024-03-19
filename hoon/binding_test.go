package hoon_test

import (
	"context"
	"testing"

	sitter "github.com/urbit-pilled/go-tree-sitter"
	"github.com/urbit-pilled/go-tree-sitter/hoon"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`=/  foo  123  foo`), hoon.GetLanguage())
	assert.NoError(err)
}
