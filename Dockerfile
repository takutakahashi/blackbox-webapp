FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && GO111MODULE=on go build -o app cmd/cmd.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app

EXPOSE 8000

CMD ["./app"]

