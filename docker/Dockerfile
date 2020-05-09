FROM golang
WORKDIR /go/src/docker

RUN apt-get update && apt-get install -y vim \
&& go get -u github.com/go-sql-driver/mysql  \
&& go get -u github.com/astaxie/beego \
&& go get -u github.com/beego/bee \
&& go get -u github.com/go-redis/redis \
&& go get -u github.com/nareix/curl \
&& go get -u github.com/onsi/ginkgo \
&& go get -u github.com/onsi/gomega \
&& go get -u github.com/shiena/ansicolor \
&& go get -u github.com/smartystreets/goconvey \
&& go get -u google.golang.org/appengine \
&& go get -u gopkg.in/robfig/cron.v2

 
EXPOSE 7699