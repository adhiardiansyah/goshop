package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var login = sessions.NewCookieStore([]byte("mysession"))

func LoginMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := login.Get(r, "mysession")
		username := session.Values["username"]
		if username == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
