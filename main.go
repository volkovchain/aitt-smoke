// Package main реализует HTTP-сервис aitt-smoke — минимальную «заглушку»,
// используемую оркестратором aitt в качестве smoke-теста окружения.
//
// Сервис регистрирует два эндпоинта:
//
//   - GET /health  — отдаёт JSON {"ok":true} с Content-Type application/json;
//   - GET /version — отдаёт строку "v1" с Content-Type text/plain.
//
// Адрес прослушивания берётся из переменной окружения LISTEN_ADDR;
// если она не задана, используется ":8081".
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// newMux собирает и возвращает *http.ServeMux с зарегистрированными
// эндпоинтами сервиса.
//
// Зарегистрированные маршруты:
//
//   - "/health"  — обработчик пишет заголовок Content-Type: application/json
//     и тело `{"ok":true}\n`. Статус ответа — 200 OK (выставляется неявно
//     первым вызовом записи в ResponseWriter).
//   - "/version" — обработчик пишет заголовок Content-Type: text/plain,
//     явно вызывает WriteHeader(http.StatusOK) и пишет тело "v1" без
//     завершающего перевода строки.
//
// Функция не имеет параметров, всегда возвращает не-nil *http.ServeMux и
// не возвращает ошибок. Используется как в main(), так и в тестах
// (см. main_test.go: TestVersionRoute), что позволяет проверять роутер
// без поднятия отдельного процесса.
//
// Пример использования в тесте:
//
//	mux := newMux()
//	srv := httptest.NewServer(mux)
//	defer srv.Close()
//	resp, _ := http.Get(srv.URL + "/version")
func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("content-type", "application/json")
		_, _ = fmt.Fprintln(w, `{"ok":true}`)
	})
	mux.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "v1")
	})
	return mux
}

// main — точка входа сервиса.
//
// Поведение:
//
//  1. Создаёт роутер через newMux().
//  2. Читает адрес прослушивания из переменной окружения LISTEN_ADDR;
//     при пустом значении использует ":8081".
//  3. Печатает в stdout строку вида "listening on <addr>".
//  4. Запускает http.ListenAndServe(addr, mux); при ошибке (включая штатное
//     http.ErrServerClosed, которого в текущей реализации не возникает,
//     так как graceful shutdown не реализован) пишет её в stderr и
//     завершает процесс кодом 1.
//
// Функция блокируется до завершения http.ListenAndServe и в норме не
// возвращает управление.
func main() {
	mux := newMux()
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8081"
	}
	fmt.Println("listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
