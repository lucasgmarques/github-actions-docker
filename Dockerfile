FROM ubuntu:latest

WORKDIR /app

ENV HOST=localhost PORT=5432

ENV USER=root PASSWORD=root DBNAME=root

COPY ./main main

EXPOSE 8000

CMD [ "./main" ]
