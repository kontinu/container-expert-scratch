package extras

import (
	"net/http"
	"fmt"
	"os"
)


func GetValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	var  the_variable string=os.Getenv("THE_VARIABLE")
	var  res string="<div style='text-align: center;'> <h1> The Variable is:<br> " + the_variable +"</h1> - Go</div>"
	fmt.Fprintf(w, res )
}
