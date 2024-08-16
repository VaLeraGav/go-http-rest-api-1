## Project-1

учебный проект на языке goland.

## Curl

**Создание пользователя**

```sh
curl -i -c cookies-create.txt -X POST  http://localhost:8080/users/create -H "Content-Type: application/json" -d '{"email":"name1@example.com", "password":"12345678"}'
```

**Авторизоваться и получить сессию**

```sh
curl -i -c cookies-auth.txt -X POST  http://localhost:8080/auth -H "Content-Type: application/json" -d '{"email":"name1@example.com", "password":"12345678"}'
```

**Получить данные о себе**

```sh
curl -i -b cookies-auth.txt -X POST  http://localhost:8080/private/whoami -H "Content-Type: application/json" -d '{"email":"name1@example.com", "password":"12345678"}'
```
