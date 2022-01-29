##
# build stage
##
FROM golang:1.17 as build

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY api/ ./api
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY main.go ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /app

##
# run stage
##
FROM alpine:3.15

RUN mkdir "/uploads"
COPY --from=build /app /

CMD ["/app"]