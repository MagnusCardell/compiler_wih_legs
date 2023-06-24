# Writing a compiler from scratch

- lexer
- parser
- Abstract Syntax Tree
- internal object system
- evaluator

To run:
- cd into with_legs folder
- > go run .

To test:
- cd into with_legs folder
- > go test ./lexer


## Lexical analysis

Tokenize source code into a form that is easier to work with

source code is input, output is recognized tokens
- lexer needs to recognize where the input stream contains characters to read the complete identifier/keyword until it encounters a non-letter character
- lexer needs distinguish between identifier and keywords 

## REPL
Read Eval Print Loop

Console/interpreter to give input to be read, evaluated, and returned. Interpreted language

## Parsing

Turn input into a data structure. Most interpreters use abstract syntax tree. Abstract because its not a full representation of the input. All input syntax is needed for the parser to understand and create the syntax tree. 

### Parser generators
Parser generators take a formal definition of a language to produce parsers as output. This parser can produce the syntax tree we want. Majority of existing parser generators ( yacc, bison, ANTLR)  use context-free grammar. A common notation of CFGs are Backus-Naur Form (BNF). Also Extended BNF (EBNF). While this can be considered a solved problem with excellent solutions and formal proofs, I attempt to write one from scratch to understand the problems. 

```
let <identifier> = <expression>;
```
expressions produce values but the whole row is a statement. The distinction between statement and expression vary for the language. For example, can functions be 