package postgresql

import (
	"fmt"
)

type Personage struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Percent     float32 `json:"percent"`
	EarnPeriod  float32 `json:"earn_period"`
	LifePeriod  float32 `json:"life_period"`
}

type UserPersonage struct {
	Id          int    `json:"id"`
	PersonageId int    `json:"personage_id"`
	BuyAt       string `json:"buy_at"`
	Active      bool   `json:"active"`
}

func (p *Postgresql) GetPersonages() ([]*Personage, error) {
	var personages []*Personage
	request := "SELECT id, name, description, percent, earn_period, life_period FROM public.personages"
	rows, err := p.db.Query(request)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var id int
			var name, description string
			var percent, earnPeriod, lifePeriod float32

			rows.Scan(&id, &name, &description, &percent, &earnPeriod, &lifePeriod)

			var personage Personage
			personage.Id = id
			personage.Name = name
			personage.Description = description
			personage.EarnPeriod = earnPeriod
			personage.LifePeriod = lifePeriod
			personages = append(personages, &personage)
		}
	}
	return personages, nil
}

func (p *Postgresql) AddPersonage(personage Personage) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.encoders(name, description, percent, earn_period, life_period) VALUES('%s', '%s', '%f', '%f', '%f') RETURNING id;", personage.Name, personage.Description, personage.Percent, personage.EarnPeriod, personage.LifePeriod)
	stmt, err := p.db.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

// вообще должен же быть метод, чтобы не update все поля, а только те, которые изменились?
func (p *Postgresql) EditPersonage(personage Personage) error {
	request := fmt.Sprintf("UPDATE public.personages SET name = '%s', description = '%s', percent = '%f', earn_period = '%f', life_period = '%f' WHERE id = %d", personage.Name, personage.Description, personage.Percent, personage.EarnPeriod, personage.LifePeriod, personage.Id)

	stmt, err := p.db.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}

// возможно надо перенести в user
func (p *Postgresql) GetUserPersonages(userId int) ([]*UserPersonage, error) {
	var userPersonages []*UserPersonage
	request := fmt.Sprintf("SELECT id, personage_id, buy_at, active FROM public.users_personages_link WHERE user_id = %d", userId)
	rows, err := p.db.Query(request)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var id, personageId int
			var buyAt string
			var active bool

			rows.Scan(&id, &personageId, &buyAt, &active)

			var userPersonage UserPersonage
			userPersonage.Id = id
			userPersonage.PersonageId = personageId
			userPersonage.BuyAt = buyAt
			userPersonage.Active = active
			userPersonages = append(userPersonages, &userPersonage)
		}
	}
	return userPersonages, nil
}

// надо подумать что делать со временем buy_at и тут же надо списывать с баланса средства!
func (p *Postgresql) BuyPersonage(userId int, personageId int) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.users_personages_link(user_id, personage_id, active) VALUES('%d', '%d', true) RETURNING id;", userId, personageId)
	stmt, err := p.db.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}
