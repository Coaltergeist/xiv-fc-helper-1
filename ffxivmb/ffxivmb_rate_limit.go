package ffxivmb

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
	elements      []*MBRequest
	executed      int
	executedTotal int
}

// MBRequest represents a request to ffxivmb
type MBRequest struct {
	url  string
	data chan *http.Response
}

func (q *queue) push(r *MBRequest) {
	q.elements = append(q.elements, r)
}

func (q *queue) pop() (r *MBRequest) {
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

func (r *MBRequest) queue() *MBRequest {
	requestQueue.push(r)
	return r
}

func (r *MBRequest) execute() {
	defer func() {
		if r := recover(); r != nil {
			l.Println("recovered in ffxivmb request Execute")
		}
	}()
	resp, err := http.Get(r.url)
	if err != nil {
		panic(err)
	}
	r.data <- resp
	close(r.data)
}

// NewMBRequest creates a request
func NewMBRequest() *MBRequest {
	return &MBRequest{
		data: make(chan *http.Response),
	}
}

// SetURL sets the URL of the request
func (r *MBRequest) SetURL(url string) {
	r.url = url
}

// Consume is a blocking operation that attempts to consume the request's response
func (r *MBRequest) Consume() *http.Response {
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
		elements: make([]*MBRequest, 0),
	}
	go processQueue()
}
