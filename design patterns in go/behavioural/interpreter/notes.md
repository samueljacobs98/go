# Interpreter

## Motivation

- Textual input needs to be processed
    - E.g., turned into linked structures
    - AST = Abstract Syntax Tree
- Some examples
    - Programming language compilers, interpreters and IDEs
    - HTML, XML and similar
    - Numeric expressions (3+4/5)
    - Regular expressions
- Turning strings into linked structures in a complicated process

An interpreter is a design pattern where a component processes structured text data. It does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

## Summary

- Barring simple cases, an interpreter acts in two stages
- Lexing turns text into a set of tokens, e.g. `3*(4+5) -> Lit[3] Star Lapren Lit[4] Plus Lit[5] Rparen`
- Parsing tokens into meaninful constructs (AST = Abstract Syntax Tree)
- Parsed data can then be traversed using the Vistor pattern