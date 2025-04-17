// Code generated from Lua.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua
import "github.com/antlr4-go/antlr/v4"

type BaseLuaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLuaVisitor) VisitChunk(ctx *ChunkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatEmptySemicolon(ctx *StatEmptySemicolonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatAssignment(ctx *StatAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatFunctionCall(ctx *StatFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatDo(ctx *StatDoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatWhile(ctx *StatWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatRepeat(ctx *StatRepeatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatIfThenElse(ctx *StatIfThenElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatNumericFor(ctx *StatNumericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatGenericFor(ctx *StatGenericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatFunction(ctx *StatFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatLocalFunction(ctx *StatLocalFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatLocalAttributeNameList(ctx *StatLocalAttributeNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitAttnamelist(ctx *AttnamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitAttrib(ctx *AttribContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStatReturn(ctx *StatReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFuncname(ctx *FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNamelist(ctx *NamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExplist(ctx *ExplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpFalse(ctx *ExpFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpVararg(ctx *ExpVarargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpTableConstructor(ctx *ExpTableConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpPrefixExp(ctx *ExpPrefixExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpTrue(ctx *ExpTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpNumber(ctx *ExpNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorUnary(ctx *ExpOperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorAnd(ctx *ExpOperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorPower(ctx *ExpOperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorAddSub(ctx *ExpOperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorStrcat(ctx *ExpOperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorComparison(ctx *ExpOperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpNil(ctx *ExpNilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorOr(ctx *ExpOperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpString(ctx *ExpStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpOperatorMulDivMod(ctx *ExpOperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpFunctionDef(ctx *ExpFunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitPrefixexp(ctx *PrefixexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFunctioncall(ctx *FunctioncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarOrExp(ctx *VarOrExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarSuffix(ctx *VarSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNameAndArgs(ctx *NameAndArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFunctiondef(ctx *FunctiondefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFuncbody(ctx *FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitParlist(ctx *ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTableconstructor(ctx *TableconstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFieldlist(ctx *FieldlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFieldsep(ctx *FieldsepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorOr(ctx *OperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorAnd(ctx *OperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorComparison(ctx *OperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorStrcat(ctx *OperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorAddSub(ctx *OperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorMulDivMod(ctx *OperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorUnary(ctx *OperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorPower(ctx *OperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}
