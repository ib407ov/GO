START
migrate -path ./shema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
migrate -path ./shema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
migrate -path ./shema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up