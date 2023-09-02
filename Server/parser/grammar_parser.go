// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Server/Interfaces"
import "Server/Environment"
import "Server/Expression"
import "Server/Instruction"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'",
		"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'",
		"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'",
		"'continue'", "'return'", "'break'", "'append'", "'removeLast'", "'remove'",
		"'at'", "'IsEmpty'", "'count'", "'func'", "'inout'", "", "", "", "'!='",
		"'=='", "'!'", "'||'", "'&&'", "'='", "'+='", "'-='", "'>='", "'<='",
		"'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'",
		"'}'", "':'", "'?'", "'.'", "','", "';'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL",
		"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN",
		"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK",
		"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC",
		"INOUT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"DOUBLEPTS", "QUESTION", "POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt",
		"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt",
		"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt",
		"appendstmt", "removelaststmt", "removestmt", "emptvecstmt", "countvecstmt",
		"accessvecstmt", "assignationvecstmt", "declarematrixstmt", "declarefuncstmt",
		"expr", "accessstmt", "increaseanddecreasestmt", "type", "listParams",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 637, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 4, 1, 68, 8, 1, 11, 1, 12, 1, 69, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 125, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 132, 8,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 143, 8,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 152, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 162, 8, 4, 1, 4, 1, 4, 3,
		4, 166, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 175, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 184, 8, 5, 1, 5, 1, 5,
		3, 5, 188, 8, 5, 1, 6, 4, 6, 191, 8, 6, 11, 6, 12, 6, 192, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 224, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 242,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 260, 8, 9, 1, 10, 4, 10, 263, 8, 10,
		11, 10, 12, 10, 264, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 273,
		8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3,
		13, 305, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 3, 15, 318, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 324, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 330, 8, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 335, 8, 15, 1, 15, 3, 15, 338, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 351,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 365, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 378, 8, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 389, 8, 16,
		1, 16, 1, 16, 3, 16, 393, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 402, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 412, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 425, 8, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 452, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 464, 8, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25,
		510, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 548, 8, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 5, 26, 585, 8, 26, 10, 26, 12, 26, 588, 9, 26, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 597, 8, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 605, 8, 28, 1, 28, 1, 28, 3, 28,
		609, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 621, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 5, 30, 632, 8, 30, 10, 30, 12, 30, 635, 9, 30, 1,
		30, 0, 2, 52, 60, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0,
		6, 1, 0, 21, 23, 1, 0, 47, 48, 1, 0, 49, 51, 2, 0, 43, 43, 45, 45, 2, 0,
		44, 44, 46, 46, 1, 0, 35, 36, 687, 0, 62, 1, 0, 0, 0, 2, 67, 1, 0, 0, 0,
		4, 124, 1, 0, 0, 0, 6, 126, 1, 0, 0, 0, 8, 165, 1, 0, 0, 0, 10, 187, 1,
		0, 0, 0, 12, 190, 1, 0, 0, 0, 14, 223, 1, 0, 0, 0, 16, 241, 1, 0, 0, 0,
		18, 259, 1, 0, 0, 0, 20, 262, 1, 0, 0, 0, 22, 268, 1, 0, 0, 0, 24, 276,
		1, 0, 0, 0, 26, 304, 1, 0, 0, 0, 28, 306, 1, 0, 0, 0, 30, 337, 1, 0, 0,
		0, 32, 392, 1, 0, 0, 0, 34, 394, 1, 0, 0, 0, 36, 405, 1, 0, 0, 0, 38, 415,
		1, 0, 0, 0, 40, 428, 1, 0, 0, 0, 42, 433, 1, 0, 0, 0, 44, 438, 1, 0, 0,
		0, 46, 444, 1, 0, 0, 0, 48, 455, 1, 0, 0, 0, 50, 509, 1, 0, 0, 0, 52, 547,
		1, 0, 0, 0, 54, 589, 1, 0, 0, 0, 56, 608, 1, 0, 0, 0, 58, 620, 1, 0, 0,
		0, 60, 622, 1, 0, 0, 0, 62, 63, 3, 2, 1, 0, 63, 64, 5, 0, 0, 1, 64, 65,
		6, 0, -1, 0, 65, 1, 1, 0, 0, 0, 66, 68, 3, 4, 2, 0, 67, 66, 1, 0, 0, 0,
		68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1,
		0, 0, 0, 71, 72, 6, 1, -1, 0, 72, 3, 1, 0, 0, 0, 73, 74, 3, 6, 3, 0, 74,
		75, 6, 2, -1, 0, 75, 125, 1, 0, 0, 0, 76, 77, 3, 8, 4, 0, 77, 78, 6, 2,
		-1, 0, 78, 125, 1, 0, 0, 0, 79, 80, 3, 10, 5, 0, 80, 81, 6, 2, -1, 0, 81,
		125, 1, 0, 0, 0, 82, 83, 3, 22, 11, 0, 83, 84, 6, 2, -1, 0, 84, 125, 1,
		0, 0, 0, 85, 86, 3, 56, 28, 0, 86, 87, 6, 2, -1, 0, 87, 125, 1, 0, 0, 0,
		88, 89, 3, 14, 7, 0, 89, 90, 6, 2, -1, 0, 90, 125, 1, 0, 0, 0, 91, 92,
		3, 24, 12, 0, 92, 93, 6, 2, -1, 0, 93, 125, 1, 0, 0, 0, 94, 95, 3, 26,
		13, 0, 95, 96, 6, 2, -1, 0, 96, 125, 1, 0, 0, 0, 97, 98, 3, 16, 8, 0, 98,
		99, 6, 2, -1, 0, 99, 125, 1, 0, 0, 0, 100, 101, 3, 28, 14, 0, 101, 102,
		6, 2, -1, 0, 102, 125, 1, 0, 0, 0, 103, 104, 3, 30, 15, 0, 104, 105, 6,
		2, -1, 0, 105, 125, 1, 0, 0, 0, 106, 107, 3, 32, 16, 0, 107, 108, 6, 2,
		-1, 0, 108, 125, 1, 0, 0, 0, 109, 110, 3, 34, 17, 0, 110, 111, 6, 2, -1,
		0, 111, 125, 1, 0, 0, 0, 112, 113, 3, 36, 18, 0, 113, 114, 6, 2, -1, 0,
		114, 125, 1, 0, 0, 0, 115, 116, 3, 38, 19, 0, 116, 117, 6, 2, -1, 0, 117,
		125, 1, 0, 0, 0, 118, 119, 3, 46, 23, 0, 119, 120, 6, 2, -1, 0, 120, 125,
		1, 0, 0, 0, 121, 122, 3, 50, 25, 0, 122, 123, 6, 2, -1, 0, 123, 125, 1,
		0, 0, 0, 124, 73, 1, 0, 0, 0, 124, 76, 1, 0, 0, 0, 124, 79, 1, 0, 0, 0,
		124, 82, 1, 0, 0, 0, 124, 85, 1, 0, 0, 0, 124, 88, 1, 0, 0, 0, 124, 91,
		1, 0, 0, 0, 124, 94, 1, 0, 0, 0, 124, 97, 1, 0, 0, 0, 124, 100, 1, 0, 0,
		0, 124, 103, 1, 0, 0, 0, 124, 106, 1, 0, 0, 0, 124, 109, 1, 0, 0, 0, 124,
		112, 1, 0, 0, 0, 124, 115, 1, 0, 0, 0, 124, 118, 1, 0, 0, 0, 124, 121,
		1, 0, 0, 0, 125, 5, 1, 0, 0, 0, 126, 127, 5, 10, 0, 0, 127, 128, 5, 52,
		0, 0, 128, 129, 3, 52, 26, 0, 129, 131, 5, 53, 0, 0, 130, 132, 5, 60, 0,
		0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133,
		134, 6, 3, -1, 0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 6, 0, 0, 136, 137, 5,
		34, 0, 0, 137, 138, 5, 56, 0, 0, 138, 139, 3, 58, 29, 0, 139, 140, 5, 40,
		0, 0, 140, 142, 3, 52, 26, 0, 141, 143, 5, 60, 0, 0, 142, 141, 1, 0, 0,
		0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 6, 4, -1, 0, 145,
		166, 1, 0, 0, 0, 146, 147, 5, 6, 0, 0, 147, 148, 5, 34, 0, 0, 148, 149,
		5, 40, 0, 0, 149, 151, 3, 52, 26, 0, 150, 152, 5, 60, 0, 0, 151, 150, 1,
		0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 6, 4, -1,
		0, 154, 166, 1, 0, 0, 0, 155, 156, 5, 6, 0, 0, 156, 157, 5, 34, 0, 0, 157,
		158, 5, 56, 0, 0, 158, 159, 3, 58, 29, 0, 159, 161, 5, 57, 0, 0, 160, 162,
		5, 60, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 164, 6, 4, -1, 0, 164, 166, 1, 0, 0, 0, 165, 135, 1, 0, 0, 0,
		165, 146, 1, 0, 0, 0, 165, 155, 1, 0, 0, 0, 166, 9, 1, 0, 0, 0, 167, 168,
		5, 14, 0, 0, 168, 169, 5, 34, 0, 0, 169, 170, 5, 56, 0, 0, 170, 171, 3,
		58, 29, 0, 171, 172, 5, 40, 0, 0, 172, 174, 3, 52, 26, 0, 173, 175, 5,
		60, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 177, 6, 5, -1, 0, 177, 188, 1, 0, 0, 0, 178, 179, 5, 14, 0, 0,
		179, 180, 5, 34, 0, 0, 180, 181, 5, 40, 0, 0, 181, 183, 3, 52, 26, 0, 182,
		184, 5, 60, 0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185,
		1, 0, 0, 0, 185, 186, 6, 5, -1, 0, 186, 188, 1, 0, 0, 0, 187, 167, 1, 0,
		0, 0, 187, 178, 1, 0, 0, 0, 188, 11, 1, 0, 0, 0, 189, 191, 3, 14, 7, 0,
		190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 6, 6, -1, 0, 195, 13,
		1, 0, 0, 0, 196, 197, 5, 11, 0, 0, 197, 198, 3, 52, 26, 0, 198, 199, 5,
		54, 0, 0, 199, 200, 3, 2, 1, 0, 200, 201, 5, 55, 0, 0, 201, 202, 6, 7,
		-1, 0, 202, 224, 1, 0, 0, 0, 203, 204, 5, 11, 0, 0, 204, 205, 3, 52, 26,
		0, 205, 206, 5, 54, 0, 0, 206, 207, 3, 2, 1, 0, 207, 208, 5, 55, 0, 0,
		208, 209, 5, 12, 0, 0, 209, 210, 5, 54, 0, 0, 210, 211, 3, 2, 1, 0, 211,
		212, 5, 55, 0, 0, 212, 213, 6, 7, -1, 0, 213, 224, 1, 0, 0, 0, 214, 215,
		5, 11, 0, 0, 215, 216, 3, 52, 26, 0, 216, 217, 5, 54, 0, 0, 217, 218, 3,
		2, 1, 0, 218, 219, 5, 55, 0, 0, 219, 220, 5, 12, 0, 0, 220, 221, 3, 12,
		6, 0, 221, 222, 6, 7, -1, 0, 222, 224, 1, 0, 0, 0, 223, 196, 1, 0, 0, 0,
		223, 203, 1, 0, 0, 0, 223, 214, 1, 0, 0, 0, 224, 15, 1, 0, 0, 0, 225, 226,
		5, 17, 0, 0, 226, 227, 3, 52, 26, 0, 227, 228, 5, 54, 0, 0, 228, 229, 5,
		19, 0, 0, 229, 230, 5, 56, 0, 0, 230, 231, 3, 2, 1, 0, 231, 232, 5, 55,
		0, 0, 232, 233, 6, 8, -1, 0, 233, 242, 1, 0, 0, 0, 234, 235, 5, 17, 0,
		0, 235, 236, 3, 52, 26, 0, 236, 237, 5, 54, 0, 0, 237, 238, 3, 18, 9, 0,
		238, 239, 5, 55, 0, 0, 239, 240, 6, 8, -1, 0, 240, 242, 1, 0, 0, 0, 241,
		225, 1, 0, 0, 0, 241, 234, 1, 0, 0, 0, 242, 17, 1, 0, 0, 0, 243, 244, 5,
		18, 0, 0, 244, 245, 3, 52, 26, 0, 245, 246, 5, 56, 0, 0, 246, 247, 3, 2,
		1, 0, 247, 248, 3, 20, 10, 0, 248, 249, 6, 9, -1, 0, 249, 260, 1, 0, 0,
		0, 250, 251, 5, 18, 0, 0, 251, 252, 3, 52, 26, 0, 252, 253, 5, 56, 0, 0,
		253, 254, 3, 2, 1, 0, 254, 255, 5, 19, 0, 0, 255, 256, 5, 56, 0, 0, 256,
		257, 3, 2, 1, 0, 257, 258, 6, 9, -1, 0, 258, 260, 1, 0, 0, 0, 259, 243,
		1, 0, 0, 0, 259, 250, 1, 0, 0, 0, 260, 19, 1, 0, 0, 0, 261, 263, 3, 18,
		9, 0, 262, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0,
		264, 265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 6, 10, -1, 0, 267,
		21, 1, 0, 0, 0, 268, 269, 5, 34, 0, 0, 269, 270, 5, 40, 0, 0, 270, 272,
		3, 52, 26, 0, 271, 273, 5, 60, 0, 0, 272, 271, 1, 0, 0, 0, 272, 273, 1,
		0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 6, 11, -1, 0, 275, 23, 1, 0, 0,
		0, 276, 277, 5, 13, 0, 0, 277, 278, 3, 52, 26, 0, 278, 279, 5, 54, 0, 0,
		279, 280, 3, 2, 1, 0, 280, 281, 5, 55, 0, 0, 281, 282, 6, 12, -1, 0, 282,
		25, 1, 0, 0, 0, 283, 284, 5, 15, 0, 0, 284, 285, 5, 34, 0, 0, 285, 286,
		5, 16, 0, 0, 286, 287, 3, 52, 26, 0, 287, 288, 5, 54, 0, 0, 288, 289, 3,
		2, 1, 0, 289, 290, 5, 55, 0, 0, 290, 291, 6, 13, -1, 0, 291, 305, 1, 0,
		0, 0, 292, 293, 5, 15, 0, 0, 293, 294, 5, 34, 0, 0, 294, 295, 5, 16, 0,
		0, 295, 296, 3, 52, 26, 0, 296, 297, 5, 58, 0, 0, 297, 298, 5, 58, 0, 0,
		298, 299, 3, 52, 26, 0, 299, 300, 5, 54, 0, 0, 300, 301, 3, 2, 1, 0, 301,
		302, 5, 55, 0, 0, 302, 303, 6, 13, -1, 0, 303, 305, 1, 0, 0, 0, 304, 283,
		1, 0, 0, 0, 304, 292, 1, 0, 0, 0, 305, 27, 1, 0, 0, 0, 306, 307, 5, 20,
		0, 0, 307, 308, 3, 52, 26, 0, 308, 309, 5, 12, 0, 0, 309, 310, 5, 54, 0,
		0, 310, 311, 3, 2, 1, 0, 311, 312, 7, 0, 0, 0, 312, 313, 5, 55, 0, 0, 313,
		314, 6, 14, -1, 0, 314, 29, 1, 0, 0, 0, 315, 317, 5, 22, 0, 0, 316, 318,
		5, 60, 0, 0, 317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 1, 0,
		0, 0, 319, 338, 6, 15, -1, 0, 320, 321, 5, 22, 0, 0, 321, 323, 3, 52, 26,
		0, 322, 324, 5, 60, 0, 0, 323, 322, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324,
		325, 1, 0, 0, 0, 325, 326, 6, 15, -1, 0, 326, 338, 1, 0, 0, 0, 327, 329,
		5, 21, 0, 0, 328, 330, 5, 60, 0, 0, 329, 328, 1, 0, 0, 0, 329, 330, 1,
		0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 338, 6, 15, -1, 0, 332, 334, 5, 23,
		0, 0, 333, 335, 5, 60, 0, 0, 334, 333, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0,
		335, 336, 1, 0, 0, 0, 336, 338, 6, 15, -1, 0, 337, 315, 1, 0, 0, 0, 337,
		320, 1, 0, 0, 0, 337, 327, 1, 0, 0, 0, 337, 332, 1, 0, 0, 0, 338, 31, 1,
		0, 0, 0, 339, 340, 5, 6, 0, 0, 340, 341, 5, 34, 0, 0, 341, 342, 5, 56,
		0, 0, 342, 343, 5, 61, 0, 0, 343, 344, 3, 58, 29, 0, 344, 345, 5, 62, 0,
		0, 345, 346, 5, 40, 0, 0, 346, 347, 5, 61, 0, 0, 347, 348, 3, 60, 30, 0,
		348, 350, 5, 62, 0, 0, 349, 351, 5, 60, 0, 0, 350, 349, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 6, 16, -1, 0, 353, 393,
		1, 0, 0, 0, 354, 355, 5, 6, 0, 0, 355, 356, 5, 34, 0, 0, 356, 357, 5, 56,
		0, 0, 357, 358, 5, 61, 0, 0, 358, 359, 3, 58, 29, 0, 359, 360, 5, 62, 0,
		0, 360, 361, 5, 40, 0, 0, 361, 362, 5, 61, 0, 0, 362, 364, 5, 62, 0, 0,
		363, 365, 5, 60, 0, 0, 364, 363, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365,
		366, 1, 0, 0, 0, 366, 367, 6, 16, -1, 0, 367, 393, 1, 0, 0, 0, 368, 369,
		5, 6, 0, 0, 369, 370, 5, 34, 0, 0, 370, 371, 5, 40, 0, 0, 371, 372, 5,
		61, 0, 0, 372, 373, 3, 58, 29, 0, 373, 374, 5, 62, 0, 0, 374, 375, 5, 52,
		0, 0, 375, 377, 5, 53, 0, 0, 376, 378, 5, 60, 0, 0, 377, 376, 1, 0, 0,
		0, 377, 378, 1, 0, 0, 0, 378, 393, 1, 0, 0, 0, 379, 380, 5, 6, 0, 0, 380,
		381, 5, 34, 0, 0, 381, 382, 5, 56, 0, 0, 382, 383, 5, 61, 0, 0, 383, 384,
		3, 58, 29, 0, 384, 385, 5, 62, 0, 0, 385, 386, 5, 40, 0, 0, 386, 388, 5,
		34, 0, 0, 387, 389, 5, 60, 0, 0, 388, 387, 1, 0, 0, 0, 388, 389, 1, 0,
		0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 6, 16, -1, 0, 391, 393, 1, 0, 0,
		0, 392, 339, 1, 0, 0, 0, 392, 354, 1, 0, 0, 0, 392, 368, 1, 0, 0, 0, 392,
		379, 1, 0, 0, 0, 393, 33, 1, 0, 0, 0, 394, 395, 5, 34, 0, 0, 395, 396,
		5, 58, 0, 0, 396, 397, 5, 24, 0, 0, 397, 398, 5, 52, 0, 0, 398, 399, 3,
		52, 26, 0, 399, 401, 5, 53, 0, 0, 400, 402, 5, 60, 0, 0, 401, 400, 1, 0,
		0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 404, 6, 17, -1,
		0, 404, 35, 1, 0, 0, 0, 405, 406, 5, 34, 0, 0, 406, 407, 5, 58, 0, 0, 407,
		408, 5, 25, 0, 0, 408, 409, 5, 52, 0, 0, 409, 411, 5, 53, 0, 0, 410, 412,
		5, 60, 0, 0, 411, 410, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 1, 0,
		0, 0, 413, 414, 6, 18, -1, 0, 414, 37, 1, 0, 0, 0, 415, 416, 5, 34, 0,
		0, 416, 417, 5, 58, 0, 0, 417, 418, 5, 26, 0, 0, 418, 419, 5, 52, 0, 0,
		419, 420, 5, 27, 0, 0, 420, 421, 5, 56, 0, 0, 421, 422, 3, 52, 26, 0, 422,
		424, 5, 53, 0, 0, 423, 425, 5, 60, 0, 0, 424, 423, 1, 0, 0, 0, 424, 425,
		1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 427, 6, 19, -1, 0, 427, 39, 1, 0,
		0, 0, 428, 429, 5, 34, 0, 0, 429, 430, 5, 58, 0, 0, 430, 431, 5, 28, 0,
		0, 431, 432, 6, 20, -1, 0, 432, 41, 1, 0, 0, 0, 433, 434, 5, 34, 0, 0,
		434, 435, 5, 58, 0, 0, 435, 436, 5, 29, 0, 0, 436, 437, 6, 21, -1, 0, 437,
		43, 1, 0, 0, 0, 438, 439, 5, 34, 0, 0, 439, 440, 5, 61, 0, 0, 440, 441,
		3, 52, 26, 0, 441, 442, 5, 62, 0, 0, 442, 443, 6, 22, -1, 0, 443, 45, 1,
		0, 0, 0, 444, 445, 5, 34, 0, 0, 445, 446, 5, 61, 0, 0, 446, 447, 3, 52,
		26, 0, 447, 448, 5, 62, 0, 0, 448, 449, 5, 40, 0, 0, 449, 451, 3, 52, 26,
		0, 450, 452, 5, 60, 0, 0, 451, 450, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452,
		453, 1, 0, 0, 0, 453, 454, 6, 23, -1, 0, 454, 47, 1, 0, 0, 0, 455, 456,
		5, 6, 0, 0, 456, 457, 5, 34, 0, 0, 457, 458, 5, 52, 0, 0, 458, 459, 5,
		56, 0, 0, 459, 460, 3, 58, 29, 0, 460, 461, 5, 53, 0, 0, 461, 463, 5, 40,
		0, 0, 462, 464, 5, 60, 0, 0, 463, 462, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0,
		464, 49, 1, 0, 0, 0, 465, 466, 5, 30, 0, 0, 466, 467, 5, 34, 0, 0, 467,
		468, 5, 52, 0, 0, 468, 469, 3, 60, 30, 0, 469, 470, 5, 53, 0, 0, 470, 471,
		5, 50, 0, 0, 471, 472, 5, 45, 0, 0, 472, 473, 3, 58, 29, 0, 473, 474, 5,
		54, 0, 0, 474, 475, 3, 2, 1, 0, 475, 476, 5, 55, 0, 0, 476, 477, 6, 25,
		-1, 0, 477, 510, 1, 0, 0, 0, 478, 479, 5, 30, 0, 0, 479, 480, 5, 34, 0,
		0, 480, 481, 5, 52, 0, 0, 481, 482, 3, 60, 30, 0, 482, 483, 5, 53, 0, 0,
		483, 484, 5, 54, 0, 0, 484, 485, 3, 2, 1, 0, 485, 486, 5, 55, 0, 0, 486,
		487, 6, 25, -1, 0, 487, 510, 1, 0, 0, 0, 488, 489, 5, 30, 0, 0, 489, 490,
		5, 34, 0, 0, 490, 491, 5, 52, 0, 0, 491, 492, 5, 53, 0, 0, 492, 493, 5,
		50, 0, 0, 493, 494, 5, 45, 0, 0, 494, 495, 3, 58, 29, 0, 495, 496, 5, 54,
		0, 0, 496, 497, 3, 2, 1, 0, 497, 498, 5, 55, 0, 0, 498, 499, 6, 25, -1,
		0, 499, 510, 1, 0, 0, 0, 500, 501, 5, 30, 0, 0, 501, 502, 5, 34, 0, 0,
		502, 503, 5, 52, 0, 0, 503, 504, 5, 53, 0, 0, 504, 505, 5, 54, 0, 0, 505,
		506, 3, 2, 1, 0, 506, 507, 5, 55, 0, 0, 507, 508, 6, 25, -1, 0, 508, 510,
		1, 0, 0, 0, 509, 465, 1, 0, 0, 0, 509, 478, 1, 0, 0, 0, 509, 488, 1, 0,
		0, 0, 509, 500, 1, 0, 0, 0, 510, 51, 1, 0, 0, 0, 511, 512, 6, 26, -1, 0,
		512, 513, 5, 37, 0, 0, 513, 514, 3, 52, 26, 12, 514, 515, 6, 26, -1, 0,
		515, 548, 1, 0, 0, 0, 516, 517, 5, 50, 0, 0, 517, 518, 3, 52, 26, 11, 518,
		519, 6, 26, -1, 0, 519, 548, 1, 0, 0, 0, 520, 521, 5, 52, 0, 0, 521, 522,
		3, 52, 26, 0, 522, 523, 5, 53, 0, 0, 523, 524, 6, 26, -1, 0, 524, 548,
		1, 0, 0, 0, 525, 526, 5, 32, 0, 0, 526, 548, 6, 26, -1, 0, 527, 528, 5,
		33, 0, 0, 528, 548, 6, 26, -1, 0, 529, 530, 5, 8, 0, 0, 530, 548, 6, 26,
		-1, 0, 531, 532, 5, 9, 0, 0, 532, 548, 6, 26, -1, 0, 533, 534, 5, 7, 0,
		0, 534, 548, 6, 26, -1, 0, 535, 536, 3, 54, 27, 0, 536, 537, 6, 26, -1,
		0, 537, 548, 1, 0, 0, 0, 538, 539, 3, 40, 20, 0, 539, 540, 6, 26, -1, 0,
		540, 548, 1, 0, 0, 0, 541, 542, 3, 42, 21, 0, 542, 543, 6, 26, -1, 0, 543,
		548, 1, 0, 0, 0, 544, 545, 3, 44, 22, 0, 545, 546, 6, 26, -1, 0, 546, 548,
		1, 0, 0, 0, 547, 511, 1, 0, 0, 0, 547, 516, 1, 0, 0, 0, 547, 520, 1, 0,
		0, 0, 547, 525, 1, 0, 0, 0, 547, 527, 1, 0, 0, 0, 547, 529, 1, 0, 0, 0,
		547, 531, 1, 0, 0, 0, 547, 533, 1, 0, 0, 0, 547, 535, 1, 0, 0, 0, 547,
		538, 1, 0, 0, 0, 547, 541, 1, 0, 0, 0, 547, 544, 1, 0, 0, 0, 548, 586,
		1, 0, 0, 0, 549, 550, 10, 19, 0, 0, 550, 551, 7, 1, 0, 0, 551, 552, 3,
		52, 26, 20, 552, 553, 6, 26, -1, 0, 553, 585, 1, 0, 0, 0, 554, 555, 10,
		18, 0, 0, 555, 556, 7, 2, 0, 0, 556, 557, 3, 52, 26, 19, 557, 558, 6, 26,
		-1, 0, 558, 585, 1, 0, 0, 0, 559, 560, 10, 17, 0, 0, 560, 561, 7, 3, 0,
		0, 561, 562, 3, 52, 26, 18, 562, 563, 6, 26, -1, 0, 563, 585, 1, 0, 0,
		0, 564, 565, 10, 16, 0, 0, 565, 566, 7, 4, 0, 0, 566, 567, 3, 52, 26, 17,
		567, 568, 6, 26, -1, 0, 568, 585, 1, 0, 0, 0, 569, 570, 10, 15, 0, 0, 570,
		571, 7, 5, 0, 0, 571, 572, 3, 52, 26, 16, 572, 573, 6, 26, -1, 0, 573,
		585, 1, 0, 0, 0, 574, 575, 10, 14, 0, 0, 575, 576, 5, 39, 0, 0, 576, 577,
		3, 52, 26, 15, 577, 578, 6, 26, -1, 0, 578, 585, 1, 0, 0, 0, 579, 580,
		10, 13, 0, 0, 580, 581, 5, 38, 0, 0, 581, 582, 3, 52, 26, 14, 582, 583,
		6, 26, -1, 0, 583, 585, 1, 0, 0, 0, 584, 549, 1, 0, 0, 0, 584, 554, 1,
		0, 0, 0, 584, 559, 1, 0, 0, 0, 584, 564, 1, 0, 0, 0, 584, 569, 1, 0, 0,
		0, 584, 574, 1, 0, 0, 0, 584, 579, 1, 0, 0, 0, 585, 588, 1, 0, 0, 0, 586,
		584, 1, 0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 53, 1, 0, 0, 0, 588, 586, 1,
		0, 0, 0, 589, 590, 5, 34, 0, 0, 590, 591, 6, 27, -1, 0, 591, 55, 1, 0,
		0, 0, 592, 593, 5, 34, 0, 0, 593, 594, 5, 41, 0, 0, 594, 596, 3, 52, 26,
		0, 595, 597, 5, 60, 0, 0, 596, 595, 1, 0, 0, 0, 596, 597, 1, 0, 0, 0, 597,
		598, 1, 0, 0, 0, 598, 599, 6, 28, -1, 0, 599, 609, 1, 0, 0, 0, 600, 601,
		5, 34, 0, 0, 601, 602, 5, 42, 0, 0, 602, 604, 3, 52, 26, 0, 603, 605, 5,
		60, 0, 0, 604, 603, 1, 0, 0, 0, 604, 605, 1, 0, 0, 0, 605, 606, 1, 0, 0,
		0, 606, 607, 6, 28, -1, 0, 607, 609, 1, 0, 0, 0, 608, 592, 1, 0, 0, 0,
		608, 600, 1, 0, 0, 0, 609, 57, 1, 0, 0, 0, 610, 611, 5, 1, 0, 0, 611, 621,
		6, 29, -1, 0, 612, 613, 5, 2, 0, 0, 613, 621, 6, 29, -1, 0, 614, 615, 5,
		3, 0, 0, 615, 621, 6, 29, -1, 0, 616, 617, 5, 4, 0, 0, 617, 621, 6, 29,
		-1, 0, 618, 619, 5, 5, 0, 0, 619, 621, 6, 29, -1, 0, 620, 610, 1, 0, 0,
		0, 620, 612, 1, 0, 0, 0, 620, 614, 1, 0, 0, 0, 620, 616, 1, 0, 0, 0, 620,
		618, 1, 0, 0, 0, 621, 59, 1, 0, 0, 0, 622, 623, 6, 30, -1, 0, 623, 624,
		3, 52, 26, 0, 624, 625, 6, 30, -1, 0, 625, 633, 1, 0, 0, 0, 626, 627, 10,
		2, 0, 0, 627, 628, 5, 59, 0, 0, 628, 629, 3, 52, 26, 0, 629, 630, 6, 30,
		-1, 0, 630, 632, 1, 0, 0, 0, 631, 626, 1, 0, 0, 0, 632, 635, 1, 0, 0, 0,
		633, 631, 1, 0, 0, 0, 633, 634, 1, 0, 0, 0, 634, 61, 1, 0, 0, 0, 635, 633,
		1, 0, 0, 0, 41, 69, 124, 131, 142, 151, 161, 165, 174, 183, 187, 192, 223,
		241, 259, 264, 272, 304, 317, 323, 329, 334, 337, 350, 364, 377, 388, 392,
		401, 411, 424, 451, 463, 509, 547, 584, 586, 596, 604, 608, 620, 633,
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

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF          = antlr.TokenEOF
	GrammarParserSTRINGS      = 1
	GrammarParserINTS         = 2
	GrammarParserFLOATS       = 3
	GrammarParserBOOLS        = 4
	GrammarParserCHARACTERS   = 5
	GrammarParserVAR          = 6
	GrammarParserNIL          = 7
	GrammarParserTRUE         = 8
	GrammarParserFALSE        = 9
	GrammarParserPRINT        = 10
	GrammarParserIF           = 11
	GrammarParserELSE         = 12
	GrammarParserWHILE        = 13
	GrammarParserLET          = 14
	GrammarParserFOR          = 15
	GrammarParserIN           = 16
	GrammarParserSWITCH       = 17
	GrammarParserCASE         = 18
	GrammarParserDEFAULT      = 19
	GrammarParserGUARD        = 20
	GrammarParserCONTINUE     = 21
	GrammarParserRETURN       = 22
	GrammarParserBREAK        = 23
	GrammarParserAPPEND       = 24
	GrammarParserREMOVELAST   = 25
	GrammarParserREMOVE       = 26
	GrammarParserAT           = 27
	GrammarParserISEMPTY      = 28
	GrammarParserCOUNT        = 29
	GrammarParserFUNC         = 30
	GrammarParserINOUT        = 31
	GrammarParserNUMBER       = 32
	GrammarParserSTRING       = 33
	GrammarParserID           = 34
	GrammarParserDIF          = 35
	GrammarParserIG_IG        = 36
	GrammarParserNOT          = 37
	GrammarParserOR           = 38
	GrammarParserAND          = 39
	GrammarParserIG           = 40
	GrammarParserIG_ADD       = 41
	GrammarParserIG_SUB       = 42
	GrammarParserMAY_IG       = 43
	GrammarParserMEN_IG       = 44
	GrammarParserMAYOR        = 45
	GrammarParserMENOR        = 46
	GrammarParserMUL          = 47
	GrammarParserDIV          = 48
	GrammarParserADD          = 49
	GrammarParserSUB          = 50
	GrammarParserPERCENT      = 51
	GrammarParserPARIZQ       = 52
	GrammarParserPARDER       = 53
	GrammarParserLLAVEIZQ     = 54
	GrammarParserLLAVEDER     = 55
	GrammarParserDOUBLEPTS    = 56
	GrammarParserQUESTION     = 57
	GrammarParserPOINT        = 58
	GrammarParserCOMA         = 59
	GrammarParserPTCOMA       = 60
	GrammarParserCORCHETEIZQ  = 61
	GrammarParserCORCHETEDER  = 62
	GrammarParserWHITESPACE   = 63
	GrammarParserCOMMENT      = 64
	GrammarParserLINE_COMMENT = 65
)

