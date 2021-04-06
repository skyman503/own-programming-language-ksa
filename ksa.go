package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"os/exec"
	"strconv"
	"log"
)

func chcekFile(e error){
	if(e != nil){
		panic(e)
	}
}


func main(){
	// input file specified
	if len(os.Args) > 1{
		//open sourcecode file
		compiledFileName := "a.out"
		if len(os.Args) > 2{
			compiledFileName = os.Args[2]
		}
		sourceCodeFileName := os.Args[1]
		sourceCodeFile, err := os.Open(sourceCodeFileName)
		
		//create new file
		outputC, err := os.Create("outputSourceCode.c")
		chcekFile(err)


		if err != nil {
			fmt.Println(err)
		} else {


			//write gerenratet C to its new file
			outputC.WriteString("#include<stdlib.h>\n#include<stdio.h>\nint main(){\nchar memory[5000];\nchar* memoryStart = &memory[0];\nchar* memoryPointer = memoryStart;\n")


			//creating scanner to read lie by line
			scanner := bufio.NewScanner(sourceCodeFile)
			for scanner.Scan() {
				//read line
				line := scanner.Text()
				
				//chcek if line is the end of the program
				if line == "!"{
					break
				}else{
					//split each line by spaces
					commandArguments := strings.Fields(line)
					memoryCellNumber := "0"
					// check if first elelemt in aray is a memory cell number
					if _, err := strconv.Atoi(commandArguments[0]); err == nil {
						//change memorycell number to string
						memoryCellNumber = commandArguments[0]
						//load operator
						operator := commandArguments[1]

						switch operator{
						//assaign
						case "=":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							//genebrate output string
							if(memoryOperatorCellNumber[0] != '('){
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryOperatorCellNumber + ");"

							}else{
								memoryOperatorCellNumber = memoryOperatorCellNumber[1:len(memoryOperatorCellNumber)-1]
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = " + memoryOperatorCellNumber + ";"
							}	
							//write to outputCode
							outputC.WriteString(outputLine+"\n")
						//add
						case "+":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							if(memoryOperatorCellNumber[0] != '('){
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") + *(memoryStart+" + memoryOperatorCellNumber + ");"
						
							}else{
								memoryOperatorCellNumber = memoryOperatorCellNumber[1:len(memoryOperatorCellNumber)-1]
								outputLine = "*(memoryStart+" + memoryCellNumber + ") += " + memoryOperatorCellNumber + ";"
							}
							//write to outputCode
							outputC.WriteString(outputLine+"\n")
						//subtract
						case "-":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							if(memoryOperatorCellNumber[0] != '('){
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") - *(memoryStart+" + memoryOperatorCellNumber + ");"
						
							}else{
								memoryOperatorCellNumber = memoryOperatorCellNumber[1:len(memoryOperatorCellNumber)-1]
								outputLine = "*(memoryStart+" + memoryCellNumber + ") -= " + memoryOperatorCellNumber + ";"
							}
							//write to outputCode
							outputC.WriteString(outputLine+"\n")
						//multiply
						case "*":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							if(memoryOperatorCellNumber[0] != '('){
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") * *(memoryStart+" + memoryOperatorCellNumber + ");"
						
							}else{
								memoryOperatorCellNumber = memoryOperatorCellNumber[1:len(memoryOperatorCellNumber)-1]
								outputLine = "*(memoryStart+" + memoryCellNumber + ") *= " + memoryOperatorCellNumber + ";"
							}
							//write to outputCode
							outputC.WriteString(outputLine+"\n")
						//divide
						case "/":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							if(memoryOperatorCellNumber[0] != '('){
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") / *(memoryStart+" + memoryOperatorCellNumber + ");"
						
							}else{
								memoryOperatorCellNumber = memoryOperatorCellNumber[1:len(memoryOperatorCellNumber)-1]
								outputLine = "*(memoryStart+" + memoryCellNumber + ") /= " + memoryOperatorCellNumber + ";"
							}
							//write to outputCode
							outputC.WriteString(outputLine+"\n")	
						//negation
						case "~":
							outputLine := "*(memoryStart+" + memoryCellNumber + ") = ~(*(memoryStart+" + memoryCellNumber + "));"
							outputC.WriteString(outputLine+"\n")
						//display
						case ".":		
							outputLine := ""
							if len(commandArguments) > 2{								
								if commandArguments[2] == "d"{
									outputLine = "printf(\"%d\", *(memoryStart+" + memoryCellNumber + "));"
								}else{
									outputLine = "printf(\"%c\", *(memoryStart+" + memoryCellNumber + "));"
								}
							}else{
								outputLine = "printf(\"%c\", *(memoryStart+" + memoryCellNumber + "));"
							}							
							outputC.WriteString(outputLine+"\n")
						//input
						case "^":
							//load 3rd argument
							memoryOperatorCellNumber := commandArguments[2]
							outputLine := ""
							if(memoryOperatorCellNumber == "d"){
								outputLine = "scanf(\"%d\", (memoryStart+" + memoryCellNumber + "));"
							}else{
								outputLine = "scanf(\"%c\", (memoryStart+" + memoryCellNumber + "));"
							}
							outputC.WriteString(outputLine+"\n")
						//pointer type
						case "@":
							innerOperator := commandArguments[2]
							switch(innerOperator){
							//assaign
							case "=":
								//load 3rd argument
								memoryOperatorCellNumber := commandArguments[3]
								outputLine := ""
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + "*(memoryStart+" + memoryOperatorCellNumber  + "));"
								//write to outputCode
								outputC.WriteString(outputLine+"\n")
							//add
							case "+":
								//load 3rd argument
								memoryOperatorCellNumber := commandArguments[3]
								outputLine := ""
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") + *(memoryStart+" + "*(memoryStart+" + memoryOperatorCellNumber  + "));"
								//write to outputCode
								outputC.WriteString(outputLine+"\n")
							//subtract
							case "-":
								//load 3rd argument
								memoryOperatorCellNumber := commandArguments[3]
								outputLine := ""
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") - *(memoryStart+" + "*(memoryStart+" + memoryOperatorCellNumber  + "));"
								//write to outputCode
								outputC.WriteString(outputLine+"\n")
							//multiply
							case "*":
								//load 3rd argument
								memoryOperatorCellNumber := commandArguments[3]
								outputLine := ""
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") * *(memoryStart+" + "*(memoryStart+" + memoryOperatorCellNumber  + "));"
								//write to outputCode
								outputC.WriteString(outputLine+"\n")
							//divide
							case "/":
								//load 3rd argument
								memoryOperatorCellNumber := commandArguments[3]
								outputLine := ""
								//genebrate output string
								outputLine = "*(memoryStart+" + memoryCellNumber + ") = *(memoryStart+" + memoryCellNumber + ") / *(memoryStart+" + "*(memoryStart+" + memoryOperatorCellNumber  + "));"
								//write to outputCode
								outputC.WriteString(outputLine+"\n")	

							}

						}


					}else{
						functionName := commandArguments[0]
						switch(functionName){
						case "if":
							outputLine := ""
							comparator := commandArguments[1]
							memoryCellNumber := commandArguments[2]
							jumpEnd := commandArguments[3]

							if comparator == "<"{
								outputLine = "if(*(memoryStart+" + memoryCellNumber +") < 0){ goto " + jumpEnd + ";}"
							}else if comparator == "="{
								outputLine = "if(*(memoryStart+" + memoryCellNumber +") == 0){ goto " + jumpEnd + ";}"
							}
							
							outputC.WriteString(outputLine + "\n")

						case "$":
							jumpEnd := commandArguments[1]
							outputLine := "goto " + jumpEnd + ";"
							outputC.WriteString(outputLine + "\n")
						default:							
							labelName := commandArguments[0]
							outputLine := labelName + ":"
							outputC.WriteString(outputLine + "\n")							
						}

					}
					
				}

			}
		}


		// savig outout code to memory
		outputC.WriteString("*(memoryStart+9999) = 0;\n}")

		sourceCodeFile.Close()
		outputC.Close()
		//compling C
		cmd := exec.Command("gcc", "outputSourceCode.c", "-o", compiledFileName)
		err = cmd.Run()
		//removing temporary c file
		e := os.Remove("outputSourceCode.c")
		if e != nil {
			log.Fatal(e)
		}

	} else {
		//input file not specified
		fmt.Print("no sourceCode file scpecified")
	}
}