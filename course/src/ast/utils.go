package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *Visitor) block() *ir.Block {
	return v.blocks[len(v.blocks)-1]
}

func (v *Visitor) pushBlock(block *ir.Block) {
	if len(v.blocks) == 0 {
		v.mainBlock = block
	} else {
		endWithBr(block, v.block())
	}
	v.blocks = append(v.blocks, block)
}

func (v *Visitor) popBlock(block *ir.Block) {
	v.blocks = v.blocks[:len(v.blocks)-1]
	if len(v.blocks) == 0 {
		v.mainBlock = nil
	}
}

func (v *Visitor) replaceBlock(block *ir.Block) {
	if len(v.blocks) >= 2 {
		endWithBr(block, v.blocks[len(v.blocks)-2])
	}
	v.blocks[len(v.blocks)-1] = block
}

func (v *Visitor) pushScope() {
	v.varScopes = append(v.varScopes, VarScope{})
}
func (v *Visitor) popScope() {
	v.varScopes = v.varScopes[:len(v.varScopes)-1]
}

func (v *Visitor) newVariable(t types.Type, name string) *ir.InstAlloca {
	defindedVar := v.variable(name)
	if defindedVar != nil {
		panic(AlreadyDefinedError(name))
	}

	ptr := v.mainBlock.NewAlloca(t)

	// assing new variable to current variable scope
	v.varScopes[len(v.varScopes)-1][name] = ptr

	return ptr
}

func (v *Visitor) variable(name string) *ir.InstAlloca {
	for i := len(v.varScopes) - 1; i >= 0; i = i - 1 {
		scope := v.varScopes[i]
		if variable, ok := scope[name]; ok {
			return variable
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
		return v.block().NewLoad(ptr.ElemType, ptr)
	}
	if ptr, ok := expr.(*ir.InstGetElementPtr); ok {
		arrType := ptr.ElemType.(*types.ArrayType)
		return v.block().NewLoad(arrType.ElemType, ptr)
	}
	return expr
}

func (v *Visitor) sameT(expr, nextExpr value.Value) (value.Value, value.Value) {
	// TODO: check size compatibility
	return v.dereference(expr), v.dereference(nextExpr)
}

func endWithBr(block, nextBlock *ir.Block) {
	if block.Term == nil {
		block.NewBr(nextBlock)
	}
}
