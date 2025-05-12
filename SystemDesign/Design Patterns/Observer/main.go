package main

import (
	"fmt"
	"time"
)

// Idea: Subject/Item which is observed has atleast 2 things State and ListOfObservers
// State is somthing which the Subject holds ownership of in which other Objects are
// interested and ListOfObservers is a list of Objects which are interested

// Subject POV

type SubjectI interface {
	Subscribe(obj Person)
	UnSubscribe(obj Person)
	NotifiyAll()
}

type IPhoneItem struct {
	count           int
	ListOfObservers map[string]*Person
	// NotifiyAll func
}

func (s *IPhoneItem) Subscribe(obj *Person) {
	s.ListOfObservers[obj.Name] = obj
	fmt.Println("Subscriber count: ", len(s.ListOfObservers))
}

func (s *IPhoneItem) UnSubscribe(obj *Person) {
	delete(s.ListOfObservers, obj.Name)
	fmt.Println("Subscriber count: ", len(s.ListOfObservers))

}

func (s *IPhoneItem) NotifiyAll() error {
	for _, obj := range s.ListOfObservers {
		go func() {
			fmt.Printf("Sending Email to %s\n", obj.Name)
		}()
	}
	return nil
}

func (s *IPhoneItem) AddIphone(cnt int) {
	s.count += cnt
	fmt.Printf("More Iphones added to inventory %d\n", cnt)
	if s.count >= 10 {
		s.UpdateAvailablity()
	}
}

// Note: The Item itself must give expose some func like this to allow the client code
// to call for notification
func (s *IPhoneItem) UpdateAvailablity() {
	fmt.Println("Sending out notification to all subscribers ....")
	s.NotifiyAll()
}

// --------------------------------------------------
// Obserser POV

type Person struct {
	Name string
}

func main() {
	item := IPhoneItem{
		count:           0,
		ListOfObservers: make(map[string]*Person),
	}
	cust1 := &Person{Name: "arunsingh"}
	cust2 := &Person{Name: "Shilpa"}

	item.AddIphone(5)

	item.Subscribe(cust1)
	item.Subscribe(cust2)

	item.AddIphone(5)

	time.Sleep(1 * time.Second)

}
