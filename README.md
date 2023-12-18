# Gofermart

Сервис Gofermart является "Накопительной системой лояльности", он включает в себя:

1. регистрацию, аутентификацию и авторизацию пользователей;
2. приём номеров заказов от зарегистрированных пользователей.
3. учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;
4. учёт и ведение накопительного счёта зарегистрированного пользователя;
5. проверка принятых номеров заказов через систему расчёта баллов лояльности;
6. начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

# Установка и запуск

Для того чтобы получить доступ к проекту вы можете склонировать на свое устройство 
командой - git clone https://github.com/fraaaaaaaaaanc/gofermart.git

Перед тем как начать работу с проектом следует установить следующие зависимости:

1. `github.com/go-chi/chi` v1.5.5
2. `github.com/go-playground/validator` v9.31.0+incompatible
3. `github.com/golang-jwt/jwt/v4` v4.5.0
4. `github.com/jackc/pgerrcode` v0.0.0-20220416144525-469b46aa5efa
5. `github.com/jackc/pgx/v5` v5.5.0
6. `github.com/pressly/goose` v2.7.0+incompatible
7. `github.com/shopspring/decimal` v1.3.1
8. `github.com/stretchr/testify` v1.8.4
9. `go.uber.org/zap` v1.26.0
10. `golang.org/x/crypto` v0.15.0

Для того чтобы запустить проект, перейдите из корневой директории в директорию cmd/accrual
и напишите команду go run main.go .

# Флаги запуска

1. `-a` Aдрес и порт запуска сервиса
2. `-b` Адрес подключения к базе данных
3. `-r` Адрес системы расчёта начислений
4. `-lf` Адрес файла для записи логов
5. `-ll` Запись логов будет происходить в консоль

# Ручки сервиса

1. `POST /api/user/register` — регистрация пользователя;
2. `POST /api/user/login` — аутентификация пользователя;
3. `POST /api/user/orders` — загрузка пользователем номера заказа для расчёта;
4. `GET /api/user/orders` — получение списка загруженных пользователем номеров заказов, статусов их обработки и 
информации о начислениях;
5. `GET /api/user/balance` — получение текущего баланса счёта баллов лояльности пользователя;GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
6. `POST /api/user/balance/withdraw` — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
7. `GET /api/user/withdrawals` — получение информации о выводе средств с накопительного счёта пользователем.

