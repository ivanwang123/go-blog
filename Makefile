.PHONY: create up down force generate air

create:
	migrate create -ext sql -dir db/migrations/ -seq add_relations

up:
	migrate -database postgres://postgres:postgres@localhost/go_blog?sslmode=disable -path db/migrations/ up

down:
	migrate -database postgres://postgres:postgres@localhost/go_blog?sslmode=disable -path db/migrations/ down

force:
	migrate -database postgres://postgres:postgres@localhost/go_blog?sslmode=disable -path db/migrations/ force 3

generate:
	modelq -db="user=postgres password=postgres dbname=go_blog sslmode=disable" -pkg=models -tables=users,posts,comments,likes -driver=postgres -schema=public -template=models.tmpl

air:
	air -c .air.toml