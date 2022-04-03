package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/aintsashqa/roles-and-permissions/ent"
	"github.com/aintsashqa/roles-and-permissions/internal/delivery"
	"github.com/aintsashqa/roles-and-permissions/internal/delivery/rest"
	"github.com/aintsashqa/roles-and-permissions/internal/services"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open(dialect.Postgres, "postgres://username:password@postgres:5432/dev-db?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	roleService := services.NewRoleServiceImpl(client)
	container := delivery.Container{
		RoleService: roleService,
	}

	handler := rest.RegisterRoutes(&container)
	if err := http.ListenAndServe(":3200", handler); err != nil {
		log.Fatalln(err)
	}
}
