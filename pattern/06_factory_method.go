package pattern

import (
	"fmt"
	"log"
)

/*Переменные определяющие название доступных фигур*/
var(
	CIRCLE = "Circle"
	SQUARE = "Square"
	TRIANGLE = "Triangle"
)

/*Интерфейс фигуры, чтобы, возвращая, не привязываться к реализации.*/
type Shape interface {
	draw()
}

/*Конкретные реализации интерфейса Shape*/
type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("Circle")
}

type Square struct{}

func (s *Square) draw() {
	fmt.Println("Square")
}


type Triange struct{}

func (s *Triange) draw() {
	fmt.Println("Triange")
}


/*Интерфейс предоставляющий фабричный метод*/
type FactoryMethodI interface {
	loadShapes(shape string) Shape
}

/*Конкретная фабрика предоставляющая реализацию интерфейса*/
type ConcreteFabric struct{}

func (c *ConcreteFabric) loadShapes(shape string) Shape {
	switch shape{
	case CIRCLE:
		return &Circle{}
	case SQUARE:
		return &Square{}
	case TRIANGLE:
		return &Triange{}
	default:
		log.Fatal("unsopported operation")
		return nil
	}
}

/*
Паттерн фабричный метод предоставляет возможность создания объектов на основе входных данных, причем он возвращает тип интерфейса, что позволяет не привязываться к конкретным реализациям. Это предоставляет клиенту возможность зависить от интерфейса. Также можно поместить фабрику под интерфейс, чтобы в случае необходимости изменить логику можно было создать другую реализацию интерфейса.
*/