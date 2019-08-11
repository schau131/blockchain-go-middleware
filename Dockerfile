FROM golang:1.11.12

RUN apt-get update; 
#	apt-get install curl -y; \
#	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \

WORKDIR /go/src/github.com/filesystem

COPY filesystem .

#RUN dep init

EXPOSE 8080

CMD ["go", "run", "main.go"]