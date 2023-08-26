// Code generated from Lex.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Lex struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LexLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexLexerInit() {
	staticData := &LexLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", "TRUE",
		"FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", "SWITCH",
		"CASE", "DEFAULT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT",
		"OR", "AND", "IG", "IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", "POINT", "WHITESPACE",
		"COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 49, 323, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 4, 19, 205, 8, 19,
		11, 19, 12, 19, 206, 1, 19, 1, 19, 4, 19, 211, 8, 19, 11, 19, 12, 19, 212,
		3, 19, 215, 8, 19, 1, 20, 1, 20, 5, 20, 219, 8, 20, 10, 20, 12, 20, 222,
		9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 228, 8, 21, 10, 21, 12, 21, 231,
		9, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 4, 46, 290, 8, 46, 11, 46, 12, 46,
		291, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 300, 8, 47, 10, 47,
		12, 47, 303, 9, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 5, 48, 314, 8, 48, 10, 48, 12, 48, 317, 9, 48, 1, 48, 1, 48,
		1, 49, 1, 49, 1, 49, 1, 301, 0, 50, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85,
		43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 0, 1, 0, 7, 1,
		0, 48, 57, 1, 0, 34, 34, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13,
		7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 329, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1,
		0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 1, 101, 1, 0, 0, 0, 3, 108,
		1, 0, 0, 0, 5, 112, 1, 0, 0, 0, 7, 118, 1, 0, 0, 0, 9, 123, 1, 0, 0, 0,
		11, 133, 1, 0, 0, 0, 13, 137, 1, 0, 0, 0, 15, 141, 1, 0, 0, 0, 17, 146,
		1, 0, 0, 0, 19, 152, 1, 0, 0, 0, 21, 158, 1, 0, 0, 0, 23, 161, 1, 0, 0,
		0, 25, 166, 1, 0, 0, 0, 27, 172, 1, 0, 0, 0, 29, 176, 1, 0, 0, 0, 31, 180,
		1, 0, 0, 0, 33, 183, 1, 0, 0, 0, 35, 190, 1, 0, 0, 0, 37, 195, 1, 0, 0,
		0, 39, 204, 1, 0, 0, 0, 41, 216, 1, 0, 0, 0, 43, 225, 1, 0, 0, 0, 45, 232,
		1, 0, 0, 0, 47, 235, 1, 0, 0, 0, 49, 238, 1, 0, 0, 0, 51, 240, 1, 0, 0,
		0, 53, 243, 1, 0, 0, 0, 55, 246, 1, 0, 0, 0, 57, 248, 1, 0, 0, 0, 59, 251,
		1, 0, 0, 0, 61, 254, 1, 0, 0, 0, 63, 257, 1, 0, 0, 0, 65, 260, 1, 0, 0,
		0, 67, 262, 1, 0, 0, 0, 69, 264, 1, 0, 0, 0, 71, 266, 1, 0, 0, 0, 73, 268,
		1, 0, 0, 0, 75, 270, 1, 0, 0, 0, 77, 272, 1, 0, 0, 0, 79, 274, 1, 0, 0,
		0, 81, 276, 1, 0, 0, 0, 83, 278, 1, 0, 0, 0, 85, 280, 1, 0, 0, 0, 87, 282,
		1, 0, 0, 0, 89, 284, 1, 0, 0, 0, 91, 286, 1, 0, 0, 0, 93, 289, 1, 0, 0,
		0, 95, 295, 1, 0, 0, 0, 97, 309, 1, 0, 0, 0, 99, 320, 1, 0, 0, 0, 101,
		102, 5, 83, 0, 0, 102, 103, 5, 116, 0, 0, 103, 104, 5, 114, 0, 0, 104,
		105, 5, 105, 0, 0, 105, 106, 5, 110, 0, 0, 106, 107, 5, 103, 0, 0, 107,
		2, 1, 0, 0, 0, 108, 109, 5, 73, 0, 0, 109, 110, 5, 110, 0, 0, 110, 111,
		5, 116, 0, 0, 111, 4, 1, 0, 0, 0, 112, 113, 5, 70, 0, 0, 113, 114, 5, 108,
		0, 0, 114, 115, 5, 111, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 116,
		0, 0, 117, 6, 1, 0, 0, 0, 118, 119, 5, 66, 0, 0, 119, 120, 5, 111, 0, 0,
		120, 121, 5, 111, 0, 0, 121, 122, 5, 108, 0, 0, 122, 8, 1, 0, 0, 0, 123,
		124, 5, 67, 0, 0, 124, 125, 5, 104, 0, 0, 125, 126, 5, 97, 0, 0, 126, 127,
		5, 114, 0, 0, 127, 128, 5, 97, 0, 0, 128, 129, 5, 99, 0, 0, 129, 130, 5,
		116, 0, 0, 130, 131, 5, 101, 0, 0, 131, 132, 5, 114, 0, 0, 132, 10, 1,
		0, 0, 0, 133, 134, 5, 118, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 114,
		0, 0, 136, 12, 1, 0, 0, 0, 137, 138, 5, 110, 0, 0, 138, 139, 5, 105, 0,
		0, 139, 140, 5, 108, 0, 0, 140, 14, 1, 0, 0, 0, 141, 142, 5, 116, 0, 0,
		142, 143, 5, 114, 0, 0, 143, 144, 5, 117, 0, 0, 144, 145, 5, 101, 0, 0,
		145, 16, 1, 0, 0, 0, 146, 147, 5, 102, 0, 0, 147, 148, 5, 97, 0, 0, 148,
		149, 5, 108, 0, 0, 149, 150, 5, 115, 0, 0, 150, 151, 5, 101, 0, 0, 151,
		18, 1, 0, 0, 0, 152, 153, 5, 112, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155,
		5, 105, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157, 5, 116, 0, 0, 157, 20,
		1, 0, 0, 0, 158, 159, 5, 105, 0, 0, 159, 160, 5, 102, 0, 0, 160, 22, 1,
		0, 0, 0, 161, 162, 5, 101, 0, 0, 162, 163, 5, 108, 0, 0, 163, 164, 5, 115,
		0, 0, 164, 165, 5, 101, 0, 0, 165, 24, 1, 0, 0, 0, 166, 167, 5, 119, 0,
		0, 167, 168, 5, 104, 0, 0, 168, 169, 5, 105, 0, 0, 169, 170, 5, 108, 0,
		0, 170, 171, 5, 101, 0, 0, 171, 26, 1, 0, 0, 0, 172, 173, 5, 108, 0, 0,
		173, 174, 5, 101, 0, 0, 174, 175, 5, 116, 0, 0, 175, 28, 1, 0, 0, 0, 176,
		177, 5, 102, 0, 0, 177, 178, 5, 111, 0, 0, 178, 179, 5, 114, 0, 0, 179,
		30, 1, 0, 0, 0, 180, 181, 5, 105, 0, 0, 181, 182, 5, 110, 0, 0, 182, 32,
		1, 0, 0, 0, 183, 184, 5, 115, 0, 0, 184, 185, 5, 119, 0, 0, 185, 186, 5,
		105, 0, 0, 186, 187, 5, 116, 0, 0, 187, 188, 5, 99, 0, 0, 188, 189, 5,
		104, 0, 0, 189, 34, 1, 0, 0, 0, 190, 191, 5, 99, 0, 0, 191, 192, 5, 97,
		0, 0, 192, 193, 5, 115, 0, 0, 193, 194, 5, 101, 0, 0, 194, 36, 1, 0, 0,
		0, 195, 196, 5, 100, 0, 0, 196, 197, 5, 101, 0, 0, 197, 198, 5, 102, 0,
		0, 198, 199, 5, 97, 0, 0, 199, 200, 5, 117, 0, 0, 200, 201, 5, 108, 0,
		0, 201, 202, 5, 116, 0, 0, 202, 38, 1, 0, 0, 0, 203, 205, 7, 0, 0, 0, 204,
		203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207,
		1, 0, 0, 0, 207, 214, 1, 0, 0, 0, 208, 210, 5, 46, 0, 0, 209, 211, 7, 0,
		0, 0, 210, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0,
		212, 213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 208, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 215, 40, 1, 0, 0, 0, 216, 220, 5, 34, 0, 0, 217, 219,
		8, 1, 0, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		223, 224, 5, 34, 0, 0, 224, 42, 1, 0, 0, 0, 225, 229, 7, 2, 0, 0, 226,
		228, 7, 3, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227,
		1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 44, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 232, 233, 5, 33, 0, 0, 233, 234, 5, 61, 0, 0, 234, 46, 1, 0, 0, 0,
		235, 236, 5, 61, 0, 0, 236, 237, 5, 61, 0, 0, 237, 48, 1, 0, 0, 0, 238,
		239, 5, 33, 0, 0, 239, 50, 1, 0, 0, 0, 240, 241, 5, 124, 0, 0, 241, 242,
		5, 124, 0, 0, 242, 52, 1, 0, 0, 0, 243, 244, 5, 38, 0, 0, 244, 245, 5,
		38, 0, 0, 245, 54, 1, 0, 0, 0, 246, 247, 5, 61, 0, 0, 247, 56, 1, 0, 0,
		0, 248, 249, 5, 43, 0, 0, 249, 250, 5, 61, 0, 0, 250, 58, 1, 0, 0, 0, 251,
		252, 5, 45, 0, 0, 252, 253, 5, 61, 0, 0, 253, 60, 1, 0, 0, 0, 254, 255,
		5, 62, 0, 0, 255, 256, 5, 61, 0, 0, 256, 62, 1, 0, 0, 0, 257, 258, 5, 60,
		0, 0, 258, 259, 5, 61, 0, 0, 259, 64, 1, 0, 0, 0, 260, 261, 5, 62, 0, 0,
		261, 66, 1, 0, 0, 0, 262, 263, 5, 60, 0, 0, 263, 68, 1, 0, 0, 0, 264, 265,
		5, 42, 0, 0, 265, 70, 1, 0, 0, 0, 266, 267, 5, 47, 0, 0, 267, 72, 1, 0,
		0, 0, 268, 269, 5, 43, 0, 0, 269, 74, 1, 0, 0, 0, 270, 271, 5, 45, 0, 0,
		271, 76, 1, 0, 0, 0, 272, 273, 5, 37, 0, 0, 273, 78, 1, 0, 0, 0, 274, 275,
		5, 40, 0, 0, 275, 80, 1, 0, 0, 0, 276, 277, 5, 41, 0, 0, 277, 82, 1, 0,
		0, 0, 278, 279, 5, 123, 0, 0, 279, 84, 1, 0, 0, 0, 280, 281, 5, 125, 0,
		0, 281, 86, 1, 0, 0, 0, 282, 283, 5, 58, 0, 0, 283, 88, 1, 0, 0, 0, 284,
		285, 5, 63, 0, 0, 285, 90, 1, 0, 0, 0, 286, 287, 5, 46, 0, 0, 287, 92,
		1, 0, 0, 0, 288, 290, 7, 4, 0, 0, 289, 288, 1, 0, 0, 0, 290, 291, 1, 0,
		0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0,
		293, 294, 6, 46, 0, 0, 294, 94, 1, 0, 0, 0, 295, 296, 5, 47, 0, 0, 296,
		297, 5, 42, 0, 0, 297, 301, 1, 0, 0, 0, 298, 300, 9, 0, 0, 0, 299, 298,
		1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 301, 299, 1, 0,
		0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 5, 42, 0, 0,
		305, 306, 5, 47, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 6, 47, 0, 0, 308,
		96, 1, 0, 0, 0, 309, 310, 5, 47, 0, 0, 310, 311, 5, 47, 0, 0, 311, 315,
		1, 0, 0, 0, 312, 314, 8, 5, 0, 0, 313, 312, 1, 0, 0, 0, 314, 317, 1, 0,
		0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0,
		317, 315, 1, 0, 0, 0, 318, 319, 6, 48, 0, 0, 319, 98, 1, 0, 0, 0, 320,
		321, 5, 92, 0, 0, 321, 322, 7, 6, 0, 0, 322, 100, 1, 0, 0, 0, 9, 0, 206,
		212, 214, 220, 229, 291, 301, 315, 1, 6, 0, 0,
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

