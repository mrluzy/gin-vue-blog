package res

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type (
	ErrorCode int
	ErrorMap  map[ErrorCode]string
)

var ErrMap = ErrorMap{}

func ReadErrorCodeJson() {
	filePath := "utils/res/error_code.json"

	byteErrorCode, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("read file failed with error: " + err.Error())
		return
	}

	err = json.Unmarshal(byteErrorCode, &ErrMap)
	if err != nil {
		log.Fatalf("unmarshal json failed with error: " + err.Error())
		return
	}
	fmt.Println(ErrMap)
}
