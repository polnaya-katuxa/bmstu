package visitor

import (
	_ "embed"
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/parser"
)

//go:embed llvm/base.ll
var baseLLVM string

//go:embed llvm/generic.ll
var genericLLVM string

//go:embed llvm/math.ll
var mathLLVM string

//go:embed llvm/logical.ll
var logicalLLVM string

//go:embed llvm/string.ll
var stringLLVM string

//go:embed llvm/compare.ll
var compareLLVM string

//go:embed llvm/table.ll
var tableLLVM string

//go:embed llvm/io.ll
var ioLLVM string

var commonLLVM = strings.Join([]string{baseLLVM, genericLLVM, mathLLVM, logicalLLVM, stringLLVM, compareLLVM, tableLLVM, ioLLVM}, "\n\n")

const mainFunc = "main"

type Visitor struct {
	*parser.BaseLuaVisitor

	declaredVarList []string
	visionScope     []int

	ErrorList []error

	IR *ir.Module

	entries   []*ir.Block
	functions []*ir.Func

	falseConst value.Value
	trueConst  value.Value
	nilConst   value.Value

	globalsCounter int

	funcs    map[string]*ir.Func
	typeDefs map[string]types.Type

	variables []map[string]value.Value

	endBlocks []*ir.Block
}

func (v *Visitor) currentEntry() *ir.Block {
	return v.entries[len(v.entries)-1]
}

func New() *Visitor {
	os.WriteFile("debug.ll", []byte(commonLLVM), 0666)

	llvm, err := asm.ParseString("", commonLLVM)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	funcs := make(map[string]*ir.Func)
	for _, f := range llvm.Funcs {
		funcs[f.Name()] = f
	}

	typeDefs := make(map[string]types.Type)
	for _, f := range llvm.TypeDefs {
		typeDefs[f.Name()] = f
	}

	return &Visitor{
		BaseLuaVisitor:  &parser.BaseLuaVisitor{},
		IR:              llvm,
		declaredVarList: make([]string, 0),
		visionScope:     make([]int, 0),
		falseConst:      constant.NewInt(types.I64, 0),
		trueConst:       constant.NewInt(types.I64, 1),
		nilConst:        constant.NewInt(types.I64, 0),
		funcs:           funcs,
		typeDefs:        typeDefs,
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitChunk(ctx *parser.ChunkContext) interface{} {
	main := v.IR.NewFunc(mainFunc, types.I64)

	v.functions = append(v.functions, main)

	v.VisitBlock(ctx.Block().(*parser.BlockContext))

	return nil
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.variables = append(v.variables, make(map[string]value.Value))
	v.visionScope = append(v.visionScope, len(v.declaredVarList))
	if len(v.variables) > 1 {
		v.variables[len(v.variables)-1] = maps.Clone(v.variables[len(v.variables)-2])
	}

	curFunc := v.functions[len(v.functions)-1]
	block := curFunc.NewBlock("")

	v.entries = append(v.entries, block)

	for _, statUntyped := range ctx.AllStat() {
		switch stat := statUntyped.(type) {
		case *parser.StatAssignmentContext:
			v.VisitStatAssignment(stat)
		case *parser.StatIfThenElseContext:
			v.VisitStatIfThenElse(stat)
		case *parser.StatFunctionCallContext:
			v.VisitStatFunctionCall(stat)
		case *parser.StatEmptySemicolonContext:
			v.VisitStatEmptySemicolon(stat)
		case *parser.StatDoContext:
			v.VisitStatDo(stat)
		case *parser.StatWhileContext:
			v.VisitStatWhile(stat)
		case *parser.StatRepeatContext:
			v.VisitStatRepeat(stat)
		case *parser.StatNumericForContext:
			v.VisitStatNumericFor(stat)
		case *parser.StatGenericForContext:
			v.VisitStatGenericFor(stat)
		case *parser.StatFunctionContext:
			v.VisitStatFunction(stat)
		case *parser.StatLocalFunctionContext:
			v.VisitStatLocalFunction(stat)
		case *parser.StatLocalAttributeNameListContext:
			v.VisitStatLocalAttributeNameList(stat)
		}
	}

	endBlock := v.entries[len(v.entries)-1]

	if ctx.Retstat() != nil {
		retCtx := ctx.Retstat().(*parser.StatReturnContext)
		explist := v.VisitExplist(retCtx.Explist().(*parser.ExplistContext)).([]expression)

		retTable := v.GetTableRet(explist)
		endBlock.NewRet(retTable)

		v.entries = v.entries[:len(v.entries)-1]

		v.declaredVarList = v.declaredVarList[:v.visionScope[len(v.visionScope)-1]]
		v.visionScope = v.visionScope[:len(v.visionScope)-1]
		v.variables = v.variables[:len(v.variables)-1]

		return block
	}

	v.declaredVarList = v.declaredVarList[:v.visionScope[len(v.visionScope)-1]]
	v.visionScope = v.visionScope[:len(v.visionScope)-1]
	v.variables = v.variables[:len(v.variables)-1]

	v.entries = v.entries[:len(v.entries)-1]

	if curFunc.Name() == mainFunc && len(v.entries) == 0 {
		endBlock.NewRet(v.nilConst)
	} else if len(v.endBlocks) > 0 {
		endBlock.NewBr(v.endBlocks[len(v.endBlocks)-1])
	} else if len(v.entries) == 0 {
		endBlock.NewRet(endBlock.NewCall(v.funcs["create_nil"]))
	}

	return block
}

func (v *Visitor) VisitStatAssignment(ctx *parser.StatAssignmentContext) interface{} {
	varlist := v.VisitVarlist(ctx.Varlist().(*parser.VarlistContext)).([]variableLeft)
	explist := v.VisitExplist(ctx.Explist().(*parser.ExplistContext)).([]expression)

	if len(explist) == 1 && len(varlist) > 1 {
		for i := range varlist {
			numConst := constant.NewInt(types.I64, int64(i))
			numPtr := constant.NewIntToPtr(
				numConst,
				types.I8Ptr,
			)
			genInt := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), numPtr)

			val := v.currentEntry().NewCall(v.funcs["lua_table_get_value_at"], explist[0].value, genInt)

			// val := v.currentEntry().NewCall(v.funcs["lua_table_get"], explist[0].value, genInt)

			if varlist[i].key != nil {
				v.currentEntry().NewCall(v.funcs["lua_table_set"], v.variables[len(v.variables)-1][varlist[i].name], varlist[i].key, val)
				continue
			}

			if v.variables[len(v.variables)-1][varlist[i].name] == nil {
				v.variables[len(v.variables)-1][varlist[i].name] = val
			} else {
				v.currentEntry().NewCall(v.funcs["copy"], val, v.variables[len(v.variables)-1][varlist[i].name])
			}
		}

		return nil
	}

	if len(varlist) != len(explist) {
		v.ErrorList = append(v.ErrorList, fmt.Errorf("len of varlist and exp list mismatch %d != %d", len(varlist), len(explist)))
		return nil
	}

	for i := range varlist {
		if varlist[i].key != nil {
			v.currentEntry().NewCall(v.funcs["lua_table_set"], v.variables[len(v.variables)-1][varlist[i].name], varlist[i].key, explist[i].value)
			continue
		}

		if v.variables[len(v.variables)-1][varlist[i].name] == nil {
			v.variables[len(v.variables)-1][varlist[i].name] = explist[i].value
		} else {
			v.currentEntry().NewCall(v.funcs["copy"], explist[i].value, v.variables[len(v.variables)-1][varlist[i].name])
		}
	}

	return nil
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	num := ctx.FLOAT()
	if num == nil {
		num = ctx.INT()
		numConst, err := constant.NewIntFromString(types.I64, num.GetText())
		if err != nil {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid int %s", num.GetText()))
		}

		numPtr := constant.NewIntToPtr(
			numConst,
			types.I8Ptr,
		)
		genInt := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), numPtr)

		return genInt
	} else {
		numConst, err := constant.NewFloatFromString(types.Double, num.GetText())
		if err != nil {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid float %s", num.GetText()))
		}

		num := v.currentEntry().NewAlloca(types.Double)
		v.currentEntry().NewStore(numConst, num)

		numPtr := v.currentEntry().NewBitCast(
			num,
			types.I8Ptr,
		)
		genFloat := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 1), numPtr)

		return genFloat
	}
}

