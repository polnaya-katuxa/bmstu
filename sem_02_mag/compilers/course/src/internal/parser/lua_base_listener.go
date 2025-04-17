// Code generated from Lua.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua
import "github.com/antlr4-go/antlr/v4"

// BaseLuaListener is a complete listener for a parse tree produced by LuaParser.
type BaseLuaListener struct{}

var _ LuaListener = &BaseLuaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLuaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLuaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLuaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLuaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseLuaListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseLuaListener) ExitChunk(ctx *ChunkContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLuaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLuaListener) ExitBlock(ctx *BlockContext) {}

// EnterStatEmptySemicolon is called when production StatEmptySemicolon is entered.
func (s *BaseLuaListener) EnterStatEmptySemicolon(ctx *StatEmptySemicolonContext) {}

// ExitStatEmptySemicolon is called when production StatEmptySemicolon is exited.
func (s *BaseLuaListener) ExitStatEmptySemicolon(ctx *StatEmptySemicolonContext) {}

// EnterStatAssignment is called when production StatAssignment is entered.
func (s *BaseLuaListener) EnterStatAssignment(ctx *StatAssignmentContext) {}

// ExitStatAssignment is called when production StatAssignment is exited.
func (s *BaseLuaListener) ExitStatAssignment(ctx *StatAssignmentContext) {}

// EnterStatFunctionCall is called when production StatFunctionCall is entered.
func (s *BaseLuaListener) EnterStatFunctionCall(ctx *StatFunctionCallContext) {}

// ExitStatFunctionCall is called when production StatFunctionCall is exited.
func (s *BaseLuaListener) ExitStatFunctionCall(ctx *StatFunctionCallContext) {}

// EnterStatDo is called when production StatDo is entered.
func (s *BaseLuaListener) EnterStatDo(ctx *StatDoContext) {}

// ExitStatDo is called when production StatDo is exited.
func (s *BaseLuaListener) ExitStatDo(ctx *StatDoContext) {}

// EnterStatWhile is called when production StatWhile is entered.
func (s *BaseLuaListener) EnterStatWhile(ctx *StatWhileContext) {}

// ExitStatWhile is called when production StatWhile is exited.
func (s *BaseLuaListener) ExitStatWhile(ctx *StatWhileContext) {}

// EnterStatRepeat is called when production StatRepeat is entered.
func (s *BaseLuaListener) EnterStatRepeat(ctx *StatRepeatContext) {}

// ExitStatRepeat is called when production StatRepeat is exited.
func (s *BaseLuaListener) ExitStatRepeat(ctx *StatRepeatContext) {}

// EnterStatIfThenElse is called when production StatIfThenElse is entered.
func (s *BaseLuaListener) EnterStatIfThenElse(ctx *StatIfThenElseContext) {}

// ExitStatIfThenElse is called when production StatIfThenElse is exited.
func (s *BaseLuaListener) ExitStatIfThenElse(ctx *StatIfThenElseContext) {}

// EnterStatNumericFor is called when production StatNumericFor is entered.
func (s *BaseLuaListener) EnterStatNumericFor(ctx *StatNumericForContext) {}

// ExitStatNumericFor is called when production StatNumericFor is exited.
func (s *BaseLuaListener) ExitStatNumericFor(ctx *StatNumericForContext) {}

// EnterStatGenericFor is called when production StatGenericFor is entered.
func (s *BaseLuaListener) EnterStatGenericFor(ctx *StatGenericForContext) {}

// ExitStatGenericFor is called when production StatGenericFor is exited.
func (s *BaseLuaListener) ExitStatGenericFor(ctx *StatGenericForContext) {}

// EnterStatFunction is called when production StatFunction is entered.
func (s *BaseLuaListener) EnterStatFunction(ctx *StatFunctionContext) {}

// ExitStatFunction is called when production StatFunction is exited.
func (s *BaseLuaListener) ExitStatFunction(ctx *StatFunctionContext) {}

// EnterStatLocalFunction is called when production StatLocalFunction is entered.
func (s *BaseLuaListener) EnterStatLocalFunction(ctx *StatLocalFunctionContext) {}

// ExitStatLocalFunction is called when production StatLocalFunction is exited.
func (s *BaseLuaListener) ExitStatLocalFunction(ctx *StatLocalFunctionContext) {}

// EnterStatLocalAttributeNameList is called when production StatLocalAttributeNameList is entered.
func (s *BaseLuaListener) EnterStatLocalAttributeNameList(ctx *StatLocalAttributeNameListContext) {}

