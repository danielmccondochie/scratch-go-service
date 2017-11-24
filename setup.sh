CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src
docker rm -f scratch-go-service
docker rmi scratch-go-service
docker build -t scratch-go-service .
docker run -d --name scratch-go-service --net=host scratch-go-service 
