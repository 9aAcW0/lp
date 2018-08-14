package models

type Pointer struct {
	InputOrExports string
	EndPoint       string
	Keys           []Key
}

type Key struct {
	Name    string
	IsArray bool
}
