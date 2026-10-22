package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	stepsAndTime := strings.Split(data, ",")
	if len(stepsAndTime) != 2 {
		return 0, 0, fmt.Errorf("parsePackage: invalid data format %q", data)
	}

	steps, err := strconv.Atoi(stepsAndTime[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parsePackage: invalid steps: %v", err)
	}
	if steps < 1 {
		return 0, 0, fmt.Errorf("parsePackage: steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(stepsAndTime[1])
	if err != nil {
		return 0, 0, fmt.Errorf("parsePackage: invalid duration: %v", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("parsePackage: duration must be positive, got %v", duration)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps < 1 {
		return ""
	}

	distanceInMeters := float64(steps) * stepLength
	distanceInKm := distanceInMeters / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceInKm, spentCalories,
	)
}
