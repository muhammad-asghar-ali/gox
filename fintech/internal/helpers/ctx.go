package helpers

import (
	"fmt"
	"net/http"
)

func GetUserIdFromCtx(r *http.Request) (string, error) {
	id := r.Context().Value("user_id")
	if id == nil {
		return "", fmt.Errorf("user ID not found")
	}

	userID := fmt.Sprintf("%v", id)
	return userID, nil
}
