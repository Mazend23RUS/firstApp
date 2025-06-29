package interfaces

type RequestValidator interface {
	Validate(interface{}) error
}
