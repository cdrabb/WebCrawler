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
    
    wait.Add(15)
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println("Starting Go Routines")
    go getPage("http://www.cpp.edu/~raheja/CS408/index.html")
    go getPage("http://www.cpp.edu/~masrinivas/cs331/")
    go getPage("https://hotpads.com/")
    go getPage("https://soundcloud.com/dangertapes/sets/gameology")
    go getPage("https://www.youtube.com/")
    go getPage("http://shityoucanafford.com/")
    go getPage("https://orcahq.com/jobs?tags=Art")
    go getPage("http://www.kilobolt.com/collision-detection-basics")
    go getPage("http://www.amazon.com/")
    go getPage("http://store.steampowered.com/")
    go getPage("http://www.gamestop.com/")
    go getPage("http://www.nba.com/lakers/?tmd=1")
    go getPage("http://www.cpp.edu/")
    go getPage("http://www.dealzon.com/")
    go getPage("http://slickdeals.net/")
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