###go：
go get golang.org/x/text  
go get golang.org/x/net/html

export GO111MODULE=on
go mod init github.com/dlstonedl/go-sample/crawler
go mod download

####docker:
docker ps -a
docker images
docker start {containerId}
docker stop {containerId}
docker rm {containerId}
docker logs {containerId}
docker rmi {imageId}

docker run -d -p 9200:9200 -e "discovery.type=single-node" elasticsearch:7.0.0

###elasticsearch:
localhost:9200

###architecture-diagram: 
starUML

