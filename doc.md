migrate -database "postgres://postgres:123456@localhost:5432/belajar_golang_restful_api?sslmode=disable" -path db/migrations up / down / force {version before dirty} / version
migrate create -ext sql -dir db/migrations -tz ASIA/JAKARTA create_table_tests