func (v *Visitor) VisitStatEmptySemicolon(ctx *parser.StatEmptySemicolonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatFunctionCall(ctx *parser.StatFunctionCallContext) interface{} {
	name := ctx.Functioncall().VarOrExp().Var_().NAME().GetText()

	explist := v.VisitExplist(ctx.Functioncall().AllNameAndArgs()[0].Args().Explist().(*parser.ExplistContext)).([]expression)
	values := make([]value.Value, len(explist))
	for i := range explist {
		values[i] = explist[i].value
	}

	return v.currentEntry().NewCall(v.funcs[name], values...)
}

func (v *Visitor) VisitStatDo(ctx *parser.StatDoContext) interface{} {
	block := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)

	v.currentEntry().NewBr(block)

	endBlock := v.functions[len(v.functions)-1].NewBlock("")
	block.NewBr(endBlock)

	v.entries[len(v.entries)-1] = endBlock

	return endBlock
}

func (v *Visitor) VisitStatWhile(ctx *parser.StatWhileContext) interface{} {
	whileCheckBlock := v.functions[len(v.functions)-1].NewBlock("")

	v.currentEntry().NewBr(whileCheckBlock)

	v.entries = append(v.entries, whileCheckBlock)
	whileExpression := v.currentEntry().NewCall(v.funcs["check"], v.VisitExp(ctx.Exp()).(value.Value))
	v.entries = v.entries[:len(v.entries)-1]

	bodyEndBlock := v.functions[len(v.functions)-1].NewBlock("")
	v.endBlocks = append(v.endBlocks, bodyEndBlock)

	whileBodyBlock := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)
	bodyEndBlock.NewBr(whileCheckBlock)

	v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

	endBlock := v.functions[len(v.functions)-1].NewBlock("")

	whileCheckBlock.NewCondBr(whileExpression, whileBodyBlock, endBlock)

	v.entries[len(v.entries)-1] = endBlock

	return endBlock
}

