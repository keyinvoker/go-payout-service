# Description

- Architecture: Domain-Driven Design (DDD)
- Migration: golang-migrate

# Migration

## Create a new migration version
```bash
migrate create -ext sql -dir migrations -seq {description}
```

## Up a version
```bash
migrate -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" -path migrations up
```

## Down a version

```bash
migrate -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" -path migrations down 1
```

# How to run

## Build the binary
```bash
go build -o go-payout-service ./cmd/api
```

## Run the binary
```bash
./go-payout-service
```


# Project Structure
```
svc-payout/
├── internal/
│   ├── domain/
│   │   ├── models/
│   │   │   └── payout.go
│   │   ├── services/
│   │   │   ├── analytics/
│   │   │   │   └── clevertap.go
│   │   │   ├── payout/
│   │   │   │   └── execution.go
│   │   │   ├── tax/
│   │   │   │   └── tax.go
│   │   │   ├── transaction/
│   │   │   │   └── transaction.go
│   │   │   ├── user/
│   │   │   │   └── user.go
│   │   │   └── user_calculator/
│   │   │       └── user_calculator.go
│   │   ├── repositories/
│   │   │   ├── api/
│   │   │   │   ├── account.go
│   │   │   │   ├── auth.go
│   │   │   │   ├── campaign.go
│   │   │   │   ├── clevertap.go
│   │   │   │   ├── core.go
│   │   │   │   ├── notification.go
│   │   │   │   ├── tax.go
│   │   │   │   └── user.go
│   │   │   └── database/
│   │   │   │   └── payout.go
│   │   └── errors/
│   │       ├── account.go
│   │       ├── analytics.go
│   │       ├── caches/
│   │       │   ├── active_lender.go
│   │       │   └── loan_investment.go
│   │       ├── common.go
│   │       ├── investment.go
│   │       ├── payout.go
│   │       ├── pubsub.go
│   │       ├── repayment.go
│   │       ├── tax.go
│   │       └── user.go
│   ├── application/
│   │   ├── usecases/
│   │   │   ├── notification/
│   │   │   │   ├── email.go
│   │   │   │   └── push_notification.go
│   │   │   └── payout/
│   │   │   │   ├── calculation.go
│   │   │   │   ├── inquiry.go
│   │   │   │   ├── portfolio.go
│   │   │   │   ├── statistics.go
│   │   │   │   └── workflow.go
│   │   ├── services/
│   │   │   ├── analytics/
│   │   │   │   └── clevertap.go
│   │   │   ├── investment/
│   │   │   │   └── investment.go
│   │   │   ├── payout/
│   │   │   │   ├── analytics.go
│   │   │   │   └── notification.go
│   │   │   ├── tax/
│   │   │   │   └── tax.go
│   │   │   ├── transaction/
│   │   │   │   └── transaction.go
│   │   │   ├── user/
│   │   │   │   └── user.go
│   │   │   └── user_calculator/
│   │   │       └── user_calculator.go
│   │   ├── dtos/
│   │   │   ├── account.go
│   │   │   ├── calculation.go
│   │   │   ├── campaign.go
│   │   │   ├── investment.go
│   │   │   ├── payout.go
│   │   │   ├── pubsub.go
│   │   │   ├── repayment.go
│   │   │   ├── response.go
│   │   │   ├── tax.go
│   │   │   ├── transaction.go
│   │   │   ├── types.go
│   │   │   └── user.go
│   │   └── tasks/
│   │       ├── analytics.go
│   │       ├── notifications.go
│   │       └── payout/
│   │           ├── calculation.go
│   │           └── execution.go
│   ├── infrastructure/
│   │   ├── api/
│   │   │   ├── resources/
│   │   │   │   └── v1/
│   │   │   │       └── payout.go
│   │   │   └── routes.go
│   │   ├── persistence/
│   │   │   └── database/
│   │   │       └── connection.go
│   │   ├── messaging/
│   │   │   └── google_pubsub/
│   │   │       ├── publisher.go
│   │   │       └── subscriber.go
│   │   ├── auth/
│   │   │   └── token_validator.go
│   │   ├── caches/
│   │   │   ├── active_lender.go
│   │   │   └── loan_investment.go
│   │   ├── extensions/
│   │   │   ├── celery.go
│   │   │   └── redis.go
│   │   └── utils/
│   │       ├── auth.go
│   │       ├── datetime.go
│   │       ├── formatter.go
│   │       ├── payout.go
│   │       ├── response.go
│   │       └── transaction.go
│   ├── config/
│   │   └── config.go
│   ├── log/
│   │   └── log.go
│   └── constants/
│       ├── analytics.go
│       ├── campaign.go
│       ├── event.go
│       ├── loan.go
│       ├── notification.go
│       ├── payout.go
│       ├── pubsub.go
│       ├── repayment.go
│       ├── task.go
│       ├── tax.go
│       ├── transaction.go
│       └── user.go
├── cmd/
│   └── akpayout/
│       └── main.go
├── migrations/
│   ├── env.py
│   ├── script.py.mako
│   └── versions/
│       └── 001_add_payout.py
├── tests/
│   ├── cli/
│   │   └── message_broker.py
│   ├── message_brokers/
│   │   └── google_pubsub/
│   │       └── subscriber.py
│   ├── services/
│   │   └── payout_calculation.py
│   ├── usecases/
│   │   └── payout.go
│   └── utils/
│       └── async_events.py
├── go.mod
├── go.sum
├── Makefile
├── alembic.ini
├── CHANGELOG.md
├── Dockerfile
├── README.md
├── poetry.lock
└── pyproject.toml
```
