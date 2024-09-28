package utils

type ValidationStep struct {
	Stmt bool
	Err  string
}

func Validate(vs []ValidationStep) (errors []string) {
	for _, validationStep := range vs {
		if !validationStep.Stmt {
			errors = append(errors, validationStep.Err)
		}
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}