FROM golang AS build

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 go build -o main .

FROM scratch

WORKDIR /app

COPY --from=build /app/config.yaml .
COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/main .

ENTRYPOINT ["./main"]