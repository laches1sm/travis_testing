package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSayHello(t *testing.T) {
	Convey(`When I have a sayHello method I should expect to get back Hello, World!`, t, func() {
		expected := `Hello, World!`
		actual := sayHello()
		So(actual, ShouldEqual, expected)
	})
}
