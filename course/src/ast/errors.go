package ast

import (
	"fmt"
)

func UnimplementedError(ruleText string) error {
	return fmt.Errorf("syntax is not implemented: %s", ruleText)
}

func NoMainError() error {
	return fmt.Errorf("main func shoud be decleared")
}

func UndefindedError(varName string) error {
	return fmt.Errorf("variable is not definded: %s", varName)
}

func AlreadyDefinedError(varName string) error {
	return fmt.Errorf("variable is already definded: %s", varName)
}
