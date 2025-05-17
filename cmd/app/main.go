package main

import (
	"ikan-nusa/internal/handler/rest"
	"ikan-nusa/internal/repository"
	"ikan-nusa/internal/service"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/config"
	"ikan-nusa/pkg/database/mariadb"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/supabase"
	"log"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	supabase := supabase.Init()
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwt, supabase)

	r := rest.NewRest(svc)
	r.MountEndpoint()
	r.Run()
}
