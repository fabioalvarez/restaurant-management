package utils

import "os"

type EnVariables struct {
	Port string
}

func NewEnVariables() EnVariables {
	return EnVariables{}
}

func GetVariables() EnVariables {
	enVar := NewEnVariables()
	enVar.Port = os.Getenv("PORT")

	if enVar.Port == "" {
		enVar.Port = "8000"
	}

	return enVar
}
