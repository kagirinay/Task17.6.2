package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandInt Генерирует случайные числа
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().Unix())
	c1 := make(chan int)
	c2 := make(chan int)
	var method int
	// Отправляет сообщения
	sender := func() {
		for {
			//С помощью этой переменной будем определять
			//в какой канал отправлять сообщение
			chance := RandInt(1, 100)
			//Если число в chance меньше либо равно 50, отправляем сообщение в канал c1,
			//если больше 50, то в канал c2
			<-time.Tick(time.Second * time.Duration(RandInt(1, 5)))
			if chance <= 50 {
				c1 <- RandInt(1, 100)
			} else {
				c2 <- RandInt(1, 100)
			}
		}
	}

	fmt.Println("Метод прослушивание каналов (1 - time.After, 2 - default): ")
	_, err := fmt.Scanln(&method)
	if err != nil {
		fmt.Println("Неверное значение, метод установлен по умолчанию на 1")
		method = 1
	}
	if method > 2 || method < 1 {
		method = 1
	}

	go sender()

	if method == 1 {
		//Пример с time.After() (ожидание ответа от каждого канала 2 сек)
		for {
			select {
			case num := <-c1:
				fmt.Println("Канал с1 принял сообщение: ", num)
			case num := <-c2:
				fmt.Println("Канал с2 принял сообщение: ", num)
			case <-time.After(2 * time.Second):
				fmt.Println("Время: ", time.Now().Format("15:04"))
			}
		}
	} else {
		//Пример с default
		for {
			select {
			case num := <-c1:
				fmt.Println("Канал с1 принял сообщение: ", num)
			case num := <-c2:
				fmt.Println("Канал с2 принял сообщение: ", num)
			default:
				fmt.Println("Время: ", time.Now().Format("15:04"))
			}
		}
	}
}