// GrammarParser rules.
const (
	GrammarParserRULE_s                       = 0
	GrammarParserRULE_block                   = 1
	GrammarParserRULE_instruction             = 2
	GrammarParserRULE_printstmt               = 3
	GrammarParserRULE_declarestmt             = 4
	GrammarParserRULE_constantstmt            = 5
	GrammarParserRULE_blockelifs              = 6
	GrammarParserRULE_ifstmt                  = 7
	GrammarParserRULE_switchstmt              = 8
	GrammarParserRULE_casestmt                = 9
	GrammarParserRULE_blockcases              = 10
	GrammarParserRULE_assignationstmt         = 11
	GrammarParserRULE_whilestmt               = 12
	GrammarParserRULE_forstmt                 = 13
	GrammarParserRULE_guardstmt               = 14
	GrammarParserRULE_transferstmt            = 15
	GrammarParserRULE_declarevectorstmt       = 16
	GrammarParserRULE_appendstmt              = 17
	GrammarParserRULE_removelaststmt          = 18
	GrammarParserRULE_removestmt              = 19
	GrammarParserRULE_emptvecstmt             = 20
	GrammarParserRULE_countvecstmt            = 21
	GrammarParserRULE_accessvecstmt           = 22
	GrammarParserRULE_assignationvecstmt      = 23
	GrammarParserRULE_declarematrixstmt       = 24
	GrammarParserRULE_declarefuncstmt         = 25
	GrammarParserRULE_expr                    = 26
	GrammarParserRULE_accessstmt              = 27
	GrammarParserRULE_increaseanddecreasestmt = 28
	GrammarParserRULE_type                    = 29
	GrammarParserRULE_listParams              = 30
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
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

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(63)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

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

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
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

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(66)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())

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

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_declarestmt returns the _declarestmt rule contexts.
	Get_declarestmt() IDeclarestmtContext

	// Get_constantstmt returns the _constantstmt rule contexts.
	Get_constantstmt() IConstantstmtContext

	// Get_assignationstmt returns the _assignationstmt rule contexts.
	Get_assignationstmt() IAssignationstmtContext

	// Get_increaseanddecreasestmt returns the _increaseanddecreasestmt rule contexts.
	Get_increaseanddecreasestmt() IIncreaseanddecreasestmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_switchstmt returns the _switchstmt rule contexts.
	Get_switchstmt() ISwitchstmtContext

	// Get_guardstmt returns the _guardstmt rule contexts.
	Get_guardstmt() IGuardstmtContext

	// Get_transferstmt returns the _transferstmt rule contexts.
	Get_transferstmt() ITransferstmtContext

	// Get_declarevectorstmt returns the _declarevectorstmt rule contexts.
	Get_declarevectorstmt() IDeclarevectorstmtContext

	// Get_appendstmt returns the _appendstmt rule contexts.
	Get_appendstmt() IAppendstmtContext

	// Get_removelaststmt returns the _removelaststmt rule contexts.
	Get_removelaststmt() IRemovelaststmtContext

	// Get_removestmt returns the _removestmt rule contexts.
	Get_removestmt() IRemovestmtContext

	// Get_assignationvecstmt returns the _assignationvecstmt rule contexts.
	Get_assignationvecstmt() IAssignationvecstmtContext

	// Get_declarefuncstmt returns the _declarefuncstmt rule contexts.
	Get_declarefuncstmt() IDeclarefuncstmtContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_declarestmt sets the _declarestmt rule contexts.
	Set_declarestmt(IDeclarestmtContext)

	// Set_constantstmt sets the _constantstmt rule contexts.
	Set_constantstmt(IConstantstmtContext)

	// Set_assignationstmt sets the _assignationstmt rule contexts.
	Set_assignationstmt(IAssignationstmtContext)

	// Set_increaseanddecreasestmt sets the _increaseanddecreasestmt rule contexts.
	Set_increaseanddecreasestmt(IIncreaseanddecreasestmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_switchstmt sets the _switchstmt rule contexts.
	Set_switchstmt(ISwitchstmtContext)

	// Set_guardstmt sets the _guardstmt rule contexts.
	Set_guardstmt(IGuardstmtContext)

	// Set_transferstmt sets the _transferstmt rule contexts.
	Set_transferstmt(ITransferstmtContext)

	// Set_declarevectorstmt sets the _declarevectorstmt rule contexts.
	Set_declarevectorstmt(IDeclarevectorstmtContext)

	// Set_appendstmt sets the _appendstmt rule contexts.
	Set_appendstmt(IAppendstmtContext)

	// Set_removelaststmt sets the _removelaststmt rule contexts.
	Set_removelaststmt(IRemovelaststmtContext)

	// Set_removestmt sets the _removestmt rule contexts.
	Set_removestmt(IRemovestmtContext)

	// Set_assignationvecstmt sets the _assignationvecstmt rule contexts.
	Set_assignationvecstmt(IAssignationvecstmtContext)

	// Set_declarefuncstmt sets the _declarefuncstmt rule contexts.
	Set_declarefuncstmt(IDeclarefuncstmtContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	Declarestmt() IDeclarestmtContext
	Constantstmt() IConstantstmtContext
	Assignationstmt() IAssignationstmtContext
	Increaseanddecreasestmt() IIncreaseanddecreasestmtContext
	Ifstmt() IIfstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Switchstmt() ISwitchstmtContext
	Guardstmt() IGuardstmtContext
	Transferstmt() ITransferstmtContext
	Declarevectorstmt() IDeclarevectorstmtContext
	Appendstmt() IAppendstmtContext
	Removelaststmt() IRemovelaststmtContext
	Removestmt() IRemovestmtContext
	Assignationvecstmt() IAssignationvecstmtContext
	Declarefuncstmt() IDeclarefuncstmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser                   antlr.Parser
	inst                     interfaces.Instruction
	_printstmt               IPrintstmtContext
	_declarestmt             IDeclarestmtContext
	_constantstmt            IConstantstmtContext
	_assignationstmt         IAssignationstmtContext
	_increaseanddecreasestmt IIncreaseanddecreasestmtContext
	_ifstmt                  IIfstmtContext
	_whilestmt               IWhilestmtContext
	_forstmt                 IForstmtContext
	_switchstmt              ISwitchstmtContext
	_guardstmt               IGuardstmtContext
	_transferstmt            ITransferstmtContext
	_declarevectorstmt       IDeclarevectorstmtContext
	_appendstmt              IAppendstmtContext
	_removelaststmt          IRemovelaststmtContext
	_removestmt              IRemovestmtContext
	_assignationvecstmt      IAssignationvecstmtContext
	_declarefuncstmt         IDeclarefuncstmtContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_declarestmt() IDeclarestmtContext { return s._declarestmt }

func (s *InstructionContext) Get_constantstmt() IConstantstmtContext { return s._constantstmt }

func (s *InstructionContext) Get_assignationstmt() IAssignationstmtContext { return s._assignationstmt }

func (s *InstructionContext) Get_increaseanddecreasestmt() IIncreaseanddecreasestmtContext {
	return s._increaseanddecreasestmt
}

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_transferstmt() ITransferstmtContext { return s._transferstmt }

func (s *InstructionContext) Get_declarevectorstmt() IDeclarevectorstmtContext {
	return s._declarevectorstmt
}

func (s *InstructionContext) Get_appendstmt() IAppendstmtContext { return s._appendstmt }

func (s *InstructionContext) Get_removelaststmt() IRemovelaststmtContext { return s._removelaststmt }

func (s *InstructionContext) Get_removestmt() IRemovestmtContext { return s._removestmt }

func (s *InstructionContext) Get_assignationvecstmt() IAssignationvecstmtContext {
	return s._assignationvecstmt
}

func (s *InstructionContext) Get_declarefuncstmt() IDeclarefuncstmtContext { return s._declarefuncstmt }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_declarestmt(v IDeclarestmtContext) { s._declarestmt = v }

func (s *InstructionContext) Set_constantstmt(v IConstantstmtContext) { s._constantstmt = v }

func (s *InstructionContext) Set_assignationstmt(v IAssignationstmtContext) { s._assignationstmt = v }

func (s *InstructionContext) Set_increaseanddecreasestmt(v IIncreaseanddecreasestmtContext) {
	s._increaseanddecreasestmt = v
}

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_switchstmt(v ISwitchstmtContext) { s._switchstmt = v }

func (s *InstructionContext) Set_guardstmt(v IGuardstmtContext) { s._guardstmt = v }

func (s *InstructionContext) Set_transferstmt(v ITransferstmtContext) { s._transferstmt = v }

func (s *InstructionContext) Set_declarevectorstmt(v IDeclarevectorstmtContext) {
	s._declarevectorstmt = v
}

func (s *InstructionContext) Set_appendstmt(v IAppendstmtContext) { s._appendstmt = v }

func (s *InstructionContext) Set_removelaststmt(v IRemovelaststmtContext) { s._removelaststmt = v }

func (s *InstructionContext) Set_removestmt(v IRemovestmtContext) { s._removestmt = v }

func (s *InstructionContext) Set_assignationvecstmt(v IAssignationvecstmtContext) {
	s._assignationvecstmt = v
}

func (s *InstructionContext) Set_declarefuncstmt(v IDeclarefuncstmtContext) { s._declarefuncstmt = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Declarestmt() IDeclarestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarestmtContext)
}

func (s *InstructionContext) Constantstmt() IConstantstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantstmtContext)
}

