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
: block EOF
    {   
        $code = $block.blk
    }
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
| ifstmt { }
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
;
//declaracion de variables
declarestmt 
: VAR ID DOUBLEPTS type IG expr //declaracion con tipo y valor
| VAR ID IG expr //declaracion con valor
| VAR ID DOUBLEPTS type QUESTION //declaracion con tipo y sin valor
;
//declaracion de constantes
constantstmt
: LET ID DOUBLEPTS type IG expr //declaracion con tipo y valor
| LET ID IG expr //declaracion con valor
;

ifstmt  
: IF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER
| IF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER ELSE LLAVEIZQ block LLAVEDER
| IF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER ELSE ifstmt
;

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
;

type
  : STRINGS      
  | INTS        
  | FLOATS          
  | BOOLS       
  | CHARACTERS             
; 


/* primitive [interfaces.Expression p]  
: NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $p = expressions.NewPrimitive($STRING.line, $STRING.pos,str[1:len(str)-1],environment.STR)
    }                        
| TRU { $p = expressions.NewPrimitive($TRU.line, $TRU.pos,true,environment.BOOLEAN) }
| FAL { $p = expressions.NewPrimitive($FAL.line, $FAL.pos,false,environment.BOOLEAN) }
; */