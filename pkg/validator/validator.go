package validator

type DTOValidator interface {
	Validate(interface{}) (errors []string)
}
