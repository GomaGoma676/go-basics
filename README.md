```bash
# for mac user
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
#create project folder
mkdir go-basics
cd go-basics
#init module
go mod init go-basics
#automatically install/uninstall external package
go mod tidy
#run main.go
go run main.go
or
go run.
```
```bash
#run with race data checker
go run -race .
#run test
go test -v .
#launch tracer view
go tool trace trace.out
```
```bash
# VS Code settings.json
 "[go]": {
   "editor.defaultFormatter": "golang.go",
   "editor.formatOnSave": true
 },
```