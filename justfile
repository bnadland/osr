set dotenv-load

up:
    docker compose up --wait
    air

down:
    docker compose down

deps:
    asdf plugin add golang
    asdf plugin add nodejs
    asdf install
    npm install
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
    go install github.com/cosmtrek/air@latest
    go install github.com/a-h/templ/cmd/templ@latest
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    go install github.com/amacneil/dbmate@latest
    asdf reshim golang

pre: css templ

css:
    npx tailwind -i ./app.css -o ./assets/app.css

fonts:
    cp ./node_modules/@fontsource/inter/files/inter-latin-400-normal.woff* ./assets/

templ:
    templ generate .

fmt:
    go fmt ./...
    templ fmt .

lint:
    go mod tidy
    golangci-lint run

test:
    go test ./...

reset:
    dbmate rollback
    dbmate migrate
    sqlc generate

psql:
    psql ${DATABASE_URL}