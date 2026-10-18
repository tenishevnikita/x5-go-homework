# Домашние задания по Go (X5 Tech)

Решения домашних заданий и практики курса «Разработка на Go».

## Структура

```text
lesson-02/
├── homework/           # домашнее задание: заготовки функций + автотесты
│   ├── 01_order_total/
│   ├── 02_order_status/
│   ├── 03_unicode_name/
│   ├── 04_safe_sum/
│   └── 05_order_retry/ # разобрана на паре
└── practice/           # простые задачи с пары, вывод через fmt.Println
    ├── README.md       # условия задач
    ├── 01_fizzbuzz/
    ├── 02_order_stats/
    └── 03_order_category/
```

Каждая папка задачи — отдельный Go-модуль со своим `go.mod`.

## Правила

- В домашних заданиях не менять сигнатуры функций и файлы `*_test.go`.
- Оценка — процент пройденных автотестов.

## Как запускать

Домашнее задание, одна задача:

```bash
cd lesson-02/homework/01_order_total
go test -v ./...
```

Все задачи урока разом:

```bash
for d in lesson-02/homework/*/; do (cd "$d" && go test ./...); done
```

Практика:

```bash
cd lesson-02/practice/01_fizzbuzz
go run .
```

Перед коммитом:

```bash
gofmt -l .
```
