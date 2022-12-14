package model

type TodoRequest struct {
	ActivityGroupId uint   `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type TodoResponse struct {
	BasicData
	ActivityGroupId uint   `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}