func (v *Visitor) VisitStatRepeat(ctx *parser.StatRepeatContext) interface{} {
	bodyEndBlock := v.functions[len(v.functions)-1].NewBlock("")
	v.endBlocks = append(v.endBlocks, bodyEndBlock)
	repeatBodyBlock := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)
	v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

	v.currentEntry().NewBr(repeatBodyBlock)

	repeatCheckBlock := v.functions[len(v.functions)-1].NewBlock("")
	bodyEndBlock.NewBr(repeatCheckBlock)

	v.entries = append(v.entries, repeatCheckBlock)
	repeatExpression := v.currentEntry().NewCall(v.funcs["check"], v.VisitExp(ctx.Exp()).(value.Value))
	v.entries = v.entries[:len(v.entries)-1]

	endBlock := v.functions[len(v.functions)-1].NewBlock("")

	repeatCheckBlock.NewCondBr(repeatExpression, repeatBodyBlock, endBlock)

	v.entries[len(v.entries)-1] = endBlock

	return endBlock
}

func (v *Visitor) VisitStatIfThenElse(ctx *parser.StatIfThenElseContext) interface{} {
	expressions := make([]value.Value, len(ctx.AllExp()))
	for i, expCtx := range ctx.AllExp() {
		expressions[i] = v.currentEntry().NewCall(v.funcs["check"], v.VisitExp(expCtx).(value.Value))
	}

	expBlocks := make([]*ir.Block, len(ctx.AllBlock()))
	expBlocks[0] = v.currentEntry()
	for i := 1; i < len(ctx.AllBlock()); i++ {
		expBlocks[i] = v.functions[len(v.functions)-1].NewBlock("")
	}

	endBlock := v.functions[len(v.functions)-1].NewBlock("")
	v.endBlocks = append(v.endBlocks, endBlock)

	blocks := make([]*ir.Block, len(ctx.AllBlock()))
	for i, blockCtx := range ctx.AllBlock() {
		blocks[i] = v.VisitBlock(blockCtx.(*parser.BlockContext)).(*ir.Block)
	}

	v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

	for i := range expressions {
		if i == len(expBlocks)-1 {
			expBlocks[i].NewCondBr(expressions[i], blocks[i], endBlock)
		} else {
			expBlocks[i].NewCondBr(expressions[i], blocks[i], expBlocks[i+1])
		}
	}

	if len(blocks) > len(expressions) {
		expBlocks[len(expBlocks)-1].NewBr(blocks[len(expBlocks)-1])
	}

	v.entries[len(v.entries)-1] = endBlock

	return endBlock
}

func (v *Visitor) VisitStatNumericFor(ctx *parser.StatNumericForContext) interface{} {
	zeroConst := constant.NewInt(types.I64, 0)
	zeroPtr := constant.NewIntToPtr(
		zeroConst,
		types.I8Ptr,
	)
	genZero := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), zeroPtr)

	expressions := make([]value.Value, len(ctx.AllExp()))
	for i, expCtx := range ctx.AllExp() {
		viz := v.VisitExp(expCtx)
		exp, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		expressions[i] = exp
	}

	if len(expressions) == 2 {
		numConst := constant.NewInt(types.I64, 1)
		numPtr := constant.NewIntToPtr(
			numConst,
			types.I8Ptr,
		)
		genInt := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), numPtr)

		expressions = append(expressions, genInt)
	}

	name := ctx.NAME()
	v.variables[len(v.variables)-1][name.GetText()] = expressions[0]
	v.declaredVarList = append(v.declaredVarList, name.GetText())

	forCheckBlock := v.functions[len(v.functions)-1].NewBlock("")
	v.currentEntry().NewBr(forCheckBlock)

	endBlock := v.functions[len(v.functions)-1].NewBlock("")
	bodyEndBlock := v.functions[len(v.functions)-1].NewBlock("")

	v.endBlocks = append(v.endBlocks, bodyEndBlock)

	forBodyBlock := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)

	v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

	v.entries = append(v.entries, forCheckBlock)
	forCheckExpr := v.currentEntry().NewCall(v.funcs["check"], v.currentEntry().NewCall(v.funcs["lt"], expressions[2], genZero))
	v.entries = v.entries[:len(v.entries)-1]

	forCheckBlockGt := v.functions[len(v.functions)-1].NewBlock("")
	v.entries = append(v.entries, forCheckBlockGt)
	forGtExpression := v.currentEntry().NewCall(v.funcs["check"], v.currentEntry().NewCall(v.funcs["lt"], expressions[0], expressions[1]))
	forCheckBlockGt.NewCondBr(forGtExpression, forBodyBlock, endBlock)
	v.entries = v.entries[:len(v.entries)-1]

	forCheckBlockLt := v.functions[len(v.functions)-1].NewBlock("")
	v.entries = append(v.entries, forCheckBlockLt)
	forLtExpression := v.currentEntry().NewCall(v.funcs["check"], v.currentEntry().NewCall(v.funcs["gt"], expressions[0], expressions[1]))
	forCheckBlockLt.NewCondBr(forLtExpression, forBodyBlock, endBlock)
	v.entries = v.entries[:len(v.entries)-1]

	forCheckBlock.NewCondBr(forCheckExpr, forCheckBlockLt, forCheckBlockGt)

	v.entries = append(v.entries, bodyEndBlock)
	forMoveExpr := bodyEndBlock.NewCall(v.funcs["add"], expressions[0], expressions[2])
	bodyEndBlock.NewCall(v.funcs["copy"], forMoveExpr, expressions[0])
	v.entries = v.entries[:len(v.entries)-1]

	bodyEndBlock.NewBr(forCheckBlock)

	v.entries[len(v.entries)-1] = endBlock

	return endBlock
}

