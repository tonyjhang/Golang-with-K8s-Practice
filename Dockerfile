FROM golang:1.16.5-alpine3.13 AS build
WORKDIR /src
COPY . /src
RUN go build .

FROM alpine:3.11.3
COPY --from=build /src /
EXPOSE 8888
ENTRYPOINT ["./goRestfulPractice"]