package recover

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		info, Error := keyring.NewInMemory().NewAccount(request.Name, request.Mnemonic, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		keyOutput, Error := keyring.Bech32KeyOutput(info)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		rest.PostProcessResponse(responseWriter, cliContext, newResponse(keyOutput))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/keys/recover", handler(cliContext)).Methods("POST")
}