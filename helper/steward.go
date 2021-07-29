package helper

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

var Token string
var stewardSeed = "000000000000000000000000Steward1"
var adminUrl = "http://aca-py:8021"
var baseWalletName = "base.agent"

func InitAcaPy() {
	stewardWallet := "steward"

	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"wallet_name": stewardWallet,
		}).
		Get(adminUrl + "/multitenancy/wallets")

	if err != nil {
		log.Error().Err(err).Msg("ERROR sending the request")
		return
	}

	log.Info().Msgf("Body:\n %v", resp)

	var wallets string = resp.String()

	if gjson.Get(wallets, "results.#").Int() == 0 {
		resp, err = client.R().
			SetBody(map[string]interface{}{
				"wallet_name": stewardWallet,
				"wallet_key": stewardWallet + ".key",
				"wallet_type": "indy",
			}).
			Post(adminUrl + "/multitenancy/wallet")

		if err != nil {
			log.Error().Err(err).Msg("ERROR sending the request")
			return
		}

		log.Info().Msgf("Body:\n %v", resp)

		Token = gjson.Get(resp.String(), "token").String()

		log.Info().Msgf("token: %v", Token)

		resp, err = client.R().
			SetBody(map[string]interface{}{
				"seed": stewardSeed,
			}).
			SetAuthToken(Token).
			Post(adminUrl + "/wallet/did/create")

		if err != nil {
			log.Error().Err(err).Msg("ERROR sending the request")
			return
		}

		log.Info().Msgf("Body:\n %v", resp)

		did := gjson.Get(resp.String(), "result.did").String()

		resp, err = client.R().
			SetQueryParams(map[string]string{
				"did": did,
			}).
			SetAuthToken(Token).
			Get(adminUrl + "/wallet/did/public")

		if err != nil {
			log.Error().Err(err).Msg("ERROR sending the request")
			return
		}

		log.Info().Msgf("Body:\n %v", resp)
	} else {
		stewardWalletId := gjson.Get(wallets, "results.0.wallet_id").String()

		resp, err := client.R().
			Post(adminUrl + "/multitenancy/wallet/" + stewardWalletId + "/token")

		if err != nil {
			log.Error().Err(err).Msg("ERROR sending the request")
			return
		}

		log.Info().Msgf("Body:\n %v", resp)

		Token = gjson.Get(resp.String(), "token").String()
	}


}
