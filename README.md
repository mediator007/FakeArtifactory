# FakeArtifactory
FakeArtifafctory - object storage, which can recieve and return files by artifactory paths.  

FE ```localhost:8080```  

### Deploy
For docker compose v2  
```docker compose build```  
```docker compose up -d```  

### API

PUT file:  
```curl -X PUT "localhost:8080/artifactory/some/random/path/" -T ./qwerty.txt```  

GET file:  
```curl "localhost:8080/artifactory/some/random/path/qwerty.txt"```  
