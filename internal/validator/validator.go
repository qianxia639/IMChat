package validator

type Validator interface {
	Validate(parma interface{}) error
}
