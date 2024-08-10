package repositories

import "cloud.google.com/go/firestore"

type UserInFirebase struct {
	client *firestore.Client
}
