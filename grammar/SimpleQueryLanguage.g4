grammar SimpleQueryLanguage;

Less: '<';
LessOrEqual: '<=';
Equal: '==';
Greater: '>';
GreaterOrEqual: '>=';
NotEqual: '!=';
And: '&&';
Or: '||';
Not: '!';

//Math operators
Add: '+';
Subtract: '-';
Multiply: '*';
Divide: '/';

LeftParen: '(';
RightParen: ')';
LeftBracket: '[';
RightBracket: ']';

Question: '?';
NullCoalescing: '??';

Dot: '.';
Comma: ',';

LeftBrace: '{';
RightBrace: '}';

Null: 'null';
False: 'false';
True: 'true';

expression
    : field
    ;

function_param
    : field
    ;

field
   : LeftParen field RightParen
   | True
   | False
   | Null
   | NUMBER
   | ESCAPED_STRING
   | array
   | object
   | name
   | field (Question? Dot name)
   | field (Question? LeftBracket index RightBracket)
   | field LeftParen (function_param (Comma function_param)*)? RightParen
   | Subtract field
   | field (Divide | Multiply) field
   | field (Add | Subtract) field
   | field Dot Dot field
   | field NullCoalescing field
   | field Equal field
   | field Less field
   | field LessOrEqual field
   | field Greater field
   | field GreaterOrEqual field
   | field NotEqual field
   | field And field
   | field Or field
   | Not field
   | field Question field (':' field)?
   ;

array
    : LeftBracket (field (Comma field)*)? RightBracket
    ;

object
    : LeftBrace (object_field (Comma object_field)*)? RightBrace
    ;

object_field
    : name ':' field
    | ESCAPED_STRING ':' field
    ;

name
   : STRING
   ;

index
   : field
   | NUMBER
   | ESCAPED_STRING
   ;

ESCAPED_STRING : ('"' ('\\' . | ~["\\])* '"') | ('\'' ('\\' . | ~['\\])* '\'');

STRING
   : [a-zA-Z_$][a-zA-Z0-9_$]*
   ;

NUMBER
   : [0-9]+(Dot[0-9]+)?
   ;

WS
   : [ \r\n\t] -> channel(HIDDEN)
   ;