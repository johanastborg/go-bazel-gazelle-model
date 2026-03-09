package phonebook

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

type PhoneBook struct {
	contacts []Contact
}

func NewPhoneBook() *PhoneBook {
	return &PhoneBook{
		contacts: []Contact{},
	}
}

func (pb *PhoneBook) AddContact(name, phone string) {
	pb.contacts = append(pb.contacts, Contact{Name: name, Phone: phone})
}

func (pb *PhoneBook) FindContact(name string) (*Contact, error) {
	for _, contact := range pb.contacts {
		if contact.Name == name {
			return &contact, nil
		}
	}
	return nil, fmt.Errorf("contact %s not found", name)
}

func (pb *PhoneBook) ListContacts() []Contact {
	return pb.contacts
}

func (pb *PhoneBook) DeleteContact(name string) error {
	for i, contact := range pb.contacts {
		if contact.Name == name {
			pb.contacts = append(pb.contacts[:i], pb.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact %s not found", name)
}
