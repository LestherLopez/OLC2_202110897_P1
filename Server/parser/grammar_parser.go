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
		"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "", "",
		"", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'+='", "'-='", "'>='",
		"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'",
		"'{'", "'}'", "':'", "'?'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL",
		"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN",
		"SWITCH", "CASE", "DEFAULT", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", "POINT", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt",
		"blockelifs", "ifstmt", "blockcases", "switchstmt", "assignationstmt",
		"whilestmt", "forstmt", "expr", "accessstmt", "increaseanddecreasestmt",
		"type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 326, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 40, 8, 1, 11, 1, 12,
		1, 41, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 73, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 102, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 118, 8, 5, 1, 6, 4, 6, 121, 8, 6, 11,
		6, 12, 6, 122, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 154, 8, 7, 1, 8, 4,
		8, 157, 8, 8, 11, 8, 12, 8, 158, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 193, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 228, 8, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 257, 8, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 294, 8, 13, 10, 13, 12, 13, 297, 9, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		312, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 324, 8, 16, 1, 16, 0, 1, 26, 17, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 5, 1, 0, 35, 36, 1, 0, 37, 39,
		2, 0, 31, 31, 33, 33, 2, 0, 32, 32, 34, 34, 1, 0, 23, 24, 347, 0, 34, 1,
		0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 74, 1, 0, 0, 0, 8, 101,
		1, 0, 0, 0, 10, 117, 1, 0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 153, 1, 0, 0,
		0, 16, 156, 1, 0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 194, 1, 0, 0, 0, 22, 199,
		1, 0, 0, 0, 24, 227, 1, 0, 0, 0, 26, 256, 1, 0, 0, 0, 28, 298, 1, 0, 0,
		0, 30, 311, 1, 0, 0, 0, 32, 323, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36,
		5, 0, 0, 1, 36, 37, 6, 0, -1, 0, 37, 1, 1, 0, 0, 0, 38, 40, 3, 4, 2, 0,
		39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1,
		0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 6, 1, -1, 0, 44, 3, 1, 0, 0, 0, 45,
		46, 3, 6, 3, 0, 46, 47, 6, 2, -1, 0, 47, 73, 1, 0, 0, 0, 48, 49, 3, 8,
		4, 0, 49, 50, 6, 2, -1, 0, 50, 73, 1, 0, 0, 0, 51, 52, 3, 10, 5, 0, 52,
		53, 6, 2, -1, 0, 53, 73, 1, 0, 0, 0, 54, 55, 3, 20, 10, 0, 55, 56, 6, 2,
		-1, 0, 56, 73, 1, 0, 0, 0, 57, 58, 3, 30, 15, 0, 58, 59, 6, 2, -1, 0, 59,
		73, 1, 0, 0, 0, 60, 61, 3, 14, 7, 0, 61, 62, 6, 2, -1, 0, 62, 73, 1, 0,
		0, 0, 63, 64, 3, 22, 11, 0, 64, 65, 6, 2, -1, 0, 65, 73, 1, 0, 0, 0, 66,
		67, 3, 24, 12, 0, 67, 68, 6, 2, -1, 0, 68, 73, 1, 0, 0, 0, 69, 70, 3, 18,
		9, 0, 70, 71, 6, 2, -1, 0, 71, 73, 1, 0, 0, 0, 72, 45, 1, 0, 0, 0, 72,
		48, 1, 0, 0, 0, 72, 51, 1, 0, 0, 0, 72, 54, 1, 0, 0, 0, 72, 57, 1, 0, 0,
		0, 72, 60, 1, 0, 0, 0, 72, 63, 1, 0, 0, 0, 72, 66, 1, 0, 0, 0, 72, 69,
		1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 5, 10, 0, 0, 75, 76, 5, 40, 0, 0,
		76, 77, 3, 26, 13, 0, 77, 78, 5, 41, 0, 0, 78, 79, 6, 3, -1, 0, 79, 7,
		1, 0, 0, 0, 80, 81, 5, 6, 0, 0, 81, 82, 5, 22, 0, 0, 82, 83, 5, 44, 0,
		0, 83, 84, 3, 32, 16, 0, 84, 85, 5, 28, 0, 0, 85, 86, 3, 26, 13, 0, 86,
		87, 6, 4, -1, 0, 87, 102, 1, 0, 0, 0, 88, 89, 5, 6, 0, 0, 89, 90, 5, 22,
		0, 0, 90, 91, 5, 28, 0, 0, 91, 92, 3, 26, 13, 0, 92, 93, 6, 4, -1, 0, 93,
		102, 1, 0, 0, 0, 94, 95, 5, 6, 0, 0, 95, 96, 5, 22, 0, 0, 96, 97, 5, 44,
		0, 0, 97, 98, 3, 32, 16, 0, 98, 99, 5, 45, 0, 0, 99, 100, 6, 4, -1, 0,
		100, 102, 1, 0, 0, 0, 101, 80, 1, 0, 0, 0, 101, 88, 1, 0, 0, 0, 101, 94,
		1, 0, 0, 0, 102, 9, 1, 0, 0, 0, 103, 104, 5, 14, 0, 0, 104, 105, 5, 22,
		0, 0, 105, 106, 5, 44, 0, 0, 106, 107, 3, 32, 16, 0, 107, 108, 5, 28, 0,
		0, 108, 109, 3, 26, 13, 0, 109, 110, 6, 5, -1, 0, 110, 118, 1, 0, 0, 0,
		111, 112, 5, 14, 0, 0, 112, 113, 5, 22, 0, 0, 113, 114, 5, 28, 0, 0, 114,
		115, 3, 26, 13, 0, 115, 116, 6, 5, -1, 0, 116, 118, 1, 0, 0, 0, 117, 103,
		1, 0, 0, 0, 117, 111, 1, 0, 0, 0, 118, 11, 1, 0, 0, 0, 119, 121, 3, 14,
		7, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0,
		122, 123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 6, 6, -1, 0, 125,
		13, 1, 0, 0, 0, 126, 127, 5, 11, 0, 0, 127, 128, 3, 26, 13, 0, 128, 129,
		5, 42, 0, 0, 129, 130, 3, 2, 1, 0, 130, 131, 5, 43, 0, 0, 131, 132, 6,
		7, -1, 0, 132, 154, 1, 0, 0, 0, 133, 134, 5, 11, 0, 0, 134, 135, 3, 26,
		13, 0, 135, 136, 5, 42, 0, 0, 136, 137, 3, 2, 1, 0, 137, 138, 5, 43, 0,
		0, 138, 139, 5, 12, 0, 0, 139, 140, 5, 42, 0, 0, 140, 141, 3, 2, 1, 0,
		141, 142, 5, 43, 0, 0, 142, 143, 6, 7, -1, 0, 143, 154, 1, 0, 0, 0, 144,
		145, 5, 11, 0, 0, 145, 146, 3, 26, 13, 0, 146, 147, 5, 42, 0, 0, 147, 148,
		3, 2, 1, 0, 148, 149, 5, 43, 0, 0, 149, 150, 5, 12, 0, 0, 150, 151, 3,
		12, 6, 0, 151, 152, 6, 7, -1, 0, 152, 154, 1, 0, 0, 0, 153, 126, 1, 0,
		0, 0, 153, 133, 1, 0, 0, 0, 153, 144, 1, 0, 0, 0, 154, 15, 1, 0, 0, 0,
		155, 157, 3, 18, 9, 0, 156, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161,
		6, 8, -1, 0, 161, 17, 1, 0, 0, 0, 162, 163, 5, 17, 0, 0, 163, 164, 3, 26,
		13, 0, 164, 165, 5, 42, 0, 0, 165, 166, 5, 18, 0, 0, 166, 167, 3, 26, 13,
		0, 167, 168, 5, 44, 0, 0, 168, 169, 3, 2, 1, 0, 169, 170, 5, 43, 0, 0,
		170, 171, 6, 9, -1, 0, 171, 193, 1, 0, 0, 0, 172, 173, 5, 17, 0, 0, 173,
		174, 3, 26, 13, 0, 174, 175, 5, 42, 0, 0, 175, 176, 3, 16, 8, 0, 176, 177,
		5, 19, 0, 0, 177, 178, 5, 44, 0, 0, 178, 179, 3, 2, 1, 0, 179, 180, 5,
		43, 0, 0, 180, 181, 6, 9, -1, 0, 181, 193, 1, 0, 0, 0, 182, 183, 5, 17,
		0, 0, 183, 184, 3, 26, 13, 0, 184, 185, 5, 42, 0, 0, 185, 186, 5, 18, 0,
		0, 186, 187, 3, 26, 13, 0, 187, 188, 5, 44, 0, 0, 188, 189, 3, 2, 1, 0,
		189, 190, 3, 16, 8, 0, 190, 191, 5, 43, 0, 0, 191, 193, 1, 0, 0, 0, 192,
		162, 1, 0, 0, 0, 192, 172, 1, 0, 0, 0, 192, 182, 1, 0, 0, 0, 193, 19, 1,
		0, 0, 0, 194, 195, 5, 22, 0, 0, 195, 196, 5, 28, 0, 0, 196, 197, 3, 26,
		13, 0, 197, 198, 6, 10, -1, 0, 198, 21, 1, 0, 0, 0, 199, 200, 5, 13, 0,
		0, 200, 201, 3, 26, 13, 0, 201, 202, 5, 42, 0, 0, 202, 203, 3, 2, 1, 0,
		203, 204, 5, 43, 0, 0, 204, 205, 6, 11, -1, 0, 205, 23, 1, 0, 0, 0, 206,
		207, 5, 15, 0, 0, 207, 208, 5, 22, 0, 0, 208, 209, 5, 16, 0, 0, 209, 210,
		3, 26, 13, 0, 210, 211, 5, 42, 0, 0, 211, 212, 3, 2, 1, 0, 212, 213, 5,
		43, 0, 0, 213, 214, 6, 12, -1, 0, 214, 228, 1, 0, 0, 0, 215, 216, 5, 15,
		0, 0, 216, 217, 5, 22, 0, 0, 217, 218, 5, 16, 0, 0, 218, 219, 3, 26, 13,
		0, 219, 220, 5, 46, 0, 0, 220, 221, 5, 46, 0, 0, 221, 222, 3, 26, 13, 0,
		222, 223, 5, 42, 0, 0, 223, 224, 3, 2, 1, 0, 224, 225, 5, 43, 0, 0, 225,
		226, 6, 12, -1, 0, 226, 228, 1, 0, 0, 0, 227, 206, 1, 0, 0, 0, 227, 215,
		1, 0, 0, 0, 228, 25, 1, 0, 0, 0, 229, 230, 6, 13, -1, 0, 230, 231, 5, 25,
		0, 0, 231, 232, 3, 26, 13, 9, 232, 233, 6, 13, -1, 0, 233, 257, 1, 0, 0,
		0, 234, 235, 5, 38, 0, 0, 235, 236, 3, 26, 13, 8, 236, 237, 6, 13, -1,
		0, 237, 257, 1, 0, 0, 0, 238, 239, 5, 40, 0, 0, 239, 240, 3, 26, 13, 0,
		240, 241, 5, 41, 0, 0, 241, 242, 6, 13, -1, 0, 242, 257, 1, 0, 0, 0, 243,
		244, 5, 20, 0, 0, 244, 257, 6, 13, -1, 0, 245, 246, 5, 21, 0, 0, 246, 257,
		6, 13, -1, 0, 247, 248, 5, 8, 0, 0, 248, 257, 6, 13, -1, 0, 249, 250, 5,
		9, 0, 0, 250, 257, 6, 13, -1, 0, 251, 252, 5, 7, 0, 0, 252, 257, 6, 13,
		-1, 0, 253, 254, 3, 28, 14, 0, 254, 255, 6, 13, -1, 0, 255, 257, 1, 0,
		0, 0, 256, 229, 1, 0, 0, 0, 256, 234, 1, 0, 0, 0, 256, 238, 1, 0, 0, 0,
		256, 243, 1, 0, 0, 0, 256, 245, 1, 0, 0, 0, 256, 247, 1, 0, 0, 0, 256,
		249, 1, 0, 0, 0, 256, 251, 1, 0, 0, 0, 256, 253, 1, 0, 0, 0, 257, 295,
		1, 0, 0, 0, 258, 259, 10, 16, 0, 0, 259, 260, 7, 0, 0, 0, 260, 261, 3,
		26, 13, 17, 261, 262, 6, 13, -1, 0, 262, 294, 1, 0, 0, 0, 263, 264, 10,
		15, 0, 0, 264, 265, 7, 1, 0, 0, 265, 266, 3, 26, 13, 16, 266, 267, 6, 13,
		-1, 0, 267, 294, 1, 0, 0, 0, 268, 269, 10, 14, 0, 0, 269, 270, 7, 2, 0,
		0, 270, 271, 3, 26, 13, 15, 271, 272, 6, 13, -1, 0, 272, 294, 1, 0, 0,
		0, 273, 274, 10, 13, 0, 0, 274, 275, 7, 3, 0, 0, 275, 276, 3, 26, 13, 14,
		276, 277, 6, 13, -1, 0, 277, 294, 1, 0, 0, 0, 278, 279, 10, 12, 0, 0, 279,
		280, 7, 4, 0, 0, 280, 281, 3, 26, 13, 13, 281, 282, 6, 13, -1, 0, 282,
		294, 1, 0, 0, 0, 283, 284, 10, 11, 0, 0, 284, 285, 5, 27, 0, 0, 285, 286,
		3, 26, 13, 12, 286, 287, 6, 13, -1, 0, 287, 294, 1, 0, 0, 0, 288, 289,
		10, 10, 0, 0, 289, 290, 5, 26, 0, 0, 290, 291, 3, 26, 13, 11, 291, 292,
		6, 13, -1, 0, 292, 294, 1, 0, 0, 0, 293, 258, 1, 0, 0, 0, 293, 263, 1,
		0, 0, 0, 293, 268, 1, 0, 0, 0, 293, 273, 1, 0, 0, 0, 293, 278, 1, 0, 0,
		0, 293, 283, 1, 0, 0, 0, 293, 288, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295,
		293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 27, 1, 0, 0, 0, 297, 295, 1,
		0, 0, 0, 298, 299, 5, 22, 0, 0, 299, 300, 6, 14, -1, 0, 300, 29, 1, 0,
		0, 0, 301, 302, 5, 22, 0, 0, 302, 303, 5, 29, 0, 0, 303, 304, 3, 26, 13,
		0, 304, 305, 6, 15, -1, 0, 305, 312, 1, 0, 0, 0, 306, 307, 5, 22, 0, 0,
		307, 308, 5, 30, 0, 0, 308, 309, 3, 26, 13, 0, 309, 310, 6, 15, -1, 0,
		310, 312, 1, 0, 0, 0, 311, 301, 1, 0, 0, 0, 311, 306, 1, 0, 0, 0, 312,
		31, 1, 0, 0, 0, 313, 314, 5, 1, 0, 0, 314, 324, 6, 16, -1, 0, 315, 316,
		5, 2, 0, 0, 316, 324, 6, 16, -1, 0, 317, 318, 5, 3, 0, 0, 318, 324, 6,
		16, -1, 0, 319, 320, 5, 4, 0, 0, 320, 324, 6, 16, -1, 0, 321, 322, 5, 5,
		0, 0, 322, 324, 6, 16, -1, 0, 323, 313, 1, 0, 0, 0, 323, 315, 1, 0, 0,
		0, 323, 317, 1, 0, 0, 0, 323, 319, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324,
		33, 1, 0, 0, 0, 14, 41, 72, 101, 117, 122, 153, 158, 192, 227, 256, 293,
		295, 311, 323,
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
	GrammarParserNUMBER       = 20
	GrammarParserSTRING       = 21
	GrammarParserID           = 22
	GrammarParserDIF          = 23
	GrammarParserIG_IG        = 24
	GrammarParserNOT          = 25
	GrammarParserOR           = 26
	GrammarParserAND          = 27
	GrammarParserIG           = 28
	GrammarParserIG_ADD       = 29
	GrammarParserIG_SUB       = 30
	GrammarParserMAY_IG       = 31
	GrammarParserMEN_IG       = 32
	GrammarParserMAYOR        = 33
	GrammarParserMENOR        = 34
	GrammarParserMUL          = 35
	GrammarParserDIV          = 36
	GrammarParserADD          = 37
	GrammarParserSUB          = 38
	GrammarParserPERCENT      = 39
	GrammarParserPARIZQ       = 40
	GrammarParserPARDER       = 41
	GrammarParserLLAVEIZQ     = 42
	GrammarParserLLAVEDER     = 43
	GrammarParserDOUBLEPTS    = 44
	GrammarParserQUESTION     = 45
	GrammarParserPOINT        = 46
	GrammarParserWHITESPACE   = 47
	GrammarParserCOMMENT      = 48
	GrammarParserLINE_COMMENT = 49
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
	GrammarParserRULE_blockcases              = 8
	GrammarParserRULE_switchstmt              = 9
	GrammarParserRULE_assignationstmt         = 10
	GrammarParserRULE_whilestmt               = 11
	GrammarParserRULE_forstmt                 = 12
	GrammarParserRULE_expr                    = 13
	GrammarParserRULE_accessstmt              = 14
	GrammarParserRULE_increaseanddecreasestmt = 15
	GrammarParserRULE_type                    = 16
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
		p.SetState(34)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(35)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(38)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(41)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)

			var _x = p.Declarestmt()

			localctx.(*InstructionContext)._declarestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarestmt().GetDec()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)

			var _x = p.Constantstmt()

			localctx.(*InstructionContext)._constantstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constantstmt().GetConst_()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)

			var _x = p.Assignationstmt()

			localctx.(*InstructionContext)._assignationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignationstmt().GetAssign()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)

			var _x = p.Increaseanddecreasestmt()

			localctx.(*InstructionContext)._increaseanddecreasestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_increaseanddecreasestmt().GetIncreasedecrease()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIft()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(63)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhile()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(66)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFor_()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(69)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwitch_()

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
		p.SetState(74)

		var _m = p.Match(GrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(77)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(84)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)

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
			p.SetState(88)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

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
			p.SetState(94)

			var _m = p.Match(GrammarParserVAR)

			localctx.(*DeclarestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)

			var _m = p.Match(GrammarParserID)

			localctx.(*DeclarestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _x = p.Type_()

			localctx.(*DeclarestmtContext)._type = _x
		}
		{
			p.SetState(98)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _x = p.Type_()

			localctx.(*ConstantstmtContext)._type = _x
		}
		{
			p.SetState(107)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)

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
			p.SetState(111)

			var _m = p.Match(GrammarParserLET)

			localctx.(*ConstantstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)

			var _m = p.Match(GrammarParserID)

			localctx.(*ConstantstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(GrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)

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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(119)

				var _x = p.Ifstmt()

				localctx.(*BlockelifsContext)._ifstmt = _x
			}
			localctx.(*BlockelifsContext).elif = append(localctx.(*BlockelifsContext).elif, localctx.(*BlockelifsContext)._ifstmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(122)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(128)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifb = _x
		}
		{
			p.SetState(130)
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
		localctx.(*IfstmtContext).SetEl(localctx.(*IfstmtContext).GetIfb().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(135)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifelseblck = _x
		}
		{
			p.SetState(137)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)

			var _x = p.Block()

			localctx.(*IfstmtContext).elseifblck = _x
		}
		{
			p.SetState(141)
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
			p.SetState(144)

			var _m = p.Match(GrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(146)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)

			var _x = p.Block()

			localctx.(*IfstmtContext).elif = _x
		}
		{
			p.SetState(148)
			p.Match(GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)

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

// IBlockcasesContext is an interface to support dynamic dispatch.
type IBlockcasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_switchstmt returns the _switchstmt rule contexts.
	Get_switchstmt() ISwitchstmtContext

	// Set_switchstmt sets the _switchstmt rule contexts.
	Set_switchstmt(ISwitchstmtContext)

	// GetCase_ returns the case_ rule context list.
	GetCase_() []ISwitchstmtContext

	// SetCase_ sets the case_ rule context list.
	SetCase_([]ISwitchstmtContext)

	// GetBlkcase returns the blkcase attribute.
	GetBlkcase() []interface{}

	// SetBlkcase sets the blkcase attribute.
	SetBlkcase([]interface{})

	// Getter signatures
	AllSwitchstmt() []ISwitchstmtContext
	Switchstmt(i int) ISwitchstmtContext

	// IsBlockcasesContext differentiates from other interfaces.
	IsBlockcasesContext()
}

type BlockcasesContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	blkcase     []interface{}
	_switchstmt ISwitchstmtContext
	case_       []ISwitchstmtContext
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

func (s *BlockcasesContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *BlockcasesContext) Set_switchstmt(v ISwitchstmtContext) { s._switchstmt = v }

func (s *BlockcasesContext) GetCase_() []ISwitchstmtContext { return s.case_ }

func (s *BlockcasesContext) SetCase_(v []ISwitchstmtContext) { s.case_ = v }

func (s *BlockcasesContext) GetBlkcase() []interface{} { return s.blkcase }

func (s *BlockcasesContext) SetBlkcase(v []interface{}) { s.blkcase = v }

func (s *BlockcasesContext) AllSwitchstmt() []ISwitchstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			len++
		}
	}

	tst := make([]ISwitchstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchstmtContext); ok {
			tst[i] = t.(ISwitchstmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockcasesContext) Switchstmt(i int) ISwitchstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
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

	return t.(ISwitchstmtContext)
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
	p.EnterRule(localctx, 16, GrammarParserRULE_blockcases)

	localctx.(*BlockcasesContext).blkcase = []interface{}{}
	var listcases []ISwitchstmtContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserSWITCH {
		{
			p.SetState(155)

			var _x = p.Switchstmt()

			localctx.(*BlockcasesContext)._switchstmt = _x
		}
		localctx.(*BlockcasesContext).case_ = append(localctx.(*BlockcasesContext).case_, localctx.(*BlockcasesContext)._switchstmt)

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listcases = localctx.(*BlockcasesContext).GetCase_()
	for _, a := range listcases {
		localctx.(*BlockcasesContext).blkcase = append(localctx.(*BlockcasesContext).blkcase, a.GetSwitch_())
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

	// GetExpc returns the expc rule contexts.
	GetExpc() IExprContext

	// GetExpv returns the expv rule contexts.
	GetExpv() IExprContext

	// GetCase_ returns the case_ rule contexts.
	GetCase_() IBlockContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetExp returns the exp rule contexts.
	GetExp() IExprContext

	// SetExpc sets the expc rule contexts.
	SetExpc(IExprContext)

	// SetExpv sets the expv rule contexts.
	SetExpv(IExprContext)

	// SetCase_ sets the case_ rule contexts.
	SetCase_(IBlockContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetExp sets the exp rule contexts.
	SetExp(IExprContext)

	// GetSwitch_ returns the switch_ attribute.
	GetSwitch_() interfaces.Instruction

	// SetSwitch_ sets the switch_ attribute.
	SetSwitch_(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	CASE() antlr.TerminalNode
	DOUBLEPTS() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Block() IBlockContext
	Blockcases() IBlockcasesContext
	DEFAULT() antlr.TerminalNode

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	switch_ interfaces.Instruction
	_SWITCH antlr.Token
	expc    IExprContext
	expv    IExprContext
	case_   IBlockContext
	_block  IBlockContext
	_expr   IExprContext
	exp     IExprContext
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

func (s *SwitchstmtContext) GetExpc() IExprContext { return s.expc }

func (s *SwitchstmtContext) GetExpv() IExprContext { return s.expv }

func (s *SwitchstmtContext) GetCase_() IBlockContext { return s.case_ }

func (s *SwitchstmtContext) Get_block() IBlockContext { return s._block }

func (s *SwitchstmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchstmtContext) GetExp() IExprContext { return s.exp }

func (s *SwitchstmtContext) SetExpc(v IExprContext) { s.expc = v }

func (s *SwitchstmtContext) SetExpv(v IExprContext) { s.expv = v }

func (s *SwitchstmtContext) SetCase_(v IBlockContext) { s.case_ = v }

func (s *SwitchstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *SwitchstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchstmtContext) SetExp(v IExprContext) { s.exp = v }

func (s *SwitchstmtContext) GetSwitch_() interfaces.Instruction { return s.switch_ }

func (s *SwitchstmtContext) SetSwitch_(v interfaces.Instruction) { s.switch_ = v }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *SwitchstmtContext) DOUBLEPTS() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLEPTS, 0)
}

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(GrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) AllExpr() []IExprContext {
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

func (s *SwitchstmtContext) Expr(i int) IExprContext {
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

func (s *SwitchstmtContext) Blockcases() IBlockcasesContext {
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

func (s *SwitchstmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
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
	p.EnterRule(localctx, 18, GrammarParserRULE_switchstmt)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext).expc = _x
		}
		{
			p.SetState(164)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(GrammarParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext).expv = _x
		}
		{
			p.SetState(167)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)

			var _x = p.Block()

			localctx.(*SwitchstmtContext).case_ = _x
			localctx.(*SwitchstmtContext)._block = _x
		}
		{
			p.SetState(169)
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
		}()), localctx.(*SwitchstmtContext).GetExpc().GetE(), localctx.(*SwitchstmtContext).GetExpv().GetE(), localctx.(*SwitchstmtContext).Get_block().GetBlk(), nil))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)

			var _m = p.Match(GrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(174)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Blockcases()
		}
		{
			p.SetState(176)
			p.Match(GrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)

			var _x = p.Block()

			localctx.(*SwitchstmtContext)._block = _x
		}
		{
			p.SetState(179)
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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(GrammarParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext).expc = _x
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
			p.Match(GrammarParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext).exp = _x
		}
		{
			p.SetState(187)
			p.Match(GrammarParserDOUBLEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Block()
		}
		{
			p.SetState(189)
			p.Blockcases()
		}
		{
			p.SetState(190)
			p.Match(GrammarParserLLAVEDER)
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
	p.EnterRule(localctx, 20, GrammarParserRULE_assignationstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)

		var _m = p.Match(GrammarParserID)

		localctx.(*AssignationstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)

		var _m = p.Match(GrammarParserIG)

		localctx.(*AssignationstmtContext)._IG = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)

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
	p.EnterRule(localctx, 22, GrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)

		var _m = p.Match(GrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(201)
		p.Match(GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 24, GrammarParserRULE_forstmt)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)

			var _x = p.expr(0)

			localctx.(*ForstmtContext)._expr = _x
		}
		{
			p.SetState(210)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(212)
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
			p.SetState(215)

			var _m = p.Match(GrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)

			var _m = p.Match(GrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).first = _x
		}
		{
			p.SetState(219)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Match(GrammarParserPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).second = _x
		}
		{
			p.SetState(222)
			p.Match(GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(224)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserNOT:
		{
			p.SetState(230)

			var _m = p.Match(GrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)

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
			p.SetState(234)

			var _m = p.Match(GrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

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
			p.SetState(238)
			p.Match(GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(240)
			p.Match(GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case GrammarParserNUMBER:
		{
			p.SetState(243)

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
			p.SetState(245)

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
			p.SetState(247)

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
			p.SetState(249)

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
			p.SetState(251)

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
			p.SetState(253)

			var _x = p.Accessstmt()

			localctx.(*ExprContext)._accessstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_accessstmt().GetAccess()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(259)

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
					p.SetState(260)

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
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(264)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&962072674304) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(265)

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
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(269)

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
					p.SetState(270)

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
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(274)

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
					p.SetState(275)

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
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(279)

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
					p.SetState(280)

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
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(284)

					var _m = p.Match(GrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(285)

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
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(289)

					var _m = p.Match(GrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(290)

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
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, GrammarParserRULE_accessstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)

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
	p.EnterRule(localctx, 30, GrammarParserRULE_increaseanddecreasestmt)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)

			var _m = p.Match(GrammarParserIG_ADD)

			localctx.(*IncreaseanddecreasestmtContext)._IG_ADD = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

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
			p.SetState(306)

			var _m = p.Match(GrammarParserID)

			localctx.(*IncreaseanddecreasestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)

			var _m = p.Match(GrammarParserIG_SUB)

			localctx.(*IncreaseanddecreasestmtContext)._IG_SUB = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

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
	p.EnterRule(localctx, 32, GrammarParserRULE_type)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSTRINGS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
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
			p.SetState(315)
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
			p.SetState(317)
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
			p.SetState(319)
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
			p.SetState(321)
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
	case 13:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
