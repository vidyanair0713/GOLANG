package main

import (
	"fmt"
	"net/http"
)

const page = `<!DOCTYPE html>
<html>
<body>
<h1> it got here</h1>

<form method = "post" action = "/formTest">
one <input name="one" label = "one"></input>
<br/>
two <input name="two" label = "two"></input>
<br/>
<input type="submit" value = "Submit this form"></input>

</body>
</html>
`

type httpReal struct {

}

const host = ""
const port = 8080


func main() {
 
	var myHandler *httpReal
	//var myServer = &http.Server{Addr: ":8080", Handler: myHandler}

	var myServer = &http.Server {Addr: fmt.Sprintf("%s:%d", host, port), Handler: myHandler}
	var err error
	err = myServer.ListenAndServe()
	if err != nil{
		fmt.Printf("error : %v\n", err)
	}

	fmt.Printf("if it got here we are in trouble...")
}

func (httpReal *httpReal) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var localPath string
	localPath = req.URL.Path

	fmt.Printf("path: %s \n", localPath)
	//fmt.Printf("Header: %v \n", req.Header)

	var one, two string

	one = req.FormValue("one")
	two = req.FormValue("two")

	fmt.Printf("one: %s \n", one)
	fmt.Printf("two: %s \n", two)


	defer req.Body.Close()
	var err error
	var writeCount int

	writeCount, err = res.Write([]byte(page))
	if err != nil {
		fmt.Printf("Error : %v \n", err)
		return
	}

	fmt.Printf("writeCount : %d \n", writeCount)



}

