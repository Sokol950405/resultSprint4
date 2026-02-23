package daysteps

import (
	"errors"
	"fmt"
	"log"
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
	//get slice from string
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, errors.New("invalid line")
	}

	//get steps from string
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("invalid steps")
	}

	//get time of activity from string
	stepsDuration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	if stepsDuration <= 0 {
		return 0, 0, errors.New("invalid time")
	}

	return steps, stepsDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	////get steps and duration from string
	steps, stepsDuration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	//get distance
	distance := float64(steps) * stepLength / mInKm

	//get calories based on the type of activity
	numberCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, stepsDuration)
	if err != nil {
		return err.Error()
	}

	//get strings
	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		steps, distance, numberCalories)

	return result
}
