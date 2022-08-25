package apiErrors

const (
	UserProfileIDRequired = "USER_PROFILE_ID_REQUIRED"
	UserProfileNotFound   = "USER_PROFILE_NOT_FOUND"
)

var (
	userProfileErrors = []apiError{
		{
			Id:      UserProfileIDRequired,
			Message: "user profile id required",
			Status:  400,
		},
		{
			Id:      UserProfileNotFound,
			Message: "user profile not found",
			Status:  404,
		},
	}
)
