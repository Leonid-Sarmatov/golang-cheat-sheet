/*
Функция предназначена для потоковой обработки.

Первая функция принимает имя файла, а возвращает канал, в
котором лежат числа прочитанные из файла. 
Файл читается построчно, для экономии памяти, а
символы, не являющиеся числом, пропускаются

Вторая принимает канал от первой, и фильтрует числа
от нечетных чисел

Третья функция суммирует числа из канала от второй функции 
*/

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

// Функция, которая реализует пошаговую многопоточную обработку данных
func SumValuesPipeline(filename string) int {
    return Sum(Filter(NumbersGen(filename)))
}

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

// Принимаем канал с числами, возвращаем канал с четными числами
func Filter(in <-chan int) <-chan int {
    res := make(chan int)
    go func() {
        defer close(res)
        for val := range in {
            if val % 2 == 0 {
                res <- val
            }
        }
    }()
    return res
}

// Функция для подсчета суммы четных чисел. Принимаем канал, возвращаем число
func Sum(in <-chan int) int {
    res := 0
    for val := range in {
        res += val
    }
    return res
}