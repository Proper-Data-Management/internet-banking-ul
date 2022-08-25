package apiErrors

const (
	ApiErrorNotFound   = "API_ERROR_NOT_FOUND"
	ApiErrorIdRequired = "API_ERROR_ID_REQUIRED"
)

var (
	apiErrorErrors = []apiError{
		{
			Id:      ApiErrorNotFound,
			Message: "API error not found",
			Status:  404,
		},
		{
			Id:      ApiErrorIdRequired,
			Message: "API error Id required ",
			Status:  400,
		},
	}
)
