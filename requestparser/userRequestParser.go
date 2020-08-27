package requestparser

import (
	"encoding/json"
	"net/http"

	"github.com/Diarselimi/go-todo/models"
)

func ParseRequest(r *http.Request) (models.User, error) {
	decoded := json.NewDecoder(r.Body)
	var u models.User

	err := decoded.Decode(&u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
