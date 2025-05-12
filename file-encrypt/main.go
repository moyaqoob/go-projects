package main

import (
	"fmt"
	"os"
	"bytes"
	"golang.org/x/term"
	"file-encrypt/filecrypt"
)


func main(){
	if len(os.Args) < 2{
		printHelp()
	}

	
	function := os.Args[1]
	switch function{
	case "help":
		printHelp();
	case "encrypt":
		encryptHandle()

	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file and decrypt to decrypt a file")
		os.Exit(1)
	}
}


func printHelp() {
    helpMessage := `
Usage: [COMMAND] [OPTIONS]
Commands:
  encrypt   Encrypt a file or text.
            Usage: encrypt [input] [output]
            Example: encrypt myfile.txt encryptedfile.txt

  decrypt   Decrypt a file or text.
            Usage: decrypt [input] [output]
            Example: decrypt encryptedfile.txt decryptedfile.txt

  help      Display this help message.
            Usage: help
            Example: help

 Options:
  -k, --key    Specify a key for encryption or decryption.
               Example: go run . encrypt <fileName>

	Notes:
	- Ensure you use the same key for encryption and decryption.
	- Supported file formats: .txt, .json, etc.
	- For any issues, contact support@example.com.
`
    fmt.Println(helpMessage)
}


func encryptHandle(){
	if len(os.Args) < 3 {
		println("missing the path to file. For more info run go run . help")
	}
	file := os.Args[2]

	if !validateFile(file){
		panic("File not found")
	}
	password := getPassword()
	println("\nEncrypting")
	filecrypt.Encrypt(file,password)	
	fmt.Println("\n file successfully protected")
}



func decryptHandle(){
	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run go run . help")
	}

	file := os.Args[2]
	if !validateFile(file){
		panic("File not found")
	}

	password := getPassword()
	fmt.Println("decrypting, please wait ")
	filecrypt.Decrypt(file,password)
	fmt.Println("\n file decrypted successfully");

}

func getPassword()[]byte{

	println("Enter the password")
	password,_ := term.ReadPassword(0)
	fmt.Println("Re Enter the password");
	confirmPassword,_ := term.ReadPassword(0)

	if !validatePassword(password,confirmPassword) {
		fmt.Println("\n Passwords doesnt match")
		return getPassword()
	}
	return password
}

//validate the password and its checks
func validatePassword(password[] byte,confirmPassword []byte)bool{
	
	if !bytes.Equal(password,confirmPassword){
		return false
	}
	return true

}

//validate the file using inbuild os functions
func validateFile(file string)bool{
	if _,err := os.Stat(file);
	os.IsNotExist(err){
		return false
	}
	return true
}



