package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrOutOfRange = errors.New("column > len of columns")
	ErrMonth = errors.New("month doesn't exists")
	month = map[string]int{
		"january": 1,
		"february": 2,
		"march": 3,
		"april": 4,
		"may": 5,
		"june": 6,
		"july": 7,
		"august": 8,
		"september": 9,
		"october": 10,
		"november": 11,
		"december": 12,
	}
)

// TODO log to file
func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	storage := []string{}
	for scanner.Scan() {
		storage = append(storage, scanner.Text())
	}

	return storage
}

func getValue(str string, col int, sep string) (string, error) {
	values := strings.Split(str, sep)
	if col >= len(values) || col < 0 {
		return "", ErrOutOfRange
	}
	return values[col], nil
}

func getUnique(rows []string) []string {
	set := make(map[string]struct{})
	res := []string{}
	for _, r := range rows {
		if _, ok := set[r]; !ok {
			res = append(res, r)
			set[r] = struct{}{}
		}
	}
	return res
}

type CompareFunc func(a, b string) (bool, error)

// True if a < b
func compareAsStrings(a, b string) (bool, error) {
	return a < b, nil
}

func compareAsNums(a, b string) (bool, error) {
	v1, err := strconv.Atoi(a)
	if err != nil {
		return false, err
	}

	v2, err := strconv.Atoi(b)
	if err != nil {
		return false, err
	}

	return v1 < v2, nil
}

func compareAsMonths(a, b string) (bool, error){
	a, b = strings.ToLower(a), strings.ToLower(b)
	m1, ok := month[a]
	if !ok{
		return false, ErrMonth
	}

	m2, ok := month[b]
	if !ok{
		return false, ErrMonth
	}

	return m1 < m2, nil
}

//TODO Last 3
func main() {
	k := flag.Int("k", 1, "")
	r := flag.Bool("r", false, "")
	n := flag.Bool("n", false, "")
	u := flag.Bool("u", false, "")
	m := flag.Bool("M", false, "")

	flag.Parse()

	rows := ReadFile(flag.Arg(0))
	//init compare strategy
	var compareStrategy = compareAsStrings
	if *n {
		compareStrategy = compareAsNums
	} else if *m{
		compareStrategy = compareAsMonths
	}

	compareFunction := func(i, j int) bool {
		r1, err := getValue(rows[i], *k-1, " ")
		if err != nil {
			log.Fatal(err.Error())
		}

		r2, err := getValue(rows[j], *k-1, " ")
		if err != nil {
			log.Fatal(err.Error())
		}

		res, err := compareStrategy(r1, r2)
		if err != nil {
			log.Fatal(err.Error())
		}

		return res == !*r

	}

	sort.Slice(rows, compareFunction)

	if *u {
		rows = getUnique(rows)
	}

	fmt.Println(rows)
}
