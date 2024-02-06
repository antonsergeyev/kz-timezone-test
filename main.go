package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	isOk := true
	now := time.Now()
	fmt.Printf("Текущее время в дефолтной таймзоне: %s\n", now.Format(time.RFC3339))

	offsetName, offset := now.Zone()
	boundsStart, boundsEnd := now.ZoneBounds()
	fmt.Printf(
		"Текущий сдвиг: %s (%d секунд), действует от %s до %s\n",
		offsetName, offset,
		boundsStart.Format(time.RFC3339), boundsEnd.Format(time.RFC3339),
	)

	timeInMarch := time.Date(2024, 3, 1, 2, 0, 0, 0, time.Local)
	marchOffsetName, marchOffset := timeInMarch.Zone()
	marchBoundsStart, marchBoundsEnd := timeInMarch.ZoneBounds()
	fmt.Printf(
		"Сдвиг в марте 2024: %s (%d секунд), действует от %s до %s\n",
		marchOffsetName, marchOffset,
		marchBoundsStart.Format(time.RFC3339), marchBoundsEnd.Format(time.RFC3339),
	)

	if marchOffset != 60*60*5 {
		isOk = false
		fmt.Printf("👎 Сдвиг в марте не равен ожидаемому. Вероятно, пакет tzdata не обновлён, либо не установлен, либо не настроена текущая таймзона.\n")
	} else {
		fmt.Printf("👍 Сдвиг в марте равен ожидаемому. Вероятно, пакет tzdata уже обновлён.\n")
	}

	almaty, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		isOk = false
		fmt.Printf("👎 Не удалось загрузить таймзону Asia/Almaty: %s. Вероятно, в ОС или образе не установлен пакет tzdata.\n", err)
	} else {
		fmt.Printf("Текущее время в таймзоне Asia/Almaty: %s\n", now.In(almaty).Format(time.RFC3339))

		if !now.In(almaty).Equal(now) {
			isOk = false
			fmt.Printf("👎 Текущее время в дефолтной таймзоне не совпадает со временем Asia/Almaty. Установите таймзону Asia/Almaty или utc+6.\n")
		}
	}

	if !isOk {
		os.Exit(-1)
	}
}
