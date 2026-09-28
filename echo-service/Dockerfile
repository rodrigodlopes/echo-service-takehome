FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /echo-service .

FROM scratch
COPY --from=build /echo-service /echo-service
EXPOSE 8080
ENTRYPOINT ["/echo-service"]
