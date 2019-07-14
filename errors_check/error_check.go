package errors_check

import (
	"log"
	"fmt"
	"os"
)

func Fatal(msg string, e error){
	if e != nil {log.Fatalf("%v: %v\n", msg, e.Error())}
}

func DefaultErrorCheck(e error) {
	if e != nil{fmt.Println(e)}
}

func DirExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory <%v> does not exist!\n", dir)
	}
}