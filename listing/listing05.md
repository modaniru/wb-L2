Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Тут опять стоит обратиться к внутреннему устройству интерфейсов. В функции test мы возвращаем указатель на customError с значением nil. Далее в функции main мы присваиваем ее переменной статического интерфейсного типа error. Из-за этого в устройстве error возникает следующее: data = nil, itab != nil, из этого следует, что ошибка не равна nil. Вот поэтому выводит error, хотя никакой ошибки нет.