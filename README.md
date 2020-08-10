# Golang Web Application

This Web App helps to user compose and send an email with Golang

Prerequisites
- GO

Run the below command to access the application 
```
git clone https://github.com/angudadevops/golang_webapp.git

go run main.go
```
Access application from your browser with below url

```
http://hostIP:8080
```

If you don't have go environment, you can use docker to use this application 

run the below command to build the docker image
```

sudo docker build -t gowebapp . --no-cache
```

Run the below command to access the web application 
```
sudo docker run -d --name goweb -p 80:8080 gowebapp
```
