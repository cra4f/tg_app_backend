package database

import (
	"fmt"
	"models"
)

func (d *Database) GetTasks() ([]*models.Personage, error) {
	var personages []*models.Personage
	request := fmt.Sprintf("SELECT id, name, description, percent, earn_period, life_period FROM public.personages;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return personages, err
	} else {
		for rows.Next() {
			var id int
			var name, description string
			var percent, earn_period, life_period float64
			rows.Scan(&id, &name, &description, &percent, &earn_period, &life_period)
			var personage models.Personage
			personage.Id = id
			personage.Name = name
			personage.Description = description
			personage.EarnPercent = percent
			personage.EarnPeriod = earn_period
			personage.LifePeriod = life_period
			personages = append(personages, &personage)
		}
	}

	return personages, nil
}
