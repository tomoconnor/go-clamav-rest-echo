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

API:
----

/ and /healthz  Both run a clamd.Ping and return OK if ClamD is contactable, error if not. 
Example:
```
curl -i http://localhost:8080/
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 09:54:30 GMT
Content-Length: 5

"OK"
```

error: 

```
curl -i http://localhost:8080/
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 10:01:15 GMT
Content-Length: 35

{"message":"Could not ping clamd"}
```

Scanning Files
--------------

`/scan` returns a simple Yes/No (actually with a message) and either HTTP Status code 200 or 451 :D 

```
curl -i -F "name=eicar" -F "file=@./eicar.com" http://localhost:8080/scan
HTTP/1.1 451 Unavailable For Legal Reasons
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 10:09:38 GMT
Content-Length: 31

{"message":"Malware detected"}
```

`/scanResponse` returns a JSON object with the information of what was found in Raw and Description fields.

```
curl -i -F "name=eicar" -F "file=@./eicar.com" http://localhost:8080/scanResponse
HTTP/1.1 451 Unavailable For Legal Reasons
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 10:09:45 GMT
Content-Length: 134

{"Raw":"stream: Win.Test.EICAR_HDB-1 FOUND","Description":"Win.Test.EICAR_HDB-1","Path":"stream","Hash":"","Size":0,"Status":"FOUND"}
```

Clean Files:
```
curl -i -F "name=clean" -F "file=@/bin/true" http://localhost:8080/scanResponse
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 10:12:04 GMT
Content-Length: 87

{"Raw":"stream: OK","Description":"","Path":"stream","Hash":"","Size":0,"Status":"OK"}
```

```
user@sugarloaf:~/git/docker-clamav$ curl -i -F "name=clean" -F "file=@/bin/true" http://localhost:8080/scan
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Fri, 02 Sep 2022 10:12:17 GMT
Content-Length: 5

"OK"
```
