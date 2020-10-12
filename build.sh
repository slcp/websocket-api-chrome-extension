env GOOS=linux go build -o bin/loopback ./api/loopback
env GOOS=linux go build -o bin/connectionmanager ./api/connectionmanager
env GOOS=linux go build -o bin/apitrigger ./api/apitrigger