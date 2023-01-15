package helloworld_test

// import (
// 	"testing"

// 	"github.com/sophielizg/harvest/api/pkg/helloworld"
// 	"github.com/sophielizg/harvest/api/tests/mocks"
// )

// type SayHelloTest struct {
// 	name     string
// 	expected interface{}
// }

// func TestSayHello(t *testing.T) {
// 	sayHelloTests := []SayHelloTest{
// 		{"test", helloworld.Success{Hello: "test"}},
// 		{"", helloworld.Failure{Good: "bye"}},
// 	}

// 	for _, tc := range sayHelloTests {
// 		actual := helloworld.SayHello(mocks.MockApp(), tc.name)

// 		if actual != tc.expected {
// 			t.Errorf("SayHello(MockApp, %s): expected %+v, actual %+v", tc.name, tc.expected, actual)
// 		}
// 	}
// }