func (v *Visitor) VisitStatGenericFor(ctx *parser.StatGenericForContext) interface{} {
	curIndexConst := constant.NewInt(types.I64, 0)
	curIndexPtr := constant.NewIntToPtr(
		curIndexConst,
		types.I8Ptr,
	)
	gencurIndex := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), curIndexPtr)

	oneConst := constant.NewInt(types.I64, 1)
	onePtr := constant.NewIntToPtr(
		oneConst,
		types.I8Ptr,
	)
	genOne := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), onePtr)

	iVariables := v.VisitNamelist(ctx.Namelist().(*parser.NamelistContext))
	variables := iVariables.([]variableLeft)
	iTable := v.VisitExplist(ctx.Explist().(*parser.ExplistContext))
	tables := iTable.([]expression)
	if len(tables) == 0 {
		v.ErrorList = append(v.ErrorList, fmt.Errorf("empty explist"))
		return nil
	}
	table := tables[0]

	var endBlock *ir.Block
	switch len(variables) {
	case 2:
		v.declaredVarList = append(v.declaredVarList, variables[0].name)
		v.declaredVarList = append(v.declaredVarList, variables[1].name)
		nilVal1 := v.currentEntry().NewCall(v.funcs["create_nil"])
		nilVal2 := v.currentEntry().NewCall(v.funcs["create_nil"])
		v.variables[len(v.variables)-1][variables[0].name] = nilVal1
		v.variables[len(v.variables)-1][variables[1].name] = nilVal2

		endBlock = v.functions[len(v.functions)-1].NewBlock("")

		bodyStartBlock := v.functions[len(v.functions)-1].NewBlock("")
		v.entries = append(v.entries, bodyStartBlock)
		forKeyExpr := bodyStartBlock.NewCall(v.funcs["lua_table_get_key_at"], table.value, gencurIndex)
		forValExpr := bodyStartBlock.NewCall(v.funcs["lua_table_get_value_at"], table.value, gencurIndex)
		bodyStartBlock.NewCall(v.funcs["copy"], forKeyExpr, v.variables[len(v.variables)-1][variables[0].name])
		bodyStartBlock.NewCall(v.funcs["copy"], forValExpr, v.variables[len(v.variables)-1][variables[1].name])
		v.entries = v.entries[:len(v.entries)-1]

		forCheckBlock := v.functions[len(v.functions)-1].NewBlock("")
		v.currentEntry().NewBr(forCheckBlock)

		bodyEndBlock := v.functions[len(v.functions)-1].NewBlock("")

		v.endBlocks = append(v.endBlocks, bodyEndBlock)

		forBodyBlock := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)

		v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

		v.entries = append(v.entries, forCheckBlock)
		forCheckExpr := v.currentEntry().NewCall(v.funcs["check"], v.currentEntry().NewCall(v.funcs["ge"], gencurIndex, v.currentEntry().NewCall(v.funcs["lua_table_len"], table.value)))
		v.entries = v.entries[:len(v.entries)-1]
		forCheckBlock.NewCondBr(forCheckExpr, endBlock, bodyStartBlock)
		bodyStartBlock.NewBr(forBodyBlock)

		v.entries = append(v.entries, bodyEndBlock)
		newIndexExpr := bodyEndBlock.NewCall(v.funcs["add"], gencurIndex, genOne)
		bodyEndBlock.NewCall(v.funcs["copy"], newIndexExpr, gencurIndex)
		v.entries = v.entries[:len(v.entries)-1]

		bodyEndBlock.NewBr(forCheckBlock)

		v.entries[len(v.entries)-1] = endBlock
	case 1:
		v.declaredVarList = append(v.declaredVarList, variables[0].name)
		nilVal1 := v.currentEntry().NewCall(v.funcs["create_nil"])
		v.variables[len(v.variables)-1][variables[0].name] = nilVal1

		bodyStartBlock := v.functions[len(v.functions)-1].NewBlock("")
		v.entries = append(v.entries, bodyStartBlock)
		forKeyExpr := bodyStartBlock.NewCall(v.funcs["lua_table_get_key_at"], table.value, gencurIndex)
		bodyStartBlock.NewCall(v.funcs["copy"], forKeyExpr, v.variables[len(v.variables)-1][variables[0].name])
		v.entries = v.entries[:len(v.entries)-1]

		forCheckBlock := v.functions[len(v.functions)-1].NewBlock("")
		v.currentEntry().NewBr(forCheckBlock)

		endBlock = v.functions[len(v.functions)-1].NewBlock("")
		bodyEndBlock := v.functions[len(v.functions)-1].NewBlock("")

		v.endBlocks = append(v.endBlocks, bodyEndBlock)

		forBodyBlock := v.VisitBlock(ctx.Block().(*parser.BlockContext)).(*ir.Block)

		v.endBlocks = v.endBlocks[:len(v.endBlocks)-1]

		v.entries = append(v.entries, forCheckBlock)
		forCheckExpr := v.currentEntry().NewCall(v.funcs["check"], v.currentEntry().NewCall(v.funcs["ge"], gencurIndex, v.currentEntry().NewCall(v.funcs["lua_table_len"], table.value)))
		v.entries = v.entries[:len(v.entries)-1]
		forCheckBlock.NewCondBr(forCheckExpr, endBlock, bodyStartBlock)
		bodyStartBlock.NewBr(forBodyBlock)

		v.entries = append(v.entries, bodyEndBlock)
		newIndexExpr := bodyEndBlock.NewCall(v.funcs["add"], gencurIndex, genOne)
		bodyEndBlock.NewCall(v.funcs["copy"], newIndexExpr, gencurIndex)
		v.entries = v.entries[:len(v.entries)-1]

		bodyEndBlock.NewBr(forCheckBlock)

		v.entries[len(v.entries)-1] = endBlock
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid variables len"))
		return nil
	}

	return endBlock
}

