FROM golang:1.19-alpine as build

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build ./main.go

FROM scratch as image

COPY --from=build /app/main .

CMD ["/main"]