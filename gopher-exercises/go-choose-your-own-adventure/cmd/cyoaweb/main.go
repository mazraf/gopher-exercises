package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/farzamalam/gopher-exercises/go-choose-your-own-adventure/cyoa"
)

// main is used to first intialize the flags then opens up the json file
// then decodes the json file into story type.
// Default handler is h, it can be called with argument story or with story and cyoa.WithTemplate to have
// custom template
func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with cyoa story")
	port := flag.Int("port", 3000, "The port to run our application on.")
	flag.Parse()
	fmt.Println("fileName : ", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Error while opening file : ", err)
		os.Exit(1)
	}
	story, err := cyoa.JsonDecode(f)
	if err != nil {
		fmt.Println("Error while decoding file : ", err)
		os.Exit(1)
	}
	fmt.Printf("Server has started on port :%d\n", *port)
	h := cyoa.NewHandler(story)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	//fmt.Printf("%+v\n", story)
}
