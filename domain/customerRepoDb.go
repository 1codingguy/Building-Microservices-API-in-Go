package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database Error")
		}
	}

	return &c, nil
}

func query(d CustomerRepositoryDb, urlQueryString string) (*sql.Rows, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	if urlQueryString == "" {
		return d.client.Query(findAllSql)
	} else {
		findAllSql = findAllSql + " " + "where status = ?"
		return d.client.Query(findAllSql, urlQueryString)
	}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	rows, err := query(d, status)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
