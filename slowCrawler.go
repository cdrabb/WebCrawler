package main

import 
(
	"fmt"
        "net/http"
	"code.google.com/p/go.net/html"
	"io"
	"sync"
)

var wait sync.WaitGroup

func main() {
    
    wait.Add(15)
    fmt.Println("Starting Go Routines")
    getPage("http://www.cpp.edu/~raheja/CS408/index.html")
    getPage("http://www.cpp.edu/~masrinivas/cs331/")
    getPage("https://hotpads.com/")
    getPage("https://soundcloud.com/dangertapes/sets/gameology")
    getPage("https://www.youtube.com/")
    getPage("http://shityoucanafford.com/")
    getPage("https://orcahq.com/jobs?tags=Art")
    getPage("http://www.kilobolt.com/collision-detection-basics")
    getPage("http://www.amazon.com/")
    getPage("http://store.steampowered.com/")
    getPage("http://www.gamestop.com/")
    getPage("http://www.nba.com/lakers/?tmd=1")
    getPage("http://www.cpp.edu/")
    getPage("http://www.dealzon.com/")
    getPage("http://slickdeals.net/")
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