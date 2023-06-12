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

func (v *Visitor) sameT(expr, nextExpr value.Value) (value.Value, value.Value) {
	if types.IsPointer(expr.Type()) {
		ptr := expr.(*ir.InstAlloca)
		expr = v.curBlock.NewLoad(ptr.ElemType, ptr)
	}
	if types.IsPointer(nextExpr.Type()) {
		ptr := nextExpr.(*ir.InstAlloca)
		nextExpr = v.curBlock.NewLoad(ptr.ElemType, ptr)
	}
	return expr, nextExpr
}
