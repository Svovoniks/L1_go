package main

import (
	"fmt"
	"strings"
)

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

var justString string

func someFuncFix1() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func someFuncFix2() {
    v := []rune(createHugeString(1 << 10))[:100]
    justString = string(v)
}

func someFuncFix3() {
	justString = createHugeString(100)
}

// Представленный код может привести к нескольким
// негативным последствиям
// 1. Созданная подстрока будет ссылатья на отрезок из целой строки
//    и при присвоении её к глобальной переменной вся строка будет хранится в
//    памяти хотя доступа к элементам за границами подстроки уже не будет
// 2. Обрезание строки с [:100] может привести к неправильному разделению строки
//    например если она содержит эмоджи
// Исправить данный код можно несколькими способами:
// 1. Можно клонировать подстроку
// 2. Можно перевести строку в []rune, обрезать его и перевести обратно в строку
//    это также решит проблему с неправильным разделением строки
// 3. Если есть такая возможность можно сразу создавать строку нужного размера

func main() {
	someFuncFix1()
    fmt.Println(justString)
	someFuncFix2()
    fmt.Println(justString)
	someFuncFix3()
    fmt.Println(justString)
}
