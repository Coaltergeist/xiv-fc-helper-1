package xivdb

import (
	"log"
	"os"
	"time"
)

const rate = time.Second / 15 // 15 requests per second
const reportRate = 5 * time.Second

var (
	l            *log.Logger
	requestQueue *queue
)

type queue struct {
	elements      []*Request
	executed      int
	executedTotal int
}

func (q *queue) push(r *Request) {
	q.elements = append(q.elements, r)
}

func (q *queue) pop() (r *Request) {
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
	l = log.New(os.Stderr, "xivdb: ", log.LstdFlags|log.Lshortfile)
	l.Println("Initializing XIVDB request queue...")
	requestQueue = &queue{
		elements: make([]*Request, 0),
	}
	go processQueue()
}
