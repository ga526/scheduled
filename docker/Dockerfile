FROM golang
WORKDIR /go/src/docker

RUN apt-get clean && apt-get update && apt-get install -y vim
RUN go get -u github.com/gpmgo/gopm \
&& gopm get -u github.com/go-sql-driver/mysql \
&& gopm get -u github.com/astaxie/beego \
&& gopm get -u github.com/beego/bee \
&& gopm get -u github.com/go-redis/redis \
&& gopm get -u github.com/nareix/curl \
&& gopm get -u github.com/onsi/ginkgo \
&& gopm get -u github.com/onsi/gomega \
&& gopm get -u github.com/shiena/ansicolor \
&& gopm get -u github.com/smartystreets/goconvey \
&& gopm get -u google.golang.org/appengine \
&& gopm get -u gopkg.in/robfig/cron.v2

EXPOSE 7699