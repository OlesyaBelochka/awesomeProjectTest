package sectionThree_refactoringType

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	s1 := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), ",", "."), ";")
	a1, _ := strconv.ParseFloat(s1[0], 64)
	a2, _ := strconv.ParseFloat(s1[1], 64)

	fmt.Println(fmt.Sprintf("%.4f", a1/a2))
}
