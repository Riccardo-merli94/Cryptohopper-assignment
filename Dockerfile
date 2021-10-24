FROM golang:1.17.2-alpine3.14 as build-env
RUN adduser -D -u 1000 hopper
RUN mkdir /hopper/ && chown hopper /hopper/
USER hopper


WORKDIR /hopper/
ADD ./app /hopper/

RUN CGO_ENABLED=0 go build -o /hopper/build .

FROM alpine:3.14

RUN adduser -D -u 1000 hopper
USER hopper


WORKDIR /

COPY --from=build-env /hopper/build /hopper

EXPOSE 8080

CMD ["/hopper"]