// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterDeclarestmt is called when entering the declarestmt production.
	EnterDeclarestmt(c *DeclarestmtContext)

	// EnterConstantstmt is called when entering the constantstmt production.
	EnterConstantstmt(c *ConstantstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitDeclarestmt is called when exiting the declarestmt production.
	ExitDeclarestmt(c *DeclarestmtContext)

	// ExitConstantstmt is called when exiting the constantstmt production.
	ExitConstantstmt(c *ConstantstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}