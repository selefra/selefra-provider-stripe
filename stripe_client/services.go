package stripe_client

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-stripe/constants"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func Connect(ctx context.Context, stripeConfig *Config) (*client.API, error) {

	stripe.SetAppInfo(&stripe.AppInfo{
		Name: constants.Selefra,
		URL:  "https://github.com/selefra/selefra-provider-stripe",
	})

	apiKey := os.Getenv(constants.STRIPEAPIKEY)

	if stripeConfig.APIKey != constants.Constants_3 {
		apiKey = stripeConfig.APIKey
	}

	if apiKey == constants.Constants_4 {

		return nil, errors.New(constants.Apikeymustbeconfigured)
	}

	config := &stripe.BackendConfig{
		MaxNetworkRetries: 10,
	}

	conn := &client.API{}
	conn.Init(apiKey, &stripe.Backends{
		API:     stripe.GetBackendWithConfig(stripe.APIBackend, config),
		Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
	})

	return conn, nil
}
