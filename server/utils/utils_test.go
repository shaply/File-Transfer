package utils_test

import (
	"fmt"
	"testing"

	"github.com/shaply/File-Transfer/server/utils"
)

func TestCodeGenerator(t *testing.T) {
	code := utils.GenerateCode()
	fmt.Println(code)
	for i := 0; i < 10; i++ {
		newCode := utils.GenerateCode()
		fmt.Println(newCode)
	}
}
