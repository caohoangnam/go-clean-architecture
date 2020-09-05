FROM golang

ARG app_env
ENV APP_ENV $app_env

# add files from Dockers client's current directory
COPY . /go-clean-architecture
WORKDIR /go-clean-architecture

RUN go get -u github.com/olivere/elastic
RUN go get -u github.com/nats-io/nats.go
RUN go get -u github.com/tinrab/retry


# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
    then \
    api; \
    else \
    go get github.com/pilu/fresh && \
    fresh; \
    fi

EXPOSE 8080
