package visitor

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/parser"
)

type Visitor struct {
	*parser.BaseLuaVisitor

	declaredVarList []string
	visionScope     []int

	ErrorList []error

	IR *ir.Module

	entries []*ir.Block

	falseConst value.Value
	trueConst  value.Value
	nilConst   value.Value

	globalsCounter int

	funcs map[string]*ir.Func

	variables []map[string]value.Value
}

func (v *Visitor) currentEntry() *ir.Block {
	return v.entries[len(v.entries)-1]
}

//go:embed common.ll
var commonLLVM string

func New() *Visitor {
	llvm, err := asm.ParseString("", commonLLVM)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// strlen := llvm.NewFunc("strlen", types.I64, ir.NewParam("", types.I8Ptr))
	// malloc := llvm.NewFunc("malloc", types.I8Ptr, ir.NewParam("", types.I64))
	// memcpy := llvm.NewFunc("memcpy", types.Void,
	// 	ir.NewParam("dest", types.I8Ptr),
	// 	ir.NewParam("src", types.I8Ptr),
	// 	ir.NewParam("len", types.I64),
	// 	ir.NewParam("align", types.I1),
	// )

	funcs := make(map[string]*ir.Func)
	for _, f := range llvm.Funcs {
		funcs[f.Name()] = f
	}

	fmt.Printf("%#v\n", funcs)

	return &Visitor{
		BaseLuaVisitor:  &parser.BaseLuaVisitor{},
		IR:              llvm,
		declaredVarList: make([]string, 0),
		visionScope:     make([]int, 0),
		falseConst:      constant.NewInt(types.I64, 0),
		trueConst:       constant.NewInt(types.I64, 1),
		nilConst:        constant.NewInt(types.I64, 0),
		funcs:           funcs,
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitChunk(ctx *parser.ChunkContext) interface{} {
	main := v.IR.NewFunc("main", types.I64)
	entry := main.NewBlock("entry")

	v.entries = append(v.entries, entry)
	v.variables = append(v.variables, make(map[string]value.Value))

	v.VisitBlock(ctx.Block().(*parser.BlockContext))

	entry.NewRet(v.nilConst)

	return nil
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	fmt.Println("block")

	v.visionScope = append(v.visionScope, len(v.declaredVarList))

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

	fmt.Println("declaredVarList", v.declaredVarList)
	fmt.Println("visionScope", v.visionScope)

	v.declaredVarList = v.declaredVarList[:v.visionScope[len(v.visionScope)-1]]
	v.visionScope = v.visionScope[:len(v.visionScope)-1]

	return nil
}

func (v *Visitor) VisitStatAssignment(ctx *parser.StatAssignmentContext) interface{} {
	fmt.Println("assignment")

	varlist := v.VisitVarlist(ctx.Varlist().(*parser.VarlistContext)).([]variableLeft)
	explist := v.VisitExplist(ctx.Explist().(*parser.ExplistContext)).([]expression)

	if len(varlist) != len(explist) {
		v.ErrorList = append(v.ErrorList, fmt.Errorf("len of varlist and exp list mismatch %d != %d", len(varlist), len(explist)))
		return nil
	}

	for i := range varlist {
		if varlist[i].key != nil {
			continue
		}

		v.variables[len(v.variables)-1][varlist[i].name] = explist[i].value
	}

	fmt.Println(varlist)
	fmt.Println(explist)

	return nil
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	fmt.Println("number")

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
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatWhile(ctx *parser.StatWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatRepeat(ctx *parser.StatRepeatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatIfThenElse(ctx *parser.StatIfThenElseContext) interface{} {
	fmt.Println("if")
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatNumericFor(ctx *parser.StatNumericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatGenericFor(ctx *parser.StatGenericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatFunction(ctx *parser.StatFunctionContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	fmt.Println("varlist")

	variables := make([]variableLeft, 0, len(ctx.AllVar_()))

	for _, variable := range ctx.AllVar_() {
		name := variable.NAME().GetText()

		if len(variable.AllVarSuffix()) > 0 {
			// [1] [""] .ss
			fmt.Println("old", variable.NAME())
			if !slices.Contains(v.declaredVarList, name) {
				v.ErrorList = append(v.ErrorList, fmt.Errorf("var %s used but not declared, line: %d, column: %d", name, variable.GetStart().GetLine(), variable.GetStart().GetColumn()))
				continue
			}

			children := variable.AllVarSuffix()[0].GetChildren()
			key := children[1]

			fmt.Printf("%#v %T\n", key, key)

			var index interface{}
			switch keyCtx := key.(type) {
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
				index = v.VisitTerminal(keyCtx)
			default:
				v.ErrorList = append(v.ErrorList, fmt.Errorf("invalid key type, line: %d, column: %d", variable.GetStart().GetLine(), variable.GetStart().GetColumn()))
				continue
			}

			variables = append(variables, variableLeft{
				name: name,
				key:  index.(value.Value),
			})
		} else {
			fmt.Println("new", variable.NAME())
			v.declaredVarList = append(v.declaredVarList, name)
			variables = append(variables, variableLeft{
				name: name,
			})
		}
	}

	return variables
}

func (v *Visitor) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	fmt.Println("namelist")

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
	fmt.Println("explist")

	expressions := make([]expression, 0, len(ctx.AllExp()))
	for _, exp := range ctx.AllExp() {
		index := v.VisitExp(exp)
		if index == nil {
			continue
		}

		fmt.Printf("%T %#v\n", exp, exp)

		expressions = append(expressions, expression{
			value: index.(value.Value),
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

func (v *Visitor) VisitExpTableConstructor(ctx *parser.ExpTableConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpPrefixExp(ctx *parser.ExpPrefixExpContext) interface{} {
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
	fmt.Println("exp number")

	return v.VisitNumber(ctx.Number().(*parser.NumberContext))
}

func (v *Visitor) VisitExpOperatorUnary(ctx *parser.ExpOperatorUnaryContext) interface{} {
	exp := v.VisitExp(ctx.Exp()).(value.Value)

	switch ctx.OperatorUnary().GetText() {
	case "-":
		return v.currentEntry().NewCall(v.funcs["neg"], exp)
	case "not":
		return v.currentEntry().NewCall(v.funcs["not"], exp)
	case "#":
		return v.currentEntry().NewCall(v.funcs["string_len"], exp)
	default:
		return nil
	}
}

func (v *Visitor) VisitExpOperatorAnd(ctx *parser.ExpOperatorAndContext) interface{} {
	fmt.Printf("AND %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
	}

	return v.currentEntry().NewCall(v.funcs["and"], values[0], values[1])
}

func (v *Visitor) VisitExpOperatorPower(ctx *parser.ExpOperatorPowerContext) interface{} {
	fmt.Printf("MULDIVMOD %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
	}

	return v.currentEntry().NewCall(v.funcs["power"], values[0], values[1])
}

func (v *Visitor) VisitExpOperatorAddSub(ctx *parser.ExpOperatorAddSubContext) interface{} {
	fmt.Printf("ADDSUB %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
	}

	operator := ctx.OperatorAddSub().(*parser.OperatorAddSubContext)
	switch operator.GetText() {
	case "+":
		return v.currentEntry().NewCall(v.funcs["add"], values[0], values[1])
	case "-":
		return v.currentEntry().NewCall(v.funcs["sub"], values[0], values[1])
	}

	return nil
}

func (v *Visitor) VisitExpOperatorStrcat(ctx *parser.ExpOperatorStrcatContext) interface{} {
	fmt.Println("strcat")
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
	}

	result := v.currentEntry().NewCall(v.funcs["concat"], values[0], values[1])
	for i := 2; i < len(values); i++ {
		result = v.currentEntry().NewCall(v.funcs["concat"], result, values[i])
	}

	return result
}

func (v *Visitor) VisitExpOperatorComparison(ctx *parser.ExpOperatorComparisonContext) interface{} {
	fmt.Printf("CMP %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
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

	return nil
}

func (v *Visitor) VisitExpNil(ctx *parser.ExpNilContext) interface{} {
	return v.currentEntry().NewCall(v.funcs["create_nil"])
}

func (v *Visitor) VisitExpOperatorOr(ctx *parser.ExpOperatorOrContext) interface{} {
	fmt.Printf("AND %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
	}

	return v.currentEntry().NewCall(v.funcs["or"], values[0], values[1])
}

func (v *Visitor) VisitExpString(ctx *parser.ExpStringContext) interface{} {
	fmt.Println("string")

	s := strings.Trim(ctx.GetText(), "\"")

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
	fmt.Printf("MULDIVMOD %#v\n", ctx)
	expressions := ctx.AllExp()
	values := make([]value.Value, len(expressions))
	for i, e := range expressions {
		values[i] = v.Visit(e).(value.Value)
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
		fmt.Println("!!!", exp)
		return exp
	}

	if ctx.Var_() != nil {
		return v.variables[len(v.variables)-1][ctx.Var_().NAME().GetText()]
	}

	fmt.Println("!PARASHA!!")

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