func (s *InstructionContext) Assignationstmt() IAssignationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignationstmtContext)
}

func (s *InstructionContext) Increaseanddecreasestmt() IIncreaseanddecreasestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncreaseanddecreasestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncreaseanddecreasestmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Transferstmt() ITransferstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransferstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransferstmtContext)
}

func (s *InstructionContext) Declarevectorstmt() IDeclarevectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarevectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarevectorstmtContext)
}

func (s *InstructionContext) Appendstmt() IAppendstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendstmtContext)
}

func (s *InstructionContext) Removelaststmt() IRemovelaststmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovelaststmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovelaststmtContext)
}

func (s *InstructionContext) Removestmt() IRemovestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovestmtContext)
}

func (s *InstructionContext) Assignationvecstmt() IAssignationvecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignationvecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignationvecstmtContext)
}

func (s *InstructionContext) Declarefuncstmt() IDeclarefuncstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarefuncstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarefuncstmtContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *GrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_instruction)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)

			var _x = p.Declarestmt()

			localctx.(*InstructionContext)._declarestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarestmt().GetDec()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)

			var _x = p.Constantstmt()

			localctx.(*InstructionContext)._constantstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constantstmt().GetConst_()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)

			var _x = p.Assignationstmt()

			localctx.(*InstructionContext)._assignationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignationstmt().GetAssign()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)

			var _x = p.Increaseanddecreasestmt()

			localctx.(*InstructionContext)._increaseanddecreasestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_increaseanddecreasestmt().GetIncreasedecrease()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIft()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhile()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(94)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFor_()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(97)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwitch_()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(100)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGua()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(103)

			var _x = p.Transferstmt()

			localctx.(*InstructionContext)._transferstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transferstmt().GetTran()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(106)

			var _x = p.Declarevectorstmt()

			localctx.(*InstructionContext)._declarevectorstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarevectorstmt().GetDecvec()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(109)

			var _x = p.Appendstmt()

			localctx.(*InstructionContext)._appendstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_appendstmt().GetApp()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(112)

			var _x = p.Removelaststmt()

			localctx.(*InstructionContext)._removelaststmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removelaststmt().GetRemovl()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(115)

			var _x = p.Removestmt()

			localctx.(*InstructionContext)._removestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removestmt().GetRemov()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(118)

			var _x = p.Assignationvecstmt()

			localctx.(*InstructionContext)._assignationvecstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignationvecstmt().GetAssignvec()

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(121)

			var _x = p.Declarefuncstmt()

			localctx.(*InstructionContext)._declarefuncstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarefuncstmt().GetDecfunc()

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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prnt   interfaces.Instruction
	_PRINT antlr.Token
	_expr  IExprContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *PrintstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *GrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_printstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)

		var _m = p.Match(GrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(129)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(130)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_expr().GetE())

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

