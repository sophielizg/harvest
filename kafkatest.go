package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sophielizg/harvest/topic"
	"github.com/sophielizg/harvest/topic/kafka"
)

type Request struct {
	Method string
	Url    *url.URL
}

func (req *Request) Encode() ([]byte, error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (req *Request) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, req)
}

func (req *Request) convertToHttp() *http.Request {
	return &http.Request{
		Method: req.Method,
		URL:    req.Url,
	}
}

type handler struct{}

func (h *handler) HandleMessage(message *topic.ConsumerMessage[Request, *Request]) error {
	req := message.Value
	client := &http.Client{}
	res, err := client.Do(req.convertToHttp())
	if err != nil {
		return errors.Join(err, message.Error())
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		// bodyBytes, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	return errors.Join(err, message.Error())
		// }
		// log.Println(string(bodyBytes))
		log.Println("request")
	}

	return message.Success()
}

func main() {
	// brokersUrl := []string{"localhost:9093"}
	topicName := "requests"
	producer, producerErr := topic.NewProducer[*Request](
		kafka.WithProducerImplementation[*Request](
			kafka.Producer(kafka.WithBrokers("localhost:9093")),
			kafka.Producer(kafka.WithTopic(topicName)),
		),
	)
	consumer, consumerErr := topic.NewConsumer(
		kafka.WithConsumerImplementation[Request](
			kafka.Consumer(kafka.WithBrokers("localhost:9093")),
			kafka.Consumer(kafka.WithTopic(topicName)),
			kafka.WithGroup("test"),
		),
	)
	if err := errors.Join(producerErr, consumerErr); err != nil {
		log.Fatal(err)
	}

	reqTopic := topic.Topic[Request, *Request]{Producer: producer, Consumer: consumer}
	defer reqTopic.Close()

	errs := consumer.Start(&handler{})
	time.Sleep(4 * time.Second)

	// url, _ := url.Parse("https://google.com")
	// req := &Request{
	// 	Method: http.MethodGet,
	// 	Url:    url,
	// }

	// err = producer.SendMessages(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	select {
	case err := <-errs:
		log.Fatal(err)
	case err := <-consumer.Errors():
		log.Fatal(err)
	}
}
