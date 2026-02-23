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
	//get slice from string
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, errors.New("invalid line")
	}

	//get steps from string
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("invalid steps")
	}

	//get type of activity from string
	typeAct := slice[1]
	if typeAct == "" {
		return 0, "", 0, errors.New("invalid typeAct")
	}

	//get time of activity from string
	stepsDuration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}
	if stepsDuration <= 0 {
		return 0, "", 0, errors.New("invalid stepsDuration")
	}

	return steps, typeAct, stepsDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	//get stride length
	strideLength := height * stepLengthCoefficient

	return float64(steps) * strideLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	//get distance
	distance := distance(steps, height)

	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	//get parameters from string
	steps, typeAct, stepsDuration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//get calories based on the type of activity
	var numberCalories float64
	switch typeAct {
	case "Ходьба":
		numberCalories, err = WalkingSpentCalories(steps, weight, height, stepsDuration)
	case "Бег":
		numberCalories, err = RunningSpentCalories(steps, weight, height, stepsDuration)
	default:
		err = errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	//get mean speed
	speed := meanSpeed(steps, height, stepsDuration)

	//get distance
	distance := distance(steps, height)

	//get strings
	//get strings
	result := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		typeAct, stepsDuration.Hours(), distance, speed, numberCalories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//checking parameters
	if steps <= 0 {
		return 0, errors.New("steps are zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight are zero")
	}
	if height <= 0 {
		return 0, errors.New("height are zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration are zero")
	}

	//get number of calories
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//get number of calories
	numberCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	//get walking spent calories
	return numberCalories * walkingCaloriesCoefficient, nil
}
