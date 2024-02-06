### Проверка tzdata и настройки часового пояса на go в связи с переходом Казахстана на UTC+5.

Локальный запуск, если установлен go:

`go run main.go`

Пример вывода:

```
Текущее время в дефолтной таймзоне: 2024-02-06T12:42:47+06:00
Диапазон действия текущей таймзоны: 2004-10-31T02:00:00+06:00 до 0001-01-01T00:00:00Z
Диапазон действия таймзоны в марте 2024: 2004-10-31T02:00:00+06:00 до 0001-01-01T00:00:00Z
👎 Диапазоны действия текущей таймзоны и таймзоны в марте совпадают. Вероятно, пакет tzdata не обновлён.
Текущее время в таймзоне Asia/Almaty: 2024-02-06T12:42:47+06:00
```

### Пример запуска в последней версии alpine, где таймзона Asia/Almaty уже обновлена

```
docker run --rm alpine:latest /bin/sh -c 'apk add --no-cache tzdata && export TZ=Asia/Almaty && wget TODO && chmod +x tz-linux && ./tz-linux'

Текущее время в дефолтной таймзоне: 2024-02-06T12:43:37+06:00
Диапазон действия текущей таймзоны: 2004-10-31T02:00:00+06:00 до 2024-02-29T23:00:00+05:00
Диапазон действия таймзоны в марте 2024: 2024-02-29T23:00:00+05:00 до 0001-01-01T00:00:00Z
👍 Диапазоны действия текущей таймзоны и таймзоны в марте не совпадают. Вероятно, пакет tzdata уже обновлён.
Текущее время в таймзоне Asia/Almaty: 2024-02-06T12:43:37+06:00
```

Если вы видите 👍, то скорее всего не нужно беспокоиться - время будет автоматически переведено на UTC+5 1 марта.
Если вы видите 👎, то скорее всего вам нужно обновить tzdata и\или настроить таймзону (`export TZ=Asia/Almaty`).
