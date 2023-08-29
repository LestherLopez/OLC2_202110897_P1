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
		"'at'", "'IsEmpty'", "'count'", "", "", "", "'!='", "'=='", "'!'", "'||'",
		"'&&'", "'='", "'+='", "'-='", "'>='", "'<='", "'>'", "'<'", "'*'",
		"'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'",
		"'.'", "','", "';'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL",
		"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN",
		"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK",
		"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "NUMBER",
		"STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD",
		"IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD",
		"SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS",
		"QUESTION", "POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt",
		"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt",
		"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt",
		"appendstmt", "removelaststmt", "removestmt", "emptvecstmt", "countvecstmt",
		"listParams", "expr", "accessstmt", "increaseanddecreasestmt", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 482, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 60, 8, 1, 11, 1, 12, 1, 61,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		111, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 140, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 156,
		8, 5, 1, 6, 4, 6, 159, 8, 6, 11, 6, 12, 6, 160, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 192, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 210, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 228, 8, 9, 1, 10, 4, 10, 231, 8, 10, 11, 10, 12,
		10, 232, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 270, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 295, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 338, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 5, 22, 381, 8, 22, 10, 22, 12, 22, 384, 9, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 413, 8, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 5, 23, 450, 8, 23, 10, 23, 12, 23, 453, 9, 23, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 3, 25, 468, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 3, 26, 480, 8, 26, 1, 26, 0, 2, 44, 46, 27, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 0, 6, 1, 0, 21, 23, 1, 0, 45, 46, 1, 0, 47,
		49, 2, 0, 41, 41, 43, 43, 2, 0, 42, 42, 44, 44, 1, 0, 33, 34, 506, 0, 54,
		1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4, 110, 1, 0, 0, 0, 6, 112, 1, 0, 0, 0,
		8, 139, 1, 0, 0, 0, 10, 155, 1, 0, 0, 0, 12, 158, 1, 0, 0, 0, 14, 191,
		1, 0, 0, 0, 16, 209, 1, 0, 0, 0, 18, 227, 1, 0, 0, 0, 20, 230, 1, 0, 0,
		0, 22, 236, 1, 0, 0, 0, 24, 241, 1, 0, 0, 0, 26, 269, 1, 0, 0, 0, 28, 271,
		1, 0, 0, 0, 30, 294, 1, 0, 0, 0, 32, 337, 1, 0, 0, 0, 34, 339, 1, 0, 0,
		0, 36, 347, 1, 0, 0, 0, 38, 354, 1, 0, 0, 0, 40, 363, 1, 0, 0, 0, 42, 367,
		1, 0, 0, 0, 44, 371, 1, 0, 0, 0, 46, 412, 1, 0, 0, 0, 48, 454, 1, 0, 0,
		0, 50, 467, 1, 0, 0, 0, 52, 479, 1, 0, 0, 0, 54, 55, 3, 2, 1, 0, 55, 56,
		5, 0, 0, 1, 56, 57, 6, 0, -1, 0, 57, 1, 1, 0, 0, 0, 58, 60, 3, 4, 2, 0,
		59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 6, 1, -1, 0, 64, 3, 1, 0, 0, 0, 65,
		66, 3, 6, 3, 0, 66, 67, 6, 2, -1, 0, 67, 111, 1, 0, 0, 0, 68, 69, 3, 8,
		4, 0, 69, 70, 6, 2, -1, 0, 70, 111, 1, 0, 0, 0, 71, 72, 3, 10, 5, 0, 72,
		73, 6, 2, -1, 0, 73, 111, 1, 0, 0, 0, 74, 75, 3, 22, 11, 0, 75, 76, 6,
		2, -1, 0, 76, 111, 1, 0, 0, 0, 77, 78, 3, 50, 25, 0, 78, 79, 6, 2, -1,
		0, 79, 111, 1, 0, 0, 0, 80, 81, 3, 14, 7, 0, 81, 82, 6, 2, -1, 0, 82, 111,
		1, 0, 0, 0, 83, 84, 3, 24, 12, 0, 84, 85, 6, 2, -1, 0, 85, 111, 1, 0, 0,
		0, 86, 87, 3, 26, 13, 0, 87, 88, 6, 2, -1, 0, 88, 111, 1, 0, 0, 0, 89,
		90, 3, 16, 8, 0, 90, 91, 6, 2, -1, 0, 91, 111, 1, 0, 0, 0, 92, 93, 3, 28,
		14, 0, 93, 94, 6, 2, -1, 0, 94, 111, 1, 0, 0, 0, 95, 96, 3, 30, 15, 0,
		96, 97, 6, 2, -1, 0, 97, 111, 1, 0, 0, 0, 98, 99, 3, 32, 16, 0, 99, 100,
		6, 2, -1, 0, 100, 111, 1, 0, 0, 0, 101, 102, 3, 34, 17, 0, 102, 103, 6,
		2, -1, 0, 103, 111, 1, 0, 0, 0, 104, 105, 3, 36, 18, 0, 105, 106, 6, 2,
		-1, 0, 106, 111, 1, 0, 0, 0, 107, 108, 3, 38, 19, 0, 108, 109, 6, 2, -1,
		0, 109, 111, 1, 0, 0, 0, 110, 65, 1, 0, 0, 0, 110, 68, 1, 0, 0, 0, 110,
		71, 1, 0, 0, 0, 110, 74, 1, 0, 0, 0, 110, 77, 1, 0, 0, 0, 110, 80, 1, 0,
		0, 0, 110, 83, 1, 0, 0, 0, 110, 86, 1, 0, 0, 0, 110, 89, 1, 0, 0, 0, 110,
		92, 1, 0, 0, 0, 110, 95, 1, 0, 0, 0, 110, 98, 1, 0, 0, 0, 110, 101, 1,
		0, 0, 0, 110, 104, 1, 0, 0, 0, 110, 107, 1, 0, 0, 0, 111, 5, 1, 0, 0, 0,
		112, 113, 5, 10, 0, 0, 113, 114, 5, 50, 0, 0, 114, 115, 3, 46, 23, 0, 115,
		116, 5, 51, 0, 0, 116, 117, 6, 3, -1, 0, 117, 7, 1, 0, 0, 0, 118, 119,
		5, 6, 0, 0, 119, 120, 5, 32, 0, 0, 120, 121, 5, 54, 0, 0, 121, 122, 3,
		52, 26, 0, 122, 123, 5, 38, 0, 0, 123, 124, 3, 46, 23, 0, 124, 125, 6,
		4, -1, 0, 125, 140, 1, 0, 0, 0, 126, 127, 5, 6, 0, 0, 127, 128, 5, 32,
		0, 0, 128, 129, 5, 38, 0, 0, 129, 130, 3, 46, 23, 0, 130, 131, 6, 4, -1,
		0, 131, 140, 1, 0, 0, 0, 132, 133, 5, 6, 0, 0, 133, 134, 5, 32, 0, 0, 134,
		135, 5, 54, 0, 0, 135, 136, 3, 52, 26, 0, 136, 137, 5, 55, 0, 0, 137, 138,
		6, 4, -1, 0, 138, 140, 1, 0, 0, 0, 139, 118, 1, 0, 0, 0, 139, 126, 1, 0,
		0, 0, 139, 132, 1, 0, 0, 0, 140, 9, 1, 0, 0, 0, 141, 142, 5, 14, 0, 0,
		142, 143, 5, 32, 0, 0, 143, 144, 5, 54, 0, 0, 144, 145, 3, 52, 26, 0, 145,
		146, 5, 38, 0, 0, 146, 147, 3, 46, 23, 0, 147, 148, 6, 5, -1, 0, 148, 156,
		1, 0, 0, 0, 149, 150, 5, 14, 0, 0, 150, 151, 5, 32, 0, 0, 151, 152, 5,
		38, 0, 0, 152, 153, 3, 46, 23, 0, 153, 154, 6, 5, -1, 0, 154, 156, 1, 0,
		0, 0, 155, 141, 1, 0, 0, 0, 155, 149, 1, 0, 0, 0, 156, 11, 1, 0, 0, 0,
		157, 159, 3, 14, 7, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160,
		158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163,
		6, 6, -1, 0, 163, 13, 1, 0, 0, 0, 164, 165, 5, 11, 0, 0, 165, 166, 3, 46,
		23, 0, 166, 167, 5, 52, 0, 0, 167, 168, 3, 2, 1, 0, 168, 169, 5, 53, 0,
		0, 169, 170, 6, 7, -1, 0, 170, 192, 1, 0, 0, 0, 171, 172, 5, 11, 0, 0,
		172, 173, 3, 46, 23, 0, 173, 174, 5, 52, 0, 0, 174, 175, 3, 2, 1, 0, 175,
		176, 5, 53, 0, 0, 176, 177, 5, 12, 0, 0, 177, 178, 5, 52, 0, 0, 178, 179,
		3, 2, 1, 0, 179, 180, 5, 53, 0, 0, 180, 181, 6, 7, -1, 0, 181, 192, 1,
		0, 0, 0, 182, 183, 5, 11, 0, 0, 183, 184, 3, 46, 23, 0, 184, 185, 5, 52,
		0, 0, 185, 186, 3, 2, 1, 0, 186, 187, 5, 53, 0, 0, 187, 188, 5, 12, 0,
		0, 188, 189, 3, 12, 6, 0, 189, 190, 6, 7, -1, 0, 190, 192, 1, 0, 0, 0,
		191, 164, 1, 0, 0, 0, 191, 171, 1, 0, 0, 0, 191, 182, 1, 0, 0, 0, 192,
		15, 1, 0, 0, 0, 193, 194, 5, 17, 0, 0, 194, 195, 3, 46, 23, 0, 195, 196,
		5, 52, 0, 0, 196, 197, 5, 19, 0, 0, 197, 198, 5, 54, 0, 0, 198, 199, 3,
		2, 1, 0, 199, 200, 5, 53, 0, 0, 200, 201, 6, 8, -1, 0, 201, 210, 1, 0,
		0, 0, 202, 203, 5, 17, 0, 0, 203, 204, 3, 46, 23, 0, 204, 205, 5, 52, 0,
		0, 205, 206, 3, 18, 9, 0, 206, 207, 5, 53, 0, 0, 207, 208, 6, 8, -1, 0,
		208, 210, 1, 0, 0, 0, 209, 193, 1, 0, 0, 0, 209, 202, 1, 0, 0, 0, 210,
		17, 1, 0, 0, 0, 211, 212, 5, 18, 0, 0, 212, 213, 3, 46, 23, 0, 213, 214,
		5, 54, 0, 0, 214, 215, 3, 2, 1, 0, 215, 216, 3, 20, 10, 0, 216, 217, 6,
		9, -1, 0, 217, 228, 1, 0, 0, 0, 218, 219, 5, 18, 0, 0, 219, 220, 3, 46,
		23, 0, 220, 221, 5, 54, 0, 0, 221, 222, 3, 2, 1, 0, 222, 223, 5, 19, 0,
		0, 223, 224, 5, 54, 0, 0, 224, 225, 3, 2, 1, 0, 225, 226, 6, 9, -1, 0,
		226, 228, 1, 0, 0, 0, 227, 211, 1, 0, 0, 0, 227, 218, 1, 0, 0, 0, 228,
		19, 1, 0, 0, 0, 229, 231, 3, 18, 9, 0, 230, 229, 1, 0, 0, 0, 231, 232,
		1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0,
		0, 0, 234, 235, 6, 10, -1, 0, 235, 21, 1, 0, 0, 0, 236, 237, 5, 32, 0,
		0, 237, 238, 5, 38, 0, 0, 238, 239, 3, 46, 23, 0, 239, 240, 6, 11, -1,
		0, 240, 23, 1, 0, 0, 0, 241, 242, 5, 13, 0, 0, 242, 243, 3, 46, 23, 0,
		243, 244, 5, 52, 0, 0, 244, 245, 3, 2, 1, 0, 245, 246, 5, 53, 0, 0, 246,
		247, 6, 12, -1, 0, 247, 25, 1, 0, 0, 0, 248, 249, 5, 15, 0, 0, 249, 250,
		5, 32, 0, 0, 250, 251, 5, 16, 0, 0, 251, 252, 3, 46, 23, 0, 252, 253, 5,
		52, 0, 0, 253, 254, 3, 2, 1, 0, 254, 255, 5, 53, 0, 0, 255, 256, 6, 13,
		-1, 0, 256, 270, 1, 0, 0, 0, 257, 258, 5, 15, 0, 0, 258, 259, 5, 32, 0,
		0, 259, 260, 5, 16, 0, 0, 260, 261, 3, 46, 23, 0, 261, 262, 5, 56, 0, 0,
		262, 263, 5, 56, 0, 0, 263, 264, 3, 46, 23, 0, 264, 265, 5, 52, 0, 0, 265,
		266, 3, 2, 1, 0, 266, 267, 5, 53, 0, 0, 267, 268, 6, 13, -1, 0, 268, 270,
		1, 0, 0, 0, 269, 248, 1, 0, 0, 0, 269, 257, 1, 0, 0, 0, 270, 27, 1, 0,
		0, 0, 271, 272, 5, 20, 0, 0, 272, 273, 3, 46, 23, 0, 273, 274, 5, 12, 0,
		0, 274, 275, 5, 52, 0, 0, 275, 276, 3, 2, 1, 0, 276, 277, 7, 0, 0, 0, 277,
		278, 5, 53, 0, 0, 278, 279, 6, 14, -1, 0, 279, 29, 1, 0, 0, 0, 280, 281,
		5, 22, 0, 0, 281, 282, 5, 58, 0, 0, 282, 295, 6, 15, -1, 0, 283, 284, 5,
		22, 0, 0, 284, 285, 3, 46, 23, 0, 285, 286, 5, 58, 0, 0, 286, 287, 6, 15,
		-1, 0, 287, 295, 1, 0, 0, 0, 288, 289, 5, 21, 0, 0, 289, 290, 5, 58, 0,
		0, 290, 295, 6, 15, -1, 0, 291, 292, 5, 23, 0, 0, 292, 293, 5, 58, 0, 0,
		293, 295, 6, 15, -1, 0, 294, 280, 1, 0, 0, 0, 294, 283, 1, 0, 0, 0, 294,
		288, 1, 0, 0, 0, 294, 291, 1, 0, 0, 0, 295, 31, 1, 0, 0, 0, 296, 297, 5,
		6, 0, 0, 297, 298, 5, 32, 0, 0, 298, 299, 5, 54, 0, 0, 299, 300, 5, 59,
		0, 0, 300, 301, 3, 52, 26, 0, 301, 302, 5, 60, 0, 0, 302, 303, 5, 38, 0,
		0, 303, 304, 5, 59, 0, 0, 304, 305, 3, 44, 22, 0, 305, 306, 5, 60, 0, 0,
		306, 307, 6, 16, -1, 0, 307, 338, 1, 0, 0, 0, 308, 309, 5, 6, 0, 0, 309,
		310, 5, 32, 0, 0, 310, 311, 5, 54, 0, 0, 311, 312, 5, 59, 0, 0, 312, 313,
		3, 52, 26, 0, 313, 314, 5, 60, 0, 0, 314, 315, 5, 38, 0, 0, 315, 316, 5,
		59, 0, 0, 316, 317, 5, 60, 0, 0, 317, 318, 6, 16, -1, 0, 318, 338, 1, 0,
		0, 0, 319, 320, 5, 6, 0, 0, 320, 321, 5, 32, 0, 0, 321, 322, 5, 38, 0,
		0, 322, 323, 5, 59, 0, 0, 323, 324, 3, 52, 26, 0, 324, 325, 5, 60, 0, 0,
		325, 326, 5, 50, 0, 0, 326, 327, 5, 51, 0, 0, 327, 338, 1, 0, 0, 0, 328,
		329, 5, 6, 0, 0, 329, 330, 5, 32, 0, 0, 330, 331, 5, 54, 0, 0, 331, 332,
		5, 59, 0, 0, 332, 333, 3, 52, 26, 0, 333, 334, 5, 60, 0, 0, 334, 335, 5,
		38, 0, 0, 335, 336, 5, 32, 0, 0, 336, 338, 1, 0, 0, 0, 337, 296, 1, 0,
		0, 0, 337, 308, 1, 0, 0, 0, 337, 319, 1, 0, 0, 0, 337, 328, 1, 0, 0, 0,
		338, 33, 1, 0, 0, 0, 339, 340, 5, 32, 0, 0, 340, 341, 5, 56, 0, 0, 341,
		342, 5, 24, 0, 0, 342, 343, 5, 50, 0, 0, 343, 344, 3, 46, 23, 0, 344, 345,
		5, 51, 0, 0, 345, 346, 6, 17, -1, 0, 346, 35, 1, 0, 0, 0, 347, 348, 5,
		32, 0, 0, 348, 349, 5, 56, 0, 0, 349, 350, 5, 25, 0, 0, 350, 351, 5, 50,
		0, 0, 351, 352, 5, 51, 0, 0, 352, 353, 6, 18, -1, 0, 353, 37, 1, 0, 0,
		0, 354, 355, 5, 32, 0, 0, 355, 356, 5, 56, 0, 0, 356, 357, 5, 26, 0, 0,
		357, 358, 5, 50, 0, 0, 358, 359, 5, 27, 0, 0, 359, 360, 5, 54, 0, 0, 360,
		361, 3, 46, 23, 0, 361, 362, 5, 51, 0, 0, 362, 39, 1, 0, 0, 0, 363, 364,
		5, 32, 0, 0, 364, 365, 5, 56, 0, 0, 365, 366, 5, 28, 0, 0, 366, 41, 1,
		0, 0, 0, 367, 368, 5, 32, 0, 0, 368, 369, 5, 56, 0, 0, 369, 370, 5, 29,
		0, 0, 370, 43, 1, 0, 0, 0, 371, 372, 6, 22, -1, 0, 372, 373, 3, 46, 23,
		0, 373, 374, 6, 22, -1, 0, 374, 382, 1, 0, 0, 0, 375, 376, 10, 2, 0, 0,
		376, 377, 5, 57, 0, 0, 377, 378, 3, 46, 23, 0, 378, 379, 6, 22, -1, 0,
		379, 381, 1, 0, 0, 0, 380, 375, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382,
		380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 45, 1, 0, 0, 0, 384, 382, 1,
		0, 0, 0, 385, 386, 6, 23, -1, 0, 386, 387, 5, 35, 0, 0, 387, 388, 3, 46,
		23, 9, 388, 389, 6, 23, -1, 0, 389, 413, 1, 0, 0, 0, 390, 391, 5, 48, 0,
		0, 391, 392, 3, 46, 23, 8, 392, 393, 6, 23, -1, 0, 393, 413, 1, 0, 0, 0,
		394, 395, 5, 50, 0, 0, 395, 396, 3, 46, 23, 0, 396, 397, 5, 51, 0, 0, 397,
		398, 6, 23, -1, 0, 398, 413, 1, 0, 0, 0, 399, 400, 5, 30, 0, 0, 400, 413,
		6, 23, -1, 0, 401, 402, 5, 31, 0, 0, 402, 413, 6, 23, -1, 0, 403, 404,
		5, 8, 0, 0, 404, 413, 6, 23, -1, 0, 405, 406, 5, 9, 0, 0, 406, 413, 6,
		23, -1, 0, 407, 408, 5, 7, 0, 0, 408, 413, 6, 23, -1, 0, 409, 410, 3, 48,
		24, 0, 410, 411, 6, 23, -1, 0, 411, 413, 1, 0, 0, 0, 412, 385, 1, 0, 0,
		0, 412, 390, 1, 0, 0, 0, 412, 394, 1, 0, 0, 0, 412, 399, 1, 0, 0, 0, 412,
		401, 1, 0, 0, 0, 412, 403, 1, 0, 0, 0, 412, 405, 1, 0, 0, 0, 412, 407,
		1, 0, 0, 0, 412, 409, 1, 0, 0, 0, 413, 451, 1, 0, 0, 0, 414, 415, 10, 16,
		0, 0, 415, 416, 7, 1, 0, 0, 416, 417, 3, 46, 23, 17, 417, 418, 6, 23, -1,
		0, 418, 450, 1, 0, 0, 0, 419, 420, 10, 15, 0, 0, 420, 421, 7, 2, 0, 0,
		421, 422, 3, 46, 23, 16, 422, 423, 6, 23, -1, 0, 423, 450, 1, 0, 0, 0,
		424, 425, 10, 14, 0, 0, 425, 426, 7, 3, 0, 0, 426, 427, 3, 46, 23, 15,
		427, 428, 6, 23, -1, 0, 428, 450, 1, 0, 0, 0, 429, 430, 10, 13, 0, 0, 430,
		431, 7, 4, 0, 0, 431, 432, 3, 46, 23, 14, 432, 433, 6, 23, -1, 0, 433,
		450, 1, 0, 0, 0, 434, 435, 10, 12, 0, 0, 435, 436, 7, 5, 0, 0, 436, 437,
		3, 46, 23, 13, 437, 438, 6, 23, -1, 0, 438, 450, 1, 0, 0, 0, 439, 440,
		10, 11, 0, 0, 440, 441, 5, 37, 0, 0, 441, 442, 3, 46, 23, 12, 442, 443,
		6, 23, -1, 0, 443, 450, 1, 0, 0, 0, 444, 445, 10, 10, 0, 0, 445, 446, 5,
		36, 0, 0, 446, 447, 3, 46, 23, 11, 447, 448, 6, 23, -1, 0, 448, 450, 1,
		0, 0, 0, 449, 414, 1, 0, 0, 0, 449, 419, 1, 0, 0, 0, 449, 424, 1, 0, 0,
		0, 449, 429, 1, 0, 0, 0, 449, 434, 1, 0, 0, 0, 449, 439, 1, 0, 0, 0, 449,
		444, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452,
		1, 0, 0, 0, 452, 47, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 5, 32,
		0, 0, 455, 456, 6, 24, -1, 0, 456, 49, 1, 0, 0, 0, 457, 458, 5, 32, 0,
		0, 458, 459, 5, 39, 0, 0, 459, 460, 3, 46, 23, 0, 460, 461, 6, 25, -1,
		0, 461, 468, 1, 0, 0, 0, 462, 463, 5, 32, 0, 0, 463, 464, 5, 40, 0, 0,
		464, 465, 3, 46, 23, 0, 465, 466, 6, 25, -1, 0, 466, 468, 1, 0, 0, 0, 467,
		457, 1, 0, 0, 0, 467, 462, 1, 0, 0, 0, 468, 51, 1, 0, 0, 0, 469, 470, 5,
		1, 0, 0, 470, 480, 6, 26, -1, 0, 471, 472, 5, 2, 0, 0, 472, 480, 6, 26,
		-1, 0, 473, 474, 5, 3, 0, 0, 474, 480, 6, 26, -1, 0, 475, 476, 5, 4, 0,
		0, 476, 480, 6, 26, -1, 0, 477, 478, 5, 5, 0, 0, 478, 480, 6, 26, -1, 0,
		479, 469, 1, 0, 0, 0, 479, 471, 1, 0, 0, 0, 479, 473, 1, 0, 0, 0, 479,
		475, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 53, 1, 0, 0, 0, 18, 61, 110,
		139, 155, 160, 191, 209, 227, 232, 269, 294, 337, 382, 412, 449, 451, 467,
		479,
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
	GrammarParserNUMBER       = 30
	GrammarParserSTRING       = 31
	GrammarParserID           = 32
	GrammarParserDIF          = 33
	GrammarParserIG_IG        = 34
	GrammarParserNOT          = 35
	GrammarParserOR           = 36
	GrammarParserAND          = 37
	GrammarParserIG           = 38
	GrammarParserIG_ADD       = 39
	GrammarParserIG_SUB       = 40
	GrammarParserMAY_IG       = 41
	GrammarParserMEN_IG       = 42
	GrammarParserMAYOR        = 43
	GrammarParserMENOR        = 44
	GrammarParserMUL          = 45
	GrammarParserDIV          = 46
	GrammarParserADD          = 47
	GrammarParserSUB          = 48
	GrammarParserPERCENT      = 49
	GrammarParserPARIZQ       = 50
	GrammarParserPARDER       = 51
	GrammarParserLLAVEIZQ     = 52
	GrammarParserLLAVEDER     = 53
	GrammarParserDOUBLEPTS    = 54
	GrammarParserQUESTION     = 55
	GrammarParserPOINT        = 56
	GrammarParserCOMA         = 57
	GrammarParserPTCOMA       = 58
	GrammarParserCORCHETEIZQ  = 59
	GrammarParserCORCHETEDER  = 60
	GrammarParserWHITESPACE   = 61
	GrammarParserCOMMENT      = 62
	GrammarParserLINE_COMMENT = 63
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
	GrammarParserRULE_listParams              = 22
	GrammarParserRULE_expr                    = 23
	GrammarParserRULE_accessstmt              = 24
	GrammarParserRULE_increaseanddecreasestmt = 25
	GrammarParserRULE_type                    = 26
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
		p.SetState(54)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(55)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(58)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(61)
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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)

			var _x = p.Declarestmt()

			localctx.(*InstructionContext)._declarestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarestmt().GetDec()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)

			var _x = p.Constantstmt()

			localctx.(*InstructionContext)._constantstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constantstmt().GetConst_()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)

			var _x = p.Assignationstmt()

			localctx.(*InstructionContext)._assignationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignationstmt().GetAssign()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)

			var _x = p.Increaseanddecreasestmt()

			localctx.(*InstructionContext)._increaseanddecreasestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_increaseanddecreasestmt().GetIncreasedecrease()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(80)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIft()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(83)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhile()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(86)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFor_()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(89)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwitch_()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(92)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGua()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(95)

			var _x = p.Transferstmt()

			localctx.(*InstructionContext)._transferstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transferstmt().GetTran()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(98)

			var _x = p.Declarevectorstmt()

			localctx.(*InstructionContext)._declarevectorstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarevectorstmt().GetDecvec()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(101)

			var _x = p.Appendstmt()

			localctx.(*InstructionContext)._appendstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_appendstmt().GetApp()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(104)

			var _x = p.Removelaststmt()

			localctx.(*InstructionContext)._removelaststmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removelaststmt().GetRemovl()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(107)

			var _x = p.Removestmt()

			localctx.(*InstructionContext)._removestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removestmt().GetRemov()

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)

		var _m = p.Match(GrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(115)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(122)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

			var _x = p.expr(0)

			localctx.(*DeclarestmtContext)._expr = _x
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
			p.SetState(126)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.expr(0)

			localctx.(*DeclarestmtContext)._expr = _x
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
			p.SetState(132)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(136)
			p.Match(GrammarParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)

			var _x = p.Type_()

			localctx.(*ConstantstmtContext)._type = _x
		}
		{
			p.SetState(145)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)

			var _x = p.expr(0)

			localctx.(*ConstantstmtContext)._expr = _x
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
			p.SetState(149)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _x = p.expr(0)

			localctx.(*ConstantstmtContext)._expr = _x
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
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(157)

				var _x = p.Ifstmt()

				localctx.(*BlockelifsContext)._ifstmt = _x
			}
			localctx.(*BlockelifsContext).elif = append(localctx.(*BlockelifsContext).elif, localctx.(*BlockelifsContext)._ifstmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(166)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifb = _x
		}
		{
			p.SetState(168)
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
			p.SetState(171)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(173)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifelseblck = _x
		}
		{
			p.SetState(175)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)

			var _x = p.Block()

			localctx.(*IfstmtContext).elseifblck = _x
		}
		{
			p.SetState(179)
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
			p.SetState(182)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(184)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _x = p.Block()

			localctx.(*IfstmtContext).elif = _x
		}
		{
			p.SetState(186)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

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
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(195)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(GrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _x = p.Block()

			localctx.(*SwitchstmtContext)._block = _x
		}
		{
			p.SetState(199)
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
			p.SetState(202)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(204)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)

			var _x = p.Casestmt()

			localctx.(*SwitchstmtContext)._casestmt = _x
		}
		{
			p.SetState(206)
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
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)

			var _m = p.Match(GrammarParserCASE)

			localctx.(*CasestmtContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)

			var _x = p.expr(0)

			localctx.(*CasestmtContext)._expr = _x
		}
		{
			p.SetState(213)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)

			var _x = p.Block()

			localctx.(*CasestmtContext)._block = _x
		}
		{
			p.SetState(215)

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
			p.SetState(218)

			var _m = p.Match(GrammarParserCASE)

			localctx.(*CasestmtContext)._CASE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)

			var _x = p.expr(0)

			localctx.(*CasestmtContext)._expr = _x
		}
		{
			p.SetState(220)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)

			var _x = p.Block()

			localctx.(*CasestmtContext).sen = _x
		}
		{
			p.SetState(222)
			p.Match(GrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)

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
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(229)

				var _x = p.Casestmt()

				localctx.(*BlockcasesContext)._casestmt = _x
			}
			localctx.(*BlockcasesContext).casedef = append(localctx.(*BlockcasesContext).casedef, localctx.(*BlockcasesContext)._casestmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)

		var _m = p.Match(GrammarParserID)

		localctx.(*AssignationstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)

		var _m = p.Match(GrammarParserIG)

		localctx.(*AssignationstmtContext)._IG = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)

		var _x = p.expr(0)

		localctx.(*AssignationstmtContext)._expr = _x
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
		p.SetState(241)

		var _m = p.Match(GrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(243)
		p.Match(GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(245)
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
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.expr(0)

			localctx.(*ForstmtContext)._expr = _x
		}
		{
			p.SetState(252)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(254)
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
			p.SetState(257)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).first = _x
		}
		{
			p.SetState(261)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).second = _x
		}
		{
			p.SetState(264)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(266)
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
		p.SetState(271)

		var _m = p.Match(GrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(273)
		p.Match(GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Match(GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(276)

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
		p.SetState(277)
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
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)

			var _m = p.Match(GrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
			p.SetState(283)

			var _m = p.Match(GrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)

			var _x = p.expr(0)

			localctx.(*TransferstmtContext)._expr = _x
		}
		{
			p.SetState(285)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
			p.SetState(288)

			var _m = p.Match(GrammarParserCONTINUE)

			localctx.(*TransferstmtContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
			p.SetState(291)

			var _m = p.Match(GrammarParserBREAK)

			localctx.(*TransferstmtContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Match(GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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

func (s *DeclarevectorstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarevectorstmtContext) Set_ID(v antlr.Token) { s._ID = v }

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
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarevectorstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)

			var _x = p.Type_()

			localctx.(*DeclarevectorstmtContext)._type = _x
		}
		{
			p.SetState(301)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

			var _x = p.listParams(0)

			localctx.(*DeclarevectorstmtContext)._listParams = _x
		}
		{
			p.SetState(305)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarevectorstmtContext).decvec = instructions.NewToDecalreVector((func() int {
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
		}()), localctx.(*DeclarevectorstmtContext).Get_type().GetT(), localctx.(*DeclarevectorstmtContext).Get_listParams().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarevectorstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarevectorstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)

			var _x = p.Type_()

			localctx.(*DeclarevectorstmtContext)._type = _x
		}
		{
			p.SetState(313)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarevectorstmtContext).decvec = instructions.NewToDecalreVector((func() int {
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
		}()), localctx.(*DeclarevectorstmtContext).Get_type().GetT(), nil)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Type_()
		}
		{
			p.SetState(324)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(328)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Type_()
		}
		{
			p.SetState(333)
			p.Match(GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(GrammarParserID)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)

		var _m = p.Match(GrammarParserID)

		localctx.(*AppendstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Match(GrammarParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)

		var _x = p.expr(0)

		localctx.(*AppendstmtContext)._expr = _x
	}
	{
		p.SetState(344)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)

		var _m = p.Match(GrammarParserID)

		localctx.(*RemovelaststmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(GrammarParserREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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

	// IsRemovestmtContext differentiates from other interfaces.
	IsRemovestmtContext()
}

type RemovestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	remov  interfaces.Instruction
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(GrammarParserREMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(GrammarParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(GrammarParserDOUBLEPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.expr(0)
	}
	{
		p.SetState(361)
		p.Match(GrammarParserPARDER)
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

// IEmptvecstmtContext is an interface to support dynamic dispatch.
type IEmptvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(363)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Match(GrammarParserISEMPTY)
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

// ICountvecstmtContext is an interface to support dynamic dispatch.
type ICountvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(367)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Match(GrammarParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Match(GrammarParserCOUNT)
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, GrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
			p.SetState(375)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(376)
				p.Match(GrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(377)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_accessstmt sets the _accessstmt rule contexts.
	Set_accessstmt(IAccessstmtContext)

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
	parser      antlr.Parser
	e           interfaces.Expression
	left        IExprContext
	op          antlr.Token
	_expr       IExprContext
	_NUMBER     antlr.Token
	_STRING     antlr.Token
	_TRUE       antlr.Token
	_FALSE      antlr.Token
	_NIL        antlr.Token
	_accessstmt IAccessstmtContext
	right       IExprContext
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

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_accessstmt(v IAccessstmtContext) { s._accessstmt = v }

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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserNOT:
		{
			p.SetState(386)

			var _m = p.Match(GrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)

			var _x = p.expr(9)

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

	case GrammarParserSUB:
		{
			p.SetState(390)

			var _m = p.Match(GrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)

			var _x = p.expr(8)

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

	case GrammarParserPARIZQ:
		{
			p.SetState(394)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(396)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case GrammarParserNUMBER:
		{
			p.SetState(399)

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

	case GrammarParserSTRING:
		{
			p.SetState(401)

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

	case GrammarParserTRUE:
		{
			p.SetState(403)

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

	case GrammarParserFALSE:
		{
			p.SetState(405)

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

	case GrammarParserNIL:
		{
			p.SetState(407)

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

	case GrammarParserID:
		{
			p.SetState(409)

			var _x = p.Accessstmt()

			localctx.(*ExprContext)._accessstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_accessstmt().GetAccess()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(415)

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
					p.SetState(416)

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

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(420)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(421)

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

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

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
					p.SetState(426)

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

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(430)

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
					p.SetState(431)

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

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(435)

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
					p.SetState(436)

					var _x = p.expr(13)

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
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(440)

					var _m = p.Match(GrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(441)

					var _x = p.expr(12)

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
				p.SetState(444)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(445)

					var _m = p.Match(GrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(446)

					var _x = p.expr(11)

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
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, GrammarParserRULE_accessstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)

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
	p.EnterRule(localctx, 50, GrammarParserRULE_increaseanddecreasestmt)
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(457)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _m = p.Match(GrammarParserIG_ADD)

			localctx.(*IncreaseanddecreasestmtContext)._IG_ADD = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)

			var _x = p.expr(0)

			localctx.(*IncreaseanddecreasestmtContext)._expr = _x
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
			p.SetState(462)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)

			var _m = p.Match(GrammarParserIG_SUB)

			localctx.(*IncreaseanddecreasestmtContext)._IG_SUB = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)

			var _x = p.expr(0)

			localctx.(*IncreaseanddecreasestmtContext)._expr = _x
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
	p.EnterRule(localctx, 52, GrammarParserRULE_type)
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSTRINGS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)
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
			p.SetState(471)
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
			p.SetState(473)
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
			p.SetState(475)
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
			p.SetState(477)
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

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 23:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
