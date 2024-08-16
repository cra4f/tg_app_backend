package postgresql

import "fmt"

type Balance struct {
	Gold float32 `json:"gold"`
}

func (p *Postgresql) GetUserBalance(user_id int) (*Balance, error) {
	var balance Balance
	request := fmt.Sprintf("SELECT gold FROM public.balances WHERE user_id = %d;", user_id)
	rows := p.db.QueryRow(request)

	var gold float32
	if err := rows.Scan(&gold); err != nil {
		return nil, err
	}

	balance.Gold = gold

	return &balance, nil
}

func (p *Postgresql) UpdateUserBalanceGold(user_id int, gold int) error {
	return nil

}
