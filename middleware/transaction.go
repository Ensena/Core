package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

type TX struct {
	UserID string
	Email  string
	UUID   string
}

// func SetNewTransaction(r *http.Request) {
// 	token := Generate()
// 	tx := TX{}
// 	tx.UserID = use
// 	r.Header.Add("UserID", user.ID)
// 	r.Header.Add("Email", user.Email)
// 	r.Header.Add("AuthorizationToken", token)
// 	if r.Header.Get("UUID") != "" {
// 		uuidWithHyphen := uuid.New()
// 		r.Header.Add("UUID", uuidWithHyphen)
// 	}
// 	SetNewTransaction(tx,r)
// }

func SetTransaction(r *http.Request, tx *TX) {

	token := Generate()
	r.Header.Add("UserID", tx.UserID)
	r.Header.Add("Email", tx.Email)
	r.Header.Add("AuthorizationToken", token)
	if r.Header.Get("UUID") == "" && tx.UUID == "" {
		uuidWithHyphen := uuid.New()
		r.Header.Add("UUID", uuidWithHyphen.String())
	} else {
		r.Header.Add("UUID", tx.UUID)

	}
}

func GetTransaction(r *http.Request) *TX {

	tx := TX{}
	tx.UserID = r.Header.Get("UserID")
	tx.Email = r.Header.Get("Email")
	tx.UUID = r.Header.Get("UUID")
	return &tx
}
