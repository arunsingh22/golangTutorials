package main

import "testing"

func TestMain(t *testing.T) {
	t.Log("Hello world") // This only runs if the test fails
	// t.Fail() // This explicitly fails the test and continues => soft fail 
	t.Errorf("this test fails because we said so") // This explicitly fails the test and continues => soft fail 
	t.Fatalf("Hard stop failure..")

	
}
