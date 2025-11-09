package services

import (
	"fmt"

	"github.com/matoous/go-nanoid"
)
func CreateNanoID()(string , error){
	fmt.Println("Criando NanoID")
	NanoID, err := gonanoid.Nanoid(13)
	if err != nil{
		return "" , err
	}
	return NanoID , nil
}