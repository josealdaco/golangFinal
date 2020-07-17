
![](https://img.shields.io/badge/ImageSize-565MB-green)
![](https://img.shields.io/badge/build%20status-100%25-green)
![](https://img.shields.io/badge/Availability-100%25-green)


<h3>RUNNING CONTAINER</h3>

In Order to build the container you will need to run this command:
docker build -t ImageName .


The tag will default to latest if you run docker image ls, in order to run our container we will need to run this command:
docker run --publish port:8080 --name ContainerName --rm ImageName
