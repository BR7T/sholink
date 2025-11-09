package functions

import (
	"fmt"
	"github.com/BR7T/sholink/database"
)

func FindLinkByNano(NanoID string)string{
	fmt.Printf("\nComeçando a procurar o NanoID:\n%v \n" , NanoID)
	db := database.InitDB()
	var originalLink string

	db.QueryRow(`SELECT "OriginalLink" FROM "links" WHERE "NanoIDLink" = $1` , NanoID).Scan(&originalLink)

	fmt.Println(originalLink)
	return originalLink
}