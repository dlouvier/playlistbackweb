package spotify

import (
	"fmt"
	"os"
)

func init() {
	c := ConfigParser()
	s := CalculateSecret(c.ClientID, c.ClientSecret)
	GetToken(c.URLGetToken, s)

	fmt.Println("Hola, has importado la libreria")
	fmt.Println(os.Getenv("TOKEN"))
}
