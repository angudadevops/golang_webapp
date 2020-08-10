FROM golang

COPY . ./
RUN go build ./main.go 

EXPOSE 8080 

CMD ["./main"]
