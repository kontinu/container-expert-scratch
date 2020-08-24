package extras

import (
	"encoding/json"
	"net/http"
	"os"
)

type Response struct {
	Msg    	string`json:"msg"`
	Host 	string`json:"host"`
  }

func KontinuResponse(w http.ResponseWriter, r *http.Request) {
	var kontinu_msg=os.Getenv("MSG")
	if len(kontinu_msg) <= 0{
		kontinu_msg="Bootcamp Experts 🤓!"
	}
	var host,_=os.Hostname()
	profile := Response{kontinu_msg, host }

	js, err := json.Marshal(profile)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
