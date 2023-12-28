/*
Дан сервер с адресом "localhost:8082", по 
запросу http://localhost:8082/mark?name=<имя студента>"
сервер возвращает оценку студента и значение ошибки

Код отправляет запрос на сервер и вычисляет среднюю успеваемость студентов,
при возникновении ошибки хоть на одном запросе, функция возвращает ошибку
*/

package main

import (
    "fmt"
    "net/http"
    "io"
    "sync"
    "strconv"
)

// Структура содержащая оценку студента и ошибку, если она возникла
type A struct {
    val int
    err error
}

func Average(names []string) (int, error) {
	// Создаем синхронизаторы
    wg := &sync.WaitGroup{}
    mx := &sync.Mutex{}
    
	// Массив с результатами 
    arr := make([]A, len(names))
    
	// Добавляем в ожидание столько горутин, сколько оценок студентов хотим получить
    wg.Add(len(arr))
    for i := 0; i < len(names); i+=1 {
        go func(k int) {
            defer wg.Done()
            mx.Lock()
            req, err := http.NewRequest(http.MethodGet, "http://localhost:8082/mark?name="+names[k], nil)
            mx.Unlock()
            if err != nil {
                arr[k].err = err
                return 
            }
            
            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                arr[k].err = err
                return 
            }
            
            body, err := io.ReadAll(resp.Body)
            if err != nil {
                arr[k].err = err
                return 
            }
            
            res, err := strconv.Atoi(string(body))
            if err != nil {
                arr[k].err = err
                return 
            }
            
            mx.Lock()
            arr[k].val = res
            mx.Unlock()
        }(i)
    }
	// Ждем окончания работы всех горутин
    wg.Wait()
    
	// Среднее арифметическое всех оценок
    res := 0
	// Флаг оцутствия ошибок
    flag := true 
	// Обходим список с результатами
    for _, v := range arr {
        res += v.val
        if v.err != nil {
        	flag = flag && false
        }
    }
    
	// Если не удалось получить хотя бы одну оценку, то возвращаем ошибку
    if !flag {
        return 0, fmt.Errorf("error")
    }
	// Возвращаем результат 
    return int(res/len(arr)), nil
}