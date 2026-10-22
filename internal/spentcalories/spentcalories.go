package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	allParameters := strings.Split(data, ",")
	if len(allParameters) != 3 {
		return 0, "", 0, fmt.Errorf("parseTraining: invalid data format %q", data)
	}

	steps, err := strconv.Atoi(allParameters[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseTraining: invalid steps: %v", err)
	}
	if steps < 1 {
		return 0, "", 0, fmt.Errorf("parseTraining: steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(allParameters[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseTraining: invalid duration: %v", err)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("parseTraining: duration must be positive, got %v", duration)
	}

	activity := allParameters[1]

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceInKm := (stepLength * float64(steps)) / mInKm

	return distanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	distanceInKm := distance(steps, height)
	avgSpeed := distanceInKm / duration.Hours()
	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var spentCalories float64

	switch activity {
	case "Бег":
		spentCalories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		spentCalories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	distanceInKm := distance(steps, height)
	avgSpeed := meanSpeed(steps, height, duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), distanceInKm, avgSpeed, spentCalories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, fmt.Errorf("RunningSpentCalories: steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("RunningSpentCalories: weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("RunningSpentCalories: height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("RunningSpentCalories: duration must be positive, got %v", duration)
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * avgSpeed * durationInMinutes) / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, fmt.Errorf("WalkingSpentCalories: steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("WalkingSpentCalories: weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("WalkingSpentCalories: height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("WalkingSpentCalories: duration must be positive, got %v", duration)
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * avgSpeed * durationInMinutes) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}
