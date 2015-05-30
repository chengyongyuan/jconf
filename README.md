# JSON based simple configuation package

Introduction
------------

The jconf package is simple json based config package in Go language.
Its main purpose is to make getting program configuation from file easy.
JSON is a simple and widely used data exchange format, json is also
human-readable, Meanwhile, it can easily processed by many lanuages.
Besides JSON, it also parse the normal apache style config file.


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
```html
[main]
IPLIST = 8.8.8.8, 1.1.1.1
Port = 80,443

[misc]
ServerName =  15T
ID = 1
```

```Go
//Initilization
 r, err := json.NewConfReader("basic.json")
 err = r.Init()

//Get Scalar Key or default value.
 servername := r.GetStr("ServerName", "default")
 id:= r.GetInt("ID", 0)

//Get Array Key or default value.
 sa := r.GetStrArray("IPLIST", []string{})
 ia := r.GetIntArray("Port", []int{})

//Apache like file with sector
 r, err := json.NewConfReader("simple_sect.json")
 err = r.Init()

//Get Scalar Key or default value.
 servername := r.GetStr("misc.ServerName", "default")

//Get Array Key or default value.
 sa := r.GetStrArray("main.IPLIST", []string{})
 ia := r.GetIntArray("main.Port", []int{})
 ```
Improvement
-----------

 Considering json nested object support? 
 such as:
 GetIntKey("key1.key2....keyN, xxx)
