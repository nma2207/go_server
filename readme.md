Тестовое задание: реализовать простой микросервис на Go
===

Использовал базу данных sqlite, и библитеку github.com/gorilla/mux" для более удобной работы с HTTP

Что было сделано:
---

Для работы с БД написан интерфейс IDatabase, который реализовал сначал тестовый класс, работающий с массивом, а потом реализовал класс, использующий SQLite базу данных

В функции main открывается база данных и разруливаются запросы

Как проверял?
---

Микросервис запускал командой

`$ go run hello.go`

Для GET/POST/PUT/DELETE запросов:

Получение продукта по id:
`$ curl -i http://localhost:8000/product/4`

Получение продуктов, отсортированных по цене:
`$ curl -i http://localhost:8000/product/sort`

Добавление:
`curl -i http://localhost:8000/product -d '{"name": "some_product", "cost": 1234}'`

Удаление:

`curl -i http://localhost:8000/product/1 -XDELETE`

Изменение:
`curl -i http://localhost:8000/product/1 -d '{"name": "some_product", "cost": 1234}' -XPUT`


