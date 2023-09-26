when go programe first start, it start with a goroutine.
every go programe has on go rotuine
think like goroutine exist in our go programme.

start with an example. here is a simple code that do a api get call

```
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	links := []string{
		"https://www.google.com/",
	}

	for _, link := range links {
		CheckLink(link)
	}
}

func CheckLink(link string) {
	_, err := *http.Get(link)*

	if err != nil {
		fmt.Print(link, "  site is down \n")
		return
	}

	fmt.Println(link + "  site is up")
}
```

when we complied and run this code, there will a goroutine start, that will excute our code line by line, but when it do a api call in **http.Get(link)** which take some time to complete, it will block the other line to excute.

unitl this **http.Get(link)** finish, main goroutine is frozen.

Although we can run multiple go routine, but only one goroutine will one given time. behind the schene, go scheduler will handle this which goroutine will run in any given time.

be default goscheduler will use only one core. this is default behavior. but we can change that. let’s say we have 3 core and 3 go routine. so what go scheduler do is give each core one goroutine to run

when we say a programme can run concurrently, what we actually mean,
the programme has abitlity to load up multiple goroutine at a time,
our programme has ability to run different things kinda same time, (not actuallly (go scheduler ))

when we talk about paralleslism, we run our programme with multiple physical CPU, multi core,
with parallesism we can say we can do multiple thing exact same time

the problem with only `go` keyword that it created multiple goroutine but our main function don’t care about child goroutine, so it exit the programme when it don’t have any other code to excute.

so solve this bug we have channel
channel is like central messeg, where we can communication between goroutine only via channel

channel is typed, mean if we creat string type channel we can only send string typed message
channel is two way messenging device. like we can send and receive message
