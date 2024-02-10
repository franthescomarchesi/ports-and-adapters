package cli_test

import (
	"fmt"
	"testing"

	"github.com/franthescomarchesi/ports_and_adapters/adapters/cli"
	mock_application "github.com/franthescomarchesi/ports_and_adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	res, err := cli.Run(serviceMock, "create", "", productMock.GetName(), productMock.GetPrice())
	require.Nil(t, err)
	require.Equal(t, res, resExpected)

	resExpected = fmt.Sprintf("Product %s has been enabled.", productMock.GetName())
	res, err = cli.Run(serviceMock, "enable", productMock.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, res, resExpected)

	resExpected = fmt.Sprintf("Product %s has been disabled.", productMock.GetName())
	res, err = cli.Run(serviceMock, "disable", productMock.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, res, resExpected)

	resExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	res, err = cli.Run(serviceMock, "get", productMock.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, res, resExpected)
}