// IDeclarestmtContext is an interface to support dynamic dispatch.
type IDeclarestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_type returns the _type rule contexts.
	Get_type() ITypeContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_type sets the _type rule contexts.
	Set_type(ITypeContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	Type_() ITypeContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	PTCOMA() antlr.TerminalNode
	QUESTION() antlr.TerminalNode

	// IsDeclarestmtContext differentiates from other interfaces.
	IsDeclarestmtContext()
}

type DeclarestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    interfaces.Instruction
	_VAR   antlr.Token
	_ID    antlr.Token
	_type  ITypeContext
	_expr  IExprContext
}

func NewEmptyDeclarestmtContext() *DeclarestmtContext {
	var p = new(DeclarestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarestmt
	return p
}

func InitEmptyDeclarestmtContext(p *DeclarestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarestmt
}

func (*DeclarestmtContext) IsDeclarestmtContext() {}

func NewDeclarestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarestmtContext {
	var p = new(DeclarestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_declarestmt

	return p
}

func (s *DeclarestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarestmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclarestmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarestmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarestmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarestmtContext) Get_type() ITypeContext { return s._type }

func (s *DeclarestmtContext) Get_expr() IExprContext { return s._expr }

func (s *DeclarestmtContext) Set_type(v ITypeContext) { s._type = v }

func (s *DeclarestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclarestmtContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarestmtContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarestmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclarestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclarestmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *DeclarestmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarestmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *DeclarestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarestmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *DeclarestmtContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(GrammarParserQUESTION, 0)
}

func (s *DeclarestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclarestmt(s)
	}
}

func (s *DeclarestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclarestmt(s)
	}
}

func (p *GrammarParser) Declarestmt() (localctx IDeclarestmtContext) {
	localctx = NewDeclarestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_declarestmt)
	var _la int

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(139)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)

			var _x = p.expr(0)

			localctx.(*DeclarestmtContext)._expr = _x
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(141)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarestmtContext).dec = instructions.NewTodeclare((func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarestmtContext).Get_type().GetT(), localctx.(*DeclarestmtContext).Get_expr().GetE(), false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)

			var _x = p.expr(0)

			localctx.(*DeclarestmtContext)._expr = _x
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(150)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarestmtContext).dec = instructions.NewTodeclare((func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarestmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*DeclarestmtContext).Get_expr().GetE(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(159)
			p.Match(GrammarParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(160)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarestmtContext).dec = instructions.NewTodeclare((func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarestmtContext).Get_type().GetT(), nil, false)

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

// IConstantstmtContext is an interface to support dynamic dispatch.
type IConstantstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_type returns the _type rule contexts.
	Get_type() ITypeContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_type sets the _type rule contexts.
	Set_type(ITypeContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetConst_ returns the const_ attribute.
	GetConst_() interfaces.Instruction

	// SetConst_ sets the const_ attribute.
	SetConst_(interfaces.Instruction)

	// Getter signatures
	LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	Type_() ITypeContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	PTCOMA() antlr.TerminalNode

	// IsConstantstmtContext differentiates from other interfaces.
	IsConstantstmtContext()
}

type ConstantstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	const_ interfaces.Instruction
	_LET   antlr.Token
	_ID    antlr.Token
	_type  ITypeContext
	_expr  IExprContext
}

func NewEmptyConstantstmtContext() *ConstantstmtContext {
	var p = new(ConstantstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_constantstmt
	return p
}

func InitEmptyConstantstmtContext(p *ConstantstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_constantstmt
}

func (*ConstantstmtContext) IsConstantstmtContext() {}

func NewConstantstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantstmtContext {
	var p = new(ConstantstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_constantstmt

	return p
}

func (s *ConstantstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantstmtContext) Get_LET() antlr.Token { return s._LET }

func (s *ConstantstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ConstantstmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *ConstantstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ConstantstmtContext) Get_type() ITypeContext { return s._type }

func (s *ConstantstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ConstantstmtContext) Set_type(v ITypeContext) { s._type = v }

func (s *ConstantstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ConstantstmtContext) GetConst_() interfaces.Instruction { return s.const_ }

func (s *ConstantstmtContext) SetConst_(v interfaces.Instruction) { s.const_ = v }

func (s *ConstantstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *ConstantstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ConstantstmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *ConstantstmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstantstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *ConstantstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstantstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *ConstantstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterConstantstmt(s)
	}
}

func (s *ConstantstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitConstantstmt(s)
	}
}

func (p *GrammarParser) Constantstmt() (localctx IConstantstmtContext) {
	localctx = NewConstantstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_constantstmt)
	var _la int

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)

			var _x = p.Type_()

			localctx.(*ConstantstmtContext)._type = _x
		}
		{
			p.SetState(171)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _x = p.expr(0)

			localctx.(*ConstantstmtContext)._expr = _x
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(173)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*ConstantstmtContext).const_ = instructions.NewTodeclare((func() int {
			if localctx.(*ConstantstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConstantstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConstantstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConstantstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConstantstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConstantstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ConstantstmtContext).Get_type().GetT(), localctx.(*ConstantstmtContext).Get_expr().GetE(), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)

			var _x = p.expr(0)

			localctx.(*ConstantstmtContext)._expr = _x
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(182)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*ConstantstmtContext).const_ = instructions.NewTodeclare((func() int {
			if localctx.(*ConstantstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConstantstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConstantstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConstantstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConstantstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConstantstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*ConstantstmtContext).Get_expr().GetE(), true)

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

// IBlockelifsContext is an interface to support dynamic dispatch.
type IBlockelifsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// GetElif returns the elif rule context list.
	GetElif() []IIfstmtContext

	// SetElif sets the elif rule context list.
	SetElif([]IIfstmtContext)

	// GetBlkef returns the blkef attribute.
	GetBlkef() []interface{}

	// SetBlkef sets the blkef attribute.
	SetBlkef([]interface{})

	// Getter signatures
	AllIfstmt() []IIfstmtContext
	Ifstmt(i int) IIfstmtContext

	// IsBlockelifsContext differentiates from other interfaces.
	IsBlockelifsContext()
}

type BlockelifsContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	blkef   []interface{}
	_ifstmt IIfstmtContext
	elif    []IIfstmtContext
}

func NewEmptyBlockelifsContext() *BlockelifsContext {
	var p = new(BlockelifsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockelifs
	return p
}

func InitEmptyBlockelifsContext(p *BlockelifsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockelifs
}

func (*BlockelifsContext) IsBlockelifsContext() {}

func NewBlockelifsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockelifsContext {
	var p = new(BlockelifsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_blockelifs

	return p
}

func (s *BlockelifsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockelifsContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *BlockelifsContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *BlockelifsContext) GetElif() []IIfstmtContext { return s.elif }

func (s *BlockelifsContext) SetElif(v []IIfstmtContext) { s.elif = v }

func (s *BlockelifsContext) GetBlkef() []interface{} { return s.blkef }

func (s *BlockelifsContext) SetBlkef(v []interface{}) { s.blkef = v }

func (s *BlockelifsContext) AllIfstmt() []IIfstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfstmtContext); ok {
			len++
		}
	}

	tst := make([]IIfstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfstmtContext); ok {
			tst[i] = t.(IIfstmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockelifsContext) Ifstmt(i int) IIfstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
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

	return t.(IIfstmtContext)
}

func (s *BlockelifsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockelifsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockelifsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBlockelifs(s)
	}
}

func (s *BlockelifsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBlockelifs(s)
	}
}

func (p *GrammarParser) Blockelifs() (localctx IBlockelifsContext) {
	localctx = NewBlockelifsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_blockelifs)

	localctx.(*BlockelifsContext).blkef = []interface{}{}
	var listifs []IIfstmtContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(189)

				var _x = p.Ifstmt()

				localctx.(*BlockelifsContext)._ifstmt = _x
			}
			localctx.(*BlockelifsContext).elif = append(localctx.(*BlockelifsContext).elif, localctx.(*BlockelifsContext)._ifstmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listifs = localctx.(*BlockelifsContext).GetElif()
	for _, a := range listifs {
		localctx.(*BlockelifsContext).blkef = append(localctx.(*BlockelifsContext).blkef, a.GetIft())

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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetIfb returns the ifb rule contexts.
	GetIfb() IBlockContext

	// GetIfelseblck returns the ifelseblck rule contexts.
	GetIfelseblck() IBlockContext

	// GetElseifblck returns the elseifblck rule contexts.
	GetElseifblck() IBlockContext

	// GetElif returns the elif rule contexts.
	GetElif() IBlockContext

	// Get_blockelifs returns the _blockelifs rule contexts.
	Get_blockelifs() IBlockelifsContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetIfb sets the ifb rule contexts.
	SetIfb(IBlockContext)

	// SetIfelseblck sets the ifelseblck rule contexts.
	SetIfelseblck(IBlockContext)

	// SetElseifblck sets the elseifblck rule contexts.
	SetElseifblck(IBlockContext)

	// SetElif sets the elif rule contexts.
	SetElif(IBlockContext)

	// Set_blockelifs sets the _blockelifs rule contexts.
	Set_blockelifs(IBlockelifsContext)

	// GetIft returns the ift attribute.
	GetIft() interfaces.Instruction

	// GetEl returns the el attribute.
	GetEl() []interface{}

	// GetElse_ returns the else_ attribute.
	GetElse_() interfaces.Instruction

	// SetIft sets the ift attribute.
	SetIft(interfaces.Instruction)

	// SetEl sets the el attribute.
	SetEl([]interface{})

	// SetElse_ sets the else_ attribute.
	SetElse_(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	Blockelifs() IBlockelifsContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	ift         interfaces.Instruction
	el          []interface{}
	else_       interfaces.Instruction
	_IF         antlr.Token
	_expr       IExprContext
	ifb         IBlockContext
	ifelseblck  IBlockContext
	elseifblck  IBlockContext
	elif        IBlockContext
	_blockelifs IBlockelifsContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) GetIfb() IBlockContext { return s.ifb }

func (s *IfstmtContext) GetIfelseblck() IBlockContext { return s.ifelseblck }

func (s *IfstmtContext) GetElseifblck() IBlockContext { return s.elseifblck }

func (s *IfstmtContext) GetElif() IBlockContext { return s.elif }

func (s *IfstmtContext) Get_blockelifs() IBlockelifsContext { return s._blockelifs }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) SetIfb(v IBlockContext) { s.ifb = v }

func (s *IfstmtContext) SetIfelseblck(v IBlockContext) { s.ifelseblck = v }

func (s *IfstmtContext) SetElseifblck(v IBlockContext) { s.elseifblck = v }

func (s *IfstmtContext) SetElif(v IBlockContext) { s.elif = v }

func (s *IfstmtContext) Set_blockelifs(v IBlockelifsContext) { s._blockelifs = v }

func (s *IfstmtContext) GetIft() interfaces.Instruction { return s.ift }

func (s *IfstmtContext) GetEl() []interface{} { return s.el }

func (s *IfstmtContext) GetElse_() interfaces.Instruction { return s.else_ }

func (s *IfstmtContext) SetIft(v interfaces.Instruction) { s.ift = v }

func (s *IfstmtContext) SetEl(v []interface{}) { s.el = v }

func (s *IfstmtContext) SetElse_(v interfaces.Instruction) { s.else_ = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
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

func (s *IfstmtContext) Block(i int) IBlockContext {
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

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfstmtContext) Blockelifs() IBlockelifsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockelifsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockelifsContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *GrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_ifstmt)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(198)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifb = _x
		}
		{
			p.SetState(200)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).SetIft(instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetIfb().GetBlk(), nil))
		localctx.(*IfstmtContext).el = localctx.(*IfstmtContext).GetIfb().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(205)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifelseblck = _x
		}
		{
			p.SetState(207)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)

			var _x = p.Block()

			localctx.(*IfstmtContext).elseifblck = _x
		}
		{
			p.SetState(211)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).SetIft(instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetIfelseblck().GetBlk(), localctx.(*IfstmtContext).GetElseifblck().GetBlk()))
		localctx.(*IfstmtContext).SetEl(localctx.(*IfstmtContext).GetIfelseblck().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(214)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(216)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)

			var _x = p.Block()

			localctx.(*IfstmtContext).elif = _x
		}
		{
			p.SetState(218)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)

			var _x = p.Blockelifs()

			localctx.(*IfstmtContext)._blockelifs = _x
		}
		localctx.(*IfstmtContext).SetIft(instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetElif().GetBlk(), localctx.(*IfstmtContext).Get_blockelifs().GetBlkef()))

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

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_casestmt returns the _casestmt rule contexts.
	Get_casestmt() ICasestmtContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_casestmt sets the _casestmt rule contexts.
	Set_casestmt(ICasestmtContext)

	// GetSwitch_ returns the switch_ attribute.
	GetSwitch_() interfaces.Instruction

	// SetSwitch_ sets the switch_ attribute.
	SetSwitch_(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	Casestmt() ICasestmtContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	switch_   interfaces.Instruction
	_SWITCH   antlr.Token
	_expr     IExprContext
	_block    IBlockContext
	_casestmt ICasestmtContext
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchstmtContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchstmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchstmtContext) Get_block() IBlockContext { return s._block }

func (s *SwitchstmtContext) Get_casestmt() ICasestmtContext { return s._casestmt }

func (s *SwitchstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *SwitchstmtContext) Set_casestmt(v ICasestmtContext) { s._casestmt = v }

func (s *SwitchstmtContext) GetSwitch_() interfaces.Instruction { return s.switch_ }

func (s *SwitchstmtContext) SetSwitch_(v interfaces.Instruction) { s.switch_ = v }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *SwitchstmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *SwitchstmtContext) Block() IBlockContext {
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

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) Casestmt() ICasestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasestmtContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *GrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_switchstmt)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(227)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(GrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)

			var _x = p.Block()

			localctx.(*SwitchstmtContext)._block = _x
		}
		{
			p.SetState(231)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).SetSwitch_(instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), nil, nil, localctx.(*SwitchstmtContext).Get_block().GetBlk()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(236)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)

			var _x = p.Casestmt()

			localctx.(*SwitchstmtContext)._casestmt = _x
		}
		{
			p.SetState(238)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).SetSwitch_(instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), nil, localctx.(*SwitchstmtContext).Get_casestmt().GetCas(), nil))

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

