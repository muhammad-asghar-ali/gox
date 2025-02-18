package common

type (
	ResponseMeta struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	APIResponse[T any] struct {
		ResponseMeta
		Data *T `json:"data,omitempty"`
	}
)

func NewSuccessResponse[T any](data T, message string) APIResponse[T] {
	return APIResponse[T]{
		ResponseMeta: ResponseMeta{
			Success: true,
			Message: message,
		},
		Data: &data,
	}
}

func NewErrorResponse(message string) ResponseMeta {
	return ResponseMeta{
		Success: false,
		Message: message,
	}
}
