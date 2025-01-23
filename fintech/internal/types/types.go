package types

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ErrResponse struct {
		Message string `json:"message"`
	}

	RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AccountResponse struct {
		Type    string `json:"type"`
		Name    string `json:"name"`
		Balance uint   `json:"balance"`
	}

	UserResponse struct {
		ID       uint               `json:"id"`
		Username string             `json:"username"`
		Email    string             `json:"email"`
		Accounts []*AccountResponse `json:"accounts"`
	}

	LoginResponse struct {
		Message string        `json:"message"`
		Token   *string       `json:"token"`
		Data    *UserResponse `json:"data"`
	}

	RegisterResponse struct {
		Message string        `json:"message"`
		Data    *UserResponse `json:"data"`
	}
)
