FROM golang:1.19.3-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY main.go ./

RUN go build -o /PrometheusAppTest

FROM gcr.io/distroless/base-debian11

COPY --from=build /PrometheusAppTest /PrometheusAppTest

ENTRYPOINT ["/PrometheusAppTest"]

