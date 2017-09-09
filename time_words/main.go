package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var words = []string{
	"---",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
	"twenty",
	"thirty",
	"fourty",
	"fifty",
	"sixty",
}

func main() {
	r := bufio.NewReader(os.Stdin)

	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	h, _ := strconv.Atoi(line)

	line, _ = r.ReadString('\n')
	line = strings.TrimSpace(line)
	m, _ := strconv.Atoi(line)

	var fphrase string
	if m <= 30 {
		fphrase = minute(m) + " past %s"
	} else {
		m = 60 - m
		h++
		if h > 12 {
			h = 1
		}
		fphrase = minute(m) + " to %s"
	}

	hSpelt := itoa(h)
	if m == 0 {
		fmt.Printf("%s o' clock\n", hSpelt)
		return
	}

	fmt.Printf(fphrase+"\n", hSpelt)
}

func minute(n int) string {
	var ph string
	switch n {
	case 1:
		ph = "%s minute"
	case 15:
		return "quarter"
	case 30:
		return "half"
	default:
		ph = "%s minutes"
	}
	return fmt.Sprintf(ph, itoa(n))
}

func itoa(n int) string {
	if n <= 20 {
		return words[n]
	}
	tn := n / 10
	tn = 18 + tn
	rn := n % 10
	if rn == 0 {
		return words[tn]
	}
	return fmt.Sprintf("%s %s", words[tn], words[rn])
}
