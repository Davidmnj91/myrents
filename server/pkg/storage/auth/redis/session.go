package redis

type Session struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`
}
