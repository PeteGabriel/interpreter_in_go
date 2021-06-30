# Interpreter in Go

This project reflects the study of the book "Interpreter in Go" writen by Thorsten Ball.

## Language used - Monkey

"Without a compiler or an interpreter a programming language is nothing more than an idea or a specification."

Monkey has the following list of features:

* C-like syntax
* variable bindings
* integers and booleans
* arithmetic expressions
* built-in functions
* first-class and higher-order functions
* closures
* a string data structure
* an array data structure
* a hash data structure

Some examples of this language usage:

```
//some variables
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

```
//hash-map
let thorsten = {"name": "Thorsten", "age": 28};
```

Or something more complex

```
let twice = fn(f, x) {
  return f(f(x));
};
let addTwo = fn(x) {
  return x + 2;
};
twice(addTwo, 2); // => 6
```

## Parts of an interpreter:

* the lexer
* the parser
* the Abstract Syntax Tree (AST)
* the internal object system
* the evaluator

## The Lexer

```
let five = 5;
let ten = 10;
let add = fn(x, y) {
  x + y;
};
let result = add(five, ten);
```

From the example above we can identify tokens like an integer, a keyword or even variable names. We'll distinguish from
types, keywords and identifiers (variable and function names) among others. We can specify these in our code by using
constants. "ILLEGAL" will denote something we are not expecting and "EOF" will mark the end of our reading process.

The lexer will ignore spaces since Monkey language does not care for them. Also `_` are supported as part of variable
names.

## The Parser

Basically, a parser turns its input into a data structure that represents the input and checks its correctness in the 
process. This component is responsible for parser errors.

As an example:

```
var input = 'if (3 * 5 > 10) { return "hello"; } else { return "goodbye"; }';
var tokens = MagicLexer.parse(input);
MagicParser.parse(tokens);
```
```javascript
{
  type: "if-statement",
  condition: {
    type: "operator-expression",
    operator: ">",
    left: {
      type: "operator-expression",
      operator: "*",
      left: { type: "integer-literal", value: 3 },
      right: { type: "integer-literal", value: 5 }
    },
    right: { type: "integer-literal", value: 10 }
  },
  consequence: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "hello" }
  },
  alternative: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "goodbye" }
  }
}
```

## Abstract Syntax Tree

For this project this data structure will be basically just composed by nodes connected to each other. This structure will 
represent the program running. Represents the syntax almost exactly and omits irrelevant details like whitespace or so 
(because these dont matter in our language). The implementation can be seen in the package [ast](./ast).


## Parsing expressions

Parsing statements is fairly straightforward. Reading from left to right and identify which keyword and from there parse
 the rest of the statement. Parsing expressions is a bit more complicated. One of the problems building a parser is the
operator precedence. The fact that expressions can be found in many different situations is also a problem that we need 
to take care of by applying a a correct parsing procedure that is understandable and extensible since the beginning. 

![diagram](https://i.imgur.com/oo9UNwR.png)

We can see below an high level representation of the flow for parsing an expression, the most complicated type of parsing in our interpreter. Because of recursivity we don't need many lines for each parsing function. 
