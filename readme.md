# Имитатор банкомата
Имитатор банкомата с наипростейшей логикой, применением
горутин и каналов и архитектурой REST API.

## Навигация
 - [Как запустить](#как-запустить)
 - [REST API эндпоинты и примеры запросов](#rest-api-эндпоинты-и-примеры-запросов)
 - [Примечания](#примечания)
 - [Контакты](#контакты)

## Как запустить
 - Клонировать репозиторий - `git clone https://github.com/dusk-chancellor/atm_imitator.git`
 - Установить нужные значения в `.env`
 - Установить все нужные внешние пакеты - `go mod tidy`
 - Запустить - `go run main.go`

## REST API эндпоинты и примеры запросов
 - `POST /accounts` - Создать новый аккаунт
   - Пример запроса: ...
 - `POST /accounts/{id}/deposit` - Депозит в аккаунт
   - Пример запроса: 
     - параметры: `id`=1
     - тело(body): 
     ```json
        {
            "amount": 100.0
        }
     ```
 - `POST /accounts/{id}/withdraw` - Забрать с аккаунта
   - Пример запроса:
     - параметры: `id`=1
     - тело(body):
     ```json
        {
            "amount": 100.0
        }
 - `GET  /accounts/{id}/balance` - Баланс аккаунта
   - Пример запроса:
     - параметры: `id`=1

## Примечания
 - в `~/docs` содержатся swagger документации

## Контакты
 - Телеграм: [@duskchancellor](https://t.me/duskchancellor)
 - Email: duskchancellor@gmail.com