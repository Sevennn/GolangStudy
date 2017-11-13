# ReadMe & Test Report
### 选择框架: Martini
### 原因: revel链接失效，根据老师推荐择优选取了martini
### 框架优势:
- 使用极其简单.
- 无侵入式的设计.
- 很好的与其他的Go语言包协同使用.
- 超赞的路径匹配和路由.
- 模块化的设计 - 容易插入功能件，也容易将其拔出来.
- 已有很多的中间件可以直接使用.
- 框架内已拥有很好的开箱即用的功能支持.
- 完全兼容http.HandlerFunc接口.
- 个人使用起来也较为顺手，类似于nodejs-express开发
### Test Part
For curl:
```
* Rebuilt URL to: http://localhost:3000/
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.38.0
> Host: localhost:3000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 13 Nov 2017 11:40:11 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
```
For ab:
```
This is ApacheBench, Version 2.3 <$Revision: 1604373 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      100
Time taken for tests:   1.633 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      129000 bytes
HTML transferred:       12000 bytes
Requests per second:    612.30 [#/sec] (mean)
Time per request:       163.317 [ms] (mean)
Time per request:       1.633 [ms] (mean, across all concurrent requests)
Transfer rate:          77.14 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0      12
Processing:     1  153 144.6    111     901
Waiting:        1  148 146.6    102     900
Total:          2  154 144.6    111     901

Percentage of the requests served within a certain time (ms)
  50%    111
  66%    184
  75%    213
  80%    238
  90%    350
  95%    456
  98%    548
  99%    634
 100%    901 (longest request)
```
每秒的请求数目为 612.30
平均请求响应时间为 163.317ms
每个请求实际相应时间的平均值 1.633ms

50％的用户响应时间小于111毫秒，最大的响应时间小于901毫秒