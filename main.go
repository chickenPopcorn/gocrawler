package main


//"bufio"
	//"fmt"
	//"golang.org/x/net/html/charset"
	//"golang.org/x/text/encoding"
	//"golang.org/x/text/transform"
	//"io"
	//"io/ioutil"
	//"net/http"
	//"regexp"


import (
"./engine"
"./zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}