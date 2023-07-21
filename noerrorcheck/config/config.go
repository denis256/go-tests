package config

func CreateEvalConfig(data string) (string, error) {
	if data == "error" {
		return "", &CfgError{"error"}
	}

	return data, nil
}

type CfgError struct {
	s string
}

func (e *CfgError) Error() string {
	return e.s
}
