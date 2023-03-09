FROM alpine:3.17.0

WORKDIR /app

COPY . ./main

EXPOSE 8080

CMD "./main"