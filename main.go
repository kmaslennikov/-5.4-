package main

import "fmt"

func main() {
	fmt.Println("Введите наменалы 3х монет")
	var coin1, coin2, coin3 int
	fmt.Scan(&coin1)
	fmt.Scan(&coin2)
	fmt.Scan(&coin3)

	fmt.Println("Введите цену товара")
	var cost int
	fmt.Scan(&cost)
	if cost <= coin1+coin2+coin3 {
		if cost == coin1 || cost == coin2 || cost == coin3 || cost == coin1+coin2 || cost == coin1+coin3 || cost == coin2+coin3 || cost == coin1+coin2+coin3 {
			fmt.Println("Пользователь может заплатить за товар без сдачи")
		} else {
			fmt.Println("Пользователь не может заплатить за товар без сдачи")
		}
	} else {
		fmt.Println("Суммы не достаточно для покупки товара")
	}
}
