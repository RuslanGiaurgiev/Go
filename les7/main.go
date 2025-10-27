package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

// Создаём структуру Car с одним лишь значчением armor
type Car struct {
	armor int
}

func (c *Car) Gas() (int, error) {
	if c.armor-10 <= 0 {
		return 0, errors.New("Армора не хватает, разгон запрещен")
	}

	speed := rand.Intn(150)

	c.armor -= 10

	return speed, nil

}

func main() {

	car := Car{
		armor: 40,
	}

	for {
		pp.Println("До:", car)
		speed, err := car.Gas()
		pp.Println("После:", car)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("Получившийся разгон:", speed)
		fmt.Println("")

	}

}
