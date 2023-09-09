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
		"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'",
		"'continue'", "'return'", "'break'", "'append'", "'removeLast'", "'remove'",
		"'at'", "'isEmpty'", "'count'", "'func'", "'inout'", "'struct'", "",
		"", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'+='", "'-='",
		"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('",
		"')'", "'{'", "'}'", "':'", "'?'", "'.'", "','", "';'", "'['", "']'",
		"'_'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL",
		"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN",
		"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK",
		"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC",
		"INOUT", "STRUCT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT",
		"OR", "AND", "IG", "IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", "POINT", "COMA", "PTCOMA",
		"CORCHETEIZQ", "CORCHETEDER", "GUION_BAJO", "AND_SIMPLE", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", "TRUE",
		"FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", "SWITCH",
		"CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", "APPEND",
		"REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC", "INOUT", "STRUCT",
		"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG",
		"IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV",
		"ADD", "SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"DOUBLEPTS", "QUESTION", "POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER",
		"GUION_BAJO", "AND_SIMPLE", "WHITESPACE", "COMMENT", "LINE_COMMENT",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 68, 461, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 4, 32, 331, 8, 32,
		11, 32, 12, 32, 332, 1, 32, 1, 32, 4, 32, 337, 8, 32, 11, 32, 12, 32, 338,
		3, 32, 341, 8, 32, 1, 33, 1, 33, 5, 33, 345, 8, 33, 10, 33, 12, 33, 348,
		9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 5, 34, 354, 8, 34, 10, 34, 12, 34, 357,
		9, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51,
		1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1,
		56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61,
		1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 4, 65, 428, 8, 65, 11,
		65, 12, 65, 429, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 5, 66, 438,
		8, 66, 10, 66, 12, 66, 441, 9, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1,
		67, 1, 67, 1, 67, 1, 67, 5, 67, 452, 8, 67, 10, 67, 12, 67, 455, 9, 67,
		1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 439, 0, 69, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99,
		50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115,
		58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129, 65, 131,
		66, 133, 67, 135, 68, 137, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9,
		10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35,
		43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 467, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0,
		0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1,
		0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0,
		125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0,
		0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 1, 139, 1, 0, 0, 0, 3, 146,
		1, 0, 0, 0, 5, 150, 1, 0, 0, 0, 7, 156, 1, 0, 0, 0, 9, 161, 1, 0, 0, 0,
		11, 171, 1, 0, 0, 0, 13, 175, 1, 0, 0, 0, 15, 179, 1, 0, 0, 0, 17, 184,
		1, 0, 0, 0, 19, 190, 1, 0, 0, 0, 21, 196, 1, 0, 0, 0, 23, 199, 1, 0, 0,
		0, 25, 204, 1, 0, 0, 0, 27, 210, 1, 0, 0, 0, 29, 214, 1, 0, 0, 0, 31, 218,
		1, 0, 0, 0, 33, 221, 1, 0, 0, 0, 35, 228, 1, 0, 0, 0, 37, 233, 1, 0, 0,
		0, 39, 241, 1, 0, 0, 0, 41, 247, 1, 0, 0, 0, 43, 256, 1, 0, 0, 0, 45, 263,
		1, 0, 0, 0, 47, 269, 1, 0, 0, 0, 49, 276, 1, 0, 0, 0, 51, 287, 1, 0, 0,
		0, 53, 294, 1, 0, 0, 0, 55, 297, 1, 0, 0, 0, 57, 305, 1, 0, 0, 0, 59, 311,
		1, 0, 0, 0, 61, 316, 1, 0, 0, 0, 63, 322, 1, 0, 0, 0, 65, 330, 1, 0, 0,
		0, 67, 342, 1, 0, 0, 0, 69, 351, 1, 0, 0, 0, 71, 358, 1, 0, 0, 0, 73, 361,
		1, 0, 0, 0, 75, 364, 1, 0, 0, 0, 77, 366, 1, 0, 0, 0, 79, 369, 1, 0, 0,
		0, 81, 372, 1, 0, 0, 0, 83, 374, 1, 0, 0, 0, 85, 377, 1, 0, 0, 0, 87, 380,
		1, 0, 0, 0, 89, 383, 1, 0, 0, 0, 91, 386, 1, 0, 0, 0, 93, 388, 1, 0, 0,
		0, 95, 390, 1, 0, 0, 0, 97, 392, 1, 0, 0, 0, 99, 394, 1, 0, 0, 0, 101,
		396, 1, 0, 0, 0, 103, 398, 1, 0, 0, 0, 105, 400, 1, 0, 0, 0, 107, 402,
		1, 0, 0, 0, 109, 404, 1, 0, 0, 0, 111, 406, 1, 0, 0, 0, 113, 408, 1, 0,
		0, 0, 115, 410, 1, 0, 0, 0, 117, 412, 1, 0, 0, 0, 119, 414, 1, 0, 0, 0,
		121, 416, 1, 0, 0, 0, 123, 418, 1, 0, 0, 0, 125, 420, 1, 0, 0, 0, 127,
		422, 1, 0, 0, 0, 129, 424, 1, 0, 0, 0, 131, 427, 1, 0, 0, 0, 133, 433,
		1, 0, 0, 0, 135, 447, 1, 0, 0, 0, 137, 458, 1, 0, 0, 0, 139, 140, 5, 83,
		0, 0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143, 5, 105,
		0, 0, 143, 144, 5, 110, 0, 0, 144, 145, 5, 103, 0, 0, 145, 2, 1, 0, 0,
		0, 146, 147, 5, 73, 0, 0, 147, 148, 5, 110, 0, 0, 148, 149, 5, 116, 0,
		0, 149, 4, 1, 0, 0, 0, 150, 151, 5, 70, 0, 0, 151, 152, 5, 108, 0, 0, 152,
		153, 5, 111, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 116, 0, 0, 155,
		6, 1, 0, 0, 0, 156, 157, 5, 66, 0, 0, 157, 158, 5, 111, 0, 0, 158, 159,
		5, 111, 0, 0, 159, 160, 5, 108, 0, 0, 160, 8, 1, 0, 0, 0, 161, 162, 5,
		67, 0, 0, 162, 163, 5, 104, 0, 0, 163, 164, 5, 97, 0, 0, 164, 165, 5, 114,
		0, 0, 165, 166, 5, 97, 0, 0, 166, 167, 5, 99, 0, 0, 167, 168, 5, 116, 0,
		0, 168, 169, 5, 101, 0, 0, 169, 170, 5, 114, 0, 0, 170, 10, 1, 0, 0, 0,
		171, 172, 5, 118, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 114, 0, 0,
		174, 12, 1, 0, 0, 0, 175, 176, 5, 110, 0, 0, 176, 177, 5, 105, 0, 0, 177,
		178, 5, 108, 0, 0, 178, 14, 1, 0, 0, 0, 179, 180, 5, 116, 0, 0, 180, 181,
		5, 114, 0, 0, 181, 182, 5, 117, 0, 0, 182, 183, 5, 101, 0, 0, 183, 16,
		1, 0, 0, 0, 184, 185, 5, 102, 0, 0, 185, 186, 5, 97, 0, 0, 186, 187, 5,
		108, 0, 0, 187, 188, 5, 115, 0, 0, 188, 189, 5, 101, 0, 0, 189, 18, 1,
		0, 0, 0, 190, 191, 5, 112, 0, 0, 191, 192, 5, 114, 0, 0, 192, 193, 5, 105,
		0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 116, 0, 0, 195, 20, 1, 0, 0,
		0, 196, 197, 5, 105, 0, 0, 197, 198, 5, 102, 0, 0, 198, 22, 1, 0, 0, 0,
		199, 200, 5, 101, 0, 0, 200, 201, 5, 108, 0, 0, 201, 202, 5, 115, 0, 0,
		202, 203, 5, 101, 0, 0, 203, 24, 1, 0, 0, 0, 204, 205, 5, 119, 0, 0, 205,
		206, 5, 104, 0, 0, 206, 207, 5, 105, 0, 0, 207, 208, 5, 108, 0, 0, 208,
		209, 5, 101, 0, 0, 209, 26, 1, 0, 0, 0, 210, 211, 5, 108, 0, 0, 211, 212,
		5, 101, 0, 0, 212, 213, 5, 116, 0, 0, 213, 28, 1, 0, 0, 0, 214, 215, 5,
		102, 0, 0, 215, 216, 5, 111, 0, 0, 216, 217, 5, 114, 0, 0, 217, 30, 1,
		0, 0, 0, 218, 219, 5, 105, 0, 0, 219, 220, 5, 110, 0, 0, 220, 32, 1, 0,
		0, 0, 221, 222, 5, 115, 0, 0, 222, 223, 5, 119, 0, 0, 223, 224, 5, 105,
		0, 0, 224, 225, 5, 116, 0, 0, 225, 226, 5, 99, 0, 0, 226, 227, 5, 104,
		0, 0, 227, 34, 1, 0, 0, 0, 228, 229, 5, 99, 0, 0, 229, 230, 5, 97, 0, 0,
		230, 231, 5, 115, 0, 0, 231, 232, 5, 101, 0, 0, 232, 36, 1, 0, 0, 0, 233,
		234, 5, 100, 0, 0, 234, 235, 5, 101, 0, 0, 235, 236, 5, 102, 0, 0, 236,
		237, 5, 97, 0, 0, 237, 238, 5, 117, 0, 0, 238, 239, 5, 108, 0, 0, 239,
		240, 5, 116, 0, 0, 240, 38, 1, 0, 0, 0, 241, 242, 5, 103, 0, 0, 242, 243,
		5, 117, 0, 0, 243, 244, 5, 97, 0, 0, 244, 245, 5, 114, 0, 0, 245, 246,
		5, 100, 0, 0, 246, 40, 1, 0, 0, 0, 247, 248, 5, 99, 0, 0, 248, 249, 5,
		111, 0, 0, 249, 250, 5, 110, 0, 0, 250, 251, 5, 116, 0, 0, 251, 252, 5,
		105, 0, 0, 252, 253, 5, 110, 0, 0, 253, 254, 5, 117, 0, 0, 254, 255, 5,
		101, 0, 0, 255, 42, 1, 0, 0, 0, 256, 257, 5, 114, 0, 0, 257, 258, 5, 101,
		0, 0, 258, 259, 5, 116, 0, 0, 259, 260, 5, 117, 0, 0, 260, 261, 5, 114,
		0, 0, 261, 262, 5, 110, 0, 0, 262, 44, 1, 0, 0, 0, 263, 264, 5, 98, 0,
		0, 264, 265, 5, 114, 0, 0, 265, 266, 5, 101, 0, 0, 266, 267, 5, 97, 0,
		0, 267, 268, 5, 107, 0, 0, 268, 46, 1, 0, 0, 0, 269, 270, 5, 97, 0, 0,
		270, 271, 5, 112, 0, 0, 271, 272, 5, 112, 0, 0, 272, 273, 5, 101, 0, 0,
		273, 274, 5, 110, 0, 0, 274, 275, 5, 100, 0, 0, 275, 48, 1, 0, 0, 0, 276,
		277, 5, 114, 0, 0, 277, 278, 5, 101, 0, 0, 278, 279, 5, 109, 0, 0, 279,
		280, 5, 111, 0, 0, 280, 281, 5, 118, 0, 0, 281, 282, 5, 101, 0, 0, 282,
		283, 5, 76, 0, 0, 283, 284, 5, 97, 0, 0, 284, 285, 5, 115, 0, 0, 285, 286,
		5, 116, 0, 0, 286, 50, 1, 0, 0, 0, 287, 288, 5, 114, 0, 0, 288, 289, 5,
		101, 0, 0, 289, 290, 5, 109, 0, 0, 290, 291, 5, 111, 0, 0, 291, 292, 5,
		118, 0, 0, 292, 293, 5, 101, 0, 0, 293, 52, 1, 0, 0, 0, 294, 295, 5, 97,
		0, 0, 295, 296, 5, 116, 0, 0, 296, 54, 1, 0, 0, 0, 297, 298, 5, 105, 0,
		0, 298, 299, 5, 115, 0, 0, 299, 300, 5, 69, 0, 0, 300, 301, 5, 109, 0,
		0, 301, 302, 5, 112, 0, 0, 302, 303, 5, 116, 0, 0, 303, 304, 5, 121, 0,
		0, 304, 56, 1, 0, 0, 0, 305, 306, 5, 99, 0, 0, 306, 307, 5, 111, 0, 0,
		307, 308, 5, 117, 0, 0, 308, 309, 5, 110, 0, 0, 309, 310, 5, 116, 0, 0,
		310, 58, 1, 0, 0, 0, 311, 312, 5, 102, 0, 0, 312, 313, 5, 117, 0, 0, 313,
		314, 5, 110, 0, 0, 314, 315, 5, 99, 0, 0, 315, 60, 1, 0, 0, 0, 316, 317,
		5, 105, 0, 0, 317, 318, 5, 110, 0, 0, 318, 319, 5, 111, 0, 0, 319, 320,
		5, 117, 0, 0, 320, 321, 5, 116, 0, 0, 321, 62, 1, 0, 0, 0, 322, 323, 5,
		115, 0, 0, 323, 324, 5, 116, 0, 0, 324, 325, 5, 114, 0, 0, 325, 326, 5,
		117, 0, 0, 326, 327, 5, 99, 0, 0, 327, 328, 5, 116, 0, 0, 328, 64, 1, 0,
		0, 0, 329, 331, 7, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0,
		332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 340, 1, 0, 0, 0, 334,
		336, 5, 46, 0, 0, 335, 337, 7, 0, 0, 0, 336, 335, 1, 0, 0, 0, 337, 338,
		1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 341, 1, 0,
		0, 0, 340, 334, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 66, 1, 0, 0, 0,
		342, 346, 5, 34, 0, 0, 343, 345, 8, 1, 0, 0, 344, 343, 1, 0, 0, 0, 345,
		348, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349,
		1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 350, 5, 34, 0, 0, 350, 68, 1, 0,
		0, 0, 351, 355, 7, 2, 0, 0, 352, 354, 7, 3, 0, 0, 353, 352, 1, 0, 0, 0,
		354, 357, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356,
		70, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 359, 5, 33, 0, 0, 359, 360,
		5, 61, 0, 0, 360, 72, 1, 0, 0, 0, 361, 362, 5, 61, 0, 0, 362, 363, 5, 61,
		0, 0, 363, 74, 1, 0, 0, 0, 364, 365, 5, 33, 0, 0, 365, 76, 1, 0, 0, 0,
		366, 367, 5, 124, 0, 0, 367, 368, 5, 124, 0, 0, 368, 78, 1, 0, 0, 0, 369,
		370, 5, 38, 0, 0, 370, 371, 5, 38, 0, 0, 371, 80, 1, 0, 0, 0, 372, 373,
		5, 61, 0, 0, 373, 82, 1, 0, 0, 0, 374, 375, 5, 43, 0, 0, 375, 376, 5, 61,
		0, 0, 376, 84, 1, 0, 0, 0, 377, 378, 5, 45, 0, 0, 378, 379, 5, 61, 0, 0,
		379, 86, 1, 0, 0, 0, 380, 381, 5, 62, 0, 0, 381, 382, 5, 61, 0, 0, 382,
		88, 1, 0, 0, 0, 383, 384, 5, 60, 0, 0, 384, 385, 5, 61, 0, 0, 385, 90,
		1, 0, 0, 0, 386, 387, 5, 62, 0, 0, 387, 92, 1, 0, 0, 0, 388, 389, 5, 60,
		0, 0, 389, 94, 1, 0, 0, 0, 390, 391, 5, 42, 0, 0, 391, 96, 1, 0, 0, 0,
		392, 393, 5, 47, 0, 0, 393, 98, 1, 0, 0, 0, 394, 395, 5, 43, 0, 0, 395,
		100, 1, 0, 0, 0, 396, 397, 5, 45, 0, 0, 397, 102, 1, 0, 0, 0, 398, 399,
		5, 37, 0, 0, 399, 104, 1, 0, 0, 0, 400, 401, 5, 40, 0, 0, 401, 106, 1,
		0, 0, 0, 402, 403, 5, 41, 0, 0, 403, 108, 1, 0, 0, 0, 404, 405, 5, 123,
		0, 0, 405, 110, 1, 0, 0, 0, 406, 407, 5, 125, 0, 0, 407, 112, 1, 0, 0,
		0, 408, 409, 5, 58, 0, 0, 409, 114, 1, 0, 0, 0, 410, 411, 5, 63, 0, 0,
		411, 116, 1, 0, 0, 0, 412, 413, 5, 46, 0, 0, 413, 118, 1, 0, 0, 0, 414,
		415, 5, 44, 0, 0, 415, 120, 1, 0, 0, 0, 416, 417, 5, 59, 0, 0, 417, 122,
		1, 0, 0, 0, 418, 419, 5, 91, 0, 0, 419, 124, 1, 0, 0, 0, 420, 421, 5, 93,
		0, 0, 421, 126, 1, 0, 0, 0, 422, 423, 5, 95, 0, 0, 423, 128, 1, 0, 0, 0,
		424, 425, 5, 38, 0, 0, 425, 130, 1, 0, 0, 0, 426, 428, 7, 4, 0, 0, 427,
		426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430,
		1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 432, 6, 65, 0, 0, 432, 132, 1, 0,
		0, 0, 433, 434, 5, 47, 0, 0, 434, 435, 5, 42, 0, 0, 435, 439, 1, 0, 0,
		0, 436, 438, 9, 0, 0, 0, 437, 436, 1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439,
		440, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 442, 1, 0, 0, 0, 441, 439,
		1, 0, 0, 0, 442, 443, 5, 42, 0, 0, 443, 444, 5, 47, 0, 0, 444, 445, 1,
		0, 0, 0, 445, 446, 6, 66, 0, 0, 446, 134, 1, 0, 0, 0, 447, 448, 5, 47,
		0, 0, 448, 449, 5, 47, 0, 0, 449, 453, 1, 0, 0, 0, 450, 452, 8, 5, 0, 0,
		451, 450, 1, 0, 0, 0, 452, 455, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 453,
		454, 1, 0, 0, 0, 454, 456, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 456, 457,
		6, 67, 0, 0, 457, 136, 1, 0, 0, 0, 458, 459, 5, 92, 0, 0, 459, 460, 7,
		6, 0, 0, 460, 138, 1, 0, 0, 0, 9, 0, 332, 338, 340, 346, 355, 429, 439,
		453, 1, 6, 0, 0,
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
	LexGUARD        = 20
	LexCONTINUE     = 21
	LexRETURN       = 22
	LexBREAK        = 23
	LexAPPEND       = 24
	LexREMOVELAST   = 25
	LexREMOVE       = 26
	LexAT           = 27
	LexISEMPTY      = 28
	LexCOUNT        = 29
	LexFUNC         = 30
	LexINOUT        = 31
	LexSTRUCT       = 32
	LexNUMBER       = 33
	LexSTRING       = 34
	LexID           = 35
	LexDIF          = 36
	LexIG_IG        = 37
	LexNOT          = 38
	LexOR           = 39
	LexAND          = 40
	LexIG           = 41
	LexIG_ADD       = 42
	LexIG_SUB       = 43
	LexMAY_IG       = 44
	LexMEN_IG       = 45
	LexMAYOR        = 46
	LexMENOR        = 47
	LexMUL          = 48
	LexDIV          = 49
	LexADD          = 50
	LexSUB          = 51
	LexPERCENT      = 52
	LexPARIZQ       = 53
	LexPARDER       = 54
	LexLLAVEIZQ     = 55
	LexLLAVEDER     = 56
	LexDOUBLEPTS    = 57
	LexQUESTION     = 58
	LexPOINT        = 59
	LexCOMA         = 60
	LexPTCOMA       = 61
	LexCORCHETEIZQ  = 62
	LexCORCHETEDER  = 63
	LexGUION_BAJO   = 64
	LexAND_SIMPLE   = 65
	LexWHITESPACE   = 66
	LexCOMMENT      = 67
	LexLINE_COMMENT = 68
)
