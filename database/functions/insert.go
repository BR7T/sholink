package functions

import (
	"github.com/BR7T/sholink/database"
	"github.com/BR7T/sholink/services"
)

func InsertLink(inputURL string)(bool, string , error){
	//verifica link que está sendo inserido
	validURL , err := services.VerifyLink(inputURL)
	if err != nil || validURL == ""{
		return false ,"", err
	}

	// database
	db := database.InitDB()
	NanoIDLink , err := services.CreateNanoID()
	if err != nil{
		return false ,"", err
	}

	_ , err = db.Query(
		`INSERT INTO links ("NanoIDLink" , "OriginalLink") VALUES ($1 , $2)`,
		NanoIDLink , validURL,
	)

	if err != nil{
		return false ,"", err
	}

	return true,NanoIDLink, nil
}