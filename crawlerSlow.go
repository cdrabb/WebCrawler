package main

import 
(
	"fmt"
        "net/http"
	"code.google.com/p/go.net/html"
	"io"
	"runtime"
	"sync"
)

var wait sync.WaitGroup

func main() {
    
    wait.Add(2)
    runtime.GOMAXPROCS(1)
    fmt.Println("Starting Go Routines")
    go getPage("http://www.cpp.edu/~raheja/CS408/index.html")
    go getPage("http://www.cpp.edu/~masrinivas/cs331/")
    wait.Wait()
    fmt.Println("Done!")
}

func getPage(url string){
    defer wait.Done()
    resp, err := http.Get(url)
    if err != nil{
        fmt.Println("An Error Occurred")
        return 
    }
	
    //body, _ := ioutil.ReadAll(resp.Body)

    //fmt.Println(string(body))

    links := getLinks(resp.Body)  

    for _, link := range(links) { 
        fmt.Println(link)
    }
    resp.Body.Close()
}

func getLinks(httpBody io.Reader) []string {
  links := make([]string, 0)
  page := html.NewTokenizer(httpBody)
  for {
    tokenType := page.Next()
    if tokenType == html.ErrorToken {
      return links
    }
    token := page.Token()
    if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
        for _, attr := range token.Attr {
	    if attr.Key == "href" {
               links = append(links, attr.Val)
            }
        }
    }
  }
}