package request

import (
	"container/heap"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/sophielizg/harvest/queue"
)

type DelayCache struct {
	requestQueue queue.Queue
	sender       *sender
	messages     delayHeap
	sync.Mutex
}

func NewDelayCache(requestQueue queue.Queue, interval time.Duration) *DelayCache {
	cache := &DelayCache{
		requestQueue: requestQueue,
	}
	if interval > 0 {
		runSender(cache, interval)
		runtime.SetFinalizer(cache, stopSender)
	}
	return cache
}

func (c *DelayCache) SendMessages(messages []*requestMessage) {
	c.Lock()
	defer c.Unlock()

	for _, message := range messages {
		delaySeconds := rand.Float32() * (message.req.RandomDelaySeconds.Max - message.req.RandomDelaySeconds.Min)
		delayDuration := time.Duration(delaySeconds * float32(time.Second))
		heap.Push(&c.messages, &delayedMessage{
			sendAfter: time.Now().Add(delayDuration),
			message:   message,
		})
	}
}

type delayedMessage struct {
	sendAfter time.Time
	message   *requestMessage
}

type delayHeap []*delayedMessage

func (h delayHeap) Len() int {
	return len(h)
}

func (h delayHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h delayHeap) Less(i, j int) bool {
	return h[i].sendAfter.Before(h[j].sendAfter)
}

func (h *delayHeap) Push(x interface{}) {
	item := x.(*delayedMessage)
	*h = append(*h, item)
}

func (h *delayHeap) Pop() interface{} {
	old, n := *h, len(*h)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type sender struct {
	interval time.Duration
	stop     chan bool
}

func runSender(cache *DelayCache, interval time.Duration) {
	s := &sender{
		interval: interval,
		stop:     make(chan bool),
	}
	s.interval = interval
	cache.sender = s
	go s.Run(cache)
}

func stopSender(cache *DelayCache) {
	cache.sender.stop <- true
}

func (s *sender) Run(cache *DelayCache) {
	ticker := time.NewTicker(s.interval)
	for {
		select {
		case <-ticker.C:
			sendReadyMessages(cache)
		case <-s.stop:
			ticker.Stop()
			return
		}
	}
}

func sendReadyMessages(cache *DelayCache) error {
	cache.Lock()
	defer cache.Unlock()

	if len(cache.messages) == 0 {
		return nil
	}

	readyMessages := []*requestMessage{}
	for {
		if len(cache.messages) > 0 && cache.messages[0].sendAfter.Before(time.Now()) {
			delayMessage := heap.Pop(&cache.messages).(delayedMessage)
			readyMessages = append(readyMessages, delayMessage.message)
		}
	}

	if len(readyMessages) > 0 {
		return cache.requestQueue.SendMessages(readyMessages)
	} else {
		return nil
	}
}
