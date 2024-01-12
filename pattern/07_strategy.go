package pattern

type NumbersOperation interface{
	apply(a, b int) int
}

type Add struct{}

func (add *Add) apply(a int, b int) int {
	return a + b
}

type Multiply struct{}

func (m *Multiply) apply(a int, b int) int {
	return a * b
}

type context struct{
	op1 NumbersOperation
	op2 NumbersOperation
	op3 NumbersOperation
}

func (c *context) OperationWithTwoNumbers(a, b int) int{
	return c.op3.apply(c.op1.apply(a, b), c.op2.apply(a, b)) // Будет в итоге следующее (a +/* b) +/* (a +/* b). Это будет в нашем конкретном случае, однако могут появиться новые операции,  которые усложнят наше выражение
}


/*
	Позволяет инкапсулировать алго логику. Нам остается лишь добавлять новую алго логику посредстов создания новых структур, реализующих нужный интерфейс.
	Пример: Калькулятор. Для каждой операции определить стратегию сложения, при нажатии использовать их.
*/
