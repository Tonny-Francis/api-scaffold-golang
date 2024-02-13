package exampleSchema

type ExampleSchema struct {
	Name  string `json:"name" validate:"required"`
	Tenat string `json:"tenat" validate:"required"`
}
