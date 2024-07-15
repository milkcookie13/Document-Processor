﻿# Document Processor

## Описание

`Document Processor` - это Go-проект для обработки и сохранения документов. Он включает API для обработки и записи документов и методы для управления базой данных. 
В качестве примера базы данных выбрана PostgreSQL.
Данный код будет работать в сервисе, читающим входные сообщения из очереди сообщений (Kafka или подобное), и
записывающем результат также в очередь. 

Документы поступают в формате:

```
message TDocument {
  string Url = 1; // URL документа
  uint64 PubDate = 2; // время публикации документа
  uint64 FetchTime = 3; // время получения обновления, уникально с Url
  string Text = 4; // текст документа
  uint64 FirstFetchTime = 5; // изначально отсутствует
}
```

# Правила обработки документов

Документы могут поступать в произвольном порядке и дублироваться. Поля внутри этих документов должны корректироваться по следующим правилам:

- **Text и FetchTime**: 
  - Берутся из документа с **наибольшим FetchTime**.

- **PubDate**:
  - Берется из документа с **наименьшим FetchTime**.

- **FirstFetchTime**:
  - Равно **минимальному FetchTime** среди всех документов.

## Структура проекта

```
dcproccer/
├── go.mod
├── go.sum
├── app/
│ ├── api.go
│ ├── documentprocessor.go
│ └── documentprocessor_internal_test.go
│ └── app/api_internal_test.go
├── store/
│ ├── config.go
│ ├── documentrepository.go
│ ├── store.go
│ └── testing.go
```

| Путь | Описание |
| --- | --- |
| `go.mod` и `go.sum` | Файлы для управления зависимостями Go |
| `app/api.go` | Содержит функцию обработки документа |
| `app/documentprocessor.go` | Содержит функции для управления очередями ввода и вывода |
| `app/documentprocessor_internal_test.go` | Тесты функций чтения/записи документов |
| `app/api_internal_test.go` | Тесты функции обработки документов |
| `store/config.go` | Настройка конфигурации базы данных |
| `store/documentrepository.go` | Содержит методы для работы с базой данных |
| `store/store.go` | Содержит функции для подключения/закрытия соедининения к базе данных на примере PostgreSQL  |
| `store/testing.go` | Тест подключения к базе данных |



## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/milkcookie13/Document-Processor.git
    ```

2. Перейдите в директорию проекта:
    ```sh
    cd Document-Processor
    ```

3. Установите зависимости:
    ```sh
    go mod tidy
    ```
4. Запустить тесты:
   ```sh
   go test ./...
   ```



