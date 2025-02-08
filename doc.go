/*

file app.int:
    [postgres]
    host=192.168.1.2
    port=5432
    db=books
    conns=5


package main

import (
	"fmt"

	"conf/conf"
)

// Структура, которая может быть разной
type ss struct {
	ID     int
	Host   string
	Db     string
	Conns  int
	Inp    string
	Zip    string
	Key3   bool
	List   string
	Passwd string
}

func main() {
	fmt.Println("Чтени переменных из файла config")

	// Пример структуры
	c := ss{}
	// можно и так
	// s1 := struct {
	// 	Host  string
	// 	Db    string
	// 	Conns int
	// }{}
	// fmt.Println(s1)

	ini, err := conf.NewConf("app.ini")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := ini.Read("[import]", &c); err == nil {
		fmt.Println("Заполненная структура import:")
		fmt.Println("Imp: ", c.Inp)
		fmt.Println("Zip: ", c.Zip)
		fmt.Println("List:", c.List)
	}
	if err := ini.Read("[postgres]", &c); err == nil {
		fmt.Println("Заполненная структура postgres:")
		fmt.Println("Host:  ", c.Host)
		fmt.Println("Db:    ", c.Db)
		fmt.Println("Passwd:", c.Passwd)
		fmt.Println("Conns: ", c.Conns)
	}

}

*/
