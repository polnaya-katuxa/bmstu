// Code generated from Lua.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LuaParser struct {
	*antlr.BaseParser
}

var LuaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func luaParserInit() {
	staticData := &LuaParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'='", "'do'", "'end'", "'while'", "'repeat'", "'until'",
		"'if'", "'then'", "'elseif'", "'else'", "'for'", "','", "'in'", "'function'",
		"'local'", "'<'", "'>'", "'return'", "'.'", "':'", "'nil'", "'false'",
		"'true'", "'...'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'or'",
		"'and'", "'<='", "'>='", "'~='", "'=='", "'..'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'//'", "'not'", "'#'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NAME", "NORMALSTRING",
		"INT", "FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
	}
	staticData.RuleNames = []string{
		"chunk", "block", "stat", "attnamelist", "attrib", "retstat", "funcname",
		"varlist", "namelist", "explist", "exp", "prefixexp", "functioncall",
		"varOrExp", "var", "varSuffix", "nameAndArgs", "args", "functiondef",
		"funcbody", "parlist", "tableconstructor", "fieldlist", "field", "fieldsep",
		"operatorOr", "operatorAnd", "operatorComparison", "operatorStrcat",
		"operatorAddSub", "operatorMulDivMod", "operatorUnary", "operatorPower",
		"number", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 402, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0, 1, 1, 5,
		1, 75, 8, 1, 10, 1, 12, 1, 78, 9, 1, 1, 1, 3, 1, 81, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 113, 8, 2, 10, 2, 12, 2, 116, 9, 2, 1, 2,
		1, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 132, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 158, 8, 2, 3, 2, 160, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 167, 8, 3, 10, 3, 12, 3, 170, 9, 3, 1, 4, 1,
		4, 1, 4, 3, 4, 175, 8, 4, 1, 5, 1, 5, 3, 5, 179, 8, 5, 1, 5, 3, 5, 182,
		8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 187, 8, 6, 10, 6, 12, 6, 190, 9, 6, 1, 6,
		1, 6, 3, 6, 194, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 199, 8, 7, 10, 7, 12, 7,
		202, 9, 7, 1, 8, 1, 8, 1, 8, 5, 8, 207, 8, 8, 10, 8, 12, 8, 210, 9, 8,
		1, 9, 1, 9, 1, 9, 5, 9, 215, 8, 9, 10, 9, 12, 9, 218, 9, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 233, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 263, 8, 10, 10, 10, 12, 10, 266, 9, 10, 1, 11, 1, 11, 5, 11, 270,
		8, 11, 10, 11, 12, 11, 273, 9, 11, 1, 12, 1, 12, 4, 12, 277, 8, 12, 11,
		12, 12, 12, 278, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 286, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 294, 8, 14, 1, 14, 5,
		14, 297, 8, 14, 10, 14, 12, 14, 300, 9, 14, 1, 15, 5, 15, 303, 8, 15, 10,
		15, 12, 15, 306, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		314, 8, 15, 1, 16, 1, 16, 3, 16, 318, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		3, 17, 324, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 329, 8, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 3, 19, 336, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 3, 20, 345, 8, 20, 1, 20, 3, 20, 348, 8, 20, 1, 21, 1,
		21, 3, 21, 352, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		360, 8, 22, 10, 22, 12, 22, 363, 9, 22, 1, 22, 3, 22, 366, 8, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 378,
		8, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 0, 1, 20, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 0, 6, 2, 0, 1, 1, 13, 13, 2, 0, 17, 18, 34, 37,
		1, 0, 39, 40, 1, 0, 41, 44, 2, 0, 40, 40, 45, 46, 1, 0, 50, 51, 427, 0,
		70, 1, 0, 0, 0, 2, 76, 1, 0, 0, 0, 4, 159, 1, 0, 0, 0, 6, 161, 1, 0, 0,
		0, 8, 174, 1, 0, 0, 0, 10, 176, 1, 0, 0, 0, 12, 183, 1, 0, 0, 0, 14, 195,
		1, 0, 0, 0, 16, 203, 1, 0, 0, 0, 18, 211, 1, 0, 0, 0, 20, 232, 1, 0, 0,
		0, 22, 267, 1, 0, 0, 0, 24, 274, 1, 0, 0, 0, 26, 285, 1, 0, 0, 0, 28, 293,
		1, 0, 0, 0, 30, 304, 1, 0, 0, 0, 32, 317, 1, 0, 0, 0, 34, 328, 1, 0, 0,
		0, 36, 330, 1, 0, 0, 0, 38, 333, 1, 0, 0, 0, 40, 347, 1, 0, 0, 0, 42, 349,
		1, 0, 0, 0, 44, 355, 1, 0, 0, 0, 46, 377, 1, 0, 0, 0, 48, 379, 1, 0, 0,
		0, 50, 381, 1, 0, 0, 0, 52, 383, 1, 0, 0, 0, 54, 385, 1, 0, 0, 0, 56, 387,
		1, 0, 0, 0, 58, 389, 1, 0, 0, 0, 60, 391, 1, 0, 0, 0, 62, 393, 1, 0, 0,
		0, 64, 395, 1, 0, 0, 0, 66, 397, 1, 0, 0, 0, 68, 399, 1, 0, 0, 0, 70, 71,
		3, 2, 1, 0, 71, 72, 5, 0, 0, 1, 72, 1, 1, 0, 0, 0, 73, 75, 3, 4, 2, 0,
		74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 81, 3, 10, 5, 0, 80,
		79, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 3, 1, 0, 0, 0, 82, 160, 5, 1, 0,
		0, 83, 84, 3, 14, 7, 0, 84, 85, 5, 2, 0, 0, 85, 86, 3, 18, 9, 0, 86, 160,
		1, 0, 0, 0, 87, 160, 3, 24, 12, 0, 88, 89, 5, 3, 0, 0, 89, 90, 3, 2, 1,
		0, 90, 91, 5, 4, 0, 0, 91, 160, 1, 0, 0, 0, 92, 93, 5, 5, 0, 0, 93, 94,
		3, 20, 10, 0, 94, 95, 5, 3, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 4, 0,
		0, 97, 160, 1, 0, 0, 0, 98, 99, 5, 6, 0, 0, 99, 100, 3, 2, 1, 0, 100, 101,
		5, 7, 0, 0, 101, 102, 3, 20, 10, 0, 102, 160, 1, 0, 0, 0, 103, 104, 5,
		8, 0, 0, 104, 105, 3, 20, 10, 0, 105, 106, 5, 9, 0, 0, 106, 114, 3, 2,
		1, 0, 107, 108, 5, 10, 0, 0, 108, 109, 3, 20, 10, 0, 109, 110, 5, 9, 0,
		0, 110, 111, 3, 2, 1, 0, 111, 113, 1, 0, 0, 0, 112, 107, 1, 0, 0, 0, 113,
		116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 119,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 11, 0, 0, 118, 120, 3, 2,
		1, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0,
		121, 122, 5, 4, 0, 0, 122, 160, 1, 0, 0, 0, 123, 124, 5, 12, 0, 0, 124,
		125, 5, 48, 0, 0, 125, 126, 5, 2, 0, 0, 126, 127, 3, 20, 10, 0, 127, 128,
		5, 13, 0, 0, 128, 131, 3, 20, 10, 0, 129, 130, 5, 13, 0, 0, 130, 132, 3,
		20, 10, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 5, 3, 0, 0, 134, 135, 3, 2, 1, 0, 135, 136, 5, 4, 0, 0,
		136, 160, 1, 0, 0, 0, 137, 138, 5, 12, 0, 0, 138, 139, 3, 16, 8, 0, 139,
		140, 5, 14, 0, 0, 140, 141, 3, 18, 9, 0, 141, 142, 5, 3, 0, 0, 142, 143,
		3, 2, 1, 0, 143, 144, 5, 4, 0, 0, 144, 160, 1, 0, 0, 0, 145, 146, 5, 15,
		0, 0, 146, 147, 3, 12, 6, 0, 147, 148, 3, 38, 19, 0, 148, 160, 1, 0, 0,
		0, 149, 150, 5, 16, 0, 0, 150, 151, 5, 15, 0, 0, 151, 152, 5, 48, 0, 0,
		152, 160, 3, 38, 19, 0, 153, 154, 5, 16, 0, 0, 154, 157, 3, 6, 3, 0, 155,
		156, 5, 2, 0, 0, 156, 158, 3, 18, 9, 0, 157, 155, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 82, 1, 0, 0, 0, 159, 83, 1, 0, 0,
		0, 159, 87, 1, 0, 0, 0, 159, 88, 1, 0, 0, 0, 159, 92, 1, 0, 0, 0, 159,
		98, 1, 0, 0, 0, 159, 103, 1, 0, 0, 0, 159, 123, 1, 0, 0, 0, 159, 137, 1,
		0, 0, 0, 159, 145, 1, 0, 0, 0, 159, 149, 1, 0, 0, 0, 159, 153, 1, 0, 0,
		0, 160, 5, 1, 0, 0, 0, 161, 162, 5, 48, 0, 0, 162, 168, 3, 8, 4, 0, 163,
		164, 5, 13, 0, 0, 164, 165, 5, 48, 0, 0, 165, 167, 3, 8, 4, 0, 166, 163,
		1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0,
		0, 0, 169, 7, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 172, 5, 17, 0, 0,
		172, 173, 5, 48, 0, 0, 173, 175, 5, 18, 0, 0, 174, 171, 1, 0, 0, 0, 174,
		175, 1, 0, 0, 0, 175, 9, 1, 0, 0, 0, 176, 178, 5, 19, 0, 0, 177, 179, 3,
		18, 9, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0,
		0, 180, 182, 5, 1, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		11, 1, 0, 0, 0, 183, 188, 5, 48, 0, 0, 184, 185, 5, 20, 0, 0, 185, 187,
		5, 48, 0, 0, 186, 184, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0,
		0, 0, 188, 189, 1, 0, 0, 0, 189, 193, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0,
		191, 192, 5, 21, 0, 0, 192, 194, 5, 48, 0, 0, 193, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 13, 1, 0, 0, 0, 195, 200, 3, 28, 14, 0, 196, 197,
		5, 13, 0, 0, 197, 199, 3, 28, 14, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1,
		0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 15, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 203, 208, 5, 48, 0, 0, 204, 205, 5, 13, 0, 0,
		205, 207, 5, 48, 0, 0, 206, 204, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208,
		206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 17, 1, 0, 0, 0, 210, 208, 1,
		0, 0, 0, 211, 216, 3, 20, 10, 0, 212, 213, 5, 13, 0, 0, 213, 215, 3, 20,
		10, 0, 214, 212, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0,
		216, 217, 1, 0, 0, 0, 217, 19, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220,
		6, 10, -1, 0, 220, 233, 5, 22, 0, 0, 221, 233, 5, 23, 0, 0, 222, 233, 5,
		24, 0, 0, 223, 233, 3, 66, 33, 0, 224, 233, 3, 68, 34, 0, 225, 233, 5,
		25, 0, 0, 226, 233, 3, 36, 18, 0, 227, 233, 3, 22, 11, 0, 228, 233, 3,
		42, 21, 0, 229, 230, 3, 62, 31, 0, 230, 231, 3, 20, 10, 7, 231, 233, 1,
		0, 0, 0, 232, 219, 1, 0, 0, 0, 232, 221, 1, 0, 0, 0, 232, 222, 1, 0, 0,
		0, 232, 223, 1, 0, 0, 0, 232, 224, 1, 0, 0, 0, 232, 225, 1, 0, 0, 0, 232,
		226, 1, 0, 0, 0, 232, 227, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0, 232, 229,
		1, 0, 0, 0, 233, 264, 1, 0, 0, 0, 234, 235, 10, 8, 0, 0, 235, 236, 3, 64,
		32, 0, 236, 237, 3, 20, 10, 8, 237, 263, 1, 0, 0, 0, 238, 239, 10, 6, 0,
		0, 239, 240, 3, 60, 30, 0, 240, 241, 3, 20, 10, 7, 241, 263, 1, 0, 0, 0,
		242, 243, 10, 5, 0, 0, 243, 244, 3, 58, 29, 0, 244, 245, 3, 20, 10, 6,
		245, 263, 1, 0, 0, 0, 246, 247, 10, 4, 0, 0, 247, 248, 3, 56, 28, 0, 248,
		249, 3, 20, 10, 4, 249, 263, 1, 0, 0, 0, 250, 251, 10, 3, 0, 0, 251, 252,
		3, 54, 27, 0, 252, 253, 3, 20, 10, 4, 253, 263, 1, 0, 0, 0, 254, 255, 10,
		2, 0, 0, 255, 256, 3, 52, 26, 0, 256, 257, 3, 20, 10, 3, 257, 263, 1, 0,
		0, 0, 258, 259, 10, 1, 0, 0, 259, 260, 3, 50, 25, 0, 260, 261, 3, 20, 10,
		2, 261, 263, 1, 0, 0, 0, 262, 234, 1, 0, 0, 0, 262, 238, 1, 0, 0, 0, 262,
		242, 1, 0, 0, 0, 262, 246, 1, 0, 0, 0, 262, 250, 1, 0, 0, 0, 262, 254,
		1, 0, 0, 0, 262, 258, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0,
		0, 0, 264, 265, 1, 0, 0, 0, 265, 21, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0,
		267, 271, 3, 26, 13, 0, 268, 270, 3, 32, 16, 0, 269, 268, 1, 0, 0, 0, 270,
		273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 23, 1,
		0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 276, 3, 26, 13, 0, 275, 277, 3, 32,
		16, 0, 276, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0,
		278, 279, 1, 0, 0, 0, 279, 25, 1, 0, 0, 0, 280, 286, 3, 28, 14, 0, 281,
		282, 5, 26, 0, 0, 282, 283, 3, 20, 10, 0, 283, 284, 5, 27, 0, 0, 284, 286,
		1, 0, 0, 0, 285, 280, 1, 0, 0, 0, 285, 281, 1, 0, 0, 0, 286, 27, 1, 0,
		0, 0, 287, 294, 5, 48, 0, 0, 288, 289, 5, 26, 0, 0, 289, 290, 3, 20, 10,
		0, 290, 291, 5, 27, 0, 0, 291, 292, 3, 30, 15, 0, 292, 294, 1, 0, 0, 0,
		293, 287, 1, 0, 0, 0, 293, 288, 1, 0, 0, 0, 294, 298, 1, 0, 0, 0, 295,
		297, 3, 30, 15, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296,
		1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 29, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 301, 303, 3, 32, 16, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0, 0,
		0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 313, 1, 0, 0, 0, 306,
		304, 1, 0, 0, 0, 307, 308, 5, 28, 0, 0, 308, 309, 3, 20, 10, 0, 309, 310,
		5, 29, 0, 0, 310, 314, 1, 0, 0, 0, 311, 312, 5, 20, 0, 0, 312, 314, 5,
		48, 0, 0, 313, 307, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 31, 1, 0, 0,
		0, 315, 316, 5, 21, 0, 0, 316, 318, 5, 48, 0, 0, 317, 315, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 3, 34, 17, 0, 320,
		33, 1, 0, 0, 0, 321, 323, 5, 26, 0, 0, 322, 324, 3, 18, 9, 0, 323, 322,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 329, 5, 27,
		0, 0, 326, 329, 3, 42, 21, 0, 327, 329, 3, 68, 34, 0, 328, 321, 1, 0, 0,
		0, 328, 326, 1, 0, 0, 0, 328, 327, 1, 0, 0, 0, 329, 35, 1, 0, 0, 0, 330,
		331, 5, 15, 0, 0, 331, 332, 3, 38, 19, 0, 332, 37, 1, 0, 0, 0, 333, 335,
		5, 26, 0, 0, 334, 336, 3, 40, 20, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1,
		0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 5, 27, 0, 0, 338, 339, 3, 2, 1,
		0, 339, 340, 5, 4, 0, 0, 340, 39, 1, 0, 0, 0, 341, 344, 3, 16, 8, 0, 342,
		343, 5, 13, 0, 0, 343, 345, 5, 25, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345,
		1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 348, 5, 25, 0, 0, 347, 341, 1, 0,
		0, 0, 347, 346, 1, 0, 0, 0, 348, 41, 1, 0, 0, 0, 349, 351, 5, 30, 0, 0,
		350, 352, 3, 44, 22, 0, 351, 350, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352,
		353, 1, 0, 0, 0, 353, 354, 5, 31, 0, 0, 354, 43, 1, 0, 0, 0, 355, 361,
		3, 46, 23, 0, 356, 357, 3, 48, 24, 0, 357, 358, 3, 46, 23, 0, 358, 360,
		1, 0, 0, 0, 359, 356, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0,
		0, 0, 361, 362, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0,
		364, 366, 3, 48, 24, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366,
		45, 1, 0, 0, 0, 367, 368, 5, 28, 0, 0, 368, 369, 3, 20, 10, 0, 369, 370,
		5, 29, 0, 0, 370, 371, 5, 2, 0, 0, 371, 372, 3, 20, 10, 0, 372, 378, 1,
		0, 0, 0, 373, 374, 5, 48, 0, 0, 374, 375, 5, 2, 0, 0, 375, 378, 3, 20,
		10, 0, 376, 378, 3, 20, 10, 0, 377, 367, 1, 0, 0, 0, 377, 373, 1, 0, 0,
		0, 377, 376, 1, 0, 0, 0, 378, 47, 1, 0, 0, 0, 379, 380, 7, 0, 0, 0, 380,
		49, 1, 0, 0, 0, 381, 382, 5, 32, 0, 0, 382, 51, 1, 0, 0, 0, 383, 384, 5,
		33, 0, 0, 384, 53, 1, 0, 0, 0, 385, 386, 7, 1, 0, 0, 386, 55, 1, 0, 0,
		0, 387, 388, 5, 38, 0, 0, 388, 57, 1, 0, 0, 0, 389, 390, 7, 2, 0, 0, 390,
		59, 1, 0, 0, 0, 391, 392, 7, 3, 0, 0, 392, 61, 1, 0, 0, 0, 393, 394, 7,
		4, 0, 0, 394, 63, 1, 0, 0, 0, 395, 396, 5, 47, 0, 0, 396, 65, 1, 0, 0,
		0, 397, 398, 7, 5, 0, 0, 398, 67, 1, 0, 0, 0, 399, 400, 5, 49, 0, 0, 400,
		69, 1, 0, 0, 0, 36, 76, 80, 114, 119, 131, 157, 159, 168, 174, 178, 181,
		188, 193, 200, 208, 216, 232, 262, 264, 271, 278, 285, 293, 298, 304, 313,
		317, 323, 328, 335, 344, 347, 351, 361, 365, 377,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LuaParserInit initializes any static state used to implement LuaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLuaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LuaParserInit() {
	staticData := &LuaParserStaticData
	staticData.once.Do(luaParserInit)
}

// NewLuaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLuaParser(input antlr.TokenStream) *LuaParser {
	LuaParserInit()
	this := new(LuaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LuaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lua.g4"

	return this
}

// LuaParser tokens.
const (
	LuaParserEOF          = antlr.TokenEOF
	LuaParserT__0         = 1
	LuaParserT__1         = 2
	LuaParserT__2         = 3
	LuaParserT__3         = 4
	LuaParserT__4         = 5
	LuaParserT__5         = 6
	LuaParserT__6         = 7
	LuaParserT__7         = 8
	LuaParserT__8         = 9
	LuaParserT__9         = 10
	LuaParserT__10        = 11
	LuaParserT__11        = 12
	LuaParserT__12        = 13
	LuaParserT__13        = 14
	LuaParserT__14        = 15
	LuaParserT__15        = 16
	LuaParserT__16        = 17
	LuaParserT__17        = 18
	LuaParserT__18        = 19
	LuaParserT__19        = 20
	LuaParserT__20        = 21
	LuaParserT__21        = 22
	LuaParserT__22        = 23
	LuaParserT__23        = 24
	LuaParserT__24        = 25
	LuaParserT__25        = 26
	LuaParserT__26        = 27
	LuaParserT__27        = 28
	LuaParserT__28        = 29
	LuaParserT__29        = 30
	LuaParserT__30        = 31
	LuaParserT__31        = 32
	LuaParserT__32        = 33
	LuaParserT__33        = 34
	LuaParserT__34        = 35
	LuaParserT__35        = 36
	LuaParserT__36        = 37
	LuaParserT__37        = 38
	LuaParserT__38        = 39
	LuaParserT__39        = 40
	LuaParserT__40        = 41
	LuaParserT__41        = 42
	LuaParserT__42        = 43
	LuaParserT__43        = 44
	LuaParserT__44        = 45
	LuaParserT__45        = 46
	LuaParserT__46        = 47
	LuaParserNAME         = 48
	LuaParserNORMALSTRING = 49
	LuaParserINT          = 50
	LuaParserFLOAT        = 51
	LuaParserCOMMENT      = 52
	LuaParserLINE_COMMENT = 53
	LuaParserWS           = 54
	LuaParserSHEBANG      = 55
)

// LuaParser rules.
const (
	LuaParserRULE_chunk              = 0
	LuaParserRULE_block              = 1
	LuaParserRULE_stat               = 2
	LuaParserRULE_attnamelist        = 3
	LuaParserRULE_attrib             = 4
	LuaParserRULE_retstat            = 5
	LuaParserRULE_funcname           = 6
	LuaParserRULE_varlist            = 7
	LuaParserRULE_namelist           = 8
	LuaParserRULE_explist            = 9
	LuaParserRULE_exp                = 10
	LuaParserRULE_prefixexp          = 11
	LuaParserRULE_functioncall       = 12
	LuaParserRULE_varOrExp           = 13
	LuaParserRULE_var                = 14
	LuaParserRULE_varSuffix          = 15
	LuaParserRULE_nameAndArgs        = 16
	LuaParserRULE_args               = 17
	LuaParserRULE_functiondef        = 18
	LuaParserRULE_funcbody           = 19
	LuaParserRULE_parlist            = 20
	LuaParserRULE_tableconstructor   = 21
	LuaParserRULE_fieldlist          = 22
	LuaParserRULE_field              = 23
	LuaParserRULE_fieldsep           = 24
	LuaParserRULE_operatorOr         = 25
	LuaParserRULE_operatorAnd        = 26
	LuaParserRULE_operatorComparison = 27
	LuaParserRULE_operatorStrcat     = 28
	LuaParserRULE_operatorAddSub     = 29
	LuaParserRULE_operatorMulDivMod  = 30
	LuaParserRULE_operatorUnary      = 31
	LuaParserRULE_operatorPower      = 32
	LuaParserRULE_number             = 33
	LuaParserRULE_string             = 34
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_chunk
	return p
}

func InitEmptyChunkContext(p *ChunkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_chunk
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(LuaParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitChunk(s)
	}
}

func (s *ChunkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitChunk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Chunk() (localctx IChunkContext) {
	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LuaParserRULE_chunk)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Block()
	}
	{
		p.SetState(71)
		p.Match(LuaParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext
	Retstat() IRetstatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) Retstat() IRetstatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetstatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LuaParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475043922282) != 0 {
		{
			p.SetState(73)
			p.Stat()
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__18 {
		{
			p.SetState(79)
			p.Retstat()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatEmptySemicolonContext struct {
	StatContext
}

func NewStatEmptySemicolonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatEmptySemicolonContext {
	var p = new(StatEmptySemicolonContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatEmptySemicolonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatEmptySemicolonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatEmptySemicolon(s)
	}
}

func (s *StatEmptySemicolonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatEmptySemicolon(s)
	}
}

func (s *StatEmptySemicolonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatEmptySemicolon(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatNumericForContext struct {
	StatContext
}

func NewStatNumericForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatNumericForContext {
	var p = new(StatNumericForContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatNumericForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatNumericForContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *StatNumericForContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *StatNumericForContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatNumericForContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatNumericForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatNumericFor(s)
	}
}

func (s *StatNumericForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatNumericFor(s)
	}
}

func (s *StatNumericForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatNumericFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatAssignmentContext struct {
	StatContext
}

func NewStatAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatAssignmentContext {
	var p = new(StatAssignmentContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatAssignmentContext) Varlist() IVarlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *StatAssignmentContext) Explist() IExplistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExplistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatAssignment(s)
	}
}

func (s *StatAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatAssignment(s)
	}
}

func (s *StatAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatFunctionContext struct {
	StatContext
}

func NewStatFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatFunctionContext {
	var p = new(StatFunctionContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFunctionContext) Funcname() IFuncnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *StatFunctionContext) Funcbody() IFuncbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *StatFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatFunction(s)
	}
}

func (s *StatFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatFunction(s)
	}
}

func (s *StatFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatFunctionCallContext struct {
	StatContext
}

func NewStatFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatFunctionCallContext {
	var p = new(StatFunctionCallContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFunctionCallContext) Functioncall() IFunctioncallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctioncallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *StatFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatFunctionCall(s)
	}
}

func (s *StatFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatFunctionCall(s)
	}
}

