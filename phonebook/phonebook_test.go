package phonebook

import (
	"testing"
)

func TestFindContact(t *testing.T) {
	pb := NewPhoneBook()
	pb.AddContact("Alice", "123-456")
	pb.AddContact("Bob", "789-012")

	t.Run("found", func(t *testing.T) {
		contact, err := pb.FindContact("Alice")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if contact.Name != "Alice" {
			t.Errorf("expected Alice, got %v", contact.Name)
		}
		if contact.Phone != "123-456" {
			t.Errorf("expected 123-456, got %v", contact.Phone)
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, err := pb.FindContact("Charlie")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		expectedErr := "contact Charlie not found"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})
}

func TestListContacts(t *testing.T) {
	pb := NewPhoneBook()
	pb.AddContact("Alice", "123-456")
	pb.AddContact("Bob", "789-012")

	contacts := pb.ListContacts()
	if len(contacts) != 2 {
		t.Fatalf("expected 2 contacts, got %d", len(contacts))
	}

	if contacts[0].Name != "Alice" {
		t.Errorf("expected Alice, got %v", contacts[0].Name)
	}
	if contacts[1].Name != "Bob" {
		t.Errorf("expected Bob, got %v", contacts[1].Name)
	}
}

func TestDeleteContact(t *testing.T) {
	pb := NewPhoneBook()
	pb.AddContact("Alice", "123-456")
	pb.AddContact("Bob", "789-012")

	t.Run("delete successful", func(t *testing.T) {
		err := pb.DeleteContact("Alice")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(pb.ListContacts()) != 1 {
			t.Errorf("expected 1 contact, got %d", len(pb.ListContacts()))
		}
	})

	t.Run("delete non-existent", func(t *testing.T) {
		err := pb.DeleteContact("Charlie")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
