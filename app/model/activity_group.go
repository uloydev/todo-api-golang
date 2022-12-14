package model

type ActivityGroupRequest struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type ActivityGroupResponse struct {
	BasicData
	Email string `json:"email"`
	Title string `json:"title"`
}
