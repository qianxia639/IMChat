package validator

type ValidateContext struct {
	v Validator
}

func NewValidateContext(v Validator) Validator {
	return &ValidateContext{v: v}
}

func (v *ValidateContext) Validate(parma interface{}) error {
	return v.v.Validate(parma)
}