// LexInit initializes any static state used to implement Lex. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLex(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexInit() {
	staticData := &LexLexerStaticData
	staticData.once.Do(lexLexerInit)
}

// NewLex produces a new lexer instance for the optional input antlr.CharStream.
func NewLex(input antlr.CharStream) *Lex {
	LexInit()
	l := new(Lex)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LexLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lex.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lex tokens.
const (
	LexSTRINGS      = 1
	LexINTS         = 2
	LexFLOATS       = 3
	LexBOOLS        = 4
	LexCHARACTERS   = 5
	LexVAR          = 6
	LexNIL          = 7
	LexTRUE         = 8
	LexFALSE        = 9
	LexPRINT        = 10
	LexIF           = 11
	LexELSE         = 12
	LexWHILE        = 13
	LexLET          = 14
	LexFOR          = 15
	LexIN           = 16
	LexSWITCH       = 17
	LexCASE         = 18
	LexDEFAULT      = 19
	LexNUMBER       = 20
	LexSTRING       = 21
	LexID           = 22
	LexDIF          = 23
	LexIG_IG        = 24
	LexNOT          = 25
	LexOR           = 26
	LexAND          = 27
	LexIG           = 28
	LexIG_ADD       = 29
	LexIG_SUB       = 30
	LexMAY_IG       = 31
	LexMEN_IG       = 32
	LexMAYOR        = 33
	LexMENOR        = 34
	LexMUL          = 35
	LexDIV          = 36
	LexADD          = 37
	LexSUB          = 38
	LexPERCENT      = 39
	LexPARIZQ       = 40
	LexPARDER       = 41
	LexLLAVEIZQ     = 42
	LexLLAVEDER     = 43
	LexDOUBLEPTS    = 44
	LexQUESTION     = 45
	LexPOINT        = 46
	LexWHITESPACE   = 47
	LexCOMMENT      = 48
	LexLINE_COMMENT = 49
)
