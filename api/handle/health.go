package handle

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	common "wallet-go/src/helper/http"
	"wallet-go/src/usecase/healthcheck"
)

func healthCheck(service healthcheck.Service) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		common.ResponseWithSuccess(writer, http.StatusOK, service.Health())
	})
}

//MakeHealthHandlers make url handlers
func MakeHealthHandlers(r *mux.Router, n negroni.Negroni, service healthcheck.Service) {
	r.Handle("/health", n.With(
		negroni.Wrap(healthCheck(service)),
	)).Methods("GET")
}
