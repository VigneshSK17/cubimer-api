FROM golang:1.20.3 

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .
RUN go mod tidy

# RUN chmod +x ./docker_entrypoint.sh
# ENTRYPOINT [ "./docker_entrypoint.sh" ]
