package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic("source file does not exist")
	}

	srcFile,err := os.Open(source)
	if err!=nil{
		panic(err.Error())
	}
    //read from the src file which was prev opened and store it in plaintext
	plainText,err:= io.ReadAll(srcFile)
	if err!=nil{
		panic(err.Error())
	}
	//store the pass in key
	key := password
	//create a nonce 
	nonce := make([]byte,12)

	if _,err:= io.ReadFull(rand.Reader,nonce); err!=nil{
		panic(err.Error())
	}
    //
	dk := pbkdf2.Key(key,nonce,4096,32,sha1.New)
	//creates a cipher block
	block,err := aes.NewCipher(dk)
	if err!=nil{
		panic(err.Error())
	}
	aesGCM,err:= cipher.NewGCM(block)
	if err!=nil{
		fmt.Println("Error creating cipher block:",err)
		return
	}

	ciphertext :=  aesGCM.Seal(nil,nonce,plainText,nil)
	ciphertext = append(ciphertext,nonce...)

	dstFile,err := os.Create(source)
	if err!=nil{
		panic(err.Error())
	}

	defer dstFile.Close()

	_,err = dstFile.Write(ciphertext)
	if err!=nil{
		panic(err.Error())
	}
}


func Decrypt(file string, password []byte) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic("file does not exist")
	}

	srcFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	ciphertext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	if len(ciphertext) < 12 {
		panic("ciphertext too short")
	}

	nonce := ciphertext[len(ciphertext)-12:]
	ciphertext = ciphertext[:len(ciphertext)-12]

	dk := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)
	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	dstFile, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(plainText)
	if err != nil {
		panic(err.Error())
	}
}


//how to encrypt
//read the file. then add the key to it,
// then block the file.seal the file
