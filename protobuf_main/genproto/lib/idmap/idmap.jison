%lex
%s COMMENT1
%s COMMENT2
%s STRING

%%
'//'              this.begin('COMMENT1');
'/*'              this.begin('COMMENT2');
'"'               this.begin('STRING');
'{'               return 'LBRACE';
'}'               return 'RBRACE';
'='               return 'EQ';
';'               return 'SEMICOLON';
<COMMENT1>[^\n]*  this.begin('INITIAL');
<COMMENT2>.*'*/'  this.begin('INITIAL');
<STRING>.*'"'     this.begin('INITIAL'); return 'STRING';
[\r?\n]+          return 'NL'
\s+               /* skip whitespace */
<<EOF>>           return 'EOF';
'pbfiles'         return 'PBFILES';
'enum'            return 'ENUM';
'include'         return 'INCLUDE';
'package'         return 'PACKAGE';
[0-9]+            return 'NUM';
\w\w+        return 'IDENT';

/lex

%start expressions

%%

expressions
    : item
    | expressions item
    ;

item
    : idDecl
    | PACKAGE IDENT { yy.setPackageName($2) }
    | PBFILES STRING { yy.setPbFiles($2.substr(0, $2.length-1)) }
    | INCLUDE STRING { yy.includeFile($2.substr(0, $2.length-1)) }
    | enum
    | EOF
    | NL
    ;

idDecl
    : NUM IDENT NUM NL+ { yy.addIdMap($1, $2, $3) }
    ;

enum
    : ENUM IDENT LBRACE NL* enumFields RBRACE { yy.addEnum($2, $5) }
    ;

enumFields
    : enumField { $$ = [$1] }
    | enumFields enumField { $$ = $1.concat($2) }
    ;

enumField
    : IDENT EQ NUM SEMICOLON？NL+ { $$ = {name: $1, id: $3} }
    ;
