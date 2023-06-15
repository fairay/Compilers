package ast

import (
	"fmt"

	"github.com/llir/llvm/ir/value"
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

func VariableSizeError(sizeExpr string) error {
	return fmt.Errorf("can not use expression %s as size, it should be constant value", sizeExpr)
}

func ExpectedPtrError(v value.Value) error {
	return fmt.Errorf("expected pointer, got %v", v)
}

func BadCondition(v value.Value) error {
	return fmt.Errorf("cannot use %v as a condition", v)
}

func NotSubscriptable(v value.Value) error {
	return fmt.Errorf("%v not subscriptable", v)
}
