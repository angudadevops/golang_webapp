FROM golang

COPY . ./

RUN go get -u github.com/go-sql-driver/mysql
RUN go build ./main.go 

EXPOSE 8080 

CMD ["./main"]
