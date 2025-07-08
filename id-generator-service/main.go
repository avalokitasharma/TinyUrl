package idgeneratorservice

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var counter uint64

func main() {
	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		val := atomic.AddUint64(&counter, 1)
		fmt.Fprintf(w, "%d", val)
	})
	fmt.Println("Global counter running on port:7001")
	http.ListenAndServe(":7001", nil)
}
