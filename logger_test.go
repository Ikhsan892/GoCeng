package goceng_test

import (
	"github.com/ikhsan892/goceng"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZapLogger(t *testing.T) {
	assert.NotNil(t, goceng.NewZapLogger())
}
