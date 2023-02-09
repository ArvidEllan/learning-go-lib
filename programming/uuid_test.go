package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuidWithHyphen(t *testing.T) {
	uuidWithHyphen := NewUuid(false)

	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")

}

func TestNewUuidWithHyphen(t *testing.T) {
	uuidWidthHyphen := NewUuid(true)

	assert.Len(t, uuidWidthHyphen, 32)
	assert.Contains(t, uuidWidthHyphen, "-")

}
