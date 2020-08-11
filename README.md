# Golang Web Application

This Web App helps to user compose and send an email with Golang

Prerequisites
- GO

# Web Applications
- [Mail Web Application](#Mail-Application)
- [Web Application with DB](#Web-Application-with-DB)

## Mail Application

Run the below command to access the application 
```
git clone https://github.com/angudadevops/golang_webapp.git

go run mail.go
```

If you don't have go environment, you can use docker to use this application 

Run the below command to access the web application 
```
sudo docker run -d --name goweb -p 80:8080 anguda/golang:web-app1
```

Access application from your browser with below url
```
http://hostIP:8080
```
## Web Application with DB

Run the below command to access the application 
```
git clone https://github.com/angudadevops/golang_webapp.git

go run main.go
```

If you don't have go environment, you can use docker to use this application 

Run the below commands to create mysql DB with Golang Web Application 

```
docker run -d -p 3306:3306 --name mysql anguda/mysql
docker run -d -p 8081:8080 --name gowebapp --link mysql anguda/golang:web-app
```
Access application from your browser with below url
```
http://hostIP:8081
```
