grammar tinyc;
primaryExpression
    :   Identifier
    |   Constant
    |   StringLiteral+
    |   '(' expression ')'
    ;
postfixExpression: primaryExpression ('[' expression ']' | funcCall)*;
funcCall: '(' (assignmentExpression (',' assignmentExpression)*)? ')';
unaryExpression: unaryOperator castExpression;
unaryOperator: '+' | '-' | '!';
castExpression
    :   postfixExpression
    |   unaryExpression
    ;
multiplicativeExpression: castExpression (('*'|'/'|'%') castExpression)*;
additiveExpression: multiplicativeExpression (('+'|'-') multiplicativeExpression)*;
relationalExpression:   additiveExpression (('<'|'>'|'<='|'>=') additiveExpression)*;
equalityExpression:   relationalExpression (('=='| '!=') relationalExpression)*;
logicalAndExpression:   equalityExpression ('&&' equalityExpression)*;
logicalOrExpression:   logicalAndExpression ( '||' logicalAndExpression)*;
assignmentExpression
    :   logicalOrExpression
    |   postfixExpression assignmentOperator assignmentExpression
    ;
assignmentOperator: '=';
expression:   assignmentExpression (',' assignmentExpression)*;
declaration:   declarationSpecifiers initDeclaratorList? ';';
declarationSpecifiers:   typeSpecifier+;
initDeclaratorList:   initDeclarator (',' initDeclarator)*;
initDeclarator:   declarator ('=' initializer)?;
typeSpecifier
    :   'char'
    |   'int'
    |   'float'
    ;
structDeclaratorList:   structDeclarator (',' structDeclarator)*;
structDeclarator
    :   declarator
    |   declarator? ':' logicalOrExpression
    ;
declarator
    :   Identifier
    |   '(' declarator ')'
    |   declarator '[' assignmentExpression? ']'
    |   declarator '(' identifierList? ')'
    ;

parameterList
    :   parameterDeclaration (',' parameterDeclaration)*
    |
    ;
parameterDeclaration:   declarationSpecifiers declarator;
identifierList:   Identifier (',' Identifier)*;
initializer
    :   assignmentExpression
    |   '{' initializerList ','? '}'
    ;
initializerList:   designation? initializer (',' designation? initializer)*;
designation:   designatorList '=';
designatorList:   designator+;
designator
    :   '[' logicalOrExpression ']'
    |   '.' Identifier
    ;
statement
    :   labeledStatement
    |   compoundStatement
    |   expressionStatement
    |   selectionStatement
    |   iterationStatement
    |   jumpStatement
    ;
labeledStatement:   Identifier ':' statement;
compoundStatement
    :   '{' blockItem* '}'
    ;
blockItem
    :   statement
    |   declaration
    ;
expressionStatement:   expression? ';';

selectionStatement
    :   'if' '(' expression ')' statement ('else' statement)?
    ;

iterationStatement
    :   While '(' expression ')' statement
    ;

jumpStatement: 'return' expression ';' ;

compilationUnit
    :   externalDeclaration+ EOF
    ;

externalDeclaration
    :   functionDefinition
    |   declaration
    |   ';' // stray ;
    ;

functionDefinition
    :   declarationSpecifiers? Identifier '(' parameterList ')' declarationList? compoundStatement
    ;

declarationList
    :   declaration+
    ;

Char : 'char';
Const : 'const';
Else : 'else';
Float : 'float';
If : 'if';
Int : 'int';
Return : 'return';
While : 'while';
LeftParen : '(';
RightParen : ')';
LeftBracket : '[';
RightBracket : ']';
LeftBrace : '{';
RightBrace : '}';
Less : '<';
LessEqual : '<=';
Greater : '>';
GreaterEqual : '>=';
Plus : '+';
Minus : '-';
Star : '*';
Div : '/';
Mod : '%';
And : '&';
Or : '|';
AndAnd : '&&';
OrOr : '||';
Caret : '^';
Not : '!';
Tilde : '~';
Question : '?';
Colon : ':';
Semi : ';';
Comma : ',';
Assign : '=';
Equal : '==';
NotEqual : '!=';

