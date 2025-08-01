package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	caloriesSpent, err := RunningSpentCalories(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	return caloriesSpent * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if weight <= 0 {
		return 0, errors.New("invalid weight: must be a positive integer")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	if averageSpeed <= 0 {
		return 0, errors.New("invalid meanSpeed: must be a positive integer")
	}

	return (weight * averageSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	fullLenStep := height * stepLengthCoefficient
	fullSteps := float64(steps) * fullLenStep

	return fullSteps / mInKm
}
