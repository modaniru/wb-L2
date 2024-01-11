package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	d := flag.String("d", " ", "seporator")
	f := flag.Int("f", 1, "fields")
	s := flag.Bool("s", false, "separated")
	flag.Parse()

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan(){
		line := scan.Text()
		fields := slices.DeleteFunc[[]string, string](strings.Split(line, *d), func(s string) bool {
			return s == ""
		})
		if len(fields) < *f || (*s && !strings.Contains(line, *d)){
			continue
		}

		fmt.Println(fields[*f - 1])
	}
}
