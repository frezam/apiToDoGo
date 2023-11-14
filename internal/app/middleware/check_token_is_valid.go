package middleware

import (
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/handlers"
	"github.com/GbSouza15/apiToDoGo/internal/authenticator"
)

func CheckTokenIsValid(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				handlers.SendResponse(401, []byte("Não autorizado"), w)
				return
			}
			handlers.SendResponse(400, []byte("Erro no servidor"), w)
			return
		}

		tokenString := c.Value

		isValid, err := authenticator.ValidatorToken(tokenString)

		if err != nil {
			handlers.SendResponse(400, []byte("Erro no servidor"), w)
			return
		}

		if !isValid {
			handlers.SendResponse(401, []byte("Não autorizado"), w)
			return
		}

		n(w, r)
	}
}
