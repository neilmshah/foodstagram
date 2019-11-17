FROM golang:latest 
EXPOSE 8080
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options
RUN go get golang.org/x/crypto/bcrypt
RUN cd /app ; go install go-login
CMD ["/app/bin/go-login"]
