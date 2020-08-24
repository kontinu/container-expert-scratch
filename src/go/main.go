package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	extras "./extras"
)


func root(w http.ResponseWriter, r *http.Request) {
	var kontinu_msg=os.Getenv("MSG")
	if len(kontinu_msg) <= 0{
		kontinu_msg="Bootcamp Experts 🤓!"
	}
	fmt.Println(r.URL.RawQuery)
	//msg,host := extras.KontinuEndpoint()
	//w.Header().Add("HostChain", host)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

%s
- Go
`,kontinu_msg )
}

func main() {

	fmt.Println("====== Starting GO Server ======")
	//http.HandleFunc("/env", extras.GetValue)
	//http.HandleFunc("/config", extras.ConfigRead)
	http.HandleFunc("/health",extras.Health)
	http.HandleFunc("/api", extras.KontinuResponse)
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
