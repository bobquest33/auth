package auth

import (
	"github.com/gorilla/mux"
)

func Serve(router *mux.Router) {
	if HANDLER_REGISTER == nil {
		panic("kidstuff/auth: HANDLER_REGISTER need to be overide by a mngr")
	}

	if DEFAULT_NOTIFICATOR == nil {
		panic("kidstuff/auth: DEFAULT_NOTIFICATOR need to be overide by a mngr")
	}

	if ID_FROM_STRING == nil {
		panic("kidstuff/auth: ID_FROM_STRING need to be overide by a mngr")
	}

	if ID_TO_STRING == nil {
		panic("kidstuff/auth: ID_TO_STRING need to be overide by a mngr")
	}

	router.Handle("/signup", HANDLER_REGISTER(SignUp, false, nil, nil))
	router.Handle("/tokens",
		HANDLER_REGISTER(GetToken, false, nil, nil))

	router.Handle("/users/{user_id}/activate",
		HANDLER_REGISTER(Activate, false, nil, nil))

	router.Handle("/users/{user_id}/password",
		HANDLER_REGISTER(UpdatePassword, true, []string{"admin"}, []string{"manage_user"})).Methods("PUT")

	router.Handle("/users/{user_id}",
		HANDLER_REGISTER(GetProfile, true, []string{"admin"}, []string{"manage_user"})).Methods("GET")

	router.Handle("/users/{user_id}",
		HANDLER_REGISTER(UpdateProfile, true, []string{"admin"}, []string{"manage_user"})).Methods("PATCH")

	router.Handle("/users",
		HANDLER_REGISTER(ListProfile, false, []string{"admin"}, []string{"manage_user"}))

}