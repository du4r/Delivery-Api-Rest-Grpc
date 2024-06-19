package models

import (
	"errors"
	"fmt"
	"mega_api/db"
)

type Costumer struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}


func GetCostumer(id int64) (costumer Costumer, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM costumers WHERE id=$1`, id)

	err = row.Scan(&costumer.Id, &costumer.Name, &costumer.Address)

	return
}

func GetAllCostumers() (costumers []Costumer, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM costumers`)
	if err != nil {
		return
	}

	for rows.Next(){
		var costumer Costumer
		err = rows.Scan(&costumer.Id, &costumer.Name, &costumer.Address)

		if err != nil{
			continue
		}

		costumers = append(costumers, costumer)
	}

	return
}

func Insert(costumer Costumer) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM costumers WHERE name=$1`, costumer.Name)
	
	var exists Costumer 
	err = row.Scan(&exists.Id, &exists.Name, &exists.Address)

	if err != nil {
		fmt.Println(err)
	}else{
		return -1, errors.New("usuario ja cadastrado!")
	}

	q := `INSERT INTO costumers (name, address) values ($1, $2) RETURNING id`
	err = conn.QueryRow(q, costumer.Name, costumer.Address).Scan(&id)

	return
}

func Update(id int64, costumer Costumer) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE costumers SET name=$1, address=$2 WHERE id=$3`, costumer.Name,costumer.Address, costumer.Id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM costumers WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}