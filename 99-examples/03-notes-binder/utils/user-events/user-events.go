package userevents

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserData(m string) string {
	fmt.Print(m)

	bufioReader := bufio.NewReader(os.Stdin)
	inputData, _ := bufioReader.ReadString('\n')

	content := strings.TrimSuffix(inputData, "\n")
	content = strings.TrimSuffix(content, "\r") // For Windows compatibility

	return content
}
