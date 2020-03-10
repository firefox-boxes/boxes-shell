mkdir dist
go build -o dist/boxes-shell.exe %* boxes-shell.go shared.go
go build -o dist/boxes-ext-native-shell.exe %* boxes-ext-native-shell.go shared.go