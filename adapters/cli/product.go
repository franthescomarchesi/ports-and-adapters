package cli

import (
	"fmt"

	"github.com/franthescomarchesi/ports_and_adapters/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var res = ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return res, err
		}
		res = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return res, err
		}
		result, err := service.Enable(product)
		if err != nil {
			return res, err
		}
		res = fmt.Sprintf("Product %s has been enabled.", result.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return res, err
		}
		result, err := service.Disable(product)
		if err != nil {
			return res, err
		}
		res = fmt.Sprintf("Product %s has been disabled.", result.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return res, err
		}
		res = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return res, nil
}
