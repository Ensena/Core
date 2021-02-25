env GOOS=linux GOARCH=amd64 go build -o core cmd/main.go
docker build -t ensena/core .
docker push ensena/core
rm core