# duty_checklist

// DEV
migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

// TEST
migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up