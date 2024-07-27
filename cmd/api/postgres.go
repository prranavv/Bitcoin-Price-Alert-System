package main

import (
	"context"
	"time"
)

func (m *DB) AddingToAlert(price int, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO public.alert(
 	price, status) VALUES ($1,$2);`
	_, err := m.SQL.ExecContext(ctx, query, price, status)
	return err
}

func (m *DB) UpdatingFromAlert(alertId int, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `UPDATE public.alert
	SET alertid=$1, status=$2
	WHERE alertid=$3`
	_, err := m.SQL.ExecContext(ctx, query, alertId, status, alertId)
	return err
}

func (m *DB) GettingFromAlert() ([]Alert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT alertid, price, status
	FROM public.alert`
	rows, err := m.SQL.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var alerts []Alert
	for rows.Next() {
		var alert Alert
		err := rows.Scan(
			&alert.AlertID,
			&alert.Price,
			&alert.Status,
		)
		if err != nil {
			return alerts, nil
		}
		alerts = append(alerts, alert)
	}
	if err := rows.Err(); err != nil {
		return alerts, err
	}
	return alerts, nil
}

func (m *DB) InsertingIntoUser(credential Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO public.users(
 username, password)
	VALUES ($1,$2);`
	_, err := m.SQL.ExecContext(ctx, query, credential.Username, credential.Password)
	return err
}

func (m *DB) GettingFromUser(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT password
	FROM public.users where username=$1;`
	var password string
	err := m.SQL.QueryRowContext(ctx, query, username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
