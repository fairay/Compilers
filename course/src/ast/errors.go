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
