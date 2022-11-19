package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	//File input
	csvFileName:= flag.String("csv","problems.csv","A csv filename in the form of \n ans and question")
	flag.Parse()

	file,err:= os.Open(*csvFileName)

	if err!=nil{
		ErrorStatus("Unable to open the sv file")
	}
	r:=csv.NewReader(file)
	lines,err:=r.ReadAll()
	if err!=nil{
		ErrorStatus("Cannot Open File")
	}
	problemBank:= parseLine(lines)

	var correctAnswer int =0

	for i,problem :=range problemBank{
		fmt.Printf("Problem %d \n Question : %s \n",i+1,problem.q)
		var answer string
		fmt.Scanf("%s \n",&answer)
		if answer==problem.a{
			fmt.Println("Correct Boy !")
			correctAnswer++
		}else{
			fmt.Println("Wrong Answer Boy !")
		}
	}

	fmt.Println("The Quiz is Over")
	fmt.Printf("Your Score: \n  %d",correctAnswer)
}


func ErrorStatus(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

// A slice is required
func parseLine(lines [][] string) []problemSet{
	
	parsedLines:= make([]problemSet,len(lines))

	for i,line:=range lines{
		parsedLines[i]=problemSet{
			q: line[0],
			a:line[1],
		}
	}

	return parsedLines
}


// Problem set type 
type problemSet struct{
	q string
	a string
}