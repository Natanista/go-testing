package services_test

import (
	"testing"

	"github.com/natanista/go-testing/services"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T){
	if(services.Sum(2,2) != 4){
		t.Error("2 + 2 must be 4")
	}
}

func TestMul(t *testing.T){
	assert.Equal(t, services.Multiply(2,2), 4)
}

func TestMulSum(t *testing.T){
	require.Equal(t, services.Multiply(2,2), 4)
}