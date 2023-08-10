package main

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"Server/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type TreeShapeListener struct {
	*parser.BaseGrammarListener
	Code []interface{}
}


func main() {
	
	app := fiber.New()
	app.Use(cors.New())
	//Entrada
	app.Post("/interpreter", func(c *fiber.Ctx) error {
		var request struct {
			Code string `json:"code"`
		}

		if err := c.BodyParser(&request); err != nil {
			return err
		}
	
		//Leyendo entrada
		input := antlr.NewInputStream(request.Code)
		lexer := parser.NewLex(input)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		//creacion de parser
		p := parser.NewGrammarParser(tokens)
		p.BuildParseTrees = true
		tree := p.S()
		//listener
		var listener *TreeShapeListener = NewTreeShapeListener()
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		Code := listener.Code
		//create ast
		var Ast environment.AST
		//ejecución
		for _, inst := range Code {
			err := inst.(interfaces.Instruction).Ejecutar(&Ast, nil)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Error executing code",
			})
		}
		}
		fmt.Println(Ast.GetPrint())
		return c.JSON(Ast.GetPrint())
		})
		app.Listen(":5000")
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}