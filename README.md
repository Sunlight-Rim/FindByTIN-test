# Тестовое задание

Решение [тестового задания](https://gist.github.com/foxcpp/0fdf9bad8504fa803e73406bbeffadb3)

## Описание

У сервиса доступны две конечные точки API: HTTP REST и gRPC. При обращении к REST API запрос адресуется на grpc-gateway, который вызывает метод микросервиса на gRPC и возвращает результат. \
Имеется проверка на корректность ИНН в запросе и обработка прочих внутренних ошибок. \
Также доступна документация swagger.

### REST

При переходе по адресу `localhost:8080/tin/<number>`, где `<number>` — это номер ИНН компании, сервис возвращает ИНН, КПП, название и ФИО руководителя компании в формате JSON.

Пример запроса:
```
curl localhost:8080/tin/7705938896 -s
```
Пример ответа:
```
{
  "tin": "7705938896",
  "tgrc": "770501001",
  "title": "ООО 'Авито'",
  "FCs": "Хамеденова Гульмай Турганбаевна"
}
```

### gRPC

gRPC-сервер принимает на вход строчный параметр ИНН и возвращает ИНН, КПП, название и ФИО руководителя компании в виде структуры протобуфера.

Формат запроса:
```
message GetTinRequest {
    string tin = 1;
}
```
Формат ответа:
```
message GetTinResponse {
    string tin = 1;     // ИНН
    string tgrc = 2;    // КПП
    string title = 3;   // Название
    string FCs = 4;     // ФИО
}
```

## Swagger UI

Swagger доступен по адресу `localhost:8080/swaggerui`. JSON-описание для него было сгенерировано protoc-компилятором из annotations.proto.

## Структура проекта

В проекте поддержана структура папок [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

- `cmd`: содержит только точку входа
- `internal/<service>/proto`: gRPC-определения сервисов в формате протофайлов
- `internal/<service>/pb`: сгенерированный протобуферный код
- `internal/<service>`: реализации микросервисных методов

## Запуск

Склонировать репозиторий
```
git clone https://github.com/Sunlight-Rim/FindByTIN-test.git
```

### Вручную

Запустить точку входа
```
go run ./cmd/main.go
```

### Docker

Собрать и запустить контейнер
```
docker build -t findbytin .
docker run -p 8080:8080 findbytin
```
