FROM ubuntu:latest

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update
RUN apt-get install -y software-properties-common

RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get update

RUN apt-get install -y lsb-release curl ca-certificates gnupg sqlite3 make golang-go git

# Install keys for postgres repository
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN curl https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

# Install postgres 13
RUN apt-get update
RUN apt-get -y install postgresql-client-13

ADD . /db-backup
WORKDIR /db-backup

RUN make
CMD ["bin/db-backup"]
