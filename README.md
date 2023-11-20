# custom-fetch
## _Overview_
custom fetch is a simple web scrapping built using Go (Golang). It is based on Go's [net/html](https://pkg.go.dev/net/http) package in Go to get the response from destination urls and additional [goquery](https://github.com/PuerkitoBio/goquery) package to read and inspect the page elements. 

## _Table of Contents_
- [Features](#features)
- [Installation](#installation)
- [Examples](#examples)

## Features
There are 2 main features :
- Fetch and store contents of fetched url to html file (.html)
- Fetch metadata of fetched URL

## Installation
Just Clone this repository to your local directory, and run main file using go command `go run main.go`, assuming you already setup Go environment. To avoid any platform environment differences, we can use docker instead to execute this program. First, we need to build the image with `docker build` command.

```sh
$ docker build -t <your-iamge-name> .
```
check list of all local image using `docker image ls` command, and you will see yout <your-image-name> created.
after that, run and access the newly created image using run command
```sh
$ docker run -it <your-created-image> 
```
if success, then you will on Terminal command line inside running container
```sh
/app # 
```

## Examples
### Help command 
you can check help command for explanation and its available parameter by adding command `-h` or `--help`
```sh
/app # ./fetch --help
``` 
this will shows output

[![JnkDbOF.md.png](https://iili.io/JnkDbOF.md.png)](https://freeimage.host/i/JnkDbOF)

### Fetch URLs and stores page contents. 
Only URLs with valid format can be fetched. 
command : 
```sh
/app # ./fetch https://www.mangasaki.com https://www.coldplay.com
```
above sample is fetched 2 urls and stored its content into 2 html file. 
[![JnkLB29.md.png](https://iili.io/JnkLB29.md.png)](https://freeimage.host/i/JnkLB29)
if one of URL in invalid, then fetch will not be processed. 

### Fetch page information Metadata
Fetch url metadata can be executed using `--metadata` or `-m` param. It will shows 4 metadata elements :
- Site URL
- num_links
- images
- time and date of last fetch data

command : 
```sh
/app # ./fetch --metadata https://www.mangasaki.com
```

as it shows output below :

[![JnkyczX.md.png](https://iili.io/JnkyczX.md.png)](https://freeimage.host/i/JnkyczX)
