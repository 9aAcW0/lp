package models

type Root struct {
	EndPoint  string              `json:"end_point"`
	Arguments map[string]Argument `json:"arguments"`
}
