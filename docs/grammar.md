### Grammar

```
code = {statement};

statement =
  variable definition
  | expression;
variable definition = IDENTIFIER, "=", expression;

expression = addition;
addition = multiplication, [("+" | "-"), addition];
multiplication = exponentiation, [("*" | "/" | "%"), multiplication];
exponentiation = atom, ["^", exponentiation];

atom =
  INTEGER NUMBER
  | FLOATING-POINT NUMBER
  | IDENTIFIER
  | function call
  | ("(", expression, ")");
function call = IDENTIFIER, "(", [expression, {",", expression}], ")";

COMMENT = ? /\/\/.*/ ?;
INTEGER NUMBER = ? /\b\d+(e[+-]?\d+)?\b/i ?;
FLOATING-POINT NUMBER = ? /\b(\.\d+|\d+\.\d*)(e[+-]?\d+)?\b/i ?;
IDENTIFIER = ? /[a-z_]\w*/i ?;
```
