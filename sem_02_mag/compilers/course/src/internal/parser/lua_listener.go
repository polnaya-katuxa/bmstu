// Code generated from Lua.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua
import "github.com/antlr4-go/antlr/v4"

// LuaListener is a complete listener for a parse tree produced by LuaParser.
type LuaListener interface {
	antlr.ParseTreeListener

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatEmptySemicolon is called when entering the StatEmptySemicolon production.
	EnterStatEmptySemicolon(c *StatEmptySemicolonContext)

	// EnterStatAssignment is called when entering the StatAssignment production.
	EnterStatAssignment(c *StatAssignmentContext)

	// EnterStatFunctionCall is called when entering the StatFunctionCall production.
	EnterStatFunctionCall(c *StatFunctionCallContext)

	// EnterStatDo is called when entering the StatDo production.
	EnterStatDo(c *StatDoContext)

	// EnterStatWhile is called when entering the StatWhile production.
	EnterStatWhile(c *StatWhileContext)

	// EnterStatRepeat is called when entering the StatRepeat production.
	EnterStatRepeat(c *StatRepeatContext)

	// EnterStatIfThenElse is called when entering the StatIfThenElse production.
	EnterStatIfThenElse(c *StatIfThenElseContext)

	// EnterStatNumericFor is called when entering the StatNumericFor production.
	EnterStatNumericFor(c *StatNumericForContext)

	// EnterStatGenericFor is called when entering the StatGenericFor production.
	EnterStatGenericFor(c *StatGenericForContext)

	// EnterStatFunction is called when entering the StatFunction production.
	EnterStatFunction(c *StatFunctionContext)

	// EnterStatLocalFunction is called when entering the StatLocalFunction production.
	EnterStatLocalFunction(c *StatLocalFunctionContext)

	// EnterStatLocalAttributeNameList is called when entering the StatLocalAttributeNameList production.
	EnterStatLocalAttributeNameList(c *StatLocalAttributeNameListContext)

	// EnterAttnamelist is called when entering the attnamelist production.
	EnterAttnamelist(c *AttnamelistContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterStatReturn is called when entering the StatReturn production.
	EnterStatReturn(c *StatReturnContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterVarlist is called when entering the varlist production.
	EnterVarlist(c *VarlistContext)

	// EnterNamelist is called when entering the namelist production.
	EnterNamelist(c *NamelistContext)

	// EnterExplist is called when entering the explist production.
	EnterExplist(c *ExplistContext)

	// EnterExpFalse is called when entering the ExpFalse production.
	EnterExpFalse(c *ExpFalseContext)

	// EnterExpVararg is called when entering the ExpVararg production.
	EnterExpVararg(c *ExpVarargContext)

	// EnterExpTableConstructor is called when entering the ExpTableConstructor production.
	EnterExpTableConstructor(c *ExpTableConstructorContext)

	// EnterExpPrefixExp is called when entering the ExpPrefixExp production.
	EnterExpPrefixExp(c *ExpPrefixExpContext)

	// EnterExpTrue is called when entering the ExpTrue production.
	EnterExpTrue(c *ExpTrueContext)

	// EnterExpNumber is called when entering the ExpNumber production.
	EnterExpNumber(c *ExpNumberContext)

	// EnterExpOperatorUnary is called when entering the ExpOperatorUnary production.
	EnterExpOperatorUnary(c *ExpOperatorUnaryContext)

	// EnterExpOperatorAnd is called when entering the ExpOperatorAnd production.
	EnterExpOperatorAnd(c *ExpOperatorAndContext)

	// EnterExpOperatorPower is called when entering the ExpOperatorPower production.
	EnterExpOperatorPower(c *ExpOperatorPowerContext)

	// EnterExpOperatorAddSub is called when entering the ExpOperatorAddSub production.
	EnterExpOperatorAddSub(c *ExpOperatorAddSubContext)

	// EnterExpOperatorStrcat is called when entering the ExpOperatorStrcat production.
	EnterExpOperatorStrcat(c *ExpOperatorStrcatContext)

	// EnterExpOperatorComparison is called when entering the ExpOperatorComparison production.
	EnterExpOperatorComparison(c *ExpOperatorComparisonContext)

	// EnterExpNil is called when entering the ExpNil production.
	EnterExpNil(c *ExpNilContext)

	// EnterExpOperatorOr is called when entering the ExpOperatorOr production.
	EnterExpOperatorOr(c *ExpOperatorOrContext)

	// EnterExpString is called when entering the ExpString production.
	EnterExpString(c *ExpStringContext)

	// EnterExpOperatorMulDivMod is called when entering the ExpOperatorMulDivMod production.
	EnterExpOperatorMulDivMod(c *ExpOperatorMulDivModContext)

	// EnterExpFunctionDef is called when entering the ExpFunctionDef production.
	EnterExpFunctionDef(c *ExpFunctionDefContext)

	// EnterPrefixexp is called when entering the prefixexp production.
	EnterPrefixexp(c *PrefixexpContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterVarOrExp is called when entering the varOrExp production.
	EnterVarOrExp(c *VarOrExpContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterVarSuffix is called when entering the varSuffix production.
	EnterVarSuffix(c *VarSuffixContext)

	// EnterNameAndArgs is called when entering the nameAndArgs production.
	EnterNameAndArgs(c *NameAndArgsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterFunctiondef is called when entering the functiondef production.
	EnterFunctiondef(c *FunctiondefContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterParlist is called when entering the parlist production.
	EnterParlist(c *ParlistContext)

	// EnterTableconstructor is called when entering the tableconstructor production.
	EnterTableconstructor(c *TableconstructorContext)

	// EnterFieldlist is called when entering the fieldlist production.
	EnterFieldlist(c *FieldlistContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldsep is called when entering the fieldsep production.
	EnterFieldsep(c *FieldsepContext)

	// EnterOperatorOr is called when entering the operatorOr production.
	EnterOperatorOr(c *OperatorOrContext)

	// EnterOperatorAnd is called when entering the operatorAnd production.
	EnterOperatorAnd(c *OperatorAndContext)

	// EnterOperatorComparison is called when entering the operatorComparison production.
	EnterOperatorComparison(c *OperatorComparisonContext)

	// EnterOperatorStrcat is called when entering the operatorStrcat production.
	EnterOperatorStrcat(c *OperatorStrcatContext)

	// EnterOperatorAddSub is called when entering the operatorAddSub production.
	EnterOperatorAddSub(c *OperatorAddSubContext)

	// EnterOperatorMulDivMod is called when entering the operatorMulDivMod production.
	EnterOperatorMulDivMod(c *OperatorMulDivModContext)

	// EnterOperatorUnary is called when entering the operatorUnary production.
	EnterOperatorUnary(c *OperatorUnaryContext)

	// EnterOperatorPower is called when entering the operatorPower production.
	EnterOperatorPower(c *OperatorPowerContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatEmptySemicolon is called when exiting the StatEmptySemicolon production.
	ExitStatEmptySemicolon(c *StatEmptySemicolonContext)

	// ExitStatAssignment is called when exiting the StatAssignment production.
	ExitStatAssignment(c *StatAssignmentContext)

	// ExitStatFunctionCall is called when exiting the StatFunctionCall production.
	ExitStatFunctionCall(c *StatFunctionCallContext)

	// ExitStatDo is called when exiting the StatDo production.
	ExitStatDo(c *StatDoContext)

	// ExitStatWhile is called when exiting the StatWhile production.
	ExitStatWhile(c *StatWhileContext)

	// ExitStatRepeat is called when exiting the StatRepeat production.
	ExitStatRepeat(c *StatRepeatContext)

	// ExitStatIfThenElse is called when exiting the StatIfThenElse production.
	ExitStatIfThenElse(c *StatIfThenElseContext)

	// ExitStatNumericFor is called when exiting the StatNumericFor production.
	ExitStatNumericFor(c *StatNumericForContext)

	// ExitStatGenericFor is called when exiting the StatGenericFor production.
	ExitStatGenericFor(c *StatGenericForContext)

	// ExitStatFunction is called when exiting the StatFunction production.
	ExitStatFunction(c *StatFunctionContext)

	// ExitStatLocalFunction is called when exiting the StatLocalFunction production.
	ExitStatLocalFunction(c *StatLocalFunctionContext)

	// ExitStatLocalAttributeNameList is called when exiting the StatLocalAttributeNameList production.
	ExitStatLocalAttributeNameList(c *StatLocalAttributeNameListContext)

	// ExitAttnamelist is called when exiting the attnamelist production.
	ExitAttnamelist(c *AttnamelistContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitStatReturn is called when exiting the StatReturn production.
	ExitStatReturn(c *StatReturnContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitVarlist is called when exiting the varlist production.
	ExitVarlist(c *VarlistContext)

	// ExitNamelist is called when exiting the namelist production.
	ExitNamelist(c *NamelistContext)

	// ExitExplist is called when exiting the explist production.
	ExitExplist(c *ExplistContext)

	// ExitExpFalse is called when exiting the ExpFalse production.
	ExitExpFalse(c *ExpFalseContext)

	// ExitExpVararg is called when exiting the ExpVararg production.
	ExitExpVararg(c *ExpVarargContext)

	// ExitExpTableConstructor is called when exiting the ExpTableConstructor production.
	ExitExpTableConstructor(c *ExpTableConstructorContext)

	// ExitExpPrefixExp is called when exiting the ExpPrefixExp production.
	ExitExpPrefixExp(c *ExpPrefixExpContext)

	// ExitExpTrue is called when exiting the ExpTrue production.
	ExitExpTrue(c *ExpTrueContext)

	// ExitExpNumber is called when exiting the ExpNumber production.
	ExitExpNumber(c *ExpNumberContext)

	// ExitExpOperatorUnary is called when exiting the ExpOperatorUnary production.
	ExitExpOperatorUnary(c *ExpOperatorUnaryContext)

	// ExitExpOperatorAnd is called when exiting the ExpOperatorAnd production.
	ExitExpOperatorAnd(c *ExpOperatorAndContext)

	// ExitExpOperatorPower is called when exiting the ExpOperatorPower production.
	ExitExpOperatorPower(c *ExpOperatorPowerContext)

	// ExitExpOperatorAddSub is called when exiting the ExpOperatorAddSub production.
	ExitExpOperatorAddSub(c *ExpOperatorAddSubContext)

	// ExitExpOperatorStrcat is called when exiting the ExpOperatorStrcat production.
	ExitExpOperatorStrcat(c *ExpOperatorStrcatContext)

	// ExitExpOperatorComparison is called when exiting the ExpOperatorComparison production.
	ExitExpOperatorComparison(c *ExpOperatorComparisonContext)

	// ExitExpNil is called when exiting the ExpNil production.
	ExitExpNil(c *ExpNilContext)

	// ExitExpOperatorOr is called when exiting the ExpOperatorOr production.
	ExitExpOperatorOr(c *ExpOperatorOrContext)

	// ExitExpString is called when exiting the ExpString production.
	ExitExpString(c *ExpStringContext)

	// ExitExpOperatorMulDivMod is called when exiting the ExpOperatorMulDivMod production.
	ExitExpOperatorMulDivMod(c *ExpOperatorMulDivModContext)

	// ExitExpFunctionDef is called when exiting the ExpFunctionDef production.
	ExitExpFunctionDef(c *ExpFunctionDefContext)

	// ExitPrefixexp is called when exiting the prefixexp production.
	ExitPrefixexp(c *PrefixexpContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitVarOrExp is called when exiting the varOrExp production.
	ExitVarOrExp(c *VarOrExpContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitVarSuffix is called when exiting the varSuffix production.
	ExitVarSuffix(c *VarSuffixContext)

	// ExitNameAndArgs is called when exiting the nameAndArgs production.
	ExitNameAndArgs(c *NameAndArgsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitFunctiondef is called when exiting the functiondef production.
	ExitFunctiondef(c *FunctiondefContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitParlist is called when exiting the parlist production.
	ExitParlist(c *ParlistContext)

	// ExitTableconstructor is called when exiting the tableconstructor production.
	ExitTableconstructor(c *TableconstructorContext)

	// ExitFieldlist is called when exiting the fieldlist production.
	ExitFieldlist(c *FieldlistContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldsep is called when exiting the fieldsep production.
	ExitFieldsep(c *FieldsepContext)

	// ExitOperatorOr is called when exiting the operatorOr production.
	ExitOperatorOr(c *OperatorOrContext)

	// ExitOperatorAnd is called when exiting the operatorAnd production.
	ExitOperatorAnd(c *OperatorAndContext)

	// ExitOperatorComparison is called when exiting the operatorComparison production.
	ExitOperatorComparison(c *OperatorComparisonContext)

	// ExitOperatorStrcat is called when exiting the operatorStrcat production.
	ExitOperatorStrcat(c *OperatorStrcatContext)

	// ExitOperatorAddSub is called when exiting the operatorAddSub production.
	ExitOperatorAddSub(c *OperatorAddSubContext)

	// ExitOperatorMulDivMod is called when exiting the operatorMulDivMod production.
	ExitOperatorMulDivMod(c *OperatorMulDivModContext)

	// ExitOperatorUnary is called when exiting the operatorUnary production.
	ExitOperatorUnary(c *OperatorUnaryContext)

	// ExitOperatorPower is called when exiting the operatorPower production.
	ExitOperatorPower(c *OperatorPowerContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}