Identifier
    :   IdentifierNondigit
        (   IdentifierNondigit
        |   Digit
        )*
    ;

fragment
IdentifierNondigit
    :   Nondigit
    |   UniversalCharacterName
    //|   // other implementation-defined characters...
    ;

fragment
Nondigit
    :   [a-zA-Z_]
    ;

fragment
Digit
    :   [0-9]
    ;

fragment
UniversalCharacterName
    :   '\\u' HexQuad
    |   '\\U' HexQuad HexQuad
    ;

fragment
HexQuad
    :   HexadecimalDigit HexadecimalDigit HexadecimalDigit HexadecimalDigit
    ;

Constant
    :   IntegerConstant
    |   FloatingConstant
    |   CharacterConstant
    ;

fragment
IntegerConstant
    :   DecimalConstant
    |   OctalConstant
    |   HexadecimalConstant
    |	BinaryConstant
    ;

fragment
BinaryConstant
	:	'0' [bB] [0-1]+
	;

fragment
DecimalConstant
    :   NonzeroDigit Digit*
    ;

fragment
OctalConstant
    :   '0' OctalDigit*
    ;

fragment
HexadecimalConstant
    :   HexadecimalPrefix HexadecimalDigit+
    ;

fragment
HexadecimalPrefix
    :   '0' [xX]
    ;

fragment
NonzeroDigit
    :   [1-9]
    ;

fragment
OctalDigit
    :   [0-7]
    ;

fragment
HexadecimalDigit
    :   [0-9a-fA-F]
    ;

fragment
FloatingConstant
    :   DecimalFloatingConstant
    |   HexadecimalFloatingConstant
    ;

fragment
DecimalFloatingConstant
    :   FractionalConstant ExponentPart? FloatingSuffix?
    |   DigitSequence ExponentPart FloatingSuffix?
    ;

fragment
HexadecimalFloatingConstant
    :   HexadecimalPrefix (HexadecimalFractionalConstant | HexadecimalDigitSequence) BinaryExponentPart FloatingSuffix?
    ;

fragment
FractionalConstant
    :   DigitSequence? '.' DigitSequence
    |   DigitSequence '.'
    ;

fragment
ExponentPart
    :   [eE] Sign? DigitSequence
    ;

fragment
Sign
    :   [+-]
    ;

DigitSequence
    :   Digit+
    ;

fragment
HexadecimalFractionalConstant
    :   HexadecimalDigitSequence? '.' HexadecimalDigitSequence
    |   HexadecimalDigitSequence '.'
    ;

fragment
BinaryExponentPart
    :   [pP] Sign? DigitSequence
    ;

fragment
HexadecimalDigitSequence
    :   HexadecimalDigit+
    ;

fragment
FloatingSuffix
    :   [flFL]
    ;

fragment
CharacterConstant:   '\'' CCharSequence '\'';

fragment
CCharSequence
    :   CChar+
    ;

fragment
CChar
    :   ~['\\\r\n]
    |   EscapeSequence
    ;

fragment
EscapeSequence
    :   SimpleEscapeSequence
    |   OctalEscapeSequence
    |   HexadecimalEscapeSequence
    |   UniversalCharacterName
    ;

fragment
SimpleEscapeSequence
    :   '\\' ['"?abfnrtv\\]
    ;

fragment
OctalEscapeSequence
    :   '\\' OctalDigit OctalDigit? OctalDigit?
    ;

fragment
HexadecimalEscapeSequence
    :   '\\x' HexadecimalDigit+
    ;

StringLiteral
    :   '"' SCharSequence? '"'
    ;

fragment
SCharSequence
    :   SChar+
    ;

fragment
SChar
    :   ~["\\\r\n]
    |   EscapeSequence
    |   '\\\n'   // Added line
    |   '\\\r\n' // Added line
    ;


LineDirective
    :   '#' Whitespace? DecimalConstant Whitespace? StringLiteral ~[\r\n]*
        -> skip
    ;

Whitespace
    :   [ \t]+
        -> skip
    ;

Newline
    :   (   '\r' '\n'?
        |   '\n'
        )
        -> skip
    ;

BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;

LineComment
    :   '//' ~[\r\n]*
        -> skip
    ;
