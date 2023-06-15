package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
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

// returns replaced block
func (v *Visitor) replaceBlock(block *ir.Block) *ir.Block {
	curBlock := v.block()
	if brTerm, ok := curBlock.Term.(*ir.TermBr); ok {
		block.NewBr(brTerm.Succs()[0])
	}
	v.blocks[len(v.blocks)-1] = block
	return curBlock
}

func endWithBr(block, nextBlock *ir.Block) {
	if block.Term == nil {
		block.NewBr(nextBlock)
	}
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

func (v *Visitor) dereference(expr value.Value) value.Value {
	if !types.IsPointer(expr.Type()) {
		return expr
	}

	if ptr, ok := expr.(*ir.InstAlloca); ok {
		return v.block().NewLoad(ptr.ElemType, ptr)
	}
	if ptr, ok := expr.(*ir.InstGetElementPtr); ok {
		arrType := baseArrType(ptr)

		arrElem := v.block().NewLoad(arrType.ElemType, ptr)
		return v.dereference(arrElem)
	}
	return expr
}

func baseArrType(ptr *ir.InstGetElementPtr) *types.ArrayType {
	arrType := ptr.ElemType.(*types.ArrayType)

	for {
		if elem, ok := arrType.ElemType.(*types.ArrayType); ok {
			arrType = elem
		} else {
			break
		}
	}
	return arrType
}

func (v *Visitor) sameT(expr, nextExpr value.Value) (value.Value, value.Value) {
	// TODO: check size compatibility
	return v.dereference(expr), v.dereference(nextExpr)
}

func (v *Visitor) castCond(cond value.Value) value.Value {
	cond = v.dereference(cond)

	if cond.Type().Equal(types.I1) {
		return cond
	}

	if cond.Type().Equal(types.I32) {
		return v.block().NewICmp(enum.IPredNE, cond, constant.NewInt(types.I32, 0))
	}

	panic(BadCondition(cond))
}