// ICasestmtContext is an interface to support dynamic dispatch.
type ICasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_blockcases returns the _blockcases rule contexts.
	Get_blockcases() IBlockcasesContext

	// GetSen returns the sen rule contexts.
	GetSen() IBlockContext

	// GetDef returns the def rule contexts.
	GetDef() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_blockcases sets the _blockcases rule contexts.
	Set_blockcases(IBlockcasesContext)

	// SetSen sets the sen rule contexts.
	SetSen(IBlockContext)

	// SetDef sets the def rule contexts.
	SetDef(IBlockContext)

	// GetCas returns the cas attribute.
	GetCas() interfaces.Instruction

	// SetCas sets the cas attribute.
	SetCas(interfaces.Instruction)

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	AllDOUBLEPTS() []antlr.TerminalNode
	DOUBLEPTS(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Blockcases() IBlockcasesContext
	DEFAULT() antlr.TerminalNode

	// IsCasestmtContext differentiates from other interfaces.
	IsCasestmtContext()
}

type CasestmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	cas         interfaces.Instruction
	_CASE       antlr.Token
	_expr       IExprContext
	_block      IBlockContext
	_blockcases IBlockcasesContext
	sen         IBlockContext
	def         IBlockContext
}

func NewEmptyCasestmtContext() *CasestmtContext {
	var p = new(CasestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_casestmt
	return p
}

func InitEmptyCasestmtContext(p *CasestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_casestmt
}

func (*CasestmtContext) IsCasestmtContext() {}

func NewCasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasestmtContext {
	var p = new(CasestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_casestmt

	return p
}

func (s *CasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CasestmtContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasestmtContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasestmtContext) Get_expr() IExprContext { return s._expr }

func (s *CasestmtContext) Get_block() IBlockContext { return s._block }

func (s *CasestmtContext) Get_blockcases() IBlockcasesContext { return s._blockcases }

func (s *CasestmtContext) GetSen() IBlockContext { return s.sen }

func (s *CasestmtContext) GetDef() IBlockContext { return s.def }

func (s *CasestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CasestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *CasestmtContext) Set_blockcases(v IBlockcasesContext) { s._blockcases = v }

func (s *CasestmtContext) SetSen(v IBlockContext) { s.sen = v }

func (s *CasestmtContext) SetDef(v IBlockContext) { s.def = v }

func (s *CasestmtContext) GetCas() interfaces.Instruction { return s.cas }

func (s *CasestmtContext) SetCas(v interfaces.Instruction) { s.cas = v }

func (s *CasestmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *CasestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CasestmtContext) AllDOUBLEPTS() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserDOUBLEPTS)
}

func (s *CasestmtContext) DOUBLEPTS(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, i)
}

func (s *CasestmtContext) AllBlock() []IBlockContext {
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

func (s *CasestmtContext) Block(i int) IBlockContext {
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

func (s *CasestmtContext) Blockcases() IBlockcasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockcasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockcasesContext)
}

func (s *CasestmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *CasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCasestmt(s)
	}
}

func (s *CasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCasestmt(s)
	}
}

func (p *GrammarParser) Casestmt() (localctx ICasestmtContext) {
	localctx = NewCasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_casestmt)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)

			var _m = p.Match(GrammarParserCASE)

			localctx.(*CasestmtContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)

			var _x = p.expr(0)

			localctx.(*CasestmtContext)._expr = _x
		}
		{
			p.SetState(245)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)

			var _x = p.Block()

			localctx.(*CasestmtContext)._block = _x
		}
		{
			p.SetState(247)

			var _x = p.Blockcases()

			localctx.(*CasestmtContext)._blockcases = _x
		}
		localctx.(*CasestmtContext).cas = instructions.NewCase((func() int {
			if localctx.(*CasestmtContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasestmtContext).Get_CASE().GetLine()
			}
		}()), (func() int {
			if localctx.(*CasestmtContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasestmtContext).Get_CASE().GetColumn()
			}
		}()), localctx.(*CasestmtContext).Get_expr().GetE(), localctx.(*CasestmtContext).Get_block().GetBlk(), localctx.(*CasestmtContext).Get_blockcases().GetBlkcase())
		fmt.Println("entro")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)

			var _m = p.Match(GrammarParserCASE)

			localctx.(*CasestmtContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.expr(0)

			localctx.(*CasestmtContext)._expr = _x
		}
		{
			p.SetState(252)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _x = p.Block()

			localctx.(*CasestmtContext).sen = _x
		}
		{
			p.SetState(254)
			p.Match(GrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)

			var _x = p.Block()

			localctx.(*CasestmtContext).def = _x
		}
		localctx.(*CasestmtContext).cas = instructions.NewCase((func() int {
			if localctx.(*CasestmtContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasestmtContext).Get_CASE().GetLine()
			}
		}()), (func() int {
			if localctx.(*CasestmtContext).Get_CASE() == nil {
				return 0
			} else {
				return localctx.(*CasestmtContext).Get_CASE().GetColumn()
			}
		}()), localctx.(*CasestmtContext).Get_expr().GetE(), localctx.(*CasestmtContext).GetSen().GetBlk(), localctx.(*CasestmtContext).GetDef().GetBlk())

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

// IBlockcasesContext is an interface to support dynamic dispatch.
type IBlockcasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_casestmt returns the _casestmt rule contexts.
	Get_casestmt() ICasestmtContext

	// Set_casestmt sets the _casestmt rule contexts.
	Set_casestmt(ICasestmtContext)

	// GetCasedef returns the casedef rule context list.
	GetCasedef() []ICasestmtContext

	// SetCasedef sets the casedef rule context list.
	SetCasedef([]ICasestmtContext)

	// GetBlkcase returns the blkcase attribute.
	GetBlkcase() []interface{}

	// SetBlkcase sets the blkcase attribute.
	SetBlkcase([]interface{})

	// Getter signatures
	AllCasestmt() []ICasestmtContext
	Casestmt(i int) ICasestmtContext

	// IsBlockcasesContext differentiates from other interfaces.
	IsBlockcasesContext()
}

type BlockcasesContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	blkcase   []interface{}
	_casestmt ICasestmtContext
	casedef   []ICasestmtContext
}

func NewEmptyBlockcasesContext() *BlockcasesContext {
	var p = new(BlockcasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockcases
	return p
}

func InitEmptyBlockcasesContext(p *BlockcasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockcases
}

func (*BlockcasesContext) IsBlockcasesContext() {}

func NewBlockcasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockcasesContext {
	var p = new(BlockcasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_blockcases

	return p
}

func (s *BlockcasesContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockcasesContext) Get_casestmt() ICasestmtContext { return s._casestmt }

func (s *BlockcasesContext) Set_casestmt(v ICasestmtContext) { s._casestmt = v }

func (s *BlockcasesContext) GetCasedef() []ICasestmtContext { return s.casedef }

func (s *BlockcasesContext) SetCasedef(v []ICasestmtContext) { s.casedef = v }

func (s *BlockcasesContext) GetBlkcase() []interface{} { return s.blkcase }

func (s *BlockcasesContext) SetBlkcase(v []interface{}) { s.blkcase = v }

func (s *BlockcasesContext) AllCasestmt() []ICasestmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICasestmtContext); ok {
			len++
		}
	}

	tst := make([]ICasestmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICasestmtContext); ok {
			tst[i] = t.(ICasestmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockcasesContext) Casestmt(i int) ICasestmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasestmtContext); ok {
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

	return t.(ICasestmtContext)
}

func (s *BlockcasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockcasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockcasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBlockcases(s)
	}
}

func (s *BlockcasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBlockcases(s)
	}
}

func (p *GrammarParser) Blockcases() (localctx IBlockcasesContext) {
	localctx = NewBlockcasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_blockcases)

	localctx.(*BlockcasesContext).blkcase = []interface{}{}
	var listcases []ICasestmtContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(261)

				var _x = p.Casestmt()

				localctx.(*BlockcasesContext)._casestmt = _x
			}
			localctx.(*BlockcasesContext).casedef = append(localctx.(*BlockcasesContext).casedef, localctx.(*BlockcasesContext)._casestmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listcases = localctx.(*BlockcasesContext).GetCasedef()
	for _, a := range listcases {
		localctx.(*BlockcasesContext).blkcase = append(localctx.(*BlockcasesContext).blkcase, a.GetCas())

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

// IAssignationstmtContext is an interface to support dynamic dispatch.
type IAssignationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_IG returns the _IG token.
	Get_IG() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_IG sets the _IG token.
	Set_IG(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAssign returns the assign attribute.
	GetAssign() interfaces.Instruction

	// SetAssign sets the assign attribute.
	SetAssign(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	PTCOMA() antlr.TerminalNode

	// IsAssignationstmtContext differentiates from other interfaces.
	IsAssignationstmtContext()
}

type AssignationstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	assign interfaces.Instruction
	_ID    antlr.Token
	_IG    antlr.Token
	_expr  IExprContext
}

func NewEmptyAssignationstmtContext() *AssignationstmtContext {
	var p = new(AssignationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignationstmt
	return p
}

func InitEmptyAssignationstmtContext(p *AssignationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignationstmt
}

func (*AssignationstmtContext) IsAssignationstmtContext() {}

func NewAssignationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignationstmtContext {
	var p = new(AssignationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_assignationstmt

	return p
}

func (s *AssignationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignationstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AssignationstmtContext) Get_IG() antlr.Token { return s._IG }

func (s *AssignationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AssignationstmtContext) Set_IG(v antlr.Token) { s._IG = v }

func (s *AssignationstmtContext) Get_expr() IExprContext { return s._expr }

func (s *AssignationstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AssignationstmtContext) GetAssign() interfaces.Instruction { return s.assign }

func (s *AssignationstmtContext) SetAssign(v interfaces.Instruction) { s.assign = v }

func (s *AssignationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AssignationstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *AssignationstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignationstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *AssignationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssignationstmt(s)
	}
}

func (s *AssignationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssignationstmt(s)
	}
}

func (p *GrammarParser) Assignationstmt() (localctx IAssignationstmtContext) {
	localctx = NewAssignationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_assignationstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)

		var _m = p.Match(GrammarParserID)

		localctx.(*AssignationstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)

		var _m = p.Match(GrammarParserIG)

		localctx.(*AssignationstmtContext)._IG = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)

		var _x = p.expr(0)

		localctx.(*AssignationstmtContext)._expr = _x
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(271)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*AssignationstmtContext).assign = instructions.NewAssignation((func() int {
		if localctx.(*AssignationstmtContext).Get_IG() == nil {
			return 0
		} else {
			return localctx.(*AssignationstmtContext).Get_IG().GetLine()
		}
	}()), (func() int {
		if localctx.(*AssignationstmtContext).Get_IG() == nil {
			return 0
		} else {
			return localctx.(*AssignationstmtContext).Get_IG().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AssignationstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AssignationstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*AssignationstmtContext).Get_expr().GetE())

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

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhile returns the while attribute.
	GetWhile() interfaces.Instruction

	// SetWhile sets the while attribute.
	SetWhile(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	while  interfaces.Instruction
	_WHILE antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhile() interfaces.Instruction { return s.while }

func (s *WhilestmtContext) SetWhile(v interfaces.Instruction) { s.while = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
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

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *GrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)

		var _m = p.Match(GrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(278)
		p.Match(GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(280)
		p.Match(GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).while = instructions.NewWhile((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

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

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetFirst returns the first rule contexts.
	GetFirst() IExprContext

	// GetSecond returns the second rule contexts.
	GetSecond() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetFirst sets the first rule contexts.
	SetFirst(IExprContext)

	// SetSecond sets the second rule contexts.
	SetSecond(IExprContext)

	// GetFor_ returns the for_ attribute.
	GetFor_() interfaces.Instruction

	// SetFor_ sets the for_ attribute.
	SetFor_(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	AllPOINT() []antlr.TerminalNode
	POINT(i int) antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	for_   interfaces.Instruction
	_FOR   antlr.Token
	_ID    antlr.Token
	_expr  IExprContext
	_block IBlockContext
	first  IExprContext
	second IExprContext
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) GetFirst() IExprContext { return s.first }

func (s *ForstmtContext) GetSecond() IExprContext { return s.second }

func (s *ForstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) SetFirst(v IExprContext) { s.first = v }

func (s *ForstmtContext) SetSecond(v IExprContext) { s.second = v }

func (s *ForstmtContext) GetFor_() interfaces.Instruction { return s.for_ }

func (s *ForstmtContext) SetFor_(v interfaces.Instruction) { s.for_ = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ForstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
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

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) AllPOINT() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserPOINT)
}

func (s *ForstmtContext) POINT(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, i)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *GrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_forstmt)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(283)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.expr(0)

			localctx.(*ForstmtContext)._expr = _x
		}
		{
			p.SetState(287)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(289)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).for_ = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ForstmtContext).Get_expr().GetE(), nil, localctx.(*ForstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).first = _x
		}
		{
			p.SetState(296)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).second = _x
		}
		{
			p.SetState(299)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(301)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).for_ = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ForstmtContext).GetFirst().GetE(), localctx.(*ForstmtContext).GetSecond().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

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

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// GetR returns the r token.
	GetR() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// SetR sets the r token.
	SetR(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGua returns the gua attribute.
	GetGua() interfaces.Instruction

	// SetGua sets the gua attribute.
	SetGua(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	BREAK() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	gua    interfaces.Instruction
	_GUARD antlr.Token
	_expr  IExprContext
	_block IBlockContext
	r      antlr.Token
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardstmtContext) GetR() antlr.Token { return s.r }

func (s *GuardstmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardstmtContext) SetR(v antlr.Token) { s.r = v }

func (s *GuardstmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardstmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardstmtContext) GetGua() interfaces.Instruction { return s.gua }

func (s *GuardstmtContext) SetGua(v interfaces.Instruction) { s.gua = v }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(GrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
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

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *GuardstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *GuardstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (p *GrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_guardstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)

		var _m = p.Match(GrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(308)
		p.Match(GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(311)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*GuardstmtContext).r = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14680064) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*GuardstmtContext).r = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(312)
		p.Match(GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardstmtContext).gua = instructions.NewGuard((func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardstmtContext).Get_expr().GetE(), localctx.(*GuardstmtContext).Get_block().GetBlk(), (func() string {
		if localctx.(*GuardstmtContext).GetR() == nil {
			return ""
		} else {
			return localctx.(*GuardstmtContext).GetR().GetText()
		}
	}()))

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

// ITransferstmtContext is an interface to support dynamic dispatch.
type ITransferstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetTran returns the tran attribute.
	GetTran() interfaces.Instruction

	// SetTran sets the tran attribute.
	SetTran(interfaces.Instruction)

	// Getter signatures
	RETURN() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode
	Expr() IExprContext
	CONTINUE() antlr.TerminalNode
	BREAK() antlr.TerminalNode

	// IsTransferstmtContext differentiates from other interfaces.
	IsTransferstmtContext()
}

type TransferstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tran      interfaces.Instruction
	_RETURN   antlr.Token
	_expr     IExprContext
	_CONTINUE antlr.Token
	_BREAK    antlr.Token
}

