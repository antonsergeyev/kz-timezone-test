package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("Текущее время в дефолтной таймзоне: %s\n", now.Format(time.RFC3339))

	boundsStart, boundsEnd := now.ZoneBounds()
	fmt.Printf("Диапазон действия текущей таймзоны: %s до %s\n", boundsStart.Format(time.RFC3339), boundsEnd.Format(time.RFC3339))

	timeInMarch := time.Date(2024, 3, 1, 2, 0, 0, 0, time.Local)
	marchBoundsStart, marchBoundsEnd := timeInMarch.ZoneBounds()
	fmt.Printf("Диапазон действия таймзоны в марте 2024: %s до %s\n", marchBoundsStart.Format(time.RFC3339), marchBoundsEnd.Format(time.RFC3339))

	if marchBoundsStart.Equal(boundsStart) {
		fmt.Printf("👎 Диапазоны действия текущей таймзоны и таймзоны в марте совпадают. Вероятно, пакет tzdata не обновлён, либо не настроена текущая таймзона.\n")
	} else {
		fmt.Printf("👍 Диапазоны действия текущей таймзоны и таймзоны в марте не совпадают. Вероятно, пакет tzdata уже обновлён.\n")
	}

	almaty, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		fmt.Printf("👎 Не удалось загрузить таймзону Asia/Almaty: %s. Вероятно, в ОС или образе не установлен пакет tzdata.\n", err)
	} else {
		fmt.Printf("Текущее время в таймзоне Asia/Almaty: %s\n", now.In(almaty).Format(time.RFC3339))

		if !now.In(almaty).Equal(now) {
			fmt.Printf("👎 Текущее время в дефолтной таймзоне не совпадает со временем Asia/Almaty! Установите таймзону Asia/Almaty или utc+6.\n")
		}
	}
}
