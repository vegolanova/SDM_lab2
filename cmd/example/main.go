package main 

import (
	//"fmt"
	"github.com/SDM_lab2"
)

func main() {
	thisList := initDList()
	thisList.Append("C")
	thisList.Append("C")
	thisList.Append("C")
	thisList.Append("C")
	thisList.Append("B")
	thisList.Append("C")
	thisList.Append("R")
	thisList.Append("C")
	thisList.Append("D")
	thisList.Append("V")
	thisList.Append("C")
	thisList.Append("C")
	thisList.displayData()
	thisList.Delete(24)
	thisList.displayData()
}