func NewEmptyTransferstmtContext() *TransferstmtContext {
	var p = new(TransferstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_transferstmt
	return p
}

func InitEmptyTransferstmtContext(p *TransferstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_transferstmt
}

func (*TransferstmtContext) IsTransferstmtContext() {}

func NewTransferstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferstmtContext {
	var p = new(TransferstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_transferstmt

	return p
}

func (s *TransferstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferstmtContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *TransferstmtContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *TransferstmtContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *TransferstmtContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *TransferstmtContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *TransferstmtContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *TransferstmtContext) Get_expr() IExprContext { return s._expr }

func (s *TransferstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *TransferstmtContext) GetTran() interfaces.Instruction { return s.tran }

func (s *TransferstmtContext) SetTran(v interfaces.Instruction) { s.tran = v }

func (s *TransferstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *TransferstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *TransferstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TransferstmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *TransferstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *TransferstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTransferstmt(s)
	}
}

func (s *TransferstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTransferstmt(s)
	}
}

func (p *GrammarParser) Transferstmt() (localctx ITransferstmtContext) {
	localctx = NewTransferstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_transferstmt)
	var _la int

	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)

			var _m = p.Match(GrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(316)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*TransferstmtContext).tran = instructions.NewReturnIn((func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetColumn()
			}
		}()), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)

			var _m = p.Match(GrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)

			var _x = p.expr(0)

			localctx.(*TransferstmtContext)._expr = _x
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(322)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*TransferstmtContext).tran = instructions.NewReturnIn((func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetColumn()
			}
		}()), localctx.(*TransferstmtContext).Get_expr().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)

			var _m = p.Match(GrammarParserCONTINUE)

			localctx.(*TransferstmtContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(328)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*TransferstmtContext).tran = instructions.NewContinue((func() int {
			if localctx.(*TransferstmtContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_CONTINUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_CONTINUE().GetColumn()
			}
		}()))

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(332)

			var _m = p.Match(GrammarParserBREAK)

			localctx.(*TransferstmtContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(333)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*TransferstmtContext).tran = instructions.NewBreak((func() int {
			if localctx.(*TransferstmtContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_BREAK().GetColumn()
			}
		}()))

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

// IDeclarevectorstmtContext is an interface to support dynamic dispatch.
type IDeclarevectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetF returns the F token.
	GetF() antlr.Token

	// GetS returns the S token.
	GetS() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetF sets the F token.
	SetF(antlr.Token)

	// SetS sets the S token.
	SetS(antlr.Token)

	// Get_type returns the _type rule contexts.
	Get_type() ITypeContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_type sets the _type rule contexts.
	Set_type(ITypeContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetDecvec returns the decvec attribute.
	GetDecvec() interfaces.Instruction

	// SetDecvec sets the decvec attribute.
	SetDecvec(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	AllCORCHETEIZQ() []antlr.TerminalNode
	CORCHETEIZQ(i int) antlr.TerminalNode
	Type_() ITypeContext
	AllCORCHETEDER() []antlr.TerminalNode
	CORCHETEDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	ListParams() IListParamsContext
	PTCOMA() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsDeclarevectorstmtContext differentiates from other interfaces.
	IsDeclarevectorstmtContext()
}

type DeclarevectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	decvec      interfaces.Instruction
	_VAR        antlr.Token
	_ID         antlr.Token
	_type       ITypeContext
	_listParams IListParamsContext
	F           antlr.Token
	S           antlr.Token
}

func NewEmptyDeclarevectorstmtContext() *DeclarevectorstmtContext {
	var p = new(DeclarevectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarevectorstmt
	return p
}

func InitEmptyDeclarevectorstmtContext(p *DeclarevectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarevectorstmt
}

func (*DeclarevectorstmtContext) IsDeclarevectorstmtContext() {}

func NewDeclarevectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarevectorstmtContext {
	var p = new(DeclarevectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_declarevectorstmt

	return p
}

func (s *DeclarevectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarevectorstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclarevectorstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarevectorstmtContext) GetF() antlr.Token { return s.F }

func (s *DeclarevectorstmtContext) GetS() antlr.Token { return s.S }

func (s *DeclarevectorstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarevectorstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarevectorstmtContext) SetF(v antlr.Token) { s.F = v }

func (s *DeclarevectorstmtContext) SetS(v antlr.Token) { s.S = v }

func (s *DeclarevectorstmtContext) Get_type() ITypeContext { return s._type }

func (s *DeclarevectorstmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *DeclarevectorstmtContext) Set_type(v ITypeContext) { s._type = v }

func (s *DeclarevectorstmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *DeclarevectorstmtContext) GetDecvec() interfaces.Instruction { return s.decvec }

func (s *DeclarevectorstmtContext) SetDecvec(v interfaces.Instruction) { s.decvec = v }

func (s *DeclarevectorstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclarevectorstmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *DeclarevectorstmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *DeclarevectorstmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *DeclarevectorstmtContext) AllCORCHETEIZQ() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCORCHETEIZQ)
}

func (s *DeclarevectorstmtContext) CORCHETEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEIZQ, i)
}

func (s *DeclarevectorstmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarevectorstmtContext) AllCORCHETEDER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCORCHETEDER)
}

func (s *DeclarevectorstmtContext) CORCHETEDER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEDER, i)
}

func (s *DeclarevectorstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *DeclarevectorstmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *DeclarevectorstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *DeclarevectorstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *DeclarevectorstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *DeclarevectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarevectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarevectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclarevectorstmt(s)
	}
}

func (s *DeclarevectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclarevectorstmt(s)
	}
}

func (p *GrammarParser) Declarevectorstmt() (localctx IDeclarevectorstmtContext) {
	localctx = NewDeclarevectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_declarevectorstmt)
	var _la int

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarevectorstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)

			var _x = p.Type_()

			localctx.(*DeclarevectorstmtContext)._type = _x
		}
		{
			p.SetState(344)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)

			var _x = p.listParams(0)

			localctx.(*DeclarevectorstmtContext)._listParams = _x
		}
		{
			p.SetState(348)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(349)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarevectorstmtContext).decvec = instructions.NewToDeclareVector((func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarevectorstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarevectorstmtContext).Get_type().GetT(), localctx.(*DeclarevectorstmtContext).Get_listParams().GetL(), "")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarevectorstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)

			var _x = p.Type_()

			localctx.(*DeclarevectorstmtContext)._type = _x
		}
		{
			p.SetState(359)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(363)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarevectorstmtContext).decvec = instructions.NewToDeclareVector((func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarevectorstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarevectorstmtContext).Get_type().GetT(), nil, "")

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Type_()
		}
		{
			p.SetState(373)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(376)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(379)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarevectorstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext).F = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)

			var _x = p.Type_()

			localctx.(*DeclarevectorstmtContext)._type = _x
		}
		{
			p.SetState(384)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext).S = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(387)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*DeclarevectorstmtContext).decvec = instructions.NewToDeclareVector((func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarevectorstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarevectorstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarevectorstmtContext).GetF() == nil {
				return ""
			} else {
				return localctx.(*DeclarevectorstmtContext).GetF().GetText()
			}
		}()), localctx.(*DeclarevectorstmtContext).Get_type().GetT(), nil, (func() string {
			if localctx.(*DeclarevectorstmtContext).GetS() == nil {
				return ""
			} else {
				return localctx.(*DeclarevectorstmtContext).GetS().GetText()
			}
		}()))

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

// IAppendstmtContext is an interface to support dynamic dispatch.
type IAppendstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetApp returns the app attribute.
	GetApp() interfaces.Instruction

	// SetApp sets the app attribute.
	SetApp(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	POINT() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsAppendstmtContext differentiates from other interfaces.
	IsAppendstmtContext()
}

type AppendstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	app    interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyAppendstmtContext() *AppendstmtContext {
	var p = new(AppendstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_appendstmt
	return p
}

func InitEmptyAppendstmtContext(p *AppendstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_appendstmt
}

func (*AppendstmtContext) IsAppendstmtContext() {}

func NewAppendstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendstmtContext {
	var p = new(AppendstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_appendstmt

	return p
}

func (s *AppendstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AppendstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AppendstmtContext) Get_expr() IExprContext { return s._expr }

func (s *AppendstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AppendstmtContext) GetApp() interfaces.Instruction { return s.app }

func (s *AppendstmtContext) SetApp(v interfaces.Instruction) { s.app = v }

func (s *AppendstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AppendstmtContext) POINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, 0)
}

func (s *AppendstmtContext) APPEND() antlr.TerminalNode {
	return s.GetToken(GrammarParserAPPEND, 0)
}

func (s *AppendstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *AppendstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppendstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *AppendstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *AppendstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAppendstmt(s)
	}
}

func (s *AppendstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAppendstmt(s)
	}
}

func (p *GrammarParser) Appendstmt() (localctx IAppendstmtContext) {
	localctx = NewAppendstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_appendstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)

		var _m = p.Match(GrammarParserID)

		localctx.(*AppendstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(GrammarParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)

		var _x = p.expr(0)

		localctx.(*AppendstmtContext)._expr = _x
	}
	{
		p.SetState(399)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(400)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*AppendstmtContext).app = instructions.NewAppend((func() int {
		if localctx.(*AppendstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendstmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AppendstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendstmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AppendstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AppendstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*AppendstmtContext).Get_expr().GetE())

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

// IRemovelaststmtContext is an interface to support dynamic dispatch.
type IRemovelaststmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetRemovl returns the removl attribute.
	GetRemovl() interfaces.Instruction

	// SetRemovl sets the removl attribute.
	SetRemovl(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	POINT() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsRemovelaststmtContext differentiates from other interfaces.
	IsRemovelaststmtContext()
}

type RemovelaststmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	removl interfaces.Instruction
	_ID    antlr.Token
}

func NewEmptyRemovelaststmtContext() *RemovelaststmtContext {
	var p = new(RemovelaststmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_removelaststmt
	return p
}

func InitEmptyRemovelaststmtContext(p *RemovelaststmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_removelaststmt
}

func (*RemovelaststmtContext) IsRemovelaststmtContext() {}

func NewRemovelaststmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovelaststmtContext {
	var p = new(RemovelaststmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_removelaststmt

	return p
}

func (s *RemovelaststmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovelaststmtContext) Get_ID() antlr.Token { return s._ID }

func (s *RemovelaststmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemovelaststmtContext) GetRemovl() interfaces.Instruction { return s.removl }

func (s *RemovelaststmtContext) SetRemovl(v interfaces.Instruction) { s.removl = v }

func (s *RemovelaststmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *RemovelaststmtContext) POINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, 0)
}

func (s *RemovelaststmtContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(GrammarParserREMOVELAST, 0)
}

func (s *RemovelaststmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *RemovelaststmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *RemovelaststmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *RemovelaststmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovelaststmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovelaststmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRemovelaststmt(s)
	}
}

func (s *RemovelaststmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRemovelaststmt(s)
	}
}

func (p *GrammarParser) Removelaststmt() (localctx IRemovelaststmtContext) {
	localctx = NewRemovelaststmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_removelaststmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)

		var _m = p.Match(GrammarParserID)

		localctx.(*RemovelaststmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(GrammarParserREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(410)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*RemovelaststmtContext).removl = instructions.NewRemoveLast((func() int {
		if localctx.(*RemovelaststmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelaststmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemovelaststmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelaststmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemovelaststmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemovelaststmtContext).Get_ID().GetText()
		}
	}()))

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

// IRemovestmtContext is an interface to support dynamic dispatch.
type IRemovestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetRemov returns the remov attribute.
	GetRemov() interfaces.Instruction

	// SetRemov sets the remov attribute.
	SetRemov(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	POINT() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AT() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsRemovestmtContext differentiates from other interfaces.
	IsRemovestmtContext()
}

type RemovestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	remov  interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyRemovestmtContext() *RemovestmtContext {
	var p = new(RemovestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_removestmt
	return p
}

func InitEmptyRemovestmtContext(p *RemovestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_removestmt
}

func (*RemovestmtContext) IsRemovestmtContext() {}

func NewRemovestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovestmtContext {
	var p = new(RemovestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_removestmt

	return p
}

func (s *RemovestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovestmtContext) Get_ID() antlr.Token { return s._ID }

func (s *RemovestmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemovestmtContext) Get_expr() IExprContext { return s._expr }

func (s *RemovestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *RemovestmtContext) GetRemov() interfaces.Instruction { return s.remov }

func (s *RemovestmtContext) SetRemov(v interfaces.Instruction) { s.remov = v }

func (s *RemovestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *RemovestmtContext) POINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, 0)
}

func (s *RemovestmtContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(GrammarParserREMOVE, 0)
}

func (s *RemovestmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *RemovestmtContext) AT() antlr.TerminalNode {
	return s.GetToken(GrammarParserAT, 0)
}

func (s *RemovestmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *RemovestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RemovestmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *RemovestmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *RemovestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRemovestmt(s)
	}
}

func (s *RemovestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRemovestmt(s)
	}
}

func (p *GrammarParser) Removestmt() (localctx IRemovestmtContext) {
	localctx = NewRemovestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_removestmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)

		var _m = p.Match(GrammarParserID)

		localctx.(*RemovestmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
		p.Match(GrammarParserREMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(GrammarParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(GrammarParserDOUBLEPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)

		var _x = p.expr(0)

		localctx.(*RemovestmtContext)._expr = _x
	}
	{
		p.SetState(422)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(423)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*RemovestmtContext).remov = instructions.NewRemove((func() int {
		if localctx.(*RemovestmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovestmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemovestmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovestmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemovestmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemovestmtContext).Get_ID().GetText()
		}
	}()), localctx.(*RemovestmtContext).Get_expr().GetE())

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

// IEmptvecstmtContext is an interface to support dynamic dispatch.
type IEmptvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetEmptyvec returns the emptyvec attribute.
	GetEmptyvec() interfaces.Expression

	// SetEmptyvec sets the emptyvec attribute.
	SetEmptyvec(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	POINT() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode

	// IsEmptvecstmtContext differentiates from other interfaces.
	IsEmptvecstmtContext()
}

type EmptvecstmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	emptyvec interfaces.Expression
	_ID      antlr.Token
}

func NewEmptyEmptvecstmtContext() *EmptvecstmtContext {
	var p = new(EmptvecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_emptvecstmt
	return p
}

func InitEmptyEmptvecstmtContext(p *EmptvecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_emptvecstmt
}

func (*EmptvecstmtContext) IsEmptvecstmtContext() {}

func NewEmptvecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptvecstmtContext {
	var p = new(EmptvecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_emptvecstmt

	return p
}

func (s *EmptvecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptvecstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *EmptvecstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *EmptvecstmtContext) GetEmptyvec() interfaces.Expression { return s.emptyvec }

func (s *EmptvecstmtContext) SetEmptyvec(v interfaces.Expression) { s.emptyvec = v }

func (s *EmptvecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *EmptvecstmtContext) POINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, 0)
}

func (s *EmptvecstmtContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(GrammarParserISEMPTY, 0)
}

func (s *EmptvecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptvecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptvecstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterEmptvecstmt(s)
	}
}

func (s *EmptvecstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitEmptvecstmt(s)
	}
}

