# Переменные
DB_URL=postgresql://postgres:2005@127.0.0.1:5432/postgres?sslmode=disable

# Команда для запуска миграций "вверх"
migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

# Команда для отката миграций "вниз"
migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

# Создание новой миграции (использование: make migrate-create name=название)
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-up migrate-down migrate-create