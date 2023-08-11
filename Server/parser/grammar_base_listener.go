// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterDeclarestmt is called when production declarestmt is entered.
func (s *BaseGrammarListener) EnterDeclarestmt(ctx *DeclarestmtContext) {}

// ExitDeclarestmt is called when production declarestmt is exited.
func (s *BaseGrammarListener) ExitDeclarestmt(ctx *DeclarestmtContext) {}

// EnterConstantstmt is called when production constantstmt is entered.
func (s *BaseGrammarListener) EnterConstantstmt(ctx *ConstantstmtContext) {}

// ExitConstantstmt is called when production constantstmt is exited.
func (s *BaseGrammarListener) ExitConstantstmt(ctx *ConstantstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterType is called when production type is entered.
func (s *BaseGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGrammarListener) ExitType(ctx *TypeContext) {}
