# Микросервис - “Альманах” (Календарь)

[![image](https://img.shields.io/badge/godoc-reference-blue)](http://godoc.org/github.com/mirecl/goalmanac)

### Задание #1 - “Заготовка” для микросервиса “Альманах” (Календарь)

**Цель:** В результате выполнения ДЗ должен получиться базовый скелет микросервиса, который будет развиваться в дальнейших ДЗ. Структура кода должна соответствовать подходу Clean Architecture. В данном задании тренируются навыки: - декомпозиции предметной области; - построения элементарной архитектуры проекта.
Завести в репозитории отдельную директорию для проекта "Календарь"
Создать внутри структуру директорий, соответствующую Clean Architecture.

Cоздать модели (структуры) календаря.

Cоздать методы бизнес логики (методы у структур) для работы с этими структурами:
- добавление событий в хранилище
- удаление событий из хранилища
- изменение событий в хранилище
- листинг событий
- пр. на усмотрение студента
Создать объекты ошибок (error sentinels) соответсвующие бизнес ошибкам, например ErrDateBusy - данное время уже занято другим событием

Реализовать хранение событий в памяти (т.е. просто складывать объекты в слайсы)

Реализовать Unit тесты проверяющие работу бизнес логики (в частности ошибки)

На данном этапе не нужно:
- Делать HTTP, GRPC и пр. интерфейсы к микросервису
- Писать .proto-файлы (это будет позже)
- Использовать СУБД

### Задание #2 - Каркас микросервиса

**Цель:** Реализовать "каркас" микросервиса, считывающий конфиг из файла, создающий логгер/логгеры с указанными уровнями детализации.
Необходимо доработать код сервиса "Календарь" из предыдущего задания, добавив в него:

* Обработку аргументов командной строки
* Чтение файла конфигурации (параметр --config в командной строке)
* Создание логгеров и настройка уровня логирования
* Создание и запуск hello-world web-сервера

Параметры, передаваемые через аргументы командной строки:
* --config - путь к конфигу

Параметры, которые должны быть в конфиге:
* http - host и port на котором должен слушать web-сервер
* log_file - путь к файлу логов
* log_level - уровень логирования (error / warn / info / debug)

**Критерии оценки:** Web-сервер на данном этапе может быть не связан с бизнес логикой календаря и должен обрабатывать только URL /hello
Web-сервер должен запускаться на host:port указанном в конфиге и каждый обработанный запрос должен выводиться в log-файл.

**Описание**
Файл конфигурации состоит из 3 разделов:
 - http (настройки http-сервера: адрес сервера/порт сервера/время остановки сервера)
 - log_http (настройка logger'а для http-сервера: уровень логирования(только для stdout)/путь для сохранения файла логирования)
 - log_event (настройка logger'а для событий Календаря: уровень логирования(только для stdout)/путь для сохранения файла логирования)

 Пример файла конфигурации:
 ```yaml
http:
  host: 127.0.0.1
  port: 8080
  shutdown: 5
log_http:
  level: info 
  path: http.log
log_event:
  level: info
  path: event.log
 ```
 Ротация файла логирования необходимо делать с помощью **Logrotate**.

 Специально созданы 2 logger'a:
  - Для событий Календаря;
  - Для http-запросов.

Создан один endpoint - /hello и к нему test.

### Задание #3 - HTTP интерфейс

**Цель:** Реализовать HTTP интерфейс для сервиса Календаря. Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md Цель данного задания - отработать навыки работы со стандартной HTTP библиотекой, поэтому технологии JSONRPC, Swagger и т.п. НЕ используются.

В директории с проектом создать отдельный пакет для Web-сервера
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области
Реализовать middleware для логирования запросов

Методы API:
- POST /create_event (создание)
- POST /update_event (обновление)
- POST /delete_event (удаление)
- GET /events_for_day (события на сегодня)
- GET /events_for_week (события на этой недели)
- GET /events_for_month (события в этом месяце)

В результате каждого запроса должен возвращаться JSON документ содержащий
либо ***{"result": "..."}*** в случае успешного выполнения метода
либо ***{"error": "..."}*** в случае ошибки бизнес-логики

**Критерии оценки:** Все методы должны быть реализованы
Бизнес логика (пакет internal/domain в примере) НЕ должен зависеть от кода HTTP сервера
В случае ошибки бизнес-логики сервер должен возвращать HTTP 200
В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400
В случае остальных ошибок сервер должен возвращать HTTP 500
Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

**Описание**

Запуск сервиса
```bash
make http
```

Имеется графический интерфейс для создания/изменения/удаления событий:
- Общий вид (список событий)
![Общий вид](https://github.com/mirecl/goalmanac/blob/master/img/1.png)
- Создание события
![Создание события](https://github.com/mirecl/goalmanac/blob/master/img/2.png)
- Изменение/удаление события
![Изменение/удаление события](https://github.com/mirecl/goalmanac/blob/master/img/3.png)


Покрытие сервиса тестами:
```bash
make test
```
```bash
?       github.com/mirecl/goalmanac     [no test files]
?       github.com/mirecl/goalmanac/cmd [no test files]
ok      github.com/mirecl/goalmanac/internal/adapters   0.004s  coverage: 92.3% of statements
ok      github.com/mirecl/goalmanac/internal/adapters/db        (cached)        coverage: 48.5% of statements
ok      github.com/mirecl/goalmanac/internal/adapters/http      0.003s
ok      github.com/mirecl/goalmanac/internal/adapters/http/validate     0.003s
ok      github.com/mirecl/goalmanac/internal/adapters/logger    (cached)        coverage: 63.0% of statements
ok      github.com/mirecl/goalmanac/internal/domain     (cached)        coverage: 100.0% of statements
?       github.com/mirecl/goalmanac/internal/domain/entities    [no test files]
?       github.com/mirecl/goalmanac/internal/domain/errors      [no test files]
?       github.com/mirecl/goalmanac/internal/domain/interfaces  [no test files]
ok      github.com/mirecl/goalmanac/internal/domain/usecases    (cached)        coverage: 44.2% of statements
```

Для валидации входных параметров (POST-запросы) использовал пакет:
```bash
github.com/xeipuuv/gojsonschema
```
Примеры JSON-Schem в папке config:
```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Событие",
  "description": "Создание события в Календаре",
  "type": "object",
  "properties": {
    "id": {
      "description": "UID События",
      "type": "string",
      "maxLength": 36,
      "minLength": 36
    },
    "user": {
      "description": "Пользователь",
      "type": "string",
      "minLength": 1
    },
    "title": {
      "description": "Заголовок события",
      "type": "string",
      "minLength": 1
    },
    "body": {
      "description": "Описание события",
      "type": "string",
      "minLength": 1
    },
    "start": {
      "description": "Время старта события",
      "type": "string",
      "format": "date-time",
      "minLength": 1
    },
    "duration": {
      "description": "Продолжительность события",
      "type": "string",
      "minLength": 1,
      "enum": [
        "10m",
        "20m",
        "30m",
        "40m",
        "50m",
        "60m"
      ]
    }
  },
  "required": [
    "user",
    "title",
    "body",
    "start",
    "duration"
  ]
}
```

Для работы с дата/время использовал:
```bash
github.com/jinzhu/now
```  

### Задание #4 - Работа с базами данных

**Цель:** Обеспечить сохранение событий календаря в СУБД Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md Цель данного занятия: отработка навыков работы СУБД, SQL, пакетами database/sql и github.com/jmoiron/sqlx
Установить базу данных (например postgres) локально (или сразу в Docker, если знаете как)
Создать базу данных и пользователей для проекта календарь
Создать схему данных (таблицы, индексы) в виде отдельного SQL файла и сохранить его в репозиторий
В проекте календарь создать отдельный пакет, отвечающий за сохранение моделей в СУБД
Настройки подключения к СУБД вынести в конфиг проекта
Изменить код приложения так, что бы обеспечить сохранение событий в СУБД

**Критерии оценки:** Должны быть созданы все необходимые таблицы и индексы.
SQL миграция должна применять с первого раза и должна быть актуальной,
т.е. все изменения которые вы делали в своей базе должны быть отражены в миграции.

**Описание**
Для запуска БД необходимо выполнить
```bash
make db
```
Для запуска Сервиса:
```bash
make http
```
Логин и пароль от БД передаются в env и считываются в сервисе:
```golang
viper.AutomaticEnv()
// Зачитываем credential для БД
viper.BindEnv("db.POSTGRES_PASSWORD", "POSTGRES_PASSWORD")
viper.BindEnv("db.POSTGRES_USER", "POSTGRES_USER")
viper.BindEnv("db.POSTGRES_DB", "POSTGRES_DB")
```

### Задание #5 - Работа с очередями

**Цель:** Реализовать "напоминания" о событиях с помощью RabbitMQ. Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md Цель данного занятия: отработка навыков работы с RabbitMQ и очередями вообще.
Установить локально очередь сообщений RabbitMQ (можно сразу в Docker если знаете как)

Создать процесс (scheduler), который периодически сканирует основную базу данных, выбирая события о которых нужно напомнить.
При запуске процесс должен подключаться к RabbitMQ и создавать все необходимые структуры (топики) в ней.
Процесс должен выбирать сообытия для которых следует отправить уведомление, сериализовать их (например в JSON) и складывать в очередь.

Создать процесс (sender), который читает сообщения из очереди и шлет уведомления.
Непосредственно отправку делать не нужно - можно просто выводить сообщения в STDOUT.

**Критерии оценки:** Код должен работать и проходить проверки go vet и golint
Настройки подключения к очереди должны быть вынесены в конфиг проекта
У преподавателя должна быть возможность скомпилировать процессы scheduler и sender с помощью Makefile
После запуска RabbitMQ и PostgreSQL процессы scheduler и sender должны запускаться без дополнительных действий

**Описание**

В конфигурацю добавились доп. параметры (./config/config.yaml).
```yaml
mq:
  host: 127.0.0.1
  port: 5672
  polling: 30s # Частота опроса БД
  period: 15m  # глубина выборки при опросе
log_mq:
  level: info
  path: mq.log
```
Значение default:
```golang
viper.SetDefault("mq", map[string]interface{}{
	"host":                  "127.0.0.1",
	"port":                  "5672",
	"RABBITMQ_DEFAULT_USER": "rabbitmq",
	"RABBITMQ_DEFAULT_PASS": "rabbitmq",
	"period":                "10m",
	"polling":               "1m",
})
// Зачитываем credential для MQ
viper.BindEnv("mq.RABBITMQ_DEFAULT_PASS", "RABBITMQ_DEFAULT_PASS")
viper.BindEnv("mq.RABBITMQ_DEFAULT_USER", "RABBITMQ_DEFAULT_USER")
```
Сервис MQ поднимается по команде:
```bash
make service
```
Файл запуска: docker-compose.yml
```yaml
mq:
    image: rabbitmq:3.7.5-management
    container_name: mq
    ports:
        - 5672:5672
        - 15672:15672
    volumes:
        - ./data/rabbitmq:/var/lib/rabbitmq/mnesia/rabbit@app-rabbitmq:cached
    environment:
        RABBITMQ_ERLANG_COOKIE: 6085e2412b6fa88647466c6a81c0cea0
        RABBITMQ_DEFAULT_USER: ${RABBITMQ_DEFAULT_USER}
        RABBITMQ_DEFAULT_PASS: ${RABBITMQ_DEFAULT_PASS}
        RABBITMQ_DEFAULT_VHOST: /
```
Логин и пароль указываются через environment.

Запуск сервисов Sender и Sheduler осуществляется через команду (при запуске http-сервера).
```bash
make http
```
Файл запуска: сmd/http.go (отдельные goroutine)
```golang
// Запускаем Sender
go mq.ServeSender()
// Запускаем Sheduler
go mq.ServeSheduler()
```

Добавилось поле в таблицу Almanac - notify (статус отправки сообщения)
- null - новое сообщение
- "1" - сообщение отправлено пользователю

Добавил файл миграции:
```sql
ALTER TABLE almanac ADD notify char(1);
create index notify_idx on almanac (starttime,notify);
```

Логика сервиса нотификации:
1) Сервис Sender сканирует БД каждые mq.polling c глубиной выборки mq.period, где notify is null;
2) Все попавшиеся сообщения отправляются в Rabbit MQ;
3) Сервис Sheduler считывает сообщения и отправляет пользователю и ставит в БД flag отправки, иначе не подтверждает отправку сообщения.

### Задание #6 - GRPC сервис

**Цель:** Создать GRPC API для сервиса календаря Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md Цель данного занятия: отработка навыков работы с GRPC, построение современного API.

- Создать отдельную директорию для Protobuf спек.
- Создать Protobuf спеки с описанием всех методов API, их объектов запросов и ответов.
- Создать отдельный директорию для кода GRPC сервера
- Сгенерировать код GRPC сервера на основе Protobuf спек (скрипт генерации сохранить в репозиторий).

**Критерии оценки:** Все методы должны быть реализованы
Бизнес логика (пакет internal/domain в примере) НЕ должен зависеть от кода GRPC сервера
GRPC-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

**Описание**

В конфигурацю добавились доп. параметры (./config/config.yaml).
```yaml
grpc:
  host: 127.0.0.1
  port: 50051
log_grpc:
  level: info
  path: info.log
```
Значение default:
```golang
viper.SetDefault("log_grpc", map[string]string{
	"level": "info",
	"path":  "grpc.log",
})
```
Для работы необходимо запустить сервис БД:
```bash
make service
```
Запуск gRPC осуществляется через команду (при запуске http-сервера):
```bash
make http
```
```golang
// Запускаем GRPC API
go serverGRPC.Serve()
```
или запускается как отдельный сервис:
```bash
make grpc
```

Proto-файл (internal/adapters/grpc/api):
```yaml
syntax = "proto3";

option go_package = "api";

import "google/protobuf/timestamp.proto";
import "google/protobuf/empty.proto";

// Методы Сервиса
service Almanac {
    rpc Create(EventCreate) returns (ResponseOK) {};
    rpc Update(EventUpdate) returns (ResponseOK) {};
    rpc Delete(EventDelete) returns (ResponseOK) {};
    rpc GetAll(google.protobuf.Empty) returns (ResponseEvents) {};
    rpc GetDayEvent(EventUser) returns (ResponseEvents) {};
    rpc GetWeekEvent(EventUser) returns (ResponseEvents) {};
    rpc GetMonthEvent(EventUser) returns (ResponseEvents) {};
}

// Создание события в календаре
message EventCreate {
    string user = 1;
    string title = 2;
    string body = 3;
    google.protobuf.Timestamp starttime = 4;
    string duration = 5;
}

// Обновление события в календаре
message EventUpdate {
    string id = 1;
    string user = 2;
    string title = 3;
    string body = 4;
    google.protobuf.Timestamp starttime = 5;
    string duration = 6;
}

// Id события - для удаления
message EventDelete {
    string id = 1;
}

// Пользователь
message EventUser {
    string user = 1;
}

// Cобытие в календаре
message Event {
    string id = 1;
    string user = 2;
    string title = 3;
    string body = 4;
    google.protobuf.Timestamp starttime = 5;
    google.protobuf.Timestamp endtime = 6;
}

// Успешный ответ - список Cобытий
message ResponseEvents {
    repeated Event result = 1;
}

// Успешный ответ от сервиса
message ResponseOK {
    string result = 1;
}
```

### Documentation
* [API Reference](http://godoc.org/github.com/mirecl/goalmanac)

### Installation

    go get github.com/mirecl/goalmanac