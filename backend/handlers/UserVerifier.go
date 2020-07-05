package handlers

import "mrnakumar.com/tamaatar/storage"

type UserVerifier interface {
	CheckCredentials(uname string, passwd string) bool
}
type userVerifierImpl struct {
	userDb storage.UserStorage
}

func CreateUserVerifier(userDb storage.UserStorage) UserVerifier {
	return userVerifierImpl{userDb: userDb}
}

func (uv userVerifierImpl) CheckCredentials(name string, password string) bool {
	return uv.userDb.Exists(name, password)
}
