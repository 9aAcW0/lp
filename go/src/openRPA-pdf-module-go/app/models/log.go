package models

import "time"

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Data struct {
	Task Log `json:"task"`
}

type LogResponse struct {
	Data      Data   `json:"data"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type Log struct {
	Model
	TaskLineID string      `json:"task_line_id" validator:"existTaskLine"`
	EndPoint   string      `json:"end_point" validator:"url"`
	Status     string      `json:"status" validate:"eq=started|eq=processing|eq=succeeded|eq=failed"`
	Input      interface{} `json:"input"`
	Roots      interface{} `json:"roots"`
	Exports    interface{} `json:"exports"`
}
