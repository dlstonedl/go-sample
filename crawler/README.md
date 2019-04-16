### go：  
go get golang.org/x/text    
go get golang.org/x/net/html  

git clone --branch 7.x https://github.com/elastic/go-elasticsearch.git $GOPATH/src/github.com/elastic/go-elasticsearch/v7  

#### docker:    
docker ps -a  
docker images  
docker start {containerId}  
docker stop {containerId}  
docker rm {containerId}  
docker logs {containerId}  
docker rmi {imageId}  

docker run -d -p 9200:9200 -e "discovery.type=single-node" elasticsearch:7.0.0  

### elasticsearch:  
http://localhost:9200  
GET http://localhost:9200/crawler/zhenai/{id}?pretty=true  
GET http://localhost:9200/crawler/_search?pretty=true  

### architecture-diagram:   
starUML  

