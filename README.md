# Go lang & Gin Web Framework first step.

:link: [The Go Programming Language](https://golang.org/)  
:link: [Gin Web Framework](https://gin-gonic.com/)  

## setup

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

## go run

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
