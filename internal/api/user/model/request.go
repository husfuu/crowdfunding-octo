package model

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailAvailableRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AvatarRequest struct {
	Image string `form:"image" binding:"required"`
}

type UserReqisterRequest struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
}

type UserUpdateRequest struct {
	UserID     int64  `json:"user_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Occupation string `json:"occupation" binding:"required"`
}
