# Changelog

Все заметные изменения проекта описаны в этом файле.

Формат основан на [Keep a Changelog](https://keepachangelog.com/ru/1.1.0/),
проект придерживается [Semantic Versioning](https://semver.org/lang/ru/).

## [Unreleased]

### Added
- `README.md` с полным описанием проекта: назначение, требования к окружению,
  установка, запуск, переменные окружения, запуск тестов и структура каталогов.
- `docs/API.md` с детальным описанием HTTP-эндпоинтов (`/health`, `/version`),
  функций пакета `main` (`newMux`, `main`) и теста `TestVersionRoute`.
- Inline-документация (godoc) к пакету `main` и функциям `newMux`, `main`.

### Текущее состояние сервиса
- HTTP-сервис на Go 1.22 без внешних зависимостей.
- Эндпоинт `GET /health` → `application/json`, тело `{"ok":true}\n`.
- Эндпоинт `GET /version` → `text/plain`, тело `v1`.
- Адрес прослушивания настраивается через переменную окружения `LISTEN_ADDR`
  (по умолчанию `:8081`).
- `Makefile` с целями `run` и `health`.
- Тест `TestVersionRoute`, проверяющий статус, `Content-Type` и тело
  эндпоинта `/version` через `httptest.NewServer`.

## История коммитов

До формального релиза проект развивался следующими коммитами в ветке `main`:

- `afeda41` — `init: hello-world service for aitt smoke tests`
  Первоначальная версия: пакет `main`, эндпоинт `/health`, `Makefile`,
  модуль `github.com/volkovchain/aitt-smoke`, README с описанием цели сервиса.
- `0fad357` — `ai: task ai/971601b1 run aitt-smoke`
  Технический коммит в рамках задачи оркестратора.
- `1458a99` — `Merge pull request #2 from volkovchain/ai/971601b1`
  («Add /version endpoint») — добавлен эндпоинт `GET /version`,
  возвращающий `v1` с `Content-Type: text/plain`, и тест `TestVersionRoute`.

[Unreleased]: https://github.com/volkovchain/aitt-smoke/commits/main