func (s *StatFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatIfThenElseContext struct {
	StatContext
}

func NewStatIfThenElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatIfThenElseContext {
	var p = new(StatIfThenElseContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatIfThenElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatIfThenElseContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *StatIfThenElseContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatIfThenElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *StatIfThenElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatIfThenElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatIfThenElse(s)
	}
}

func (s *StatIfThenElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatIfThenElse(s)
	}
}

func (s *StatIfThenElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatIfThenElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatLocalFunctionContext struct {
	StatContext
}

func NewStatLocalFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatLocalFunctionContext {
	var p = new(StatLocalFunctionContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatLocalFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatLocalFunctionContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *StatLocalFunctionContext) Funcbody() IFuncbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *StatLocalFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatLocalFunction(s)
	}
}

func (s *StatLocalFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatLocalFunction(s)
	}
}

func (s *StatLocalFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatLocalFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatRepeatContext struct {
	StatContext
}

func NewStatRepeatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatRepeatContext {
	var p = new(StatRepeatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatRepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatRepeatContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatRepeatContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatRepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatRepeat(s)
	}
}

func (s *StatRepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatRepeat(s)
	}
}

func (s *StatRepeatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatRepeat(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatWhileContext struct {
	StatContext
}

func NewStatWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatWhileContext {
	var p = new(StatWhileContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatWhileContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatWhileContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatWhile(s)
	}
}

func (s *StatWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatWhile(s)
	}
}

func (s *StatWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatGenericForContext struct {
	StatContext
}

func NewStatGenericForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatGenericForContext {
	var p = new(StatGenericForContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatGenericForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatGenericForContext) Namelist() INamelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *StatGenericForContext) Explist() IExplistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExplistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatGenericForContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatGenericForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatGenericFor(s)
	}
}

func (s *StatGenericForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatGenericFor(s)
	}
}

func (s *StatGenericForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatGenericFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatDoContext struct {
	StatContext
}

func NewStatDoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatDoContext {
	var p = new(StatDoContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatDoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatDoContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatDoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatDo(s)
	}
}

func (s *StatDoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatDo(s)
	}
}

func (s *StatDoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatDo(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatLocalAttributeNameListContext struct {
	StatContext
}

func NewStatLocalAttributeNameListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatLocalAttributeNameListContext {
	var p = new(StatLocalAttributeNameListContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatLocalAttributeNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatLocalAttributeNameListContext) Attnamelist() IAttnamelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttnamelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttnamelistContext)
}

func (s *StatLocalAttributeNameListContext) Explist() IExplistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExplistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatLocalAttributeNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatLocalAttributeNameList(s)
	}
}

func (s *StatLocalAttributeNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatLocalAttributeNameList(s)
	}
}

func (s *StatLocalAttributeNameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatLocalAttributeNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LuaParserRULE_stat)
	var _la int

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatEmptySemicolonContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(LuaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStatAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Varlist()
		}
		{
			p.SetState(84)
			p.Match(LuaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Explist()
		}

	case 3:
		localctx = NewStatFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Functioncall()
		}

	case 4:
		localctx = NewStatDoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.Match(LuaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Block()
		}
		{
			p.SetState(90)
			p.Match(LuaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStatWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.Match(LuaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.exp(0)
		}
		{
			p.SetState(94)
			p.Match(LuaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Block()
		}
		{
			p.SetState(96)
			p.Match(LuaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStatRepeatContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.Match(LuaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Block()
		}
		{
			p.SetState(100)
			p.Match(LuaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.exp(0)
		}

	case 7:
		localctx = NewStatIfThenElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(103)
			p.Match(LuaParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.exp(0)
		}
		{
			p.SetState(105)
			p.Match(LuaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Block()
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__9 {
			{
				p.SetState(107)
				p.Match(LuaParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(108)
				p.exp(0)
			}
			{
				p.SetState(109)
				p.Match(LuaParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(110)
				p.Block()
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__10 {
			{
				p.SetState(117)
				p.Match(LuaParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(118)
				p.Block()
			}

		}
		{
			p.SetState(121)
			p.Match(LuaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewStatNumericForContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(123)
			p.Match(LuaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(LuaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.exp(0)
		}
		{
			p.SetState(127)
			p.Match(LuaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.exp(0)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__12 {
			{
				p.SetState(129)
				p.Match(LuaParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(130)
				p.exp(0)
			}

		}
		{
			p.SetState(133)
			p.Match(LuaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Block()
		}
		{
			p.SetState(135)
			p.Match(LuaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewStatGenericForContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(137)
			p.Match(LuaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Namelist()
		}
		{
			p.SetState(139)
			p.Match(LuaParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Explist()
		}
		{
			p.SetState(141)
			p.Match(LuaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Block()
		}
		{
			p.SetState(143)
			p.Match(LuaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewStatFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(145)
			p.Match(LuaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Funcname()
		}
		{
			p.SetState(147)
			p.Funcbody()
		}

	case 11:
		localctx = NewStatLocalFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(149)
			p.Match(LuaParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(LuaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Funcbody()
		}

	case 12:
		localctx = NewStatLocalAttributeNameListContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(153)
			p.Match(LuaParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Attnamelist()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__1 {
			{
				p.SetState(155)
				p.Match(LuaParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(156)
				p.Explist()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttnamelistContext is an interface to support dynamic dispatch.
type IAttnamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode
	AllAttrib() []IAttribContext
	Attrib(i int) IAttribContext

	// IsAttnamelistContext differentiates from other interfaces.
	IsAttnamelistContext()
}

type AttnamelistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttnamelistContext() *AttnamelistContext {
	var p = new(AttnamelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_attnamelist
	return p
}

func InitEmptyAttnamelistContext(p *AttnamelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_attnamelist
}

func (*AttnamelistContext) IsAttnamelistContext() {}

func NewAttnamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttnamelistContext {
	var p = new(AttnamelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_attnamelist

	return p
}

func (s *AttnamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *AttnamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *AttnamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *AttnamelistContext) AllAttrib() []IAttribContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttribContext); ok {
			len++
		}
	}

	tst := make([]IAttribContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttribContext); ok {
			tst[i] = t.(IAttribContext)
			i++
		}
	}

	return tst
}

func (s *AttnamelistContext) Attrib(i int) IAttribContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *AttnamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttnamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttnamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterAttnamelist(s)
	}
}

func (s *AttnamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitAttnamelist(s)
	}
}

func (s *AttnamelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitAttnamelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Attnamelist() (localctx IAttnamelistContext) {
	localctx = NewAttnamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LuaParserRULE_attnamelist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(LuaParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Attrib()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__12 {
		{
			p.SetState(163)
			p.Match(LuaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Attrib()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_attrib
	return p
}

func InitEmptyAttribContext(p *AttribContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_attrib
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (s *AttribContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitAttrib(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Attrib() (localctx IAttribContext) {
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LuaParserRULE_attrib)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__16 {
		{
			p.SetState(171)
			p.Match(LuaParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(LuaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetstatContext is an interface to support dynamic dispatch.
type IRetstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRetstatContext differentiates from other interfaces.
	IsRetstatContext()
}

type RetstatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatContext() *RetstatContext {
	var p = new(RetstatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_retstat
	return p
}

func InitEmptyRetstatContext(p *RetstatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_retstat
}

func (*RetstatContext) IsRetstatContext() {}

func NewRetstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatContext {
	var p = new(RetstatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_retstat

	return p
}

func (s *RetstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatContext) CopyAll(ctx *RetstatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RetstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatReturnContext struct {
	RetstatContext
}

func NewStatReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatReturnContext {
	var p = new(StatReturnContext)

	InitEmptyRetstatContext(&p.RetstatContext)
	p.parser = parser
	p.CopyAll(ctx.(*RetstatContext))

	return p
}

func (s *StatReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatReturnContext) Explist() IExplistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExplistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStatReturn(s)
	}
}

func (s *StatReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStatReturn(s)
	}
}

func (s *StatReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Retstat() (localctx IRetstatContext) {
	localctx = NewRetstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LuaParserRULE_retstat)
	var _la int

	localctx = NewStatReturnContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(LuaParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4328778482352128) != 0 {
		{
			p.SetState(177)
			p.Explist()
		}

	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 {
		{
			p.SetState(180)
			p.Match(LuaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_funcname
	return p
}

func InitEmptyFuncnameContext(p *FuncnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_funcname
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *FuncnameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (s *FuncnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFuncname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LuaParserRULE_funcname)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(LuaParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__19 {
		{
			p.SetState(184)
			p.Match(LuaParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__20 {
		{
			p.SetState(191)
			p.Match(LuaParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_() []IVarContext
	Var_(i int) IVarContext

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varlist
	return p
}

func InitEmptyVarlistContext(p *VarlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varlist
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVar_() []IVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarContext); ok {
			len++
		}
	}

	tst := make([]IVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarContext); ok {
			tst[i] = t.(IVarContext)
			i++
		}
	}

	return tst
}

func (s *VarlistContext) Var_(i int) IVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarlist(s)
	}
}

func (s *VarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarlist(s)
	}
}

func (s *VarlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Varlist() (localctx IVarlistContext) {
	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LuaParserRULE_varlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Var_()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__12 {
		{
			p.SetState(196)
			p.Match(LuaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Var_()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_namelist
	return p
}

func InitEmptyNamelistContext(p *NamelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_namelist
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *NamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNamelist(s)
	}
}

func (s *NamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNamelist(s)
	}
}

func (s *NamelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNamelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LuaParserRULE_namelist)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(LuaParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(204)
				p.Match(LuaParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(205)
				p.Match(LuaParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExplistContext is an interface to support dynamic dispatch.
type IExplistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsExplistContext differentiates from other interfaces.
	IsExplistContext()
}

type ExplistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExplistContext() *ExplistContext {
	var p = new(ExplistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_explist
	return p
}

func InitEmptyExplistContext(p *ExplistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_explist
}

func (*ExplistContext) IsExplistContext() {}

func NewExplistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExplistContext {
	var p = new(ExplistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_explist

	return p
}

func (s *ExplistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExplistContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExplistContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExplistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExplistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExplist(s)
	}
}

func (s *ExplistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExplist(s)
	}
}

func (s *ExplistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExplist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Explist() (localctx IExplistContext) {
	localctx = NewExplistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuaParserRULE_explist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.exp(0)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__12 {
		{
			p.SetState(212)
			p.Match(LuaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.exp(0)
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyAll(ctx *ExpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpFalseContext struct {
	ExpContext
}

func NewExpFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFalseContext {
	var p = new(ExpFalseContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpFalse(s)
	}
}

func (s *ExpFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpFalse(s)
	}
}

func (s *ExpFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpVarargContext struct {
	ExpContext
}

func NewExpVarargContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpVarargContext {
	var p = new(ExpVarargContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpVarargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpVarargContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpVararg(s)
	}
}

func (s *ExpVarargContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpVararg(s)
	}
}

func (s *ExpVarargContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpVararg(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpTableConstructorContext struct {
	ExpContext
}

func NewExpTableConstructorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpTableConstructorContext {
	var p = new(ExpTableConstructorContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpTableConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpTableConstructorContext) Tableconstructor() ITableconstructorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableconstructorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ExpTableConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpTableConstructor(s)
	}
}

func (s *ExpTableConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpTableConstructor(s)
	}
}

func (s *ExpTableConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpTableConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpPrefixExpContext struct {
	ExpContext
}

func NewExpPrefixExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpPrefixExpContext {
	var p = new(ExpPrefixExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpPrefixExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpPrefixExpContext) Prefixexp() IPrefixexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixexpContext)
}

func (s *ExpPrefixExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpPrefixExp(s)
	}
}

func (s *ExpPrefixExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpPrefixExp(s)
	}
}

func (s *ExpPrefixExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpPrefixExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpTrueContext struct {
	ExpContext
}

func NewExpTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpTrueContext {
	var p = new(ExpTrueContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpTrue(s)
	}
}

func (s *ExpTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpTrue(s)
	}
}

func (s *ExpTrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpNumberContext struct {
	ExpContext
}

func NewExpNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNumberContext {
	var p = new(ExpNumberContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNumberContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ExpNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpNumber(s)
	}
}

func (s *ExpNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpNumber(s)
	}
}

func (s *ExpNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorUnaryContext struct {
	ExpContext
}

func NewExpOperatorUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorUnaryContext {
	var p = new(ExpOperatorUnaryContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorUnaryContext) OperatorUnary() IOperatorUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorUnaryContext)
}

func (s *ExpOperatorUnaryContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorUnary(s)
	}
}

func (s *ExpOperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorUnary(s)
	}
}

func (s *ExpOperatorUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorAndContext struct {
	ExpContext
}

func NewExpOperatorAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorAndContext {
	var p = new(ExpOperatorAndContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorAndContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorAndContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorAndContext) OperatorAnd() IOperatorAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorAndContext)
}

func (s *ExpOperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorAnd(s)
	}
}

func (s *ExpOperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorAnd(s)
	}
}

func (s *ExpOperatorAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorPowerContext struct {
	ExpContext
}

func NewExpOperatorPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorPowerContext {
	var p = new(ExpOperatorPowerContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorPowerContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorPowerContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorPowerContext) OperatorPower() IOperatorPowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorPowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorPowerContext)
}

func (s *ExpOperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorPower(s)
	}
}

func (s *ExpOperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorPower(s)
	}
}

func (s *ExpOperatorPowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorPower(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorAddSubContext struct {
	ExpContext
}

func NewExpOperatorAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorAddSubContext {
	var p = new(ExpOperatorAddSubContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorAddSubContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorAddSubContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorAddSubContext) OperatorAddSub() IOperatorAddSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorAddSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorAddSubContext)
}

func (s *ExpOperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorAddSub(s)
	}
}

func (s *ExpOperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorAddSub(s)
	}
}

func (s *ExpOperatorAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorStrcatContext struct {
	ExpContext
}

func NewExpOperatorStrcatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorStrcatContext {
	var p = new(ExpOperatorStrcatContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorStrcatContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorStrcatContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorStrcatContext) OperatorStrcat() IOperatorStrcatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStrcatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorStrcatContext)
}

func (s *ExpOperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorStrcat(s)
	}
}

func (s *ExpOperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorStrcat(s)
	}
}

func (s *ExpOperatorStrcatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorStrcat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorComparisonContext struct {
	ExpContext
}

func NewExpOperatorComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorComparisonContext {
	var p = new(ExpOperatorComparisonContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorComparisonContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorComparisonContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorComparisonContext) OperatorComparison() IOperatorComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorComparisonContext)
}

func (s *ExpOperatorComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorComparison(s)
	}
}

func (s *ExpOperatorComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorComparison(s)
	}
}

func (s *ExpOperatorComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpNilContext struct {
	ExpContext
}

func NewExpNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNilContext {
	var p = new(ExpNilContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpNil(s)
	}
}

func (s *ExpNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpNil(s)
	}
}

func (s *ExpNilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorOrContext struct {
	ExpContext
}

func NewExpOperatorOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorOrContext {
	var p = new(ExpOperatorOrContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorOrContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorOrContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorOrContext) OperatorOr() IOperatorOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorOrContext)
}

func (s *ExpOperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorOr(s)
	}
}

func (s *ExpOperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorOr(s)
	}
}

func (s *ExpOperatorOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpStringContext struct {
	ExpContext
}

func NewExpStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpStringContext {
	var p = new(ExpStringContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpStringContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ExpStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpString(s)
	}
}

func (s *ExpStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpString(s)
	}
}

func (s *ExpStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpOperatorMulDivModContext struct {
	ExpContext
}

func NewExpOperatorMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpOperatorMulDivModContext {
	var p = new(ExpOperatorMulDivModContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpOperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOperatorMulDivModContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpOperatorMulDivModContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpOperatorMulDivModContext) OperatorMulDivMod() IOperatorMulDivModContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorMulDivModContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorMulDivModContext)
}

func (s *ExpOperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpOperatorMulDivMod(s)
	}
}

func (s *ExpOperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpOperatorMulDivMod(s)
	}
}

func (s *ExpOperatorMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpOperatorMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpFunctionDefContext struct {
	ExpContext
}

func NewExpFunctionDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFunctionDefContext {
	var p = new(ExpFunctionDefContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpFunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFunctionDefContext) Functiondef() IFunctiondefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctiondefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *ExpFunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExpFunctionDef(s)
	}
}

func (s *ExpFunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExpFunctionDef(s)
	}
}

func (s *ExpFunctionDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpFunctionDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *LuaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, LuaParserRULE_exp, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__21:
		localctx = NewExpNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(220)
			p.Match(LuaParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__22:
		localctx = NewExpFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(221)
			p.Match(LuaParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__23:
		localctx = NewExpTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(222)
			p.Match(LuaParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserINT, LuaParserFLOAT:
		localctx = NewExpNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Number()
		}

	case LuaParserNORMALSTRING:
		localctx = NewExpStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(224)
			p.String_()
		}

	case LuaParserT__24:
		localctx = NewExpVarargContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(225)
			p.Match(LuaParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__14:
		localctx = NewExpFunctionDefContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(226)
			p.Functiondef()
		}

	case LuaParserT__25, LuaParserNAME:
		localctx = NewExpPrefixExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(227)
			p.Prefixexp()
		}

	case LuaParserT__29:
		localctx = NewExpTableConstructorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Tableconstructor()
		}

	case LuaParserT__39, LuaParserT__44, LuaParserT__45:
		localctx = NewExpOperatorUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(229)
			p.OperatorUnary()
		}
		{
			p.SetState(230)
			p.exp(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpOperatorPowerContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(235)
					p.OperatorPower()
				}
				{
					p.SetState(236)
					p.exp(8)
				}

			case 2:
				localctx = NewExpOperatorMulDivModContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(239)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(240)
					p.exp(7)
				}

			case 3:
				localctx = NewExpOperatorAddSubContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(243)
					p.OperatorAddSub()
				}
				{
					p.SetState(244)
					p.exp(6)
				}

			case 4:
				localctx = NewExpOperatorStrcatContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(247)
					p.OperatorStrcat()
				}
				{
					p.SetState(248)
					p.exp(4)
				}

			case 5:
				localctx = NewExpOperatorComparisonContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(251)
					p.OperatorComparison()
				}
				{
					p.SetState(252)
					p.exp(4)
				}

			case 6:
				localctx = NewExpOperatorAndContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(255)
					p.OperatorAnd()
				}
				{
					p.SetState(256)
					p.exp(3)
				}

			case 7:
				localctx = NewExpOperatorOrContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(259)
					p.OperatorOr()
				}
				{
					p.SetState(260)
					p.exp(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixexpContext is an interface to support dynamic dispatch.
type IPrefixexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarOrExp() IVarOrExpContext
	AllNameAndArgs() []INameAndArgsContext
	NameAndArgs(i int) INameAndArgsContext

	// IsPrefixexpContext differentiates from other interfaces.
	IsPrefixexpContext()
}

type PrefixexpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixexpContext() *PrefixexpContext {
	var p = new(PrefixexpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_prefixexp
	return p
}

func InitEmptyPrefixexpContext(p *PrefixexpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_prefixexp
}

func (*PrefixexpContext) IsPrefixexpContext() {}

func NewPrefixexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixexpContext {
	var p = new(PrefixexpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_prefixexp

	return p
}

func (s *PrefixexpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixexpContext) VarOrExp() IVarOrExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *PrefixexpContext) AllNameAndArgs() []INameAndArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameAndArgsContext); ok {
			len++
		}
	}

	tst := make([]INameAndArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameAndArgsContext); ok {
			tst[i] = t.(INameAndArgsContext)
			i++
		}
	}

	return tst
}

func (s *PrefixexpContext) NameAndArgs(i int) INameAndArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameAndArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *PrefixexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterPrefixexp(s)
	}
}

func (s *PrefixexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitPrefixexp(s)
	}
}

func (s *PrefixexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitPrefixexp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Prefixexp() (localctx IPrefixexpContext) {
	localctx = NewPrefixexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LuaParserRULE_prefixexp)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.VarOrExp()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(268)
				p.NameAndArgs()
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarOrExp() IVarOrExpContext
	AllNameAndArgs() []INameAndArgsContext
	NameAndArgs(i int) INameAndArgsContext

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_functioncall
	return p
}

func InitEmptyFunctioncallContext(p *FunctioncallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_functioncall
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) VarOrExp() IVarOrExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *FunctioncallContext) AllNameAndArgs() []INameAndArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameAndArgsContext); ok {
			len++
		}
	}

	tst := make([]INameAndArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameAndArgsContext); ok {
			tst[i] = t.(INameAndArgsContext)
			i++
		}
	}

	return tst
}

func (s *FunctioncallContext) NameAndArgs(i int) INameAndArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameAndArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (s *FunctioncallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFunctioncall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Functioncall() (localctx IFunctioncallContext) {
	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LuaParserRULE_functioncall)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.VarOrExp()
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(275)
				p.NameAndArgs()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarOrExpContext is an interface to support dynamic dispatch.
type IVarOrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	Exp() IExpContext

	// IsVarOrExpContext differentiates from other interfaces.
	IsVarOrExpContext()
}

type VarOrExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrExpContext() *VarOrExpContext {
	var p = new(VarOrExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varOrExp
	return p
}

func InitEmptyVarOrExpContext(p *VarOrExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varOrExp
}

func (*VarOrExpContext) IsVarOrExpContext() {}

func NewVarOrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrExpContext {
	var p = new(VarOrExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varOrExp

	return p
}

func (s *VarOrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrExpContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *VarOrExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarOrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarOrExp(s)
	}
}

func (s *VarOrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarOrExp(s)
	}
}

func (s *VarOrExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarOrExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VarOrExp() (localctx IVarOrExpContext) {
	localctx = NewVarOrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LuaParserRULE_varOrExp)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Var_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(LuaParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.exp(0)
		}
		{
			p.SetState(283)
			p.Match(LuaParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Exp() IExpContext
	AllVarSuffix() []IVarSuffixContext
	VarSuffix(i int) IVarSuffixContext

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *VarContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarContext) AllVarSuffix() []IVarSuffixContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarSuffixContext); ok {
			len++
		}
	}

	tst := make([]IVarSuffixContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarSuffixContext); ok {
			tst[i] = t.(IVarSuffixContext)
			i++
		}
	}

	return tst
}

func (s *VarContext) VarSuffix(i int) IVarSuffixContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSuffixContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarSuffixContext)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LuaParserRULE_var)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		{
			p.SetState(287)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__25:
		{
			p.SetState(288)
			p.Match(LuaParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.exp(0)
		}
		{
			p.SetState(290)
			p.Match(LuaParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.VarSuffix()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(295)
				p.VarSuffix()
			}

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarSuffixContext is an interface to support dynamic dispatch.
type IVarSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	NAME() antlr.TerminalNode
	AllNameAndArgs() []INameAndArgsContext
	NameAndArgs(i int) INameAndArgsContext

	// IsVarSuffixContext differentiates from other interfaces.
	IsVarSuffixContext()
}

type VarSuffixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSuffixContext() *VarSuffixContext {
	var p = new(VarSuffixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varSuffix
	return p
}

func InitEmptyVarSuffixContext(p *VarSuffixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_varSuffix
}

func (*VarSuffixContext) IsVarSuffixContext() {}

func NewVarSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSuffixContext {
	var p = new(VarSuffixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varSuffix

	return p
}

func (s *VarSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSuffixContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarSuffixContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *VarSuffixContext) AllNameAndArgs() []INameAndArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameAndArgsContext); ok {
			len++
		}
	}

	tst := make([]INameAndArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameAndArgsContext); ok {
			tst[i] = t.(INameAndArgsContext)
			i++
		}
	}

	return tst
}

func (s *VarSuffixContext) NameAndArgs(i int) INameAndArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameAndArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *VarSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarSuffix(s)
	}
}

func (s *VarSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarSuffix(s)
	}
}

func (s *VarSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VarSuffix() (localctx IVarSuffixContext) {
	localctx = NewVarSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LuaParserRULE_varSuffix)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562951096369152) != 0 {
		{
			p.SetState(301)
			p.NameAndArgs()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__27:
		{
			p.SetState(307)
			p.Match(LuaParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.exp(0)
		}
		{
			p.SetState(309)
			p.Match(LuaParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__19:
		{
			p.SetState(311)
			p.Match(LuaParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameAndArgsContext is an interface to support dynamic dispatch.
type INameAndArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Args() IArgsContext
	NAME() antlr.TerminalNode

	// IsNameAndArgsContext differentiates from other interfaces.
	IsNameAndArgsContext()
}

type NameAndArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAndArgsContext() *NameAndArgsContext {
	var p = new(NameAndArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_nameAndArgs
	return p
}

func InitEmptyNameAndArgsContext(p *NameAndArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_nameAndArgs
}

func (*NameAndArgsContext) IsNameAndArgsContext() {}

func NewNameAndArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAndArgsContext {
	var p = new(NameAndArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_nameAndArgs

	return p
}

func (s *NameAndArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAndArgsContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NameAndArgsContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *NameAndArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAndArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAndArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNameAndArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) NameAndArgs() (localctx INameAndArgsContext) {
	localctx = NewNameAndArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LuaParserRULE_nameAndArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__20 {
		{
			p.SetState(315)
			p.Match(LuaParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(319)
		p.Args()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Explist() IExplistContext
	Tableconstructor() ITableconstructorContext
	String_() IStringContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Explist() IExplistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExplistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ArgsContext) Tableconstructor() ITableconstructorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableconstructorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ArgsContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LuaParserRULE_args)
	var _la int

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(LuaParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4328778482352128) != 0 {
			{
				p.SetState(322)
				p.Explist()
			}

		}
		{
			p.SetState(325)
			p.Match(LuaParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LuaParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.String_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcbody() IFuncbodyContext

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_functiondef
	return p
}

func InitEmptyFunctiondefContext(p *FunctiondefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_functiondef
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) Funcbody() IFuncbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFunctiondef(s)
	}
}

func (s *FunctiondefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFunctiondef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Functiondef() (localctx IFunctiondefContext) {
	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LuaParserRULE_functiondef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(LuaParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Funcbody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	Parlist() IParlistContext

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_funcbody
	return p
}

func InitEmptyFuncbodyContext(p *FuncbodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_funcbody
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncbodyContext) Parlist() IParlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (s *FuncbodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFuncbody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LuaParserRULE_funcbody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(LuaParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__24 || _la == LuaParserNAME {
		{
			p.SetState(334)
			p.Parlist()
		}

	}
	{
		p.SetState(337)
		p.Match(LuaParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Block()
	}
	{
		p.SetState(339)
		p.Match(LuaParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParlistContext is an interface to support dynamic dispatch.
type IParlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Namelist() INamelistContext

	// IsParlistContext differentiates from other interfaces.
	IsParlistContext()
}

type ParlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParlistContext() *ParlistContext {
	var p = new(ParlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_parlist
	return p
}

func InitEmptyParlistContext(p *ParlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_parlist
}

func (*ParlistContext) IsParlistContext() {}

func NewParlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParlistContext {
	var p = new(ParlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_parlist

	return p
}

func (s *ParlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParlistContext) Namelist() INamelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ParlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterParlist(s)
	}
}

func (s *ParlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitParlist(s)
	}
}

func (s *ParlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitParlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Parlist() (localctx IParlistContext) {
	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LuaParserRULE_parlist)
	var _la int

	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.Namelist()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__12 {
			{
				p.SetState(342)
				p.Match(LuaParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(343)
				p.Match(LuaParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case LuaParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(LuaParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableconstructorContext is an interface to support dynamic dispatch.
type ITableconstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fieldlist() IFieldlistContext

	// IsTableconstructorContext differentiates from other interfaces.
	IsTableconstructorContext()
}

type TableconstructorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableconstructorContext() *TableconstructorContext {
	var p = new(TableconstructorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_tableconstructor
	return p
}

func InitEmptyTableconstructorContext(p *TableconstructorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_tableconstructor
}

func (*TableconstructorContext) IsTableconstructorContext() {}

func NewTableconstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableconstructorContext {
	var p = new(TableconstructorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_tableconstructor

	return p
}

func (s *TableconstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableconstructorContext) Fieldlist() IFieldlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldlistContext)
}

func (s *TableconstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableconstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableconstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterTableconstructor(s)
	}
}

func (s *TableconstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitTableconstructor(s)
	}
}

func (s *TableconstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTableconstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Tableconstructor() (localctx ITableconstructorContext) {
	localctx = NewTableconstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LuaParserRULE_tableconstructor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(LuaParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4328778750787584) != 0 {
		{
			p.SetState(350)
			p.Fieldlist()
		}

	}
	{
		p.SetState(353)
		p.Match(LuaParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldlistContext is an interface to support dynamic dispatch.
type IFieldlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllFieldsep() []IFieldsepContext
	Fieldsep(i int) IFieldsepContext

	// IsFieldlistContext differentiates from other interfaces.
	IsFieldlistContext()
}

type FieldlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldlistContext() *FieldlistContext {
	var p = new(FieldlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_fieldlist
	return p
}

func InitEmptyFieldlistContext(p *FieldlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_fieldlist
}

func (*FieldlistContext) IsFieldlistContext() {}

func NewFieldlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldlistContext {
	var p = new(FieldlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldlist

	return p
}

func (s *FieldlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldlistContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldlistContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldlistContext) AllFieldsep() []IFieldsepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsepContext); ok {
			len++
		}
	}

	tst := make([]IFieldsepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsepContext); ok {
			tst[i] = t.(IFieldsepContext)
			i++
		}
	}

	return tst
}

func (s *FieldlistContext) Fieldsep(i int) IFieldsepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsepContext)
}

func (s *FieldlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFieldlist(s)
	}
}

func (s *FieldlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFieldlist(s)
	}
}

func (s *FieldlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFieldlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Fieldlist() (localctx IFieldlistContext) {
	localctx = NewFieldlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LuaParserRULE_fieldlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Field()
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(356)
				p.Fieldsep()
			}
			{
				p.SetState(357)
				p.Field()
			}

		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 || _la == LuaParserT__12 {
		{
			p.SetState(364)
			p.Fieldsep()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	NAME() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LuaParserRULE_field)
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(LuaParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.exp(0)
		}
		{
			p.SetState(369)
			p.Match(LuaParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(LuaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(LuaParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(LuaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsepContext is an interface to support dynamic dispatch.
type IFieldsepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFieldsepContext differentiates from other interfaces.
	IsFieldsepContext()
}

type FieldsepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsepContext() *FieldsepContext {
	var p = new(FieldsepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_fieldsep
	return p
}

func InitEmptyFieldsepContext(p *FieldsepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_fieldsep
}

func (*FieldsepContext) IsFieldsepContext() {}

func NewFieldsepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsepContext {
	var p = new(FieldsepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldsep

	return p
}

func (s *FieldsepContext) GetParser() antlr.Parser { return s.parser }
func (s *FieldsepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFieldsep(s)
	}
}

func (s *FieldsepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFieldsep(s)
	}
}

func (s *FieldsepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFieldsep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Fieldsep() (localctx IFieldsepContext) {
	localctx = NewFieldsepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LuaParserRULE_fieldsep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__0 || _la == LuaParserT__12) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorOrContext is an interface to support dynamic dispatch.
type IOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorOrContext differentiates from other interfaces.
	IsOperatorOrContext()
}

type OperatorOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorOrContext() *OperatorOrContext {
	var p = new(OperatorOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorOr
	return p
}

func InitEmptyOperatorOrContext(p *OperatorOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorOr
}

func (*OperatorOrContext) IsOperatorOrContext() {}

func NewOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorOrContext {
	var p = new(OperatorOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorOr

	return p
}

func (s *OperatorOrContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorOr(s)
	}
}

func (s *OperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorOr(s)
	}
}

func (s *OperatorOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorOr() (localctx IOperatorOrContext) {
	localctx = NewOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LuaParserRULE_operatorOr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(LuaParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorAndContext is an interface to support dynamic dispatch.
type IOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorAndContext differentiates from other interfaces.
	IsOperatorAndContext()
}

type OperatorAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAndContext() *OperatorAndContext {
	var p = new(OperatorAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAnd
	return p
}

func InitEmptyOperatorAndContext(p *OperatorAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAnd
}

func (*OperatorAndContext) IsOperatorAndContext() {}

func NewOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAndContext {
	var p = new(OperatorAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAnd

	return p
}

func (s *OperatorAndContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorAnd(s)
	}
}

func (s *OperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorAnd(s)
	}
}

func (s *OperatorAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorAnd() (localctx IOperatorAndContext) {
	localctx = NewOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LuaParserRULE_operatorAnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(LuaParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorComparisonContext is an interface to support dynamic dispatch.
type IOperatorComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorComparisonContext differentiates from other interfaces.
	IsOperatorComparisonContext()
}

type OperatorComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorComparisonContext() *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorComparison
	return p
}

func InitEmptyOperatorComparisonContext(p *OperatorComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorComparison
}

func (*OperatorComparisonContext) IsOperatorComparisonContext() {}

func NewOperatorComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorComparison

	return p
}

func (s *OperatorComparisonContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorComparison(s)
	}
}

func (s *OperatorComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorComparison(s)
	}
}

func (s *OperatorComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorComparison() (localctx IOperatorComparisonContext) {
	localctx = NewOperatorComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LuaParserRULE_operatorComparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&257698430976) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorStrcatContext is an interface to support dynamic dispatch.
type IOperatorStrcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorStrcatContext differentiates from other interfaces.
	IsOperatorStrcatContext()
}

type OperatorStrcatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorStrcatContext() *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorStrcat
	return p
}

func InitEmptyOperatorStrcatContext(p *OperatorStrcatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorStrcat
}

func (*OperatorStrcatContext) IsOperatorStrcatContext() {}

func NewOperatorStrcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorStrcat

	return p
}

func (s *OperatorStrcatContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStrcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorStrcat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorStrcat() (localctx IOperatorStrcatContext) {
	localctx = NewOperatorStrcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LuaParserRULE_operatorStrcat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(LuaParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorAddSubContext is an interface to support dynamic dispatch.
type IOperatorAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorAddSubContext differentiates from other interfaces.
	IsOperatorAddSubContext()
}

type OperatorAddSubContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAddSubContext() *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAddSub
	return p
}

func InitEmptyOperatorAddSubContext(p *OperatorAddSubContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAddSub
}

func (*OperatorAddSubContext) IsOperatorAddSubContext() {}

func NewOperatorAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAddSub

	return p
}

func (s *OperatorAddSubContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorAddSub() (localctx IOperatorAddSubContext) {
	localctx = NewOperatorAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LuaParserRULE_operatorAddSub)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__38 || _la == LuaParserT__39) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorMulDivModContext is an interface to support dynamic dispatch.
type IOperatorMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorMulDivModContext differentiates from other interfaces.
	IsOperatorMulDivModContext()
}

type OperatorMulDivModContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorMulDivModContext() *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorMulDivMod
	return p
}

func InitEmptyOperatorMulDivModContext(p *OperatorMulDivModContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorMulDivMod
}

func (*OperatorMulDivModContext) IsOperatorMulDivModContext() {}

func NewOperatorMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorMulDivMod

	return p
}

func (s *OperatorMulDivModContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorMulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorMulDivMod() (localctx IOperatorMulDivModContext) {
	localctx = NewOperatorMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LuaParserRULE_operatorMulDivMod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorUnaryContext is an interface to support dynamic dispatch.
type IOperatorUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorUnaryContext differentiates from other interfaces.
	IsOperatorUnaryContext()
}

type OperatorUnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorUnaryContext() *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorUnary
	return p
}

func InitEmptyOperatorUnaryContext(p *OperatorUnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorUnary
}

func (*OperatorUnaryContext) IsOperatorUnaryContext() {}

func NewOperatorUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorUnary

	return p
}

func (s *OperatorUnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorUnary() (localctx IOperatorUnaryContext) {
	localctx = NewOperatorUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LuaParserRULE_operatorUnary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&106652627894272) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorPowerContext is an interface to support dynamic dispatch.
type IOperatorPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorPowerContext differentiates from other interfaces.
	IsOperatorPowerContext()
}

type OperatorPowerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPowerContext() *OperatorPowerContext {
	var p = new(OperatorPowerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorPower
	return p
}

func InitEmptyOperatorPowerContext(p *OperatorPowerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_operatorPower
}

func (*OperatorPowerContext) IsOperatorPowerContext() {}

func NewOperatorPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPowerContext {
	var p = new(OperatorPowerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorPower

	return p
}

func (s *OperatorPowerContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorPower(s)
	}
}

func (s *OperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorPower(s)
	}
}

func (s *OperatorPowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorPower() (localctx IOperatorPowerContext) {
	localctx = NewOperatorPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LuaParserRULE_operatorPower)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(LuaParserT__46)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(LuaParserINT, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LuaParserFLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LuaParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserINT || _la == LuaParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NORMALSTRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LuaParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) NORMALSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserNORMALSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LuaParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(LuaParserNORMALSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LuaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LuaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
