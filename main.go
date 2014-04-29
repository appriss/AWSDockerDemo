package main 

import (
	"net/http"
	"fmt"
)

var content = "<html><head><title>DockerDemo</title></head><body><h1>DockerDemo</h1></br></hr><p>This is a <a href='http://docker.io'>docker</a> demo application written in <a href='http://golang.org'>Googles Go language</a>, deployed to <a href='http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/create_deploy_docker.html'> AWS Beanstalk</a>!</p></body></html>"

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", content)
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}
}