func (v *Visitor) VisitStatFunction(ctx *parser.StatFunctionContext) interface{} {
	name := v.VisitFuncname(ctx.Funcname().(*parser.FuncnameContext)).(string)

	funcBodyCtx := ctx.Funcbody().(*parser.FuncbodyContext)
	params := funcBodyCtx.Parlist().(*parser.ParlistContext).Namelist().(*parser.NamelistContext).AllNAME()

	v.variables = append(v.variables, make(map[string]value.Value))
	oldEntries := v.entries
	v.entries = nil

	funcParams := make([]*ir.Param, len(params))
	for i := range params {
		param := ir.NewParam(params[i].GetText(), types.NewPointer(v.typeDefs["Generic"]))
		v.declaredVarList = append(v.declaredVarList, params[i].GetText())
		v.variables[len(v.variables)-1][params[i].GetText()] = param
		funcParams[i] = param
	}

	f := v.IR.NewFunc(name, types.NewPointer(v.typeDefs["Generic"]), funcParams...)
	v.funcs[name] = f
	v.functions = append(v.functions, f)

	_ = v.VisitBlock(funcBodyCtx.Block().(*parser.BlockContext)).(*ir.Block)

	v.declaredVarList = v.declaredVarList[:len(v.declaredVarList)-len(params)]
	v.variables = v.variables[:len(v.variables)-1]
	v.functions = v.functions[:len(v.functions)-1]
	v.entries = oldEntries

	return nil
}

func (v *Visitor) VisitStatLocalFunction(ctx *parser.StatLocalFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatLocalAttributeNameList(ctx *parser.StatLocalAttributeNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttnamelist(ctx *parser.AttnamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttrib(ctx *parser.AttribContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatReturn(ctx *parser.StatReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	variables := make([]variableLeft, 0, len(ctx.AllVar_()))
	for _, variable := range ctx.AllVar_() {
		name := variable.NAME().GetText()

		if len(variable.AllVarSuffix()) > 0 {
			if !slices.Contains(v.declaredVarList, name) {
				v.ErrorList = append(v.ErrorList, fmt.Errorf("var %s used but not declared, line: %d, column: %d", name, variable.GetStart().GetLine(), variable.GetStart().GetColumn()))
				continue
			}

			children := variable.AllVarSuffix()[0].GetChildren()
			// spew.Dump(variable.AllVarSuffix())
			// spew.Dump(variable.AllVarSuffix()[0].GetChildren())
			key := children[1]

			var index interface{}
			switch keyCtx := key.(type) {
			case *parser.ExpPrefixExpContext:
				index = v.VisitExpPrefixExp(keyCtx)
			case *parser.ExpTrueContext:
				index = v.VisitExpTrue(keyCtx)
			case *parser.ExpFalseContext:
				index = v.VisitExpFalse(keyCtx)
			case *parser.ExpStringContext:
				index = v.VisitExpString(keyCtx)
			case *parser.ExpNumberContext:
				index = v.VisitExpNumber(keyCtx)
			case *parser.ExpOperatorPowerContext:
				index = v.VisitExpOperatorPower(keyCtx)
			case *parser.ExpOperatorUnaryContext:
				index = v.VisitExpOperatorUnary(keyCtx)
			case *parser.ExpOperatorMulDivModContext:
				index = v.VisitExpOperatorMulDivMod(keyCtx)
			case *parser.ExpOperatorAddSubContext:
				index = v.VisitExpOperatorAddSub(keyCtx)
			case *parser.ExpOperatorStrcatContext:
				index = v.VisitExpOperatorStrcat(keyCtx)
			case *parser.ExpOperatorComparisonContext:
				index = v.VisitExpOperatorComparison(keyCtx)
			case *parser.ExpOperatorAndContext:
				index = v.VisitExpOperatorAnd(keyCtx)
			case *parser.ExpOperatorOrContext:
				index = v.VisitExpOperatorOr(keyCtx)
			case *antlr.TerminalNodeImpl:
				index = v.VisitTerminalNode(keyCtx)
			default:
				v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid key type, line: %d, column: %d", variable.GetStart().GetLine(), variable.GetStart().GetColumn()))
				continue
			}

			val, ok := index.(value.Value)
			if !ok {
				v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
				return nil
			}

			variables = append(variables, variableLeft{
				name: name,
				key:  val,
			})
		} else {
			v.declaredVarList = append(v.declaredVarList, name)
			variables = append(variables, variableLeft{
				name: name,
			})
		}
	}

	return variables
}

func (v *Visitor) VisitTerminalNode(ctx *antlr.TerminalNodeImpl) interface{} {
	tm := ctx.GetText() + "\x00"
	strType := types.NewArray(uint64(len(tm)), types.I8)

	globalStr := v.IR.NewGlobalDef(fmt.Sprintf(".str%d", v.globalsCounter), constant.NewCharArray([]byte(tm)))

	zero := constant.NewInt(types.I32, 0)
	strPtr := constant.NewGetElementPtr(
		strType,
		globalStr,
		zero,
		zero,
	)

	v.globalsCounter++

	genStr := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 2), strPtr)

	return genStr
}

func (v *Visitor) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	variables := make([]variableLeft, 0, len(ctx.AllNAME()))
	for _, n := range ctx.AllNAME() {
		if n.GetText() == "," {
			continue
		}

		variables = append(variables, variableLeft{
			name: n.GetText(),
		})
	}

	return variables
}

func (v *Visitor) VisitExplist(ctx *parser.ExplistContext) interface{} {
	expressions := make([]expression, 0, len(ctx.AllExp()))
	for _, exp := range ctx.AllExp() {
		index := v.VisitExp(exp)
		if index == nil {
			continue
		}

		val, ok := index.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		expressions = append(expressions, expression{
			value: val,
		})
	}

	return expressions
}

func (v *Visitor) VisitExpFalse(ctx *parser.ExpFalseContext) interface{} {
	boolConst := constant.NewInt(types.I8, 0)
	boolPtr := constant.NewIntToPtr(
		boolConst,
		types.I8Ptr,
	)
	genBool := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 3), boolPtr)

	return genBool
}

