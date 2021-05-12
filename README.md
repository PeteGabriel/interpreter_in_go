# Interpreter in Go


This project reflects the study of the book "Interpreter in Go" writen by Thorsten Ball.

## Language used - Monkey

"Without a compiler or an interpreter a programming
language is nothing more than an idea or a specification."

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

From the example above we can identify tokens like an integer, a keyword or even variable names. We'll distinguish 
from types, keywords and identifiers (variable and function names) among others. We can specify these in our code
by using constants. "ILLEGAL" will denote something we are not expecting and "EOF" will mark the end of our reading
process.


## The Parser

tbc