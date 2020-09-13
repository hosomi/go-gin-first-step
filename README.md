# Go lang & Gin Web Framework first step.

:link: [The Go Programming Language](https://golang.org/)  
:link: [Gin Web Framework](https://gin-gonic.com/)  

## 1. basic start of the first step.

## 1.1 basic setup

environment:  

- Windows10
- PowerShell
- Visual Studio Code :link: [Visual Studio Code – コード エディター | Microsoft Azure](https://azure.microsoft.com/ja-jp/products/visual-studio-code/)
- Go lang -> msi package(:link: [Download and install - The Go Programming Language](https://golang.org/doc/install?download=go1.15.2.windows-amd64.msi)) 


```powershell
PS > go version
go version go1.15.2 windows/amd64

PS > mkdir go-gin-first-step
PS > cd go-gin-first-step

PS go-gin-first-step> go mod init gin_test
go: creating new go.mod: module gin_test

PS go-gin-first-step> go get -u github.com/gin-gonic/gin
go: downloading github.com/gin-gonic/gin v1.6.3
go: github.com/gin-gonic/gin upgrade => v1.6.3
go: downloading github.com/gin-contrib/sse v0.1.0
go: downloading github.com/golang/protobuf v1.3.3
go: downloading gopkg.in/yaml.v2 v2.2.8
go: downloading github.com/json-iterator/go v1.1.9
go: downloading github.com/mattn/go-isatty v0.0.12
go: downloading github.com/go-playground/validator/v10 v10.2.0
go: downloading github.com/ugorji/go v1.1.7
go: downloading golang.org/x/sys v0.0.0-20200116001909-b77594299b42
go: downloading github.com/ugorji/go/codec v1.1.7
go: downloading github.com/leodido/go-urn v1.2.0
go: downloading github.com/go-playground/universal-translator v0.17.0
go: downloading github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
go: downloading github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421
go: downloading github.com/go-playground/locales v0.13.0
go: github.com/modern-go/reflect2 upgrade => v1.0.1
go: github.com/golang/protobuf upgrade => v1.4.2
go: github.com/modern-go/concurrent upgrade => v0.0.0-20180306012644-bacd9c7ef1dd
go: gopkg.in/yaml.v2 upgrade => v2.3.0
go: github.com/go-playground/validator/v10 upgrade => v10.3.0
go: github.com/json-iterator/go upgrade => v1.1.10
go: golang.org/x/sys upgrade => v0.0.0-20200909081042-eff7692f9009
go: downloading gopkg.in/yaml.v2 v2.3.0
go: downloading github.com/json-iterator/go v1.1.10
go: downloading github.com/golang/protobuf v1.4.2
go: downloading github.com/go-playground/validator/v10 v10.3.0
go: downloading golang.org/x/sys v0.0.0-20200909081042-eff7692f9009
go: downloading google.golang.org/protobuf v1.23.0
go: downloading github.com/modern-go/reflect2 v1.0.1
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: google.golang.org/protobuf upgrade => v1.25.0
go: downloading google.golang.org/protobuf v1.25.0

PS go-gin-first-step> code .
```

## 1.2 go run

```powershell
PS go-gin-first-step> go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func2 (4 handlers)
[GIN-debug] Loaded HTML Templates (3):
        - fileupload.html
        - index.html
        -

[GIN-debug] GET    /templates                --> main.main.func3 (4 handlers)
[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] GET    /upload                   --> main.main.func4 (4 handlers)
[GIN-debug] POST   /upload                   --> main.main.func5 (4 handlers)
[GIN-debug] Listening and serving HTTP on :3000
```


``http:localhost:3000`` :  

| path       | method | description
| ---------- | ------ | ---------------
| /          | GET    | hello world json
| /templates | GET    | template hello world
| /static/*  | GET    | engine.Static (``http://localhost:3000/static/gin.png``)  *gin.png 
| /upload    | GET    | file upload front
| /upload    | POST   | file upload /images 

*gin.png ：:link: [gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/gin-gonic/gin)   


---

[history](commits/master) :  
1. [setup](commit/8714097b1800dc2ca73734414dd8c4f969281d69)  
2. [hello world!](commit/5a2f66c5e9110f25ecc8f5070978bc10bdc784d0)
3. [templates](commit/842f245406260292afc302f6bdc98e6d90e25798)
4. [engine static](commit/aee3207b389d775f4305cdf0593ec71b63af2a4e)
5. [fileupload & here document](commit/71289644bafd9970b8a68e7fc57238c1ea868990)

　  

basic end. 

　  
　  

---

## 2. first step next REST API configuration.

## 2.1 setup database

environment:  
- MySQL 8.x(docker)

```powershell
go-gin-first-step
│  docker-compose.yml
│
└─docker
   └─mysql
       │  Dockerfile
       │
       ├─conf.d
       │      my.cnf
       │
       ├─data
       │
       └─initdb.d
              create_db.sql
```

1. docker-compose up -d
2. docker-compose ps
3. mysql -uroot -prootpassword -h 127.0.0.1 -P 3306

---

logs: 

1. ``docker-compose up -d``:  

```powershell
PS go-gin-first-step> docker-compose up -d
Creating network "go-gin-first-step_default" with the default driver
Step 1/2 : FROM mysql:8
8: Pulling from library/mysql
d121f8d1c412: Pull complete
f3cebc0b4691: Pull complete
1862755a0b37: Pull complete
489b44f3dbb4: Pull complete
690874f836db: Pull complete
baa8be383ffb: Pull complete
55356608b4ac: Pull complete
dd35ceccb6eb: Pull complete
429b35712b19: Pull complete
162d8291095c: Pull complete
5e500ef7181b: Pull complete
af7528e958b6: Pull complete
Digest: sha256:e1bfe11693ed2052cb3b4e5fa356c65381129e87e38551c6cd6ec532ebe0e808
Status: Downloaded newer image for mysql:8
 ---> e1d7dc9731da
Step 2/2 : EXPOSE 3306
 ---> Running in 62e45392cc93
Removing intermediate container 62e45392cc93
 ---> 2bd982e2e863

Successfully built 2bd982e2e863
Successfully tagged go-gin-first-step_mysql:latest
WARNING: Image for service mysql was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating mysql ... done
```


2. ``docker-compose ps``:  

```powershell
Name              Command             State                 Ports
-------------------------------------------------------------------------------
mysql   docker-entrypoint.sh mysqld   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
```

3. ``mysql -uroot -prootpassword -h 127.0.0.1 -P 3306``:  

```powershell
PS go-gin-first-step> mysql -uroot -prootpassword -h 127.0.0.1 -P 3306
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 9
Server version: 8.0.21 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test               |
+--------------------+
5 rows in set (0.01 sec)

mysql> exit
Bye

PS go-gin-first-step>
```


