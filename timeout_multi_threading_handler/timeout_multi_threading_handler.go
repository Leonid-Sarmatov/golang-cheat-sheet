/*
Код представляет простейшее API, в котором к серверу 1 приходит запрос клиента,
сервер 1 отсправляет запрос на сервер 2 и ждет от него ответ определенное время.
Если сервер 2 в течении таймаута даст отклик, то сервер 1 отправит его клиенту, 
иначе отправит код ошибки.
*/

package timeout_multi_threading_handler

import (
    "fmt"
    "time"
    "io"
    "net/http"
)

var (
    limit time.Duration
)

// Спавн первого сервера
func ServerRun() {
    mux := http.NewServeMux()
    mux.HandleFunc("/readSource", Handler)
    http.ListenAndServe(":8080", mux)
}

// Функция для опроса второго сервера и записи отклика в канал
func myRequest(ch chan *http.Response) {
	// Создаем запрос на нужный адрес
    req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/provideData", nil)
    if err != nil {
        fmt.Println(err.Error())
    }

	// Создаем клиент для выполнения запроса
    client := &http.Client{}
	// Выполняем запрос
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err.Error())
    }

	// Записываем отклик в канал
    ch<-resp
}

// Обработчик запросов на сервер по нужному маршруту
func Handler(w http.ResponseWriter, r *http.Request) {
	// Создаем канал, которуй предназначен для передачи отклика от второго сервера
    ch := make(chan *http.Response)
	// Спавним первый сервер
    go myRequest(ch)
    
	// Ждем результата
    for {
        select {
			// Если в канал попало значение
            case resp:=<-ch:
				// Считываем тело запроса
                body, err := io.ReadAll(resp.Body)
                if err != nil {
                    return
                }
				// Записываем тело запроса
            	w.Write(body)
				// Записываем код ошибки
            	w.WriteHeader(resp.StatusCode)
            	return
			// Если лимит ожидания превышен
            case <-time.After(limit):
				// Отправляем ошибку о невозможности выполнить запрос 
            	w.WriteHeader(http.StatusServiceUnavailable)
            	return
        }
    }
}

// Точка входа
func StartServer(maxTimeout time.Duration) {
	// Запоминаем максимальное время ожидания
    limit = maxTimeout
	// Запускаем первый сервер
    go ServerRun() 
}