func (v *Visitor) VisitExpVararg(ctx *parser.ExpVarargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) GetTableRet(exp []expression) value.Value {
	if len(exp) == 1 {
		return exp[0].value
	}

	table := v.currentEntry().NewCall(v.funcs["lua_table_new"])

	count := 0
	for _, e := range exp {
		numConst := constant.NewInt(types.I64, int64(count))
		numPtr := constant.NewIntToPtr(
			numConst,
			types.I8Ptr,
		)
		genInt := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), numPtr)

		v.currentEntry().NewCall(v.funcs["lua_table_set"], table, genInt, e.value)
		count++
	}

	return table
}

func (v *Visitor) VisitExpTableConstructor(ctx *parser.ExpTableConstructorContext) interface{} {
	table := v.currentEntry().NewCall(v.funcs["lua_table_new"])
	if ctx.Tableconstructor().Fieldlist() == nil {
		return table
	}

	count := 0
	for _, f := range ctx.Tableconstructor().Fieldlist().AllField() {
		fieldExp := make([]value.Value, 0, len(f.AllExp()))
		for _, e := range f.AllExp() {
			fieldExp = append(fieldExp, v.VisitExp(e).(value.Value))
		}

		switch len(fieldExp) {
		case 2:
			v.currentEntry().NewCall(v.funcs["lua_table_set"], table, fieldExp[0], fieldExp[1])
		case 1:
			tm := f.NAME()
			if tm == nil {
				numConst := constant.NewInt(types.I64, int64(count))
				numPtr := constant.NewIntToPtr(
					numConst,
					types.I8Ptr,
				)
				genInt := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 0), numPtr)

				v.currentEntry().NewCall(v.funcs["lua_table_set"], table, genInt, fieldExp[0])
				count++
			} else {
				s := f.NAME().GetText() + "\x00"
				strType := types.NewArray(uint64(len(s)), types.I8)

				globalStr := v.IR.NewGlobalDef(fmt.Sprintf(".str%d", v.globalsCounter), constant.NewCharArray([]byte(s)))

				zero := constant.NewInt(types.I32, 0)
				strPtr := constant.NewGetElementPtr(
					strType,
					globalStr,
					zero,
					zero,
				)

				v.globalsCounter++

				genStr := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 2), strPtr)

				v.currentEntry().NewCall(v.funcs["lua_table_set"], table, genStr, fieldExp[0])
			}
		default:
			v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid table constructor"))
			return nil
		}
	}

	return table
}

func (v *Visitor) VisitExpPrefixExp(ctx *parser.ExpPrefixExpContext) interface{} {
	if len(ctx.Prefixexp().AllNameAndArgs()) > 0 {
		name := ctx.Prefixexp().VarOrExp().Var_().NAME().GetText()

		explist := v.VisitExplist(ctx.Prefixexp().AllNameAndArgs()[0].Args().Explist().(*parser.ExplistContext)).([]expression)
		values := make([]value.Value, len(explist))
		for i := range explist {
			values[i] = explist[i].value
		}

		return v.currentEntry().NewCall(v.funcs[name], values...)
	}

	return v.VisitVarOrExp(ctx.Prefixexp().VarOrExp().(*parser.VarOrExpContext))
}

func (v *Visitor) VisitExpTrue(ctx *parser.ExpTrueContext) interface{} {
	boolConst := constant.NewInt(types.I8, 1)
	boolPtr := constant.NewIntToPtr(
		boolConst,
		types.I8Ptr,
	)
	genBool := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 3), boolPtr)

	return genBool
}

func (v *Visitor) VisitExpNumber(ctx *parser.ExpNumberContext) interface{} {
	return v.VisitNumber(ctx.Number().(*parser.NumberContext))
}

