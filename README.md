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


### Documentation
* [API Reference](http://godoc.org/github.com/mirecl/goalmanac)

### Installation

    go get github.com/mirecl/goalmanac