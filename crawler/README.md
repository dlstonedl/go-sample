### go：  
go get golang.org/x/text    
go get golang.org/x/net/html  
go get github.com/olivere/elastic/  

#### docker:    
docker ps -a  
docker images  
docker start {containerId}  
docker stop {containerId}  
docker rm {containerId}  
docker logs {containerId}  
docker rmi {imageId}  

docker pull elasticsearch:6.7.1
docker run -d -p 9200:9200 elasticsearch:6.7.1
http://localhost:9200  

docker pull kibana:6.7.1
http://localhost:5601

docker-compose up &

### elasticsearch:  
http://localhost:9200  
GET http://localhost:9200/crawler/zhenai/{id}?pretty=true  
GET http://localhost:9200/crawler/_search?pretty=true  

### architecture-diagram:   
starUML  

