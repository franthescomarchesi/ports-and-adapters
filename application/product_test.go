package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/franthescomarchesi/ports_and_adapters/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product.", err.Error())
}

