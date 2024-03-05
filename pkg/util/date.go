package util

import (
	"app-controller/pkg/model"
	"time"
)

func checkDeadline(task model.Task) (string , error){
	currentTime := time.Now()
	InStartDate, err := time.Parse("2006-01-02", task.StartDate)
	if err != nil {
		return "" , err
	}
	startDate := InStartDate.AddDate(0, 0, 1)

	InEndDate, err := time.Parse("2006-01-02", task.EndDate)
	if err != nil {
		return "" , err
	}
	endDate := InEndDate.AddDate(0, 0, 1)

	if currentTime.Before(startDate) {
		return "" , nil
	} else if currentTime.After(endDate) {
		return "late" , nil
	} else {
		daysUntilDeadline := int(endDate.Sub(currentTime).Hours() / 24)

		if daysUntilDeadline <= 1 {
			return "critical" , nil
		} else if daysUntilDeadline <= 3 {
			return "close" , nil
		}
	}

	return "process" , nil
}