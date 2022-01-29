FROM debian:buster-slim

RUN apt-get update
RUN apt-get install golang -y

ADD go-conways-game-of-life /srv/go-conways-game-of-life
ADD conways.db /srv/conways.db

WORKDIR /srv
EXPOSE 8080

CMD ["./go-conways-game-of-life"]
