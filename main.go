package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	phpPathFlag := flag.String("php-path", "", "путь к исполняемому файлу PHP, если также нужно проверить таймзоны в PHP")
	flag.Parse()

	isOk := true
	now := time.Now()
	fmt.Printf("Текущее время в дефолтной таймзоне: %s\n", now.Format(time.RFC3339))

	offsetName, offset := now.Zone()
	boundsStart, boundsEnd := now.ZoneBounds()
	fmt.Printf(
		"Текущий смещение: %s (%d секунд), действует от %s до %s\n",
		offsetName, offset,
		boundsStart.Format(time.RFC3339), boundsEnd.Format(time.RFC3339),
	)

	timeInMarch := time.Date(2024, 3, 1, 2, 0, 0, 0, time.Local)
	marchOffsetName, marchOffset := timeInMarch.Zone()
	marchBoundsStart, marchBoundsEnd := timeInMarch.ZoneBounds()
	fmt.Printf(
		"Смещение в марте 2024: %s (%d секунд), действует от %s до %s\n",
		marchOffsetName, marchOffset,
		marchBoundsStart.Format(time.RFC3339), marchBoundsEnd.Format(time.RFC3339),
	)

	if marchOffset != 60*60*5 {
		isOk = false
		fmt.Printf("👎 Смещение в марте не равен ожидаемому. Вероятно, пакет tzdata не обновлён, либо не установлен, либо не настроена текущая таймзона.\n")
	} else {
		fmt.Printf("👍 Смещение в марте равен ожидаемому. Вероятно, пакет tzdata уже обновлён.\n")
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

	if *phpPathFlag != "" {
		phpInfo, phpErr := checkPhpTimezone(*phpPathFlag)
		if phpErr != nil {
			isOk = false
			fmt.Printf(phpErr.Error() + "\n")
		}

		fmt.Println(phpInfo)
	}

	if !isOk {
		os.Exit(-1)
	}
}

// Проверяет, что для дат после марта 2024 в Asia/Almaty PHP вернёт оффсет, равный 5 часам
func checkPhpTimezone(phpPath string) (string, error) {
	cmd := exec.Command(
		phpPath,
		"-r",
		`echo (new \DateTime('2024-03-02T00:00:00', new \DateTimeZone('Asia/Almaty')))->getOffset();`,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("👎 Не удалось получить информацию о таймзоне PHP %s: %w", phpPath, err)
	}

	expected := "18000"

	if string(out) != expected {
		return "", fmt.Errorf(
			`👎 Смещение в марте на PHP %s не равен ожидаемому: получили %s, ожидалось %s секунд. `+
				`Нужно обновить расширение timezonedb: https://serverpilot.io/docs/how-to-update-the-php-timezonedb-version/`,
			phpPath,
			expected,
			string(out),
		)
	} else {
		return fmt.Sprintf("👍 Смещение в марте на PHP %s равен ожидаемому: %s", phpPath, string(out)), nil
	}
}