func (p *GrammarParser) Emptvecstmt() (localctx IEmptvecstmtContext) {
	localctx = NewEmptvecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_emptvecstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)

		var _m = p.Match(GrammarParserID)

		localctx.(*EmptvecstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(430)
		p.Match(GrammarParserISEMPTY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*EmptvecstmtContext).emptyvec = expressions.NewEmptyVector((func() int {
		if localctx.(*EmptvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*EmptvecstmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*EmptvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*EmptvecstmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*EmptvecstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*EmptvecstmtContext).Get_ID().GetText()
		}
	}()))

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

// ICountvecstmtContext is an interface to support dynamic dispatch.
type ICountvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetCount returns the count attribute.
	GetCount() interfaces.Expression

	// SetCount sets the count attribute.
	SetCount(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	POINT() antlr.TerminalNode
	COUNT() antlr.TerminalNode

	// IsCountvecstmtContext differentiates from other interfaces.
	IsCountvecstmtContext()
}

type CountvecstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	count  interfaces.Expression
	_ID    antlr.Token
}

func NewEmptyCountvecstmtContext() *CountvecstmtContext {
	var p = new(CountvecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_countvecstmt
	return p
}

func InitEmptyCountvecstmtContext(p *CountvecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_countvecstmt
}

func (*CountvecstmtContext) IsCountvecstmtContext() {}

func NewCountvecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountvecstmtContext {
	var p = new(CountvecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_countvecstmt

	return p
}

func (s *CountvecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CountvecstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *CountvecstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CountvecstmtContext) GetCount() interfaces.Expression { return s.count }

func (s *CountvecstmtContext) SetCount(v interfaces.Expression) { s.count = v }

func (s *CountvecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *CountvecstmtContext) POINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPOINT, 0)
}

func (s *CountvecstmtContext) COUNT() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOUNT, 0)
}

func (s *CountvecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountvecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountvecstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCountvecstmt(s)
	}
}

func (s *CountvecstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCountvecstmt(s)
	}
}

func (p *GrammarParser) Countvecstmt() (localctx ICountvecstmtContext) {
	localctx = NewCountvecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_countvecstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)

		var _m = p.Match(GrammarParserID)

		localctx.(*CountvecstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)
		p.Match(GrammarParserCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CountvecstmtContext).count = expressions.NewCount((func() int {
		if localctx.(*CountvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CountvecstmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CountvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CountvecstmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CountvecstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CountvecstmtContext).Get_ID().GetText()
		}
	}()))

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

// IAccessvecstmtContext is an interface to support dynamic dispatch.
type IAccessvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAccessvec returns the accessvec attribute.
	GetAccessvec() interfaces.Expression

	// SetAccessvec sets the accessvec attribute.
	SetAccessvec(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	CORCHETEIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORCHETEDER() antlr.TerminalNode

	// IsAccessvecstmtContext differentiates from other interfaces.
	IsAccessvecstmtContext()
}

type AccessvecstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	accessvec interfaces.Expression
	_ID       antlr.Token
	_expr     IExprContext
}

func NewEmptyAccessvecstmtContext() *AccessvecstmtContext {
	var p = new(AccessvecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_accessvecstmt
	return p
}

func InitEmptyAccessvecstmtContext(p *AccessvecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_accessvecstmt
}

func (*AccessvecstmtContext) IsAccessvecstmtContext() {}

func NewAccessvecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessvecstmtContext {
	var p = new(AccessvecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_accessvecstmt

	return p
}

func (s *AccessvecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessvecstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AccessvecstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AccessvecstmtContext) Get_expr() IExprContext { return s._expr }

func (s *AccessvecstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AccessvecstmtContext) GetAccessvec() interfaces.Expression { return s.accessvec }

func (s *AccessvecstmtContext) SetAccessvec(v interfaces.Expression) { s.accessvec = v }

func (s *AccessvecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AccessvecstmtContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEIZQ, 0)
}

func (s *AccessvecstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AccessvecstmtContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEDER, 0)
}

func (s *AccessvecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessvecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessvecstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAccessvecstmt(s)
	}
}

func (s *AccessvecstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAccessvecstmt(s)
	}
}

func (p *GrammarParser) Accessvecstmt() (localctx IAccessvecstmtContext) {
	localctx = NewAccessvecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrammarParserRULE_accessvecstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)

		var _m = p.Match(GrammarParserID)

		localctx.(*AccessvecstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Match(GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)

		var _x = p.expr(0)

		localctx.(*AccessvecstmtContext)._expr = _x
	}
	{
		p.SetState(441)
		p.Match(GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*AccessvecstmtContext).accessvec = expressions.NewAccessVector((func() int {
		if localctx.(*AccessvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AccessvecstmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AccessvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AccessvecstmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AccessvecstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AccessvecstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*AccessvecstmtContext).Get_expr().GetE())

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

// IAssignationvecstmtContext is an interface to support dynamic dispatch.
type IAssignationvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetExpprim returns the expprim rule contexts.
	GetExpprim() IExprContext

	// GetExpsegundo returns the expsegundo rule contexts.
	GetExpsegundo() IExprContext

	// SetExpprim sets the expprim rule contexts.
	SetExpprim(IExprContext)

	// SetExpsegundo sets the expsegundo rule contexts.
	SetExpsegundo(IExprContext)

	// GetAssignvec returns the assignvec attribute.
	GetAssignvec() interfaces.Instruction

	// SetAssignvec sets the assignvec attribute.
	SetAssignvec(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	CORCHETEIZQ() antlr.TerminalNode
	CORCHETEDER() antlr.TerminalNode
	IG() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PTCOMA() antlr.TerminalNode

	// IsAssignationvecstmtContext differentiates from other interfaces.
	IsAssignationvecstmtContext()
}

type AssignationvecstmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	assignvec  interfaces.Instruction
	_ID        antlr.Token
	expprim    IExprContext
	expsegundo IExprContext
}

func NewEmptyAssignationvecstmtContext() *AssignationvecstmtContext {
	var p = new(AssignationvecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignationvecstmt
	return p
}

func InitEmptyAssignationvecstmtContext(p *AssignationvecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignationvecstmt
}

func (*AssignationvecstmtContext) IsAssignationvecstmtContext() {}

func NewAssignationvecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignationvecstmtContext {
	var p = new(AssignationvecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_assignationvecstmt

	return p
}

func (s *AssignationvecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignationvecstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AssignationvecstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AssignationvecstmtContext) GetExpprim() IExprContext { return s.expprim }

func (s *AssignationvecstmtContext) GetExpsegundo() IExprContext { return s.expsegundo }

func (s *AssignationvecstmtContext) SetExpprim(v IExprContext) { s.expprim = v }

func (s *AssignationvecstmtContext) SetExpsegundo(v IExprContext) { s.expsegundo = v }

func (s *AssignationvecstmtContext) GetAssignvec() interfaces.Instruction { return s.assignvec }

func (s *AssignationvecstmtContext) SetAssignvec(v interfaces.Instruction) { s.assignvec = v }

func (s *AssignationvecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AssignationvecstmtContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEIZQ, 0)
}

func (s *AssignationvecstmtContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserCORCHETEDER, 0)
}

func (s *AssignationvecstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *AssignationvecstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AssignationvecstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AssignationvecstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *AssignationvecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignationvecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignationvecstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssignationvecstmt(s)
	}
}

func (s *AssignationvecstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssignationvecstmt(s)
	}
}

func (p *GrammarParser) Assignationvecstmt() (localctx IAssignationvecstmtContext) {
	localctx = NewAssignationvecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrammarParserRULE_assignationvecstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)

		var _m = p.Match(GrammarParserID)

		localctx.(*AssignationvecstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(446)

		var _x = p.expr(0)

		localctx.(*AssignationvecstmtContext).expprim = _x
	}
	{
		p.SetState(447)
		p.Match(GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
		p.Match(GrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)

		var _x = p.expr(0)

		localctx.(*AssignationvecstmtContext).expsegundo = _x
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(450)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*AssignationvecstmtContext).assignvec = instructions.NewAssignationVector((func() int {
		if localctx.(*AssignationvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AssignationvecstmtContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AssignationvecstmtContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AssignationvecstmtContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AssignationvecstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AssignationvecstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*AssignationvecstmtContext).GetExpprim().GetE(), localctx.(*AssignationvecstmtContext).GetExpsegundo().GetE())

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

// IDeclarematrixstmtContext is an interface to support dynamic dispatch.
type IDeclarematrixstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDecmatrix returns the decmatrix attribute.
	GetDecmatrix() interfaces.Instruction

	// SetDecmatrix sets the decmatrix attribute.
	SetDecmatrix(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	Type_() ITypeContext
	PARDER() antlr.TerminalNode
	IG() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsDeclarematrixstmtContext differentiates from other interfaces.
	IsDeclarematrixstmtContext()
}

type DeclarematrixstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	decmatrix interfaces.Instruction
}

func NewEmptyDeclarematrixstmtContext() *DeclarematrixstmtContext {
	var p = new(DeclarematrixstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarematrixstmt
	return p
}

func InitEmptyDeclarematrixstmtContext(p *DeclarematrixstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarematrixstmt
}

func (*DeclarematrixstmtContext) IsDeclarematrixstmtContext() {}

func NewDeclarematrixstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarematrixstmtContext {
	var p = new(DeclarematrixstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_declarematrixstmt

	return p
}

func (s *DeclarematrixstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarematrixstmtContext) GetDecmatrix() interfaces.Instruction { return s.decmatrix }

func (s *DeclarematrixstmtContext) SetDecmatrix(v interfaces.Instruction) { s.decmatrix = v }

func (s *DeclarematrixstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclarematrixstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclarematrixstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *DeclarematrixstmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *DeclarematrixstmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarematrixstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *DeclarematrixstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG, 0)
}

func (s *DeclarematrixstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *DeclarematrixstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarematrixstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarematrixstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclarematrixstmt(s)
	}
}

func (s *DeclarematrixstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclarematrixstmt(s)
	}
}

func (p *GrammarParser) Declarematrixstmt() (localctx IDeclarematrixstmtContext) {
	localctx = NewDeclarematrixstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GrammarParserRULE_declarematrixstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(GrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.Match(GrammarParserDOUBLEPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(459)
		p.Type_()
	}
	{
		p.SetState(460)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.Match(GrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserPTCOMA {
		{
			p.SetState(462)
			p.Match(GrammarParserPTCOMA)
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

// IDeclarefuncstmtContext is an interface to support dynamic dispatch.
type IDeclarefuncstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Get_type returns the _type rule contexts.
	Get_type() ITypeContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// Set_type sets the _type rule contexts.
	Set_type(ITypeContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetDecfunc returns the decfunc attribute.
	GetDecfunc() interfaces.Instruction

	// SetDecfunc sets the decfunc attribute.
	SetDecfunc(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	PARDER() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	Type_() ITypeContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsDeclarefuncstmtContext differentiates from other interfaces.
	IsDeclarefuncstmtContext()
}

type DeclarefuncstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	decfunc     interfaces.Instruction
	_ID         antlr.Token
	_listParams IListParamsContext
	_type       ITypeContext
	_block      IBlockContext
}

func NewEmptyDeclarefuncstmtContext() *DeclarefuncstmtContext {
	var p = new(DeclarefuncstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarefuncstmt
	return p
}

func InitEmptyDeclarefuncstmtContext(p *DeclarefuncstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declarefuncstmt
}

func (*DeclarefuncstmtContext) IsDeclarefuncstmtContext() {}

func NewDeclarefuncstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarefuncstmtContext {
	var p = new(DeclarefuncstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_declarefuncstmt

	return p
}

func (s *DeclarefuncstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarefuncstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarefuncstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarefuncstmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *DeclarefuncstmtContext) Get_type() ITypeContext { return s._type }

func (s *DeclarefuncstmtContext) Get_block() IBlockContext { return s._block }

func (s *DeclarefuncstmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *DeclarefuncstmtContext) Set_type(v ITypeContext) { s._type = v }

func (s *DeclarefuncstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *DeclarefuncstmtContext) GetDecfunc() interfaces.Instruction { return s.decfunc }

func (s *DeclarefuncstmtContext) SetDecfunc(v interfaces.Instruction) { s.decfunc = v }

func (s *DeclarefuncstmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GrammarParserFUNC, 0)
}

func (s *DeclarefuncstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclarefuncstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *DeclarefuncstmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *DeclarefuncstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *DeclarefuncstmtContext) SUB() antlr.TerminalNode {
	return s.GetToken(GrammarParserSUB, 0)
}

func (s *DeclarefuncstmtContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserMAYOR, 0)
}

func (s *DeclarefuncstmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarefuncstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *DeclarefuncstmtContext) Block() IBlockContext {
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

func (s *DeclarefuncstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *DeclarefuncstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarefuncstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarefuncstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclarefuncstmt(s)
	}
}

func (s *DeclarefuncstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclarefuncstmt(s)
	}
}

func (p *GrammarParser) Declarefuncstmt() (localctx IDeclarefuncstmtContext) {
	localctx = NewDeclarefuncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GrammarParserRULE_declarefuncstmt)
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.Match(GrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarefuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)

			var _x = p.listParams(0)

			localctx.(*DeclarefuncstmtContext)._listParams = _x
		}
		{
			p.SetState(469)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Match(GrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(GrammarParserMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)

			var _x = p.Type_()

			localctx.(*DeclarefuncstmtContext)._type = _x
		}
		{
			p.SetState(473)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)

			var _x = p.Block()

			localctx.(*DeclarefuncstmtContext)._block = _x
		}
		{
			p.SetState(475)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarefuncstmtContext).decfunc = instructions.NewToDeclareFunction((func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarefuncstmtContext).Get_listParams().GetL(), localctx.(*DeclarefuncstmtContext).Get_type().GetT(), localctx.(*DeclarefuncstmtContext).Get_block().GetBlk(), 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Match(GrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarefuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)

			var _x = p.listParams(0)

			localctx.(*DeclarefuncstmtContext)._listParams = _x
		}
		{
			p.SetState(482)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)

			var _x = p.Block()

			localctx.(*DeclarefuncstmtContext)._block = _x
		}
		{
			p.SetState(485)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarefuncstmtContext).decfunc = instructions.NewToDeclareFunction((func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarefuncstmtContext).Get_listParams().GetL(), environment.NULL, localctx.(*DeclarefuncstmtContext).Get_block().GetBlk(), 2)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.Match(GrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(489)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarefuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Match(GrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(GrammarParserMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)

			var _x = p.Type_()

			localctx.(*DeclarefuncstmtContext)._type = _x
		}
		{
			p.SetState(495)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)

			var _x = p.Block()

			localctx.(*DeclarefuncstmtContext)._block = _x
		}
		{
			p.SetState(497)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarefuncstmtContext).decfunc = instructions.NewToDeclareFunction((func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarefuncstmtContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*DeclarefuncstmtContext).Get_type().GetT(), localctx.(*DeclarefuncstmtContext).Get_block().GetBlk(), 3)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(500)
			p.Match(GrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarefuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)

			var _x = p.Block()

			localctx.(*DeclarefuncstmtContext)._block = _x
		}
		{
			p.SetState(506)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if localctx.(*DeclarefuncstmtContext).Get_block().GetBlk() != nil {
			localctx.(*DeclarefuncstmtContext).decfunc = instructions.NewToDeclareFunction((func() int {
				if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*DeclarefuncstmtContext).Get_ID().GetLine()
				}
			}()), (func() int {
				if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*DeclarefuncstmtContext).Get_ID().GetColumn()
				}
			}()), (func() string {
				if localctx.(*DeclarefuncstmtContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*DeclarefuncstmtContext).Get_ID().GetText()
				}
			}()), nil, environment.NULL, localctx.(*DeclarefuncstmtContext).Get_block().GetBlk(), 4)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_accessstmt returns the _accessstmt rule contexts.
	Get_accessstmt() IAccessstmtContext

	// Get_emptvecstmt returns the _emptvecstmt rule contexts.
	Get_emptvecstmt() IEmptvecstmtContext

	// Get_countvecstmt returns the _countvecstmt rule contexts.
	Get_countvecstmt() ICountvecstmtContext

	// Get_accessvecstmt returns the _accessvecstmt rule contexts.
	Get_accessvecstmt() IAccessvecstmtContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_accessstmt sets the _accessstmt rule contexts.
	Set_accessstmt(IAccessstmtContext)

	// Set_emptvecstmt sets the _emptvecstmt rule contexts.
	Set_emptvecstmt(IEmptvecstmtContext)

	// Set_countvecstmt sets the _countvecstmt rule contexts.
	Set_countvecstmt(ICountvecstmtContext)

	// Set_accessvecstmt sets the _accessvecstmt rule contexts.
	Set_accessvecstmt(IAccessvecstmtContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	NOT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	SUB() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NIL() antlr.TerminalNode
	Accessstmt() IAccessstmtContext
	Emptvecstmt() IEmptvecstmtContext
	Countvecstmt() ICountvecstmtContext
	Accessvecstmt() IAccessvecstmtContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	PERCENT() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	e              interfaces.Expression
	left           IExprContext
	op             antlr.Token
	_expr          IExprContext
	_NUMBER        antlr.Token
	_STRING        antlr.Token
	_TRUE          antlr.Token
	_FALSE         antlr.Token
	_NIL           antlr.Token
	_accessstmt    IAccessstmtContext
	_emptvecstmt   IEmptvecstmtContext
	_countvecstmt  ICountvecstmtContext
	_accessvecstmt IAccessvecstmtContext
	right          IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *ExprContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *ExprContext) Get_NIL() antlr.Token { return s._NIL }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *ExprContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *ExprContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_accessstmt() IAccessstmtContext { return s._accessstmt }

