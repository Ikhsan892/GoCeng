package goceng_test

import (
	"github.com/ikhsan892/goceng"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBIsNotNIl(t *testing.T) {
	g := goceng.New()
	g.Start()
	assert.NotNil(t, g)
}

func TestSettingsIsNotNull(t *testing.T) {
	g := goceng.New()

	assert.NotEmpty(t, g.Settings().DB.URL)
}
