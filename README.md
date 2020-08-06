##Building Image

```
docker build -t system-api-img .
```


##Running container

For run the container you need use the next template:

```
docker run -p 9090:5000 --name system-api-con -e DbUser=[database user] -e  DbPassword=[database password] -e DbName=[database name] -d system-api-img
```

Here an example:


```
enviroment variables for database
 
 name : events
 user :  logmaster
 host : host.docker.internal
 port : 5432
 password : 9psql%Ple1

docker run -p 9090:5000 --name system-api-con -e DbHost=host.docker.internal -e DbPort=5432 -e DbUser=logmaster -e  DbPassword=9psql%Ple1 -e DbName=events -d system-api-img
```

##TODO

* benchmark test
* load charge definition





