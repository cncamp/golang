package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrease(t *testing.T) {
	t.Log("Start testing")
	result := increase(1, 2)
	assert.Equal(t, result, 3)
}
