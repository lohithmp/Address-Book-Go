package models

import (
    "address-book/contacts"
    "errors"
)

type AddressBook interface {
    AddContact(contact *contacts.Contact) contacts.Contact
    ListContacts() []contacts.Contact
    RemoveContact(name string) error
}

type addressBookImpl struct {
    contacts []contacts.Contact
}

func NewAddressBook() AddressBook {
    return &addressBookImpl{}
}

func (ab *addressBookImpl) AddContact(contact *contacts.Contact) contacts.Contact {
    ab.contacts = append(ab.contacts, *contact)
    return *contact
}

func (ab *addressBookImpl) ListContacts() []contacts.Contact {
    return ab.contacts
}

func (ab *addressBookImpl) RemoveContact(name string) error {
    for i, contact := range ab.contacts {
        if contact.Name == name {
            ab.contacts = append(ab.contacts[:i], ab.contacts[i+1:]...)
            return nil
        }
    }
    return errors.New("contact not found")
}
