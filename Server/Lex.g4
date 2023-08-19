lexer grammar Lex;


//types
STRINGS: 'String';
INTS: 'Int';
FLOATS: 'Float';
BOOLS: 'Bool';
CHARACTERS: 'Character';

//keywords
VAR: 'var';
NIL: 'nil';
TRUE: 'true';
FALSE: 'false';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
LET:   'let';
FOR: 'for';
IN: 'in';
//STRING: '"'~['"]*'"' | '\''~["']*'\''; con comilla simple
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z])[a-zA-Z0-9_]*;

// symbols

DIF:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:          '=';
IG_ADD:      '+=';
IG_SUB:      '-=';
MAY_IG:     '>=';
MEN_IG:     '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PERCENT:        '%';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
DOUBLEPTS:      ':';
QUESTION:       '?';
// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

