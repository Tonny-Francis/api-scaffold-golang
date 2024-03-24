FROM golang:1.22 as build

WORKDIR /app

ARG CONFIGMAP_CONTENT
ENV OUTPUT_ENV_FILE ./.env
ENV KUBERNETES_CONFIGMAP_PATH ./configmap.yml

RUN printf "%s" "$CONFIGMAP_CONTENT" > $KUBERNETES_CONFIGMAP_PATH

COPY go.mod ./
COPY go.sum ./
COPY Makefile ./
COPY pkg ./pkg
COPY cmd ./cmd
COPY internal ./internal

RUN make test

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exec cmd/api/main.go

RUN rm -f $KUBERNETES_CONFIGMAP_PATH
RUN rm -f $OUTPUT_ENV_FILE

FROM alpine:3

COPY --from=build /app/exec /exec

EXPOSE 8000

ENTRYPOINT ["/exec"]
