package main


import(
	"fmt"
	"flag"
	"net/http"
	"log"
)

// Serve the flat files!
func fileServer(port int) {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	stringPort := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(stringPort , nil)
	if err != nil {
		log.Fatal(err)
	}


}


func main(){
	portPointer := flag.Int("p", 8000, "Port must be a valid integer!")
	flag.Parse()
	fmt.Println("Starting Server on:", *portPointer)

	fileServer(*portPointer)
}