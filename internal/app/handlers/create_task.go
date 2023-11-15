package handlers

import (
	"encoding/json"
	"errors"
	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/config"
	"github.com/GbSouza15/apiToDoGo/internal/database/queries"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
)

var (
	ErrRequestReadRequest = errors.New("erro ao ler o corpo da requisição")
	ErrJSONDecodeError    = errors.New("erro ao decodificação do JSON")
	ErrNotAuthorized      = errors.New("não autorizado")
	ErrServerError        = errors.New("erro no servidor")
)

func (h Handler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendResponse(404, []byte(ErrRequestReadRequest.Error()), w)
		return
	}

	var newTask models.TaskCreate
	var taskId = uuid.NewString()

	if err := json.Unmarshal(body, &newTask); err != nil {
		SendResponse(500, []byte(ErrJSONDecodeError.Error()), w)
		return
	}

	c, err := r.Cookie("token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			SendResponse(401, []byte(ErrNotAuthorized.Error()), w)
			return
		}
		SendResponse(400, []byte(ErrServerError.Error()), w)
		return
	}

	tknStr := c.Value
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			SendResponse(401, []byte(ErrNotAuthorized.Error()), w)
			return
		}
		SendResponse(400, []byte(ErrServerError.Error()), w)
		return
	}
	if !tkn.Valid {
		SendResponse(401, []byte(ErrNotAuthorized.Error()), w)
		return
	}

	errCreateTask := queries.CreateTask(h.DB, taskId, newTask.Title, newTask.Description, claims.UserId)
	if errCreateTask != nil {
		//fmt.Println(errCreateTask)
		//fmt.Println(newTask)
		SendResponse(500, []byte(queries.ErrCreateTasks.Error()), w)
		return
	}

	SendResponse(201, []byte("Tarefa criada com sucesso."), w)
}
