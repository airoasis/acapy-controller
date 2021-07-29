package model

type Wallet struct {
	WalletName	string `json:"wallet_name" validate:"required,min=4""`
	Label 		string `json:"label" validate:"required,min=4""`
	ImageUrl	string `json:"image_url"`
}
