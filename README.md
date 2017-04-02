# go-interpreter
An interpreter written in Go as described in [Writing An Interpreter In Go](https://interpreterbook.com/). The book is great and I highly recommend it.

## Planned Feature Additions
* Float literals
* Character literals
* Variable redefining
* Increment and decrement operators
* ~~AND and OR operators~~
* Escape characters
* ~~String comparison~~
* ~~Concatenate basic data types to strings~~
* String index expressions
* Ternary statements
* ~~Less Than or Equal and Greater Than or Equal operators~~
* Single and multi-line comments
* Else-If statements
* Switch statements
* While loops
* For loops
* Do-While loops
* Break and continue statements
* Dynamic list structure
* Emojis

### Todo
* Extend test coverage for builtins
* Extend tests to check for more types
* Read input from a file

## Setup
1. Install Go ```sudo apt-get install golang-go```
2. Clone this repo
3. Run the set up script ```./env_setup.sh```

The setup script I've written works best with Visual Studio Code.  
If you dont' have vscode installed the script will still set your GOPATH, but I recommend installing direnv to manage your GOPATH automatically.  

### Setup direnv
1. Install direnv ```sudo apt-get install direnv```
2. Edit your ```.bashrc``` and add ```eval "$(direnv hook bash)"``` to the end
3. Activate root directory of this repo ```direnv allow```
