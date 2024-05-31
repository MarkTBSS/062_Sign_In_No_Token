```
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get github.com/jmoiron/sqlx
go get github.com/jackc/pgx/v5/stdlib

console
25/05/2024 [127.0.0.1] 200 - GET /v1
25/05/2024 [127.0.0.1] 404 - GET /v1/ee
201 - POST /v1/users/signup

postman
{
    "user": {
        "id": "U000003",
        "email": "test@test.com",
        "username": "test",
        "role_id": 1
    },
    "token": null
}
```