func (v *Visitor) VisitExpOperatorUnary(ctx *parser.ExpOperatorUnaryContext) interface{} {
	iExp := v.VisitExp(ctx.Exp())
	exp, ok := iExp.(value.Value)
	if !ok {
		v.ErrorList = append(v.ErrorList, fmt.Errorf("nil expression"))
		return nil
	}

	switch ctx.OperatorUnary().GetText() {
	case "-":
		return v.currentEntry().NewCall(v.funcs["neg"], exp)
	case "not":
		return v.currentEntry().NewCall(v.funcs["not"], exp)
	case "#":
		return v.currentEntry().NewCall(v.funcs["string_len"], exp)
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid unary operation"))
		return nil
	}
}

func (v *Visitor) VisitExpOperatorAnd(ctx *parser.ExpOperatorAndContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	return v.currentEntry().NewCall(v.funcs["and"], values[0], values[1])
}

func (v *Visitor) VisitExpOperatorPower(ctx *parser.ExpOperatorPowerContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	return v.currentEntry().NewCall(v.funcs["power"], values[0], values[1])
}

func (v *Visitor) VisitExpOperatorAddSub(ctx *parser.ExpOperatorAddSubContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	operator := ctx.OperatorAddSub().(*parser.OperatorAddSubContext)
	switch operator.GetText() {
	case "+":
		return v.currentEntry().NewCall(v.funcs["add"], values[0], values[1])
	case "-":
		return v.currentEntry().NewCall(v.funcs["sub"], values[0], values[1])
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid add sub operation"))
		return nil
	}
}

func (v *Visitor) VisitExpOperatorStrcat(ctx *parser.ExpOperatorStrcatContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	result := v.currentEntry().NewCall(v.funcs["concat"], values[0], values[1])
	for i := 2; i < len(values); i++ {
		result = v.currentEntry().NewCall(v.funcs["concat"], result, values[i])
	}

	return result
}

func (v *Visitor) VisitExpOperatorComparison(ctx *parser.ExpOperatorComparisonContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	operator := ctx.OperatorComparison().(*parser.OperatorComparisonContext)
	switch operator.GetText() {
	case "==":
		return v.currentEntry().NewCall(v.funcs["equal"], values[0], values[1])
	case "~=":
		return v.currentEntry().NewCall(v.funcs["nequal"], values[0], values[1])
	case ">":
		return v.currentEntry().NewCall(v.funcs["gt"], values[0], values[1])
	case ">=":
		return v.currentEntry().NewCall(v.funcs["ge"], values[0], values[1])
	case "<":
		return v.currentEntry().NewCall(v.funcs["lt"], values[0], values[1])
	case "<=":
		return v.currentEntry().NewCall(v.funcs["le"], values[0], values[1])
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid cmp operation"))
		return nil
	}
}

func (v *Visitor) VisitExpNil(ctx *parser.ExpNilContext) interface{} {
	return v.currentEntry().NewCall(v.funcs["create_nil"])
}

func (v *Visitor) VisitExpOperatorOr(ctx *parser.ExpOperatorOrContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	return v.currentEntry().NewCall(v.funcs["or"], values[0], values[1])
}

func parseUnicodeEscape(s string) (string, error) {
	reUnicode := regexp.MustCompile(`\\u\{([0-9a-fA-F]+)\}`)
	reDecimal := regexp.MustCompile(`\\([0-2][0-9]{2}|[0-9]{1,2})`)
	reStandard := regexp.MustCompile(`\\(x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8}|[abfnrtv\\'"])`)

	replaced := reUnicode.ReplaceAllStringFunc(s, func(match string) string {
		hexPart := reUnicode.FindStringSubmatch(match)[1]
		code, _ := strconv.ParseInt(hexPart, 16, 32)
		if code <= 0xFFFF {
			return fmt.Sprintf(`\u%04X`, code)
		}
		return fmt.Sprintf(`\U%08X`, code)
	})

	replaced = reDecimal.ReplaceAllStringFunc(replaced, func(match string) string {
		digits := reDecimal.FindStringSubmatch(match)[1]
		code, _ := strconv.ParseInt(digits, 10, 32)
		switch {
		case code <= 0xFF:
			return fmt.Sprintf(`\x%02X`, code)
		case code <= 0xFFFF:
			return fmt.Sprintf(`\u%04X`, code)
		default:
			return fmt.Sprintf(`\U%08X`, code)
		}
	})

	unquoted, err := strconv.Unquote(replaced)
	if err != nil {
		if repaired := reStandard.ReplaceAllString(replaced, `\\$1`); repaired != replaced {
			return strconv.Unquote(`"` + repaired + `"`)
		}
		return "", err
	}
	return unquoted, nil
}

func (v *Visitor) VisitExpString(ctx *parser.ExpStringContext) interface{} {
	s, _ := parseUnicodeEscape(ctx.GetText())

	s = s + "\x00"
	strType := types.NewArray(uint64(len(s)), types.I8)

	globalStr := v.IR.NewGlobalDef(fmt.Sprintf(".str%d", v.globalsCounter), constant.NewCharArray([]byte(s)))

	zero := constant.NewInt(types.I32, 0)
	strPtr := constant.NewGetElementPtr(
		strType,
		globalStr,
		zero,
		zero,
	)

	v.globalsCounter++

	genStr := v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 2), strPtr)

	return genStr
}

func (v *Visitor) VisitExpOperatorMulDivMod(ctx *parser.ExpOperatorMulDivModContext) interface{} {
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		viz := v.Visit(e)
		val, ok := viz.(value.Value)
		if !ok {
			v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
			return nil
		}

		values[i] = val
	}

	operator := ctx.OperatorMulDivMod().(*parser.OperatorMulDivModContext)
	switch operator.GetText() {
	case "*":
		return v.currentEntry().NewCall(v.funcs["mul"], values[0], values[1])
	case "/":
		return v.currentEntry().NewCall(v.funcs["div"], values[0], values[1])
	case "//":
		return v.currentEntry().NewCall(v.funcs["mod"], values[0], values[1])
	case "%":
		return v.currentEntry().NewCall(v.funcs["rem"], values[0], values[1])
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid operation"))
		return nil
	}
}

func (v *Visitor) VisitExpFunctionDef(ctx *parser.ExpFunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExp(ctx parser.IExpContext) interface{} {
	var index interface{}
	switch expCtx := ctx.(type) {
	case *parser.ExpNilContext:
		index = v.VisitExpNil(expCtx)
	case *parser.ExpTrueContext:
		index = v.VisitExpTrue(expCtx)
	case *parser.ExpFalseContext:
		index = v.VisitExpFalse(expCtx)
	case *parser.ExpStringContext:
		index = v.VisitExpString(expCtx)
	case *parser.ExpNumberContext:
		index = v.VisitExpNumber(expCtx)
	case *parser.ExpOperatorPowerContext:
		index = v.VisitExpOperatorPower(expCtx)
	case *parser.ExpOperatorUnaryContext:
		index = v.VisitExpOperatorUnary(expCtx)
	case *parser.ExpOperatorMulDivModContext:
		index = v.VisitExpOperatorMulDivMod(expCtx)
	case *parser.ExpOperatorAddSubContext:
		index = v.VisitExpOperatorAddSub(expCtx)
	case *parser.ExpOperatorStrcatContext:
		index = v.VisitExpOperatorStrcat(expCtx)
	case *parser.ExpOperatorComparisonContext:
		index = v.VisitExpOperatorComparison(expCtx)
	case *parser.ExpOperatorAndContext:
		index = v.VisitExpOperatorAnd(expCtx)
	case *parser.ExpOperatorOrContext:
		index = v.VisitExpOperatorOr(expCtx)
	case *parser.ExpVarargContext:
		index = v.VisitExpVararg(expCtx)
	case *parser.ExpFunctionDefContext:
		index = v.VisitExpFunctionDef(expCtx)
	case *parser.ExpPrefixExpContext:
		index = v.VisitExpPrefixExp(expCtx)
	case *parser.ExpTableConstructorContext:
		index = v.VisitExpTableConstructor(expCtx)
	default:
		v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid exp type, line: %d, column: %d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()))
		return nil
	}

	return index
}

func (v *Visitor) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	if ctx.Exp() != nil {
		exp := v.VisitExp(ctx.Exp())
		return exp
	}

	if ctx.Var_() != nil {
		if len(ctx.Var_().AllVarSuffix()) == 0 {
			variable, ok := v.variables[len(v.variables)-1][ctx.Var_().NAME().GetText()]
			if !ok {
				v.ErrorList = append(v.ErrorList, fmt.Errorf("var %s used but not declared, line: %d, column: %d", ctx.Var_().NAME().GetText(), ctx.Var_().GetStart().GetLine(), ctx.Var_().GetStart().GetColumn()))
				return nil
			}

			return variable
		}

		variable := v.variables[len(v.variables)-1][ctx.Var_().NAME().GetText()]
		suffix := ctx.Var_().AllVarSuffix()[0].GetChildren()[1]
		var key value.Value
		switch suffixWithType := suffix.(type) {
		case parser.IExpContext:
			viz := v.VisitExp(suffixWithType)
			val, ok := viz.(value.Value)
			if !ok {
				v.ErrorList = append(v.ErrorList, fmt.Errorf("nil value"))
				return nil
			}

			key = val
		case *antlr.TerminalNodeImpl:
			s := suffixWithType.GetText() + "\x00"
			strType := types.NewArray(uint64(len(s)), types.I8)

			globalStr := v.IR.NewGlobalDef(fmt.Sprintf(".str%d", v.globalsCounter), constant.NewCharArray([]byte(s)))

			zero := constant.NewInt(types.I32, 0)
			strPtr := constant.NewGetElementPtr(
				strType,
				globalStr,
				zero,
				zero,
			)

			v.globalsCounter++

			key = v.currentEntry().NewCall(v.funcs["create"], constant.NewInt(types.I32, 2), strPtr)
		default:
			v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid var suffix"))
			return nil
		}

		return v.currentEntry().NewCall(v.funcs["lua_table_get"], variable, key)
	}

	return nil
}

func (v *Visitor) VisitVar(ctx *parser.VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgs(ctx *parser.ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParlist(ctx *parser.ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitField(ctx *parser.FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

// not needed
func (v *Visitor) VisitString(ctx *parser.StringContext) interface{} {
	return v.VisitChildren(ctx)
}
