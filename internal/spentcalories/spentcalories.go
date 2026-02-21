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
		err := errors.New("invalid line")
		log.Println(err)
		return 0, "", 0, err
	}

	//get steps from string
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Println(err)
		return 0, "", 0, err
	}
	if steps <= 0 {
		err := errors.New("invalid steps")
		log.Println(err)
		return 0, "", 0, err
	}

	//get type of activity from string
	typeAct := slice[1]
	if typeAct == "" {
		err := errors.New("invalid typeAct")
		log.Println(err)
		return 0, "", 0, err
	}

	//get time of activity from string
	stepsDuration, err := time.ParseDuration(slice[2])
	if err != nil {
		log.Println(err)
		return 0, "", 0, err
	}
	if stepsDuration <= 0 {
		err := errors.New("invalid stepsDuration")
		log.Println(err)
		return 0, "", 0, err
	}

	return steps, typeAct, stepsDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	//get stride length
	strideLength := height * stepLengthCoefficient

	return float64(steps) * strideLength / float64(mInKm)
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
	//checking a character in a string
	if !strings.ContainsRune(data, ',') {
		err := errors.New("invalid line")
		log.Println(err)
		return "", err
	}
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
		log.Println(err)
		return "", err
	}

	//get mean speed
	speed := meanSpeed(steps, height, stepsDuration)

	//get distance
	distance := distance(steps, height)

	//get strings
	strTypeAct := fmt.Sprintf("Тип тренировки: %s", typeAct)
	strDuration := fmt.Sprintf("Длительность: %.2f ч.", stepsDuration.Hours())
	strDistance := fmt.Sprintf("Дистанция: %.2f км.", distance)
	strSpeed := fmt.Sprintf("Скорость: %.2f км/ч", speed)
	strCalories := fmt.Sprintf("Сожгли калорий: %.2f", numberCalories)

	//concatenate rows
	lines := []string{
		strTypeAct,
		strDuration,
		strDistance,
		strSpeed,
		strCalories,
	}
	result := strings.Join(lines, "\n") + "\n"
	log.Println(result)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//checking parameters
	if steps <= 0 {
		err := errors.New("steps are zero")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("weight are zero")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("height are zero")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("duration are zero")
		log.Println(err)
		return 0, err
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
		log.Println(err)
		return 0, err
	}

	//get walking spent calories
	return numberCalories * walkingCaloriesCoefficient, nil
}
