package user

import "fmt"

type User struct {
	name    string
	age     int
	premium bool
}

func Printinfo() {

	users := [5]User{
		User{
			name:    "Артур",
			age:     19,
			premium: false,
		},
		User{
			name:    "Альберт",
			age:     19,
			premium: true,
		},
		User{
			name:    "Зухро",
			age:     19,
			premium: false,
		},
		User{
			name:    "Диана",
			age:     19,
			premium: false,
		},
		User{
			name:    "Данила",
			age:     19,
			premium: true,
		},
	}
	for i := 0; i < 5; i++ {
		if users[i].premium {
			fmt.Println(users[i].name, "Красавчик")

		} else {
			fmt.Println(users[i].name, "Тебе надо качаться мэн")
		}

	}
}
