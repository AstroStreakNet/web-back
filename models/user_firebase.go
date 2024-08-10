package models

import "time"

type UserFirebase struct {
	ID          string    `firestore:"Document ID"`
	CreatedAt   time.Time `firestore:"created_at"`
	DisplayName string    `firestore:"display_name"`
	Email       string    `firestore:"email"`
	Password    string    `firestore:"password"`
	FirstName   string    `firestore:"first_name"`
	LastName    string    `firestore:"last_name"`
	Role        string    `firestore:"role"`
	Images      []string  `firestore:"images"`
}

func userFromFirebase(firebase *UserFirebase) *User {
	// convert id to uint

	// convert images to uint

	return &User{}
}
