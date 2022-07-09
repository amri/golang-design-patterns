package main

import "fmt"

type IMessaging interface {
	imbueMessage(string)
	setMessage(string)
	sendMessage()
}

//"abstract" class with abstract method (from the interface) and concrete method (itself)
type messenger struct {
	data      string
	messenger IMessaging
}

func (m *messenger) prepareAndSendMessage(message string) error {
	//m.setMessage(message)
	m.messenger.imbueMessage(message)
	m.messenger.sendMessage()
	return nil
}

//specific implementor with embedded type
type sms struct {
	messenger
}
type email struct {
	messenger
}

func (s *sms) imbueMessage(message string) {
	s.data = s.data + message + ", from sms"
}

func (s *sms) sendMessage() {
	fmt.Println(s.data)
}

func (s *sms) setMessage(message string) {
	s.data = message
}

func (s *email) imbueMessage(message string) {
	s.data = s.data + message + ", from sms"
}

func (s *email) setMessage(message string) {
	s.data = message
}

func (s *email) sendMessage() {
	fmt.Println(s.data)
}

func main() {
	smsMessenger := new(sms)
	o1 := messenger{messenger: smsMessenger}

	emailMessenger := new(email)
	o2 := messenger{messenger: emailMessenger}
	err := o1.prepareAndSendMessage("something")
	err = o2.prepareAndSendMessage("somethingElse")
	if err != nil {
		return
	}

}
