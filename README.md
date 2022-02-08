# GolangWeather

Бот погоды на Go.

### Настройка

Скопируйте файл конфигурации и добавьте в него токен для OpenWeather API + токен своего Telegram-бота

```bash
cp .env.dist .env
```

### Запуск

Для запуска бота воспользуйтесь алиасом:

```bash
make up
```

Или же

```bash
go run cmd/bot.go
```
