package formating

import (
	"encoding/json"
	"bytes"
	"strings"
	"fmt"
	e "../errors_check"
)

func PrettyJsonPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func LastIndexOfChar (s, c string) int {
	i := strings.LastIndex(s, c)
	return i
}

func PrintResponseUrl(b []byte) {
	body,err := PrettyJsonPrint(b)
	e.DefaultErrorCheck(err)
	fmt.Println(string(body))
}

