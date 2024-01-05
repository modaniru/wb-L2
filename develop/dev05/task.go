package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	A := flag.Int("A", 0, "count after equals")
	B := flag.Int("B", 0, "count before equals")
	C := flag.Int("C", 0, "count between equals")
	c := flag.Bool("c", false, "count")
	ignoreCase := flag.Bool("i", false, "ignore case")
	n := flag.Bool("n", false, "display line num")
	F := flag.Bool("F", false, "fixed")
	v := flag.Bool("v", false, "invert")

	flag.Parse()
	str := flag.Arg(0)
	f := flag.Arg(1)

	if *ignoreCase {
		str = strings.ToLower(str)
	}

	if *F {
		str = regexp.QuoteMeta(str)
	}

	r, err := os.Open(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer r.Close()

	count := 0
	scan := bufio.NewScanner(r)
	lines := []string{}
	indexes := []int{}
	for i := 0; scan.Scan(); i++ {
		temp := scan.Text()
		line := temp
		if *ignoreCase {
			line = strings.ToLower(line)
		}

		ok, err := regexp.MatchString(str, line)
		if err != nil {
			log.Fatal(err.Error())
		}

		lines = append(lines, temp)

		if ok {
			count++
			indexes = append(indexes, i)
		}
	}

	if *c {
		fmt.Println(count)
		return
	}

	if *C != 0 {
		A = C
		B = C
	}

	if len(indexes) == 0 {
		return
	}

	if *v {
		index := 0
		for i, line := range lines {
			if index < len(indexes) && indexes[index] == i {
				index++
				continue
			}
			if *n {
				fmt.Printf("%d. ", i+1)
			}
			fmt.Println(line)
		}
		return
	}

	for _, i := range indexes[:len(indexes)-1] {
		PrintLines(i-*B, i+*A, *n, lines)
		fmt.Println("------------")
	}
	PrintLines(indexes[len(indexes)-1]-*B, indexes[len(indexes)-1]+*A, *n, lines)
}

func PrintLines(start, end int, n bool, lines []string) {
	if start < 0 {
		start = 0
	}

	if end >= len(lines) {
		end = len(lines) - 1
	}

	for j := start; j <= end; j++ {
		if n {
			fmt.Printf("%d. ", j+1)
		}
		fmt.Println(lines[j])
	}
}
