package main

import (
	"fmt"
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	name := "Vinicius"
	expected := fmt.Sprintf("Hello %s", name)
	result := sayHello(name)

	if result != expected {
		t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, got -> %v", name, expected, result)
	}
}

func Test_SayGoodBye(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Bye Bye %s!", name)
	result := sayGoodBye(name)

	if result != expected {
		t.Errorf("\"sayGoodBye('%s')\" FAILED, expected -> %v and got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayGoodBye('%s')\" SUCCEDED, expected -> %v and got -> %v", name, expected, result)
	}
}
