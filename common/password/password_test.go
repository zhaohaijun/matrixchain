package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountPassword(t *testing.T) {
	var password, err = GetAccountPassword()
	assert.Nil(t, password)
	assert.NotNil(t, err)
	password, err = GetPassword()
	assert.Nil(t, password)
	assert.NotNil(t, err)
	password, err = GetConfirmedPassword()
	assert.Nil(t, password)
	assert.NotNil(t, err)
}
