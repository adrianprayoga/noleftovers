package models

import (
	"database/sql"
	"fmt"
)

type Measure struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type MeasureService struct {
	DB *sql.DB
}

func (service *MeasureService) GetMeasures() ([]Measure, error) {

	rows, err := service.DB.Query(`SELECT id, name, active FROM measure`)
	if err != nil {
		fmt.Println("Error getting recipe list")
		return nil, fmt.Errorf("error getting measure list: %w", err)
	}
	defer rows.Close()

	var measures []Measure
	for rows.Next() {
		measure := Measure{}
		err = rows.Scan(&measure.Id, &measure.Name, &measure.Active)
		if err != nil {
			fmt.Println("Error reading recipe")
			return nil, fmt.Errorf("error reading measure: %w", err)
		}

		measures = append(measures, measure)
	}

	return measures, nil
}
