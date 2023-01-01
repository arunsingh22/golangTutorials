package main

type (
	Db interface {
		Store(s something) error
	}

	Service struct {
		EmailSender interface {
			Send(email Email) error
		}
	}
)

//New – initializes our service object with the dependencies
func New(db *Db, emailSender *EmailSender) *Service {
	return &Service{
		Db:          db,
		EmailSender: emailSender,
	}
}

func (s *Service) doTheThing() error {
	//store on the db
	s.Db
	//send email
	s.EmailSender
}

func main() {

}
