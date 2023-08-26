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

	// EnterBlockelifs is called when entering the blockelifs production.
	EnterBlockelifs(c *BlockelifsContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterBlockcases is called when entering the blockcases production.
	EnterBlockcases(c *BlockcasesContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterAssignationstmt is called when entering the assignationstmt production.
	EnterAssignationstmt(c *AssignationstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAccessstmt is called when entering the accessstmt production.
	EnterAccessstmt(c *AccessstmtContext)

	// EnterIncreaseanddecreasestmt is called when entering the increaseanddecreasestmt production.
	EnterIncreaseanddecreasestmt(c *IncreaseanddecreasestmtContext)

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

	// ExitBlockelifs is called when exiting the blockelifs production.
	ExitBlockelifs(c *BlockelifsContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitBlockcases is called when exiting the blockcases production.
	ExitBlockcases(c *BlockcasesContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitAssignationstmt is called when exiting the assignationstmt production.
	ExitAssignationstmt(c *AssignationstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAccessstmt is called when exiting the accessstmt production.
	ExitAccessstmt(c *AccessstmtContext)

	// ExitIncreaseanddecreasestmt is called when exiting the increaseanddecreasestmt production.
	ExitIncreaseanddecreasestmt(c *IncreaseanddecreasestmtContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
