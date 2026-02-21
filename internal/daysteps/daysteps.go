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
	//checking a character in a string
	if !strings.ContainsRune(data, ',') {
		err := errors.New("")
		log.Println(err)
		return 0, 0, err
	}
	//get slice from string
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		err := errors.New("")
		log.Println(err)
		return 0, 0, err
	}

	//get steps from string
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if steps <= 0 {
		err := errors.New("")
		log.Println(err)
		return 0, 0, err
	}

	//get time of activity from string
	stepsDuration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	if stepsDuration <= 0 {
		err := errors.New("")
		log.Println(err)
		return 0, 0, err
	}

	return steps, stepsDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	////get steps and duration from string
	steps, stepsDuration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	if steps <= 0 || stepsDuration <= 0 {
		err := errors.New("")
		log.Println(err)
		return ""
	}
	//get distance
	distance := float64(steps) * stepLength / float64(mInKm)

	//get calories based on the type of activity
	numberCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, stepsDuration)
	if err != nil {
		return err.Error()
	}

	//get strings
	strSteps := fmt.Sprintf("Количество шагов: %d.", steps)
	strDistance := fmt.Sprintf("Дистанция составила %.2f км.", distance)
	strCalories := fmt.Sprintf("Вы сожгли %.2f ккал.", numberCalories)

	//concatenate rows
	lines := []string{
		strSteps,
		strDistance,
		strCalories,
	}
	result := strings.Join(lines, "\n") + "\n"
	return result
}
