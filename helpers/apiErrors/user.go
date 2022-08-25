package apiErrors

const (
	UserExist            = "USER_EXIST"
	UserEmailRequired    = "USER_EMAIL_REQUIRED"
	UserEmailInvalid     = "USER_EMAIL_INVALID"
	UserEmailMin         = "USER_EMAIL_MIN"
	UserEmailMax         = "USER_EMAIL_MAX"
	UserNotLogined       = "USER_NOT_LOGINED"
	UserRoleMin          = "USER_ROLE_MIN"
	UserRoleMax          = "USER_ROLE_MAX"
	UserPasswordRequired = "USER_PASSWORD_REQUIRED"
	UserNotFound         = "USER_NOT_FOUND"
	UserIdInValid        = "USER_ID_INVALID"
	UserIdParamRequired  = "USER_ID_PARAM_REQUIRED"
	UserWrongPassword    = "USER_WRONG_PASSWORD"
	UserUnauthorized     = "USER_UNAUTHORIZED"
)

var (
	userErrors = []apiError{
		{
			Id:      UserIdParamRequired,
			Message: "userId in parameter required",
			Status:  400,
		},
		{
			Id:      UserIdInValid,
			Message: "userId must objectId",
			Status:  400,
		},
		{
			Id:      UserNotFound,
			Message: "This user not found",
			Status:  404,
		},
		{
			Id:      UserExist,
			Message: "This user has been exist!",
			Status:  400,
		},
		{
			Id:      UserEmailRequired,
			Message: "Email is required",
			Status:  400,
		},
		{
			Id:      UserEmailInvalid,
			Message: "Email invalid",
			Status:  400,
		},
		{
			Id:      UserEmailMin,
			Message: "Email min length is 3",
			Status:  400,
		},
		{
			Id:      UserEmailMax,
			Message: "Email max length is 50",
			Status:  400,
		},
		{
			Id:      UserNotLogined,
			Message: "You must login to do this",
			Status:  401,
		},
		{
			Id:      UserRoleMin,
			Message: "Role min is 0 as public user",
			Status:  400,
		},
		{
			Id:      UserRoleMax,
			Message: "Role max is 2 as admin",
			Status:  400,
		},
		{
			Id:      UserPasswordRequired,
			Message: "Password is required",
			Status:  400,
		},
		{
			Id:      UserWrongPassword,
			Message: "Password is incorrect",
			Status:  400,
		},
		{
			Id:      UserUnauthorized,
			Message: "unauthorized",
			Status:  401,
		},
	}
)
