package helpler

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


//加密算法
func HashAndSalt(pwd string) string {
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//解密算法
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwds := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwds)
	if err != nil {
		return false
	}
	return true
}
//jwt 生成

