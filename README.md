# FurryShop (demo)

This project is a demo API for a small shop. It uses:

- Gin
- GORM (Postgres)
- Redis (optional cache)
- DDD-like layered structure (repository/service/handlers)
- Swagger docs via swaggo

Run:

```pwsh
$Env:DATABASE_DSN = "host=localhost user=postgres password=postgres dbname=furryshop port=5432 sslmode=disable"
$Env:REDIS_ADDR = "localhost:6379"
go run .
```

Tests:

```pwsh
go test ./... -v
```

GitHub Copilot - ака пидарас

Кратко — где точка входа, структура и как данные/запросы текут в проекте:

1) Точка входа
- main: main.go. Именно там:
    - создаётся GORM DB: `db.NewGorm`,
    - выполняется AutoMigrate всех моделей (создаются 10 таблиц: User, Category, Product, Fetish, ProductFetish, UserFetish, Like, Recommendation, Notification, Review) — смотрите список миграции в main.go,
    - создаётся Redis-клиент: `cache.NewRedisClient`,
    - создаются GORM-репозитории: `repo.NewGormProductRepository`, `repo.NewGormFetishRepository`, `repo.NewGormLikeRepository`, `repo.NewGormNotificationRepository`, `repo.NewGormRecommendationRepository` (реализации в repository),
    - создаются сервисы: например `service.NewProductService`,
    - вызывается `router.SetupRouter(...)` и запускается Gin-сервер.

2) Строение (высокоуровневые папки)
- main.go — запуск приложения и wiring (DI).
- internal/db — подключение к БД и инициализация GORM (см. internal/db/gorm.go).
- internal/cache — Redis client/обёртки (см. internal/cache).
- internal/repository — GORM-модели и реализации репозиториев (модели: `UserModel`, `CategoryModel`, `ProductModel`, `FetishModel`, `ProductFetishModel`, `UserFetishModel`, `LikeModel`, `RecommendationModel`, `NotificationModel`, `ReviewModel`) — файлы в этом пакете содержат CRUD-реализации.
    - миграция моделей выполняется из main (см. AutoMigrate call в main.go).
- internal/service — бизнес-логика, оркестрация репозиториев (так называемая layer service; см. internal/service).
- handlers — HTTP handlers (CRUD endpoints). Каждый handler должен вызывать сервис или репозиторий.
    - примеры: category_handler.go, product handlers и т.д.
- router — регистрация маршрутов Gin: router.go. Функция SetupRouter принимает сервис и несколько репозиториев для wiring.
- docs — сгенерённый swagger: docs.go и swagger.json.
- tests — тесты (unit / integration) — вы перемещали тесты в tests и добавляли интеграционный seed-тест (in-memory sqlite) для заполнения данных.

3) Как запрос проходит через систему (алгоритм)
- HTTP request -> Gin router ([router/router.go]) -> соответствующий handler в handlers/*.
- Handler:
    - парсит параметры/тело запроса,
    - вызывает метод сервиса (internal/service) или напрямую репозитория (internal/repository) для CRUD,
    - сервис использует репозитории (GORM) и при необходимости Redis cache (internal/cache),
    - репозиторий выполняет операции через GORM -> DB (Postgres в production).
- Ответ формируется в handler и возвращается клиенту.

4) Модели и связи (основное)
- Product связан с Category (category_id).
- Fetish — отдельная таблица; many-to-many Product <-> Fetish через `ProductFetishModel`.
- User может иметь фетиши (`UserFetishModel`), ставить лайки (`LikeModel`), получать уведомления (`NotificationModel`), получать рекомендации (`RecommendationModel`) и оставлять отзывы (`ReviewModel`).
- Все модели мигрируются в AutoMigrate (см. main.go).

5) Swagger / документация
- Комментарии godoc + swag добавлены в handlers и моделях; сгенерированная документация находится в docs.go и swagger.json. Открывается через endpoint `/swagger/*any` (роутер).

6) Тесты и наполнение БД
- В тестах добавлен интеграционный тест, который использует in-memory SQLite (pure-Go driver в тестовой среде) для AutoMigrate и seed данных (см. `tests/integration/*`).
- Unit-тесты расположены в tests (вы переместили `_test.go` туда); они используют mock/in-memory репозитории или тестовую sqlite для проверки репозиториев/handlers.

7) Замечания / потенциальные проблемы, которые надо решить
- В коде есть места с несоответствием DI: в router.go `SetupRouter` принимает ряд репозиториев, но некоторые handlers (например ReviewHandler) имеют конструктор без репозитория — это можно унифицировать: либо внедрять `repo.ReviewRepository`, либо дать handler умолчание/мок.
- Тестирование с SQLite: проект использует pure-Go драйверы в тестах (`github.com/glebarez/sqlite` / modernc.org/sqlite) для Windows без CGO — убедитесь, что go.mod содержит корректную версию.
- DTO vs модели: сейчас Swagger ссылается на `repository.*Model` — лучше выделить отдельные request/response DTO в `service` или handlers для чистоты API и корректной документации.

8) Рекомендуемые следующие шаги (если нужно)
- Унифицировать DI: внедрить репозитории во все handlers через конструкторы и передавать их из main -> router.
- Добавить полноценные unit-tests для handlers и сервисов, и интеграционные тесты через docker-compose (Postgres + Redis).
- Разделить DTO и модели и обновить swagger-аннотации, чтобы документация показывала только публичные схемы API.
- Реализовать/проверить ReviewRepository (сейчас есть сообщение об undefined constructor) и поправить wiring в router/main.

Если нужно, могу:
- показать конкретную диаграмму вызовов с ссылками на функции (например: main -> db.NewGorm -> repo.NewGormProductRepository -> service.NewProductService -> handlers.NewProductHandler -> router routes),
- или прямо внести правки: унифицировать ReviewHandler DI / добавить недостающий repo.NewReviewRepository и тесты.

Что хотите дальше — исправить DI для ReviewHandler или сгенерировать подробную карту вызовов (функция -> файл -> где создаётся)?
