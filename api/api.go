package api

import (
    "address-book/models"
    "address-book/contacts"
    "encoding/json"
    "fmt"
    "net/http"
)

func AddContact(addressBook models.AddressBook) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        var contact contacts.Contact
        if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        addedContact := addressBook.AddContact(&contact)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(addedContact)
    }
}

func ListContacts(addressBook models.AddressBook) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        contactList := addressBook.ListContacts()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(contactList)
    }
}

func RemoveContact(addressBook models.AddressBook) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        name := r.URL.Query().Get("name")
        if name == "" {
            http.Error(w, "Name is required", http.StatusBadRequest)
            return
        }

        if err := addressBook.RemoveContact(name); err != nil {
            http.Error(w, "Contact not found", http.StatusNotFound)
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Contact removed")
    }
}
