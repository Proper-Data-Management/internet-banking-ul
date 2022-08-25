## Архитектура проекта

```
.
├── arch.md
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── helpers               // сторонние вспомогательные мини-утилиты для обработки Fiber ошибок
│  └── apiErrors
│     ├── common.go
│     ├── error.go
│     ├── init.go
│     ├── response.go
│     ├── user.go
│     └── user_profile.go
├── internal              // вся бизнес часть (т.е. логика и пр.)
│  ├── config
│  │  └── config.go       // основной конфигурационный файл
│  ├── consts             // всопомогательные константы
│  │  └── consts.go
│  ├── handlers           // контроллеры и регистраторы
│  │  ├── customer
│  │  │  ├── controller.go
│  │  │  └── customer.go
│  │  └── response.go     // вспомогательная модель для работы с ответами
│  ├── middles            // вспомогательные мини-утилиты для работы с контекстом Fiber (в основном локальный форк из Fiber'a)
│  │  ├── attributes.go
│  │  ├── config_recovery.go
│  │  ├── contextholder.go
│  │  ├── errors.go
│  │  └── fiber-zap.go
│  ├── modules            // любая бизнес часть подключается в виде модуля, где у каждого свои модели/дто/репы/сервис итд
│  │  ├── customer
│  │  │  ├── dto          // dto модели и преобразование
│  │  │  │  └── customer.go 
│  │  │  ├── entities     // основные модели
│  │  │  │  └── customer.go
│  │  │  ├── repositories // запросы в базу данных
│  │  │  │  ├── customer.go
│  │  │  │  └── repositories.go
│  │  │  └── services     // бизнес логика
│  │  │     └── customer.go
│  │  └── entities        // общие модели для работы с фильтрами итд итп
│  │     └── base_pagination_filter.go
│  ├── server
│  │  └── router.go       // инициализация сервера
│  └── utils              // вспомогательные утилиты
│     ├── context_holder.go
│     ├── context_request_info.go
│     ├── gorm.go
│     └── user_agent_parser.go
├── LICENSE
├── main.go
├── Makefile
├── modules               // сторонние библиотеки
│  ├── daylight           // сторонняя библиотека для с PID процессами
│  ├── logger             // сторонняя библиотека для работы с логами
│  └── squirrel           // сторонняя библиотека для генерации SQL запросов
├── README.md
└── tools                 // сторонние вспомогательные мини-утилиты
```