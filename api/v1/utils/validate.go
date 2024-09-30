package utils

type ValidationStep struct {
	Stmt  bool
	Error string
}

func Validate(validationSteps []ValidationStep) []string {
	var Errors []string
	for _, el := range validationSteps {
		if !el.Stmt {
			Errors = append(Errors, el.Error)
		}
	}
	
	if len(Errors) == 0 {
		return nil
	} else {
		return Errors
	}
}
