package pattern

import "fmt"

//Интерфейс команды, который мы будем встраивать
type LightCommand interface{
	execute(light *Light)
}

//Команда включения света
type turnOnLight struct{}

func (t *turnOnLight) execute(light *Light) {
	fmt.Println("turn on")
}

//Команда включения света циклом (включено-выключено-включено)
type turnOnCycleLight struct{}

func (t *turnOnCycleLight) execute(light *Light) {
	fmt.Println("turn on - off - on - ... - on - ...")
}

//Структура света, которая содержит метод выполнения какой-то команды.
type Light struct{
	turnOnCommand LightCommand
}

//Метод, который выполняет метод команды. Передаем структуру в метод команды, чтобы можно было, если что, обращаться к полям объекта
func (l *Light) turnOn(){
	l.turnOnCommand.execute(l)
}

//Какой-то другой метод
func (l *Light) turnOff(){
	fmt.Println("off")
}

/*
Данный паттерн предоставляет писать новые методы(команды) для структуры, не изменяя при этом исходных код. Достаточно лишь создать структуру имплементирующую нужные методы. Таким образом мы соблюдаем Open Close принцип. Из недостатков - может повыситься сложность кода, если использовать его везде.
*/
func main(){
	cycle := turnOnCycleLight{}
	on := turnOnLight{}

	l := Light{turnOnCommand: &on}
	l.turnOn()
	l.turnOnCommand = &cycle
	l.turnOn()
	l.turnOff()
}