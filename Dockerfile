FROM golang:1.16.3-buster
RUN go get -u github.com/beego/bee
RUN chmod -R ugo+rw /tmp
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/mathapp
#ARG GROUP_ID
#ARG USER_ID
RUN groupadd --gid 2021 app && useradd -m -l --uid 1202 --gid 2021 $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
#USER $APP_USER
WORKDIR $APP_HOME
EXPOSE 8010
CMD ["bee", "run"]
