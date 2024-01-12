package pattern

import "fmt"

type BrakeBehav interface {
	brake()
}

type BrakeWithLog struct{}

func (b *BrakeWithLog) brake() {
	fmt.Println("brake with rollback")
	fmt.Println("[LOG] brake")
}

type BrakeWithoutLog struct{}

func (b *BrakeWithoutLog) brake() {
	fmt.Println("*brake without log")
}

type CallToMicro struct{
	behav BrakeBehav
}

func (c *CallToMicro) call(){
	c.behav.brake()
}

/*
	Паттерн состояние позволяет изменять поведение структур прибигая к OCP. Можем менять поведение структуры добавляя новые реализации интерфейса поведения и вставляя их в структуру.
	Пример: Допустим система логирования может менять способ вывода в зависимости от уровня(состояние) логирования.
*/
