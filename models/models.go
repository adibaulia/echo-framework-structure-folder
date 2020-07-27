package models

import "echo-framework-structure/config"

type (
	modelsIntance struct {
		oke string
	}

	someStruct struct {
	}
)

func NewModels(c *config.Connection) *modelsIntance {
	return &modelsIntance{c.C}
}

func (d *modelsIntance) GetHelloWorldString() string {
	return "hello world!"
}
