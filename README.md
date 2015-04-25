# JSON based simple configuation package

Introduction
------------

The jconf package is simple json based config package in Go language.
Its main purpose is to make getting program configuation from file easy.
JSON is a simple and widely used data exchange format, json is also
human-readable, Meanwhile, it can easily processed by many lanuages.


Installation and usage
----------------------

To install it, run:
go get chengyongyuan/jconf

License
-------
The jconf package is liscensed under MIT Liscese.

Example
-------

//basic.json
```json
{
    "ServerName": "testserver",
    "IPLIST": ["127.0.0.1", "192.168.0.1", "192.168.0.3"],
    "Port": [80, 443, 14000],
    "ID": 8888
}
```
```Go
//Initilization
 err := Init("basic.json")

 //Get Scalar Key
 servername := GetStr("ServerName", "default")
 id:= GetInt("ID", 0)

 //Get Array Key
 sa := GetStrArray("IPLIST", []string{})
 ia := GetIntArray("Port", []int{})
 ```
Improvement
-----------

 Considering json nest object support? 
 such as:
 GetIntKey("key1.key2....keyn, defaulint)
