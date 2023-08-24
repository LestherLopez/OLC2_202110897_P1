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
| increaseanddecreasestmt {$inst = $increaseanddecreasestmt.increasedecrease}
| whilestmt {$inst = $whilestmt.while}
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
   
blockelifs returns [[]interface{} blkef]
@init{
$blkef = []interface{}{}
    var listifs []IIfstmtContext
  }
: elif+=ifstmt+
    {
        listifs = localctx.(*BlockelifsContext).GetElif()
        for _, a := range listifs {
            $blkef = append($blkef, a.GetIft())
        }
    }
;

ifstmt  returns [interfaces.Instruction ift, []interface{} el, interfaces.Instruction else]
: IF expr  LLAVEIZQ ifb=block LLAVEDER { $ift = instructions.NewIf($IF.line, $IF.pos, $expr.e, $ifb.blk, nil); 
                                        $el = $ifb.blk;}
| IF expr  LLAVEIZQ ifelseblck=block LLAVEDER ELSE LLAVEIZQ elseifblck=block LLAVEDER { $ift = instructions.NewIf($IF.line, $IF.pos, $expr.e, $ifelseblck.blk, $elseifblck.blk); 
                                                                                        $el = $ifelseblck.blk;
                                                                                        }
| IF expr  LLAVEIZQ elif=block LLAVEDER ELSE  blockelifs { $ift = instructions.NewIf($IF.line, $IF.pos, $expr.e, $elif.blk, $blockelifs.blkef);}
;


 

assignationstmt returns [interfaces.Instruction assign]
: ID IG expr {$assign = instructions.NewAssignation($IG.line, $IG.pos,  $ID.text, $expr.e)}
;

whilestmt returns [interfaces.Instruction while]
: WHILE expr LLAVEIZQ block LLAVEDER {$while = instructions.NewWhile($WHILE.line, $WHILE.pos,  $expr.e, $block.blk)}
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
| op = SUB  left=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e,  $op.text, nil)}
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

increaseanddecreasestmt returns [interfaces.Instruction increasedecrease]
: ID IG_ADD expr {$increasedecrease = instructions.NewIncreaseDecrease($ID.line, $ID.pos, $ID.text, $IG_ADD.text, $expr.e)}
| ID IG_SUB expr {$increasedecrease = instructions.NewIncreaseDecrease($ID.line, $ID.pos, $ID.text, $IG_SUB.text, $expr.e)}
;


type returns [environment.TipoExpresion t]
  : STRINGS {$t = environment.STRING}      
  | INTS {$t = environment.INTEGER}         
  | FLOATS {$t = environment.FLOAT}  
  | BOOLS       {$t = environment.BOOLEAN }  
  | CHARACTERS    {$t = environment.CHARACTER}  
;