func (s *ExprContext) Get_emptvecstmt() IEmptvecstmtContext { return s._emptvecstmt }

func (s *ExprContext) Get_countvecstmt() ICountvecstmtContext { return s._countvecstmt }

func (s *ExprContext) Get_accessvecstmt() IAccessvecstmtContext { return s._accessvecstmt }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_accessstmt(v IAccessstmtContext) { s._accessstmt = v }

func (s *ExprContext) Set_emptvecstmt(v IEmptvecstmtContext) { s._emptvecstmt = v }

func (s *ExprContext) Set_countvecstmt(v ICountvecstmtContext) { s._countvecstmt = v }

func (s *ExprContext) Set_accessvecstmt(v IAccessvecstmtContext) { s._accessvecstmt = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(GrammarParserSUB, 0)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserPARDER, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GrammarParserNUMBER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *ExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserTRUE, 0)
}

func (s *ExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserFALSE, 0)
}

func (s *ExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(GrammarParserNIL, 0)
}

func (s *ExprContext) Accessstmt() IAccessstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessstmtContext)
}

func (s *ExprContext) Emptvecstmt() IEmptvecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptvecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptvecstmtContext)
}

func (s *ExprContext) Countvecstmt() ICountvecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountvecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountvecstmtContext)
}

func (s *ExprContext) Accessvecstmt() IAccessvecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessvecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessvecstmtContext)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIV, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(GrammarParserADD, 0)
}

func (s *ExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPERCENT, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(GrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(GrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(512)

			var _m = p.Match(GrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)

			var _x = p.expr(12)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()), localctx.(*ExprContext).GetLeft().GetE())

	case 2:
		{
			p.SetState(516)

			var _m = p.Match(GrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)

			var _x = p.expr(11)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()), nil)

	case 3:
		{
			p.SetState(520)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(522)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 4:
		{
			p.SetState(525)

			var _m = p.Match(GrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 5:
		{
			p.SetState(527)

			var _m = p.Match(GrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 6:
		{
			p.SetState(529)

			var _m = p.Match(GrammarParserTRUE)

			localctx.(*ExprContext)._TRUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRUE().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 7:
		{
			p.SetState(531)

			var _m = p.Match(GrammarParserFALSE)

			localctx.(*ExprContext)._FALSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FALSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FALSE().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case 8:
		{
			p.SetState(533)

			var _m = p.Match(GrammarParserNIL)

			localctx.(*ExprContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetColumn()
			}
		}()), nil, environment.NULL)

	case 9:
		{
			p.SetState(535)

			var _x = p.Accessstmt()

			localctx.(*ExprContext)._accessstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_accessstmt().GetAccess()

	case 10:
		{
			p.SetState(538)

			var _x = p.Emptvecstmt()

			localctx.(*ExprContext)._emptvecstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_emptvecstmt().GetEmptyvec()

	case 11:
		{
			p.SetState(541)

			var _x = p.Countvecstmt()

			localctx.(*ExprContext)._countvecstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_countvecstmt().GetCount()

	case 12:
		{
			p.SetState(544)

			var _x = p.Accessvecstmt()

			localctx.(*ExprContext)._accessvecstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_accessvecstmt().GetAccessvec()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(584)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(549)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(550)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserMUL || _la == GrammarParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(551)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(554)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(555)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3940649673949184) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(556)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(559)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(560)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserMAY_IG || _la == GrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(561)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(564)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(565)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserMEN_IG || _la == GrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(566)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(569)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(570)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserDIF || _la == GrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(571)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(574)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(575)

					var _m = p.Match(GrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(576)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(579)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(580)

					var _m = p.Match(GrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(581)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(588)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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

// IAccessstmtContext is an interface to support dynamic dispatch.
type IAccessstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetAccess returns the access attribute.
	GetAccess() interfaces.Expression

	// SetAccess sets the access attribute.
	SetAccess(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode

	// IsAccessstmtContext differentiates from other interfaces.
	IsAccessstmtContext()
}

type AccessstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	access interfaces.Expression
	op     antlr.Token
}

func NewEmptyAccessstmtContext() *AccessstmtContext {
	var p = new(AccessstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_accessstmt
	return p
}

func InitEmptyAccessstmtContext(p *AccessstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_accessstmt
}

func (*AccessstmtContext) IsAccessstmtContext() {}

func NewAccessstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessstmtContext {
	var p = new(AccessstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_accessstmt

	return p
}

func (s *AccessstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessstmtContext) GetOp() antlr.Token { return s.op }

func (s *AccessstmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *AccessstmtContext) GetAccess() interfaces.Expression { return s.access }

func (s *AccessstmtContext) SetAccess(v interfaces.Expression) { s.access = v }

func (s *AccessstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AccessstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAccessstmt(s)
	}
}

func (s *AccessstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAccessstmt(s)
	}
}

func (p *GrammarParser) Accessstmt() (localctx IAccessstmtContext) {
	localctx = NewAccessstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GrammarParserRULE_accessstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)

		var _m = p.Match(GrammarParserID)

		localctx.(*AccessstmtContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*AccessstmtContext).access = expressions.NewAccess((func() int {
		if localctx.(*AccessstmtContext).GetOp() == nil {
			return 0
		} else {
			return localctx.(*AccessstmtContext).GetOp().GetLine()
		}
	}()), (func() int {
		if localctx.(*AccessstmtContext).GetOp() == nil {
			return 0
		} else {
			return localctx.(*AccessstmtContext).GetOp().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AccessstmtContext).GetOp() == nil {
			return ""
		} else {
			return localctx.(*AccessstmtContext).GetOp().GetText()
		}
	}()))

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

// IIncreaseanddecreasestmtContext is an interface to support dynamic dispatch.
type IIncreaseanddecreasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_IG_ADD returns the _IG_ADD token.
	Get_IG_ADD() antlr.Token

	// Get_IG_SUB returns the _IG_SUB token.
	Get_IG_SUB() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_IG_ADD sets the _IG_ADD token.
	Set_IG_ADD(antlr.Token)

	// Set_IG_SUB sets the _IG_SUB token.
	Set_IG_SUB(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetIncreasedecrease returns the increasedecrease attribute.
	GetIncreasedecrease() interfaces.Instruction

	// SetIncreasedecrease sets the increasedecrease attribute.
	SetIncreasedecrease(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IG_ADD() antlr.TerminalNode
	Expr() IExprContext
	PTCOMA() antlr.TerminalNode
	IG_SUB() antlr.TerminalNode

	// IsIncreaseanddecreasestmtContext differentiates from other interfaces.
	IsIncreaseanddecreasestmtContext()
}

type IncreaseanddecreasestmtContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	increasedecrease interfaces.Instruction
	_ID              antlr.Token
	_IG_ADD          antlr.Token
	_expr            IExprContext
	_IG_SUB          antlr.Token
}

func NewEmptyIncreaseanddecreasestmtContext() *IncreaseanddecreasestmtContext {
	var p = new(IncreaseanddecreasestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_increaseanddecreasestmt
	return p
}

func InitEmptyIncreaseanddecreasestmtContext(p *IncreaseanddecreasestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_increaseanddecreasestmt
}

func (*IncreaseanddecreasestmtContext) IsIncreaseanddecreasestmtContext() {}

func NewIncreaseanddecreasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncreaseanddecreasestmtContext {
	var p = new(IncreaseanddecreasestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_increaseanddecreasestmt

	return p
}

func (s *IncreaseanddecreasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncreaseanddecreasestmtContext) Get_ID() antlr.Token { return s._ID }

func (s *IncreaseanddecreasestmtContext) Get_IG_ADD() antlr.Token { return s._IG_ADD }

func (s *IncreaseanddecreasestmtContext) Get_IG_SUB() antlr.Token { return s._IG_SUB }

func (s *IncreaseanddecreasestmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *IncreaseanddecreasestmtContext) Set_IG_ADD(v antlr.Token) { s._IG_ADD = v }

func (s *IncreaseanddecreasestmtContext) Set_IG_SUB(v antlr.Token) { s._IG_SUB = v }

func (s *IncreaseanddecreasestmtContext) Get_expr() IExprContext { return s._expr }

func (s *IncreaseanddecreasestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IncreaseanddecreasestmtContext) GetIncreasedecrease() interfaces.Instruction {
	return s.increasedecrease
}

func (s *IncreaseanddecreasestmtContext) SetIncreasedecrease(v interfaces.Instruction) {
	s.increasedecrease = v
}

func (s *IncreaseanddecreasestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IncreaseanddecreasestmtContext) IG_ADD() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG_ADD, 0)
}

func (s *IncreaseanddecreasestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncreaseanddecreasestmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserPTCOMA, 0)
}

func (s *IncreaseanddecreasestmtContext) IG_SUB() antlr.TerminalNode {
	return s.GetToken(GrammarParserIG_SUB, 0)
}

func (s *IncreaseanddecreasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncreaseanddecreasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncreaseanddecreasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIncreaseanddecreasestmt(s)
	}
}

func (s *IncreaseanddecreasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIncreaseanddecreasestmt(s)
	}
}

func (p *GrammarParser) Increaseanddecreasestmt() (localctx IIncreaseanddecreasestmtContext) {
	localctx = NewIncreaseanddecreasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrammarParserRULE_increaseanddecreasestmt)
	var _la int

	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(592)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)

			var _m = p.Match(GrammarParserIG_ADD)

			localctx.(*IncreaseanddecreasestmtContext)._IG_ADD = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)

			var _x = p.expr(0)

			localctx.(*IncreaseanddecreasestmtContext)._expr = _x
		}
		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(595)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*IncreaseanddecreasestmtContext).increasedecrease = instructions.NewIncreaseDecrease((func() int {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*IncreaseanddecreasestmtContext).Get_IG_ADD() == nil {
				return ""
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_IG_ADD().GetText()
			}
		}()), localctx.(*IncreaseanddecreasestmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(600)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)

			var _m = p.Match(GrammarParserIG_SUB)

			localctx.(*IncreaseanddecreasestmtContext)._IG_SUB = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)

			var _x = p.expr(0)

			localctx.(*IncreaseanddecreasestmtContext)._expr = _x
		}
		p.SetState(604)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserPTCOMA {
			{
				p.SetState(603)
				p.Match(GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*IncreaseanddecreasestmtContext).increasedecrease = instructions.NewIncreaseDecrease((func() int {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*IncreaseanddecreasestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*IncreaseanddecreasestmtContext).Get_IG_SUB() == nil {
				return ""
			} else {
				return localctx.(*IncreaseanddecreasestmtContext).Get_IG_SUB().GetText()
			}
		}()), localctx.(*IncreaseanddecreasestmtContext).Get_expr().GetE())

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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t attribute.
	GetT() environment.TipoExpresion

	// SetT sets the t attribute.
	SetT(environment.TipoExpresion)

	// Getter signatures
	STRINGS() antlr.TerminalNode
	INTS() antlr.TerminalNode
	FLOATS() antlr.TerminalNode
	BOOLS() antlr.TerminalNode
	CHARACTERS() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t      environment.TipoExpresion
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetT() environment.TipoExpresion { return s.t }

func (s *TypeContext) SetT(v environment.TipoExpresion) { s.t = v }

func (s *TypeContext) STRINGS() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRINGS, 0)
}

func (s *TypeContext) INTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserINTS, 0)
}

func (s *TypeContext) FLOATS() antlr.TerminalNode {
	return s.GetToken(GrammarParserFLOATS, 0)
}

func (s *TypeContext) BOOLS() antlr.TerminalNode {
	return s.GetToken(GrammarParserBOOLS, 0)
}

func (s *TypeContext) CHARACTERS() antlr.TerminalNode {
	return s.GetToken(GrammarParserCHARACTERS, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GrammarParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrammarParserRULE_type)
	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSTRINGS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(610)
			p.Match(GrammarParserSTRINGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypeContext).t = environment.STRING

	case GrammarParserINTS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(612)
			p.Match(GrammarParserINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypeContext).t = environment.INTEGER

	case GrammarParserFLOATS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(614)
			p.Match(GrammarParserFLOATS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypeContext).t = environment.FLOAT

	case GrammarParserBOOLS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(616)
			p.Match(GrammarParserBOOLS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypeContext).t = environment.BOOLEAN

	case GrammarParserCHARACTERS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(618)
			p.Match(GrammarParserCHARACTERS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypeContext).t = environment.CHARACTER

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

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *GrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *GrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, GrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(623)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(633)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_listParams)
			p.SetState(626)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(627)
				p.Match(GrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(628)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(635)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
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

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 30:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