// ExitStatLocalAttributeNameList is called when production StatLocalAttributeNameList is exited.
func (s *BaseLuaListener) ExitStatLocalAttributeNameList(ctx *StatLocalAttributeNameListContext) {}

// EnterAttnamelist is called when production attnamelist is entered.
func (s *BaseLuaListener) EnterAttnamelist(ctx *AttnamelistContext) {}

// ExitAttnamelist is called when production attnamelist is exited.
func (s *BaseLuaListener) ExitAttnamelist(ctx *AttnamelistContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseLuaListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseLuaListener) ExitAttrib(ctx *AttribContext) {}

// EnterStatReturn is called when production StatReturn is entered.
func (s *BaseLuaListener) EnterStatReturn(ctx *StatReturnContext) {}

// ExitStatReturn is called when production StatReturn is exited.
func (s *BaseLuaListener) ExitStatReturn(ctx *StatReturnContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseLuaListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseLuaListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterVarlist is called when production varlist is entered.
func (s *BaseLuaListener) EnterVarlist(ctx *VarlistContext) {}

// ExitVarlist is called when production varlist is exited.
func (s *BaseLuaListener) ExitVarlist(ctx *VarlistContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseLuaListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseLuaListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExplist is called when production explist is entered.
func (s *BaseLuaListener) EnterExplist(ctx *ExplistContext) {}

// ExitExplist is called when production explist is exited.
func (s *BaseLuaListener) ExitExplist(ctx *ExplistContext) {}

// EnterExpFalse is called when production ExpFalse is entered.
func (s *BaseLuaListener) EnterExpFalse(ctx *ExpFalseContext) {}

// ExitExpFalse is called when production ExpFalse is exited.
func (s *BaseLuaListener) ExitExpFalse(ctx *ExpFalseContext) {}

// EnterExpVararg is called when production ExpVararg is entered.
func (s *BaseLuaListener) EnterExpVararg(ctx *ExpVarargContext) {}

// ExitExpVararg is called when production ExpVararg is exited.
func (s *BaseLuaListener) ExitExpVararg(ctx *ExpVarargContext) {}

// EnterExpTableConstructor is called when production ExpTableConstructor is entered.
func (s *BaseLuaListener) EnterExpTableConstructor(ctx *ExpTableConstructorContext) {}

// ExitExpTableConstructor is called when production ExpTableConstructor is exited.
func (s *BaseLuaListener) ExitExpTableConstructor(ctx *ExpTableConstructorContext) {}

// EnterExpPrefixExp is called when production ExpPrefixExp is entered.
func (s *BaseLuaListener) EnterExpPrefixExp(ctx *ExpPrefixExpContext) {}

// ExitExpPrefixExp is called when production ExpPrefixExp is exited.
func (s *BaseLuaListener) ExitExpPrefixExp(ctx *ExpPrefixExpContext) {}

// EnterExpTrue is called when production ExpTrue is entered.
func (s *BaseLuaListener) EnterExpTrue(ctx *ExpTrueContext) {}

// ExitExpTrue is called when production ExpTrue is exited.
func (s *BaseLuaListener) ExitExpTrue(ctx *ExpTrueContext) {}

// EnterExpNumber is called when production ExpNumber is entered.
func (s *BaseLuaListener) EnterExpNumber(ctx *ExpNumberContext) {}

// ExitExpNumber is called when production ExpNumber is exited.
func (s *BaseLuaListener) ExitExpNumber(ctx *ExpNumberContext) {}

// EnterExpOperatorUnary is called when production ExpOperatorUnary is entered.
func (s *BaseLuaListener) EnterExpOperatorUnary(ctx *ExpOperatorUnaryContext) {}

// ExitExpOperatorUnary is called when production ExpOperatorUnary is exited.
func (s *BaseLuaListener) ExitExpOperatorUnary(ctx *ExpOperatorUnaryContext) {}

// EnterExpOperatorAnd is called when production ExpOperatorAnd is entered.
func (s *BaseLuaListener) EnterExpOperatorAnd(ctx *ExpOperatorAndContext) {}

// ExitExpOperatorAnd is called when production ExpOperatorAnd is exited.
func (s *BaseLuaListener) ExitExpOperatorAnd(ctx *ExpOperatorAndContext) {}

// EnterExpOperatorPower is called when production ExpOperatorPower is entered.
func (s *BaseLuaListener) EnterExpOperatorPower(ctx *ExpOperatorPowerContext) {}

// ExitExpOperatorPower is called when production ExpOperatorPower is exited.
func (s *BaseLuaListener) ExitExpOperatorPower(ctx *ExpOperatorPowerContext) {}

// EnterExpOperatorAddSub is called when production ExpOperatorAddSub is entered.
func (s *BaseLuaListener) EnterExpOperatorAddSub(ctx *ExpOperatorAddSubContext) {}

// ExitExpOperatorAddSub is called when production ExpOperatorAddSub is exited.
func (s *BaseLuaListener) ExitExpOperatorAddSub(ctx *ExpOperatorAddSubContext) {}

// EnterExpOperatorStrcat is called when production ExpOperatorStrcat is entered.
func (s *BaseLuaListener) EnterExpOperatorStrcat(ctx *ExpOperatorStrcatContext) {}

// ExitExpOperatorStrcat is called when production ExpOperatorStrcat is exited.
func (s *BaseLuaListener) ExitExpOperatorStrcat(ctx *ExpOperatorStrcatContext) {}

// EnterExpOperatorComparison is called when production ExpOperatorComparison is entered.
func (s *BaseLuaListener) EnterExpOperatorComparison(ctx *ExpOperatorComparisonContext) {}

// ExitExpOperatorComparison is called when production ExpOperatorComparison is exited.
func (s *BaseLuaListener) ExitExpOperatorComparison(ctx *ExpOperatorComparisonContext) {}

// EnterExpNil is called when production ExpNil is entered.
func (s *BaseLuaListener) EnterExpNil(ctx *ExpNilContext) {}

// ExitExpNil is called when production ExpNil is exited.
func (s *BaseLuaListener) ExitExpNil(ctx *ExpNilContext) {}

// EnterExpOperatorOr is called when production ExpOperatorOr is entered.
func (s *BaseLuaListener) EnterExpOperatorOr(ctx *ExpOperatorOrContext) {}

// ExitExpOperatorOr is called when production ExpOperatorOr is exited.
func (s *BaseLuaListener) ExitExpOperatorOr(ctx *ExpOperatorOrContext) {}

// EnterExpString is called when production ExpString is entered.
func (s *BaseLuaListener) EnterExpString(ctx *ExpStringContext) {}

// ExitExpString is called when production ExpString is exited.
func (s *BaseLuaListener) ExitExpString(ctx *ExpStringContext) {}

// EnterExpOperatorMulDivMod is called when production ExpOperatorMulDivMod is entered.
func (s *BaseLuaListener) EnterExpOperatorMulDivMod(ctx *ExpOperatorMulDivModContext) {}

// ExitExpOperatorMulDivMod is called when production ExpOperatorMulDivMod is exited.
func (s *BaseLuaListener) ExitExpOperatorMulDivMod(ctx *ExpOperatorMulDivModContext) {}

// EnterExpFunctionDef is called when production ExpFunctionDef is entered.
func (s *BaseLuaListener) EnterExpFunctionDef(ctx *ExpFunctionDefContext) {}

// ExitExpFunctionDef is called when production ExpFunctionDef is exited.
func (s *BaseLuaListener) ExitExpFunctionDef(ctx *ExpFunctionDefContext) {}

// EnterPrefixexp is called when production prefixexp is entered.
func (s *BaseLuaListener) EnterPrefixexp(ctx *PrefixexpContext) {}

// ExitPrefixexp is called when production prefixexp is exited.
func (s *BaseLuaListener) ExitPrefixexp(ctx *PrefixexpContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseLuaListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseLuaListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterVarOrExp is called when production varOrExp is entered.
func (s *BaseLuaListener) EnterVarOrExp(ctx *VarOrExpContext) {}

// ExitVarOrExp is called when production varOrExp is exited.
func (s *BaseLuaListener) ExitVarOrExp(ctx *VarOrExpContext) {}

// EnterVar is called when production var is entered.
func (s *BaseLuaListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BaseLuaListener) ExitVar(ctx *VarContext) {}

// EnterVarSuffix is called when production varSuffix is entered.
func (s *BaseLuaListener) EnterVarSuffix(ctx *VarSuffixContext) {}

// ExitVarSuffix is called when production varSuffix is exited.
func (s *BaseLuaListener) ExitVarSuffix(ctx *VarSuffixContext) {}

// EnterNameAndArgs is called when production nameAndArgs is entered.
func (s *BaseLuaListener) EnterNameAndArgs(ctx *NameAndArgsContext) {}

// ExitNameAndArgs is called when production nameAndArgs is exited.
func (s *BaseLuaListener) ExitNameAndArgs(ctx *NameAndArgsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseLuaListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseLuaListener) ExitArgs(ctx *ArgsContext) {}

// EnterFunctiondef is called when production functiondef is entered.
func (s *BaseLuaListener) EnterFunctiondef(ctx *FunctiondefContext) {}

// ExitFunctiondef is called when production functiondef is exited.
func (s *BaseLuaListener) ExitFunctiondef(ctx *FunctiondefContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseLuaListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseLuaListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterParlist is called when production parlist is entered.
func (s *BaseLuaListener) EnterParlist(ctx *ParlistContext) {}

// ExitParlist is called when production parlist is exited.
func (s *BaseLuaListener) ExitParlist(ctx *ParlistContext) {}

// EnterTableconstructor is called when production tableconstructor is entered.
func (s *BaseLuaListener) EnterTableconstructor(ctx *TableconstructorContext) {}

// ExitTableconstructor is called when production tableconstructor is exited.
func (s *BaseLuaListener) ExitTableconstructor(ctx *TableconstructorContext) {}

// EnterFieldlist is called when production fieldlist is entered.
func (s *BaseLuaListener) EnterFieldlist(ctx *FieldlistContext) {}

// ExitFieldlist is called when production fieldlist is exited.
func (s *BaseLuaListener) ExitFieldlist(ctx *FieldlistContext) {}

// EnterField is called when production field is entered.
func (s *BaseLuaListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLuaListener) ExitField(ctx *FieldContext) {}

// EnterFieldsep is called when production fieldsep is entered.
func (s *BaseLuaListener) EnterFieldsep(ctx *FieldsepContext) {}

// ExitFieldsep is called when production fieldsep is exited.
func (s *BaseLuaListener) ExitFieldsep(ctx *FieldsepContext) {}

// EnterOperatorOr is called when production operatorOr is entered.
func (s *BaseLuaListener) EnterOperatorOr(ctx *OperatorOrContext) {}

// ExitOperatorOr is called when production operatorOr is exited.
func (s *BaseLuaListener) ExitOperatorOr(ctx *OperatorOrContext) {}

// EnterOperatorAnd is called when production operatorAnd is entered.
func (s *BaseLuaListener) EnterOperatorAnd(ctx *OperatorAndContext) {}

// ExitOperatorAnd is called when production operatorAnd is exited.
func (s *BaseLuaListener) ExitOperatorAnd(ctx *OperatorAndContext) {}

// EnterOperatorComparison is called when production operatorComparison is entered.
func (s *BaseLuaListener) EnterOperatorComparison(ctx *OperatorComparisonContext) {}

// ExitOperatorComparison is called when production operatorComparison is exited.
func (s *BaseLuaListener) ExitOperatorComparison(ctx *OperatorComparisonContext) {}

// EnterOperatorStrcat is called when production operatorStrcat is entered.
func (s *BaseLuaListener) EnterOperatorStrcat(ctx *OperatorStrcatContext) {}

// ExitOperatorStrcat is called when production operatorStrcat is exited.
func (s *BaseLuaListener) ExitOperatorStrcat(ctx *OperatorStrcatContext) {}

// EnterOperatorAddSub is called when production operatorAddSub is entered.
func (s *BaseLuaListener) EnterOperatorAddSub(ctx *OperatorAddSubContext) {}

// ExitOperatorAddSub is called when production operatorAddSub is exited.
func (s *BaseLuaListener) ExitOperatorAddSub(ctx *OperatorAddSubContext) {}

// EnterOperatorMulDivMod is called when production operatorMulDivMod is entered.
func (s *BaseLuaListener) EnterOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// ExitOperatorMulDivMod is called when production operatorMulDivMod is exited.
func (s *BaseLuaListener) ExitOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// EnterOperatorUnary is called when production operatorUnary is entered.
func (s *BaseLuaListener) EnterOperatorUnary(ctx *OperatorUnaryContext) {}

// ExitOperatorUnary is called when production operatorUnary is exited.
func (s *BaseLuaListener) ExitOperatorUnary(ctx *OperatorUnaryContext) {}

// EnterOperatorPower is called when production operatorPower is entered.
func (s *BaseLuaListener) EnterOperatorPower(ctx *OperatorPowerContext) {}

// ExitOperatorPower is called when production operatorPower is exited.
func (s *BaseLuaListener) ExitOperatorPower(ctx *OperatorPowerContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLuaListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLuaListener) ExitNumber(ctx *NumberContext) {}

// EnterString is called when production string is entered.
func (s *BaseLuaListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseLuaListener) ExitString(ctx *StringContext) {}
