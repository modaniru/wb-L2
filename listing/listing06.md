Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
3, 2, 3
```

1. В функции main мы создаем слайс с длинной 3 и капасити 3.
2. Передаем копию слайса в функцию modifySlice(). В функции мы изменяем 0 элемент слайса, и так как мы передали копию, то в копии содержится указатель на начало массива, поэтому это отразится на исходном массиве. После чего мы пытаемся добавить массиву новый элемент, но у копии происходит переполнение из-за чего выделяется новый участок памяти под новый массив, из-за чего у копии меняется ссылка на первый элемент массива.
