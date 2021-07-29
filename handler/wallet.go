package handler

import (
	"example.com/controller/model"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateWallet(c echo.Context) error {
	adminUrl := "http://aca-py:8021"
	client := resty.New()

	wallet := new(model.Wallet)
	if err := c.Bind(wallet); err != nil {
		return err
	}

	resp, err := client.R().
		SetBody(map[string]interface{}{
			"wallet_name": wallet.WalletName,
			"wallet_key": uuid.New(),
			"wallet_type": "indy",
			"label": wallet.Label,
			"image_url": wallet.ImageUrl,
		}).
		Post(adminUrl + "/multitenancy/wallet")

	if err != nil {
		log.Info("ERROR sending the request", err)
		return err
	}

	return c.JSONBlob(resp.StatusCode(), resp.Body())
}