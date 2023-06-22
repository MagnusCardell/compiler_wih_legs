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
