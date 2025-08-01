package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		return errors.New("invalid input: expected at least two comma-separated values")
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("invalid steps count: must be a positive integer")
	}

	ds.Steps = steps

	parsedTime, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return err
	}
	if parsedTime <= 0 {
		return errors.New("invalid duration format: must be a positive integer")
	}

	ds.Duration = parsedTime

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	caloriesSpent, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		spentenergy.Distance(ds.Steps, ds.Height),
		caloriesSpent,
	), nil
}
