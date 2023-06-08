package main

import (
	"fmt"
	"net/http"
	"Kazbek2303.CleanArchTask/pkg/store/postgres"
	"Kazbek2303.CleanArchTask/services/contact/internal/delivery"
	"Kazbek2303.CleanArchTask/services/contact/internal/repository"
	usecase "Kazbek2303.CleanArchTask/services/contact/internal/useCase"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "kazbek",
		Password: "1234",
		DbName:   "kazbek_DB",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("started succesfully")
	http.ListenAndServe("localhost:4000", delivery.Mux)
}
