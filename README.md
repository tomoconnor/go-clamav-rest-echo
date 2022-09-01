## go-clamav-rest-echo

ClamAV proxy based on https://github.com/asmith030/go-clamav-rest/ but reimplemented using the Echo v4 Framework. 

Usage
======

Usage of go-clamav-rest-echo: 

Set Environment Variables: 
- CLAMD_HOST = localhost
- CLAMD_PORT = 3310
- LISTEN_PORT = 8080

Docker:
----------

If clamav is in a separate container: 

`docker run -e CLAMD_HOST=clamav --link=clamav -p 8080:8080 devopstom/go-clamav-rest-echo:latest`

To run clamAV's scanner container, I'm using: 
https://github.com/cabinetoffice/docker-clamav

`docker run --name clamav -v ${PWD}/clamdata:/var/lib/clamav -d -p 3310:3310 quay.io/ukhomeofficedigital/clamav:latest`


Docker Hub:
-----------
https://hub.docker.com/r/devopstom/go-clamav-rest-echo