package models

type JSON struct {
	Version    string   `json:"version" validate:"required"`
	TaskLineID string   `json:"task_line_id"`
	Input      Input    `json:"input" validate:"required"`
	Exports    []Export `json:"exports" validate:"required"`
	Roots      []Root   `json:"roots" validate:"required"`
}
