// Code generated from Lua.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LuaParser.
type LuaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LuaParser#chunk.
	VisitChunk(ctx *ChunkContext) interface{}

	// Visit a parse tree produced by LuaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LuaParser#StatEmptySemicolon.
	VisitStatEmptySemicolon(ctx *StatEmptySemicolonContext) interface{}

	// Visit a parse tree produced by LuaParser#StatAssignment.
	VisitStatAssignment(ctx *StatAssignmentContext) interface{}

	// Visit a parse tree produced by LuaParser#StatFunctionCall.
	VisitStatFunctionCall(ctx *StatFunctionCallContext) interface{}

	// Visit a parse tree produced by LuaParser#StatDo.
	VisitStatDo(ctx *StatDoContext) interface{}

	// Visit a parse tree produced by LuaParser#StatWhile.
	VisitStatWhile(ctx *StatWhileContext) interface{}

	// Visit a parse tree produced by LuaParser#StatRepeat.
	VisitStatRepeat(ctx *StatRepeatContext) interface{}

	// Visit a parse tree produced by LuaParser#StatIfThenElse.
	VisitStatIfThenElse(ctx *StatIfThenElseContext) interface{}

	// Visit a parse tree produced by LuaParser#StatNumericFor.
	VisitStatNumericFor(ctx *StatNumericForContext) interface{}

	// Visit a parse tree produced by LuaParser#StatGenericFor.
	VisitStatGenericFor(ctx *StatGenericForContext) interface{}

	// Visit a parse tree produced by LuaParser#StatFunction.
	VisitStatFunction(ctx *StatFunctionContext) interface{}

	// Visit a parse tree produced by LuaParser#StatLocalFunction.
	VisitStatLocalFunction(ctx *StatLocalFunctionContext) interface{}

	// Visit a parse tree produced by LuaParser#StatLocalAttributeNameList.
	VisitStatLocalAttributeNameList(ctx *StatLocalAttributeNameListContext) interface{}

	// Visit a parse tree produced by LuaParser#attnamelist.
	VisitAttnamelist(ctx *AttnamelistContext) interface{}

	// Visit a parse tree produced by LuaParser#attrib.
	VisitAttrib(ctx *AttribContext) interface{}

	// Visit a parse tree produced by LuaParser#StatReturn.
	VisitStatReturn(ctx *StatReturnContext) interface{}

	// Visit a parse tree produced by LuaParser#funcname.
	VisitFuncname(ctx *FuncnameContext) interface{}

	// Visit a parse tree produced by LuaParser#varlist.
	VisitVarlist(ctx *VarlistContext) interface{}

	// Visit a parse tree produced by LuaParser#namelist.
	VisitNamelist(ctx *NamelistContext) interface{}

	// Visit a parse tree produced by LuaParser#explist.
	VisitExplist(ctx *ExplistContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpFalse.
	VisitExpFalse(ctx *ExpFalseContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpVararg.
	VisitExpVararg(ctx *ExpVarargContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpTableConstructor.
	VisitExpTableConstructor(ctx *ExpTableConstructorContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpPrefixExp.
	VisitExpPrefixExp(ctx *ExpPrefixExpContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpTrue.
	VisitExpTrue(ctx *ExpTrueContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpNumber.
	VisitExpNumber(ctx *ExpNumberContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorUnary.
	VisitExpOperatorUnary(ctx *ExpOperatorUnaryContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorAnd.
	VisitExpOperatorAnd(ctx *ExpOperatorAndContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorPower.
	VisitExpOperatorPower(ctx *ExpOperatorPowerContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorAddSub.
	VisitExpOperatorAddSub(ctx *ExpOperatorAddSubContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorStrcat.
	VisitExpOperatorStrcat(ctx *ExpOperatorStrcatContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorComparison.
	VisitExpOperatorComparison(ctx *ExpOperatorComparisonContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpNil.
	VisitExpNil(ctx *ExpNilContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorOr.
	VisitExpOperatorOr(ctx *ExpOperatorOrContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpString.
	VisitExpString(ctx *ExpStringContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpOperatorMulDivMod.
	VisitExpOperatorMulDivMod(ctx *ExpOperatorMulDivModContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpFunctionDef.
	VisitExpFunctionDef(ctx *ExpFunctionDefContext) interface{}

	// Visit a parse tree produced by LuaParser#prefixexp.
	VisitPrefixexp(ctx *PrefixexpContext) interface{}

	// Visit a parse tree produced by LuaParser#functioncall.
	VisitFunctioncall(ctx *FunctioncallContext) interface{}

	// Visit a parse tree produced by LuaParser#varOrExp.
	VisitVarOrExp(ctx *VarOrExpContext) interface{}

	// Visit a parse tree produced by LuaParser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by LuaParser#varSuffix.
	VisitVarSuffix(ctx *VarSuffixContext) interface{}

	// Visit a parse tree produced by LuaParser#nameAndArgs.
	VisitNameAndArgs(ctx *NameAndArgsContext) interface{}

	// Visit a parse tree produced by LuaParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by LuaParser#functiondef.
	VisitFunctiondef(ctx *FunctiondefContext) interface{}

	// Visit a parse tree produced by LuaParser#funcbody.
	VisitFuncbody(ctx *FuncbodyContext) interface{}

	// Visit a parse tree produced by LuaParser#parlist.
	VisitParlist(ctx *ParlistContext) interface{}

	// Visit a parse tree produced by LuaParser#tableconstructor.
	VisitTableconstructor(ctx *TableconstructorContext) interface{}

	// Visit a parse tree produced by LuaParser#fieldlist.
	VisitFieldlist(ctx *FieldlistContext) interface{}

	// Visit a parse tree produced by LuaParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by LuaParser#fieldsep.
	VisitFieldsep(ctx *FieldsepContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorOr.
	VisitOperatorOr(ctx *OperatorOrContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorAnd.
	VisitOperatorAnd(ctx *OperatorAndContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorComparison.
	VisitOperatorComparison(ctx *OperatorComparisonContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorStrcat.
	VisitOperatorStrcat(ctx *OperatorStrcatContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorAddSub.
	VisitOperatorAddSub(ctx *OperatorAddSubContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorMulDivMod.
	VisitOperatorMulDivMod(ctx *OperatorMulDivModContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorUnary.
	VisitOperatorUnary(ctx *OperatorUnaryContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorPower.
	VisitOperatorPower(ctx *OperatorPowerContext) interface{}

	// Visit a parse tree produced by LuaParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by LuaParser#string.
	VisitString(ctx *StringContext) interface{}
}
