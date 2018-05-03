# zip-lang

[![Build Status](https://travis-ci.org/mzrimsek/zip-lang.svg?branch=master)](https://travis-ci.org/mzrimsek/zip-lang)

An interpreter written in Go as described in [Writing An Interpreter In Go](https://interpreterbook.com/). The book is great and I highly recommend it.

It's called Zip because there's no practical purpose to this language other than as a learning exercise in both Go and learning how interpreters work.

## Supported Features

* Builtin Data Types
  * Integer
  * Float
  * String
  * Boolean
  * Character
  * Array
  * Hash
* Flow Control
  * If-Else
  * While
* Variables (Scoped to their block)
* First-class Functions
  * With/without Arguments
  * Closures
* IO
  * Put
  * Putln

## Setup

1. Install Go ```sudo apt-get install golang-go```
2. Clone this repo
3. Install dependencies ```make install```

### Visual Studio Code

* Run the set up script ```./env_setup.sh``` 

### direnv

1. Install direnv ```sudo apt-get install direnv```
2. Edit your ```.bashrc``` and add ```eval "$(direnv hook bash)"``` to the end
3. Activate in root directory of this repo ```direnv allow```

## Testing
* ```make test```

## Starting the REPL
* ```make run```

## Planned Feature Additions
* ~~Float literals~~
* ~~Character literals~~
* ~~Variable assignment~~
* ~~Assignment shortcut operators~~
* ~~Prefix Increment and Decrement operators~~
* ~~Postfix Increment and Decrement operators~~
* ~~Remainder operator~~
* ~~Power operator~~
* ~~AND and OR operators~~
* ~~Less Than or Equal and Greater Than or Equal operators~~
* Escape characters
* ~~String comparison~~
* ~~String multiplication operator~~
* ~~Concatenate basic data types to strings~~
* String interpolation
* String index expressions
* Append to hash structure
* Set array value at index
* Set hash value at key
* Ternary statements
* Single and multi-line comments
* Else-If statements
* Switch statements
* ~~While loops~~
* Until loop
* For loops
* Foreach loops
* Do-While loops
* Break and continue statements
* List comprehensions

## Tenative Planned Features
* Package import
* Global constant variables
* Emojis

## Todo
* Extend test coverage for builtins
* Extend tests to check for more types
* ~~Read input from a file~~
* Add more builtin functions
  * Array and String includes function
  * Array all, any, etc type operators
* Add a wiki with language documentation
* Convert tests to use an assertion library
