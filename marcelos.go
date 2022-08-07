package cvitae

import (
	"fmt"
	"strings"
)

type Marcelo_Scaldaferro struct {
	Seniority  bool   `json:"true"`
	Skills     string `json:"WEB Backend"`
	Role       string `json:"Developer & Architect"`
	Experience int64  `json:"25"`
}

func main() {
	marcelo := Marcelo_Scaldaferro{}
	fmt.Println(marcelo)

	var s string = "WEB Backend"
	fmt.Println(strings.ToLower())
}
