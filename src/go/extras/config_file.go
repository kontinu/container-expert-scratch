package extras

import (
	"net/http"
	"fmt"
	"os"
	"bufio"
	"log"

)

type Config struct {
	name string
	level  string
}

func ConfigRead(w http.ResponseWriter, r *http.Request){
	var  configFile string=os.Getenv("CONFIG_FILE")
	if configFile == ""{
		configFile = "./example.json"
	}
	file, err := os.Open(configFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	var text=""
	for _, eachline := range txtlines {
		fmt.Println(eachline)
		text+=eachline+"<br>"
	}

	var  res string="<div style='text-align: center;'> <h1> The Variable is:<br> </h1>"+ text + "<br> - Go</div>"
	fmt.Fprintf(w, res )
}
