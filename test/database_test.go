package test

import (
	"testing"

	"github.com/sorfian/go-contact-management-api/app"
	"github.com/stretchr/testify/assert"
)

var db = app.Connect()

func TestConnect(t *testing.T) {
	assert.NotNil(t, db)
}
