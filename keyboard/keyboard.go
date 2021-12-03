//Package keyboard 可以读取用户在键盘上的输入
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat 将用户在键盘上输入的字符去除前后空格，并转换成float64类型
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
