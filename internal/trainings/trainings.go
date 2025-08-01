package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 3 {
		return errors.New("invalid input: expected at least two comma-separated values")
	}

	t.TrainingType = dataSlice[1]

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("invalid steps count: must be a positive integer")
	}
	t.Steps = steps

	parsedTime, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return err
	}
	if parsedTime <= 0 {
		return errors.New("invalid duration format: must be a positive integer")
	}
	t.Duration = parsedTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var caloriesSpent float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		caloriesSpent, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		caloriesSpent, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки" + t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
		caloriesSpent,
	), nil
}
