# KSA - personal programing language
## Simple Assembly programing language

KSA is a personal project, making my dream of my own programing language come true.

## Features

- Fast to develop
- Easy to learn
- Easy to make Your personal modifications


## Tech

Compiler for ksa is written in Go, after compiling a temporary C file is created, then gcc compiles temporary C file, and any leftovers after C compilations are being deleted. 

## Usage
After building GO file type in Your terminal:
```
ksa.exe [nameOfTheSourceCodeFile] [nameOfTheExecutableOutput](optional)
```

## Installation
1. Install GO on Your PC
2. Make sure You have on gcc Your PC
3. Clone this repository
4. In Your repo directory open terminal and type
``` 
go build ksa.go
```
5. (Optional) Add ksa.exe to PATH

## Syntax
1. There must be a white caracter beetwen each operands
2. Each instruction must be written in new line

### Assign a value to a variable (Note, that You need to add () around Your value, if You don't, whatever is on the right side of equal-sign will be interpreted as variable name) :
```
x = (10)
y = ('a')
```
### Ccopy one value to the other variable:
```
z = x
```
### Mathematical operations on Your variable [+ - / *].
Where:
1. ``` x + (2) ``` will add 2 to the x and override existing x
2. ``` x + y```will add the value of y to the x and override existing x
3. Same goes for all of the other signs
4. ``` x ~``` will nagete the bits of current x value and override existing x

### Jump to certain code fragment
1. Line that performs the jump: ```$ nameOfTheJumpPlace```
2. Line that the jump will be performed to: ```*nameOfTheJumpPlace```(Note: no space bettwen * and name of the place)

### If statements
1. Start line by typing ```if```
2. Then follow with either ```<``` or ```=```, depedning if You want to check if the value is less than or equal to 0
3. After that type the variable name that You are testing ```x```
4. And, at the very end, type the name of the place You want to jump if condition is true ```*endLoopPlace```(Note: that if the condition is false lines after it will be executed, and no jump will happen)

Examples:
```
if = x finishProgram
```
```
if < y begigingOfTheProgram
```

### I/O
I/O instruction must start with the variable name it will be performed on
Then, for input ```^``` and for output ```.```
And after that, mode in which You want to use I/O ```c``` when reading/displing characters and ```d``` when using numbers (Note: that You cannot display number of character that is not part of any variable)
Example:
```
x . d
```
```
y ^ c
```
### Comments
Start the comment line with ```//```
Example:
```//This is the coment line, the code here won't be executed```

### End program
To end executing, make a new line with single ```!```
Example:
```
!
```

## Demo programs
### Display numbers from 1 to 10 (each in new line)
```
howManyNumbers = (10)
newLine = (10)
newLine . c
counter = howManyNumbers
//loop will go until couter is 0
*beginningOfTheLoop
if = counter endOfTheLoop
tmp = howManyNumbers
tmp - counter
tmp + (1)
tmp . d
newLine . c
counter - (1)
$ beginningOfTheLoop
*endOfTheLoop
//end of the program
!
```

### Ask User for a number and then display the same number squared
```
//Asking fo a number
number ^ d
number * number
//Displaying the number squared
number . d
!
```
