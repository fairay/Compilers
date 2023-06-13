package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func getVariableByName(block *ir.Block, name string) *ir.InstAlloca {
	for _, inst := range block.Insts {
		if inst, ok := inst.(*ir.InstAlloca); ok {
			varName := inst.LocalName
			if varName == name {
				return inst
			}
		}
	}
	return nil
}

func toPtr(expr value.Value) *ir.InstAlloca {
	if types.IsPointer(expr.Type()) {
		return expr.(*ir.InstAlloca)
	} else {
		panic(ExpectedPtrError(expr))
	}
}

func (v *Visitor) dereference(expr value.Value) value.Value {
	if !types.IsPointer(expr.Type()) {
		return expr
	}

	if ptr, ok := expr.(*ir.InstAlloca); ok {
		return v.curBlock.NewLoad(ptr.ElemType, ptr)
	}
	if ptr, ok := expr.(*ir.InstGetElementPtr); ok {
		arrType := ptr.ElemType.(*types.ArrayType)
		return v.curBlock.NewLoad(arrType.ElemType, ptr)
	}
	return expr
}

func (v *Visitor) sameT(expr, nextExpr value.Value) (value.Value, value.Value) {
	// TODO: check size compatibility
	return v.dereference(expr), v.dereference(nextExpr)
}
