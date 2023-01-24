FROM ubuntu:latest

WORKDIR /app

EXPOSE 8000

ENV DB_HOST=localhost DB_PORT=5432

ENV DB_USER=root DB_PASS=root DB_NAME=root

COPY ./main main

CMD [ "./main" ]
