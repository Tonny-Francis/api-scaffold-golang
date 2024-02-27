FROM golang:1.22 as build

WORKDIR /app

ARG CONFIGMAP_CONTENT
ENV OUTPUT_ENV_FILE ./.env
ENV KUBERNETES_CONFIGMAP_PATH ./configmap.yml

RUN printf "%s" "$CONFIGMAP_CONTENT" > $KUBERNETES_CONFIGMAP_PATH

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exec cmd/api/main.go

FROM alpine:3

COPY --from=build /app/exec /exec
COPY --from=build /app/env/.env /env/.env

EXPOSE 8000

ENTRYPOINT ["/exec"]
