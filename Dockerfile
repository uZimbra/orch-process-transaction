FROM golang:1.21.3-alpine AS base
WORKDIR /src/app
COPY go.mod go.sum ./
COPY . .
RUN apk add --no-cache librdkafka-dev build-base
RUN go build -tags musl -o orch-process-transactions ./cmd

FROM alpine:edge
WORKDIR /www
COPY --from=base /src/app/orch-process-transactions .
COPY --from=base /src/app/profile-local.yaml .
COPY --from=base /src/app/profile-development.yaml .
CMD ["./orch-process-transactions"]