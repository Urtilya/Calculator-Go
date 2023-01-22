package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	diverr  = errors.New("division by zero")
	romeerr = errors.New("invald input syntax")
	zeroerr = errors.New("result below one")
)

func removespace(str []string) []string {
	out := make([]string, 0)
	for _, element := range str {
		if element != "" {
			out = append(out, element)
		}
	}
	return out
}

func readformula() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	sep := strings.Split(text, " ")
	sep = removespace(sep)
	if len(sep) == 3 {
		return sep, err
	} else {
		return sep, romeerr
	}

}

func inarb(form []string) (a int, b int, err error) {
	ROME_NUM := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	a = ROME_NUM[form[0]]
	b = ROME_NUM[form[2]]
	if a == 0 || b == 0 {
		err = romeerr
	}
	return
}

func inrome(a int) (string, error) {
	ROME_NUM := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	ROME_DEC := map[int]string{
		1:  "X",
		2:  "XX",
		3:  "XXX",
		4:  "XL",
		5:  "L",
		6:  "LX",
		7:  "LXX",
		8:  "LXXX",
		9:  "XC",
		10: "C",
	}
	if a < 1 {
		return "", zeroerr
	}
	var ret string
	if a >= 10 {
		ret = ret + ROME_DEC[a/10]
		a = a % 10
	}
	return ret + ROME_NUM[a], nil
}

func calculate(a int, b int, sighn string) (int, error) {
	if (a > 10) || (a < 1) || (b > 10) || (b < 1) {
		return 0, romeerr
	}
	switch sighn {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, diverr
		}
		return a / b, nil
	}
	return 0, romeerr
}

func main() {
	var err error
	sep, err := readformula()
	if err != nil {
		panic(err.Error())
	}
	var rome bool
	var a, b int
	if sep[0][0] == 'I' || sep[0][0] == 'V' || sep[0][0] == 'X' {
		rome = true
		a, b, err = inarb(sep)
		if err != nil {
			panic(romeerr.Error())
		}
	} else {
		rome = false
		a, err = strconv.Atoi(sep[0])
		if err != nil {
			panic(romeerr.Error())
		}
		b, err = strconv.Atoi(sep[2])
		if err != nil {
			panic(romeerr.Error())
		}
	}
	res, err := calculate(a, b, sep[1])
	if err != nil {
		panic(err.Error())
	}
	if rome {
		resrome, err := inrome(res)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(resrome)
	} else {
		fmt.Println(res)
	}
}
