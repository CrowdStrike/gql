We are using [gqlparser](https://github.com/vektah/gqlparser) for parsing GraphQL schema.
Lexer in gqlparser drops all the comments when it returns the ast. Since we are using comments to enable/disable lint rules
we are using a slightly modified (only change being not dropping the comments) lexer from gqlparser.

If in future gqlparser supports this functionality or there is another lexer which can give us access to comments in schema, 
we would remove this copied code.  