package handle

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"wallet-go/api/presenter"
	common "wallet-go/src/helper/http"
	"wallet-go/src/usecase/healthcheck"
)

func healthCheck(service healthcheck.Service) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		common.ResponseWithSuccess(writer, http.StatusOK, presenter.Health{
			Message: service.Execute(),
		})
	})
}

//MakeHealthHandlers make url handlers
func MakeHealthHandlers(r *mux.Router, n negroni.Negroni, service healthcheck.Service) {
	r.Handle("/health", n.With(
		negroni.Wrap(healthCheck(service)),
	)).Methods("GET")
}
