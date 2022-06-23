package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (s *service) GenerateFile(params *GenerateFileParams) error {
	year, month, day := time.Now().Date()
	hour, minute, second := time.Now().Clock()
	title := strings.Join([]string{
		strconv.Itoa(year),
		strconv.Itoa(int(month)),
		strconv.Itoa(day),
	}, "-") + " " + strings.Join([]string{
		strconv.Itoa(hour),
		strconv.Itoa(minute),
		strconv.Itoa(second),
	}, ":")

	fmt.Println(title)
	return nil
}
