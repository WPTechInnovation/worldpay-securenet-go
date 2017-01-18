package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-securenet-lib-go/sdk/client"
	"github.com/wptechinnovation/worldpay-securenet-lib-go/sdk/service/cardnotpresent"
	"github.com/wptechinnovation/worldpay-securenet-lib-go/sdk/service/tokenization"
	"github.com/wptechinnovation/worldpay-securenet-lib-go/sdk/types"
)

// Flags
var flagSecureNetID string
var flagSecureKey string
var flagHTTPProxy string
var flagPublicKey string

// App Vars
var snSDK client.Client

func init() {

	flag.StringVar(&flagSecureNetID, "securenetid", "", "SecureNet ID")
	flag.StringVar(&flagSecureKey, "securekey", "", "Secure Key")
	flag.StringVar(&flagHTTPProxy, "proxy", "", "HTTP Proxy e.g. http://hostname:port")
	flag.StringVar(&flagPublicKey, "publickey", "", "SecureNet public key")
}

func main() {

	flag.Parse()

	initLog()

	apiEndpoint := "https://gwapi.demo.securenet.com/api"
	appVersion := "0.1"
	secureNetID := flagSecureNetID
	secureKey := flagSecureKey

	_snSDK, err := client.New(apiEndpoint, appVersion, secureNetID, secureKey, flagHTTPProxy)
	snSDK = _snSDK
	errorCheck(err, "client.New()", true)

	// Create a card with test data
	card := types.Card{}
	card.Number = "5314501686810737"
	card.ExpirationDate = "05/19"
	card.Address = &types.Address{}
	card.Address.Company = "Yacero"
	card.Address.Line1 = "01438 Fieldstone Way"
	card.Address.City = "Lancaster"
	card.Address.State = "Pennsylvania"
	card.Address.Zip = "17605"
	card.Address.Country = "United States"
	card.Address.Phone = "717.615.2865"

	// Create developer application object
	devApp := &types.DeveloperApplication{}
	devApp.DeveloperID = 12345678
	devApp.Version = "0.1"

	// Attempt to tokenize the card
	reqTokenize := tokenization.TokenizeCardRequest{}
	reqTokenize.Card = &card
	reqTokenize.AddToVault = false
	reqTokenize.PublicKey = flagPublicKey
	reqTokenize.DeveloperApplication = devApp

	respTokenize, err := snSDK.TokenizationService().TokenizeCard(&reqTokenize)
	errorCheck(err, "TokenizeCard()", true)

	fmt.Printf("TokenResponse: %+v\n", respTokenize)

	reqChargeToken := cardnotpresent.ChargeTokenRequest{}
	reqChargeToken.Amount = 125.00
	reqChargeToken.DeveloperApplication = devApp
	reqChargeToken.PaymentVaultToken = &types.PaymentVaultToken{}
	reqChargeToken.PaymentVaultToken.CustomerID = "22313"
	reqChargeToken.PaymentVaultToken.PaymentMethodID = respTokenize.Token
	reqChargeToken.PaymentVaultToken.PaymentType = types.CreditCard
	reqChargeToken.PaymentVaultToken.PublicKey = flagPublicKey

	respCharge, err := snSDK.CardNotPresentService().ChargeUsingToken(&reqChargeToken)

	errorCheck(err, "chargeUsingToken()", true)

	fmt.Printf("ChargeResponse Object: %+v\n", respCharge)
}

func errorCheck(err error, hint string, exitOnErr bool) {

	if err != nil {

		fmt.Printf("Error: %s - %s", hint, err.Error())

		if exitOnErr {

			os.Exit(1)
		}
	}
}

func initLog() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
