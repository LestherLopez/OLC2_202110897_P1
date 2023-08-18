grammar Grammar; 


options {
  tokenVocab = Lex;
}

@header {
    import "Server/Interfaces"
    import "Server/Environment"
    import "Server/Expression"
    import "Server/Instruction"
    import "strings"
}


s returns [[]interface{} code]
: block EOF {$code = $block.blk}
;

block returns [[]interface{} blk]
@init{
$blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: printstmt { $inst = $printstmt.prnt}
| declarestmt {$inst = $declarestmt.dec}
| constantstmt {$inst = $constantstmt.const}
| ifstmt { $inst = $ifstmt.ift }
| assignationstmt {$inst = $assignationstmt.assign}
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line, $PRINT.pos, $expr.e)}
;
//declaracion de variables
//lin int, col int, id_var string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool
declarestmt returns [interfaces.Instruction dec]
: VAR ID DOUBLEPTS type IG expr {$dec = instructions.NewTodeclare($VAR.line, $VAR.pos, $ID.text, $type.t, $expr.e, false)}//declaracion con tipo y valor
| VAR ID IG expr {$dec = instructions.NewTodeclare($VAR.line, $VAR.pos, $ID.text, environment.NULL, $expr.e, false)}//declaracion con valor
| VAR ID DOUBLEPTS type QUESTION {$dec = instructions.NewTodeclare($VAR.line, $VAR.pos, $ID.text, $type.t, nil, false)}//declaracion con tipo y sin valor
;
//declaracion de constantes
constantstmt returns [interfaces.Instruction const]
: LET ID DOUBLEPTS type IG expr {$const = instructions.NewTodeclare($LET.line, $LET.pos, $ID.text, $type.t, $expr.e, true)} //declaracion con tipo y valor
| LET ID IG expr {$const = instructions.NewTodeclare($LET.line, $LET.pos, $ID.text, environment.NULL, $expr.e, true)}//declaracion con valor
;

//(lin, col, exp_conditional, sentence, sentence else)
ifstmt  returns [interfaces.Instruction ift]
: IF PARIZQ expr PARDER LLAVEIZQ instruction LLAVEDER { $ift = instructions.NewIf($IF.line, $IF.pos, $expr.e, $instruction.inst, nil) }
| IF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER ELSE LLAVEIZQ block LLAVEDER 
| IF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER ELSE ifstmt
;

assignationstmt returns [interfaces.Instruction assign]
: ID IG expr {$assign = instructions.NewAssignation($IG.line, $IG.pos,  $ID.text, $expr.e)}
;

whilestmt returns [interfaces.Instruction while]
: WHILE expr LLAVEIZQ block LLAVEDER
;

forstmt returns [interfaces.Instruction for]
: FOR ID IN  expr LLAVEIZQ block LLAVEDER
;
//----------------------------EXPRESIONES---------------------
expr returns [interfaces.Expression e]
: left=expr op=(MUL|DIV) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB|PERCENT) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| op= NOT   left=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $left.e) }
| PARIZQ expr PARDER { $e = $expr.e } 
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRUE { $e = expressions.NewPrimitive($TRUE.line, $TRUE.pos, true, environment.BOOLEAN) }
| FALSE { $e = expressions.NewPrimitive($FALSE.line, $FALSE.pos, false, environment.BOOLEAN) }
| NIL { $e = expressions.NewPrimitive($NIL.line, $NIL.pos, nil, environment.NULL) }
| accessstmt {$e = $accessstmt.access}
;

accessstmt returns [interfaces.Expression access]
: op=ID {$access = expressions.NewAccess($op.line, $op.pos, $op.text)}
;

type returns [environment.TipoExpresion t]
  : STRINGS {$t = environment.STRING}      
  | INTS {$t = environment.INTEGER}         
  | FLOATS {$t = environment.FLOAT}  
  | BOOLS       {$t = environment.BOOLEAN }  
  | CHARACTERS    {$t = environment.CHARACTER}  
;

