# Notes API

Простой REST API на Go для управления заметками с использованием SQLite и документацией Swagger.

---

## 🚀 Быстрый старт

### 1. Клонируй репозиторий

```bash
git clone https://github.com/твой_логин/notes-api.git
cd notes-api
```

### 2. Установи зависимости

```bash
go mod download
```

### 3. Сгенерируй Swagger-документацию и запусти сервер

```bash
make run
```

Сервер стартует на [http://localhost:8080](http://localhost:8080)

---

## 📚 API

Документация с возможностью тестирования доступна по адресу:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🔧 Структура проекта

- `main.go` — точка входа, настройка роутера и Swagger  
- `handlers/` — HTTP обработчики (CRUD)  
- `models/` — модели данных  
- `storage/` — инициализация базы и работа с БД  
- `docs/` — сгенерированные Swagger-файлы  

---

## 🛠 Используемые технологии

- [Go](https://golang.org/)  
- [GORM](https://gorm.io/) — ORM для работы с SQLite  
- [gorilla/mux](https://github.com/gorilla/mux) — роутер  
- [swaggo](https://github.com/swaggo/swag) — генерация документации Swagger  
- [SQLite](https://www.sqlite.org/index.html) — база данных