package main

import (
    "address-book/api"
    "address-book/models"
    "log"
    "net/http"
)

func main() {
    addressBook := models.NewAddressBook()

    http.HandleFunc("/contact", api.AddContact(addressBook))          
    http.HandleFunc("/contacts", api.ListContacts(addressBook))      
    http.HandleFunc("/contact/remove", api.RemoveContact(addressBook))

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
