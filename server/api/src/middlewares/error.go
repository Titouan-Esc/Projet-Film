package middlewares

import (
	"fmt"
	"runtime"
)

type TechnicalError struct {
	Code    int
	Libelle string
}

var (
	ErrMovieExist = fmt.Errorf("movie already exist")
)

func ServiceFonctionalError(libelle string, code int) *TechnicalError {
	return &TechnicalError{
		Code:    code,
		Libelle: libelle,
	}
}

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
