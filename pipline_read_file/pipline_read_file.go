/*
Функция предназначена для потоковой обработки.
Функция принимает имя файла, а возвращает канал, в
котором лежат числа прочитанные из файла. 

Файл читается построчно, для экономии памяти, а
символы, не являющиеся числом, пропускаются
*/

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

// Функция возвращает канал только для чтения
func NumbersGen(filename string) <-chan int {
	// Создаем канал
    res := make(chan int)

	// Запускаем поток обработки
    go func() {
        defer close(res)
        
		// Открываем файл
        file, err := os.Open(filename)
        if err != nil {
            fmt.Println("err1")
            return
        }
        defer file.Close()
        
		// Создаем сканнер
        scanner := bufio.NewScanner(file)
		// Считываем данные из файла построчно
        for scanner.Scan() {
            line := scanner.Text()
			// Переводим строку в массив из символов
            arrLine := strings.Split(line, "")
			// Обходим массив символов
            for _, v := range arrLine {
                x, err := strconv.Atoi(v)
				// Если удалось преобразовать символ в число, то записываем это число в канал
                if err != nil {
                    //fmt.Println("err2")
                } else {
                    res <- x
                }
            }
        }
        
    }()
	
	// Возвращаем канал
    return res
}