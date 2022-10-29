package routes

import (
	"net/http"

	"github.com/danilo-lopes/simple-bank/src/controllers"
)

var accountsRoute = []Route{
	{
		URI:                    "/accounts",
		Method:                 http.MethodPost,
		Function:               controllers.CreateAccount,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/accounts/{userID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetAccount,
		AuthenticationRequired: false,
	},
}
