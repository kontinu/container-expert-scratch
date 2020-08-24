package extras

import (
	"net/http"
	"fmt"
)


//? FRONT
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[200] healthy")
	fmt.Fprintf(w, "healthy")
}
