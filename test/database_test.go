package test

import (
	"testing"

	"github.com/sorfian/go-todo-list/app"
	"github.com/stretchr/testify/assert"
)

var db = app.Connect()

func TestConnect(t *testing.T) {
	assert.NotNil(t, db)
}
