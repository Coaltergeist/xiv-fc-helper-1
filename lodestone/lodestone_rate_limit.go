package lodestone

import (
	"log"
	"net/http"
	"os"
	"time"
)

const rate = time.Second / 2        // 2 requests per second
const reportRate = 20 * time.Second // 20 seconds

var (
	l            *log.Logger
	requestQueue *queue
)

type queue struct {
	elements      []*LodestoneRequest
	executed      int
	executedTotal int
}

type LodestoneRequest struct {
	url  string
	data chan *http.Response
}

func (q *queue) push(r *LodestoneRequest) {
	q.elements = append(q.elements, r)
}

func (q *queue) pop() (r *LodestoneRequest) {
	if len(q.elements) == 0 {
		return nil
	}
	r = q.elements[0]
	q.elements = q.elements[1:]
	return
}

func (q *queue) len() int {
	return len(q.elements)
}

func (r *LodestoneRequest) queue() *LodestoneRequest {
	requestQueue.push(r)
	return r
}

func (r *LodestoneRequest) execute() {
	defer func() {
		if r := recover(); r != nil {
			l.Println("recovered in lodestone request Execute")
		}
	}()
	resp, err := http.Get(r.url)
	if err != nil {
		panic(err)
	}
	r.data <- resp
	close(r.data)
}

// NewLodestoneRequest creates a request
func NewLodestoneRequest() *LodestoneRequest {
	return &LodestoneRequest{
		data: make(chan *http.Response),
	}
}

// SetURL sets the URL of the request
func (r *LodestoneRequest) SetURL(url string) {
	r.url = url
}

// Consume is a blocking operation that attempts to consume the request's response
func (r *LodestoneRequest) Consume() *http.Response {
	return <-r.data
}

func processQueue() {
	requestLimit := time.NewTicker(rate)
	go func() {
		for range requestLimit.C {
			if requestQueue.len() > 0 {
				req := requestQueue.pop()
				requestQueue.executed++
				requestQueue.executedTotal++
				go req.execute()
			}
		}
	}()

	statsLimit := time.NewTicker(reportRate)
	go func() {
		for range statsLimit.C {
			l.Printf("Handled %d requests | %.2f per second | %d total\n", requestQueue.executed, float32(requestQueue.executed/5.0), requestQueue.executedTotal)
			requestQueue.executed = 0
		}
	}()
}

func init() {
	l = log.New(os.Stderr, "lodestone: ", log.LstdFlags|log.Lshortfile)
	l.Println("Initializing Lodestone request queue...")
	requestQueue = &queue{
		elements: make([]*LodestoneRequest, 0),
	}
	go processQueue()
}
