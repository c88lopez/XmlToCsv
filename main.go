package main

import (
	"fmt"
	"log"
	"os"

	"github.com/clbanning/mxj"
)

// const xmlFilePath = "xmlSamples/jampp_ztt_br.xml"
const xmlFilePath = "https://s3.amazonaws.com/teste-products.dinda.com.br/Jampp.xml"

func main() {

	// response, err := http.Get(xmlFilePath)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer response.Body.Close()

	// mv, err := mxj.NewMapXmlReader(response.Body)

	sampleXMLReader, err := os.Open("./xmlSamples/usersSocial.xml")
	if err != nil {
		log.Panic(err)
	}
	defer sampleXMLReader.Close()

	mv, err := mxj.NewMapXmlReader(sampleXMLReader)

	fmt.Printf("mv: %s", mv)
	// xmlDecoder := xml.NewDecoder(response.Body)

	// for {
	// 	currentToken, _ := xmlDecoder.Token()
	// 	if currentToken == nil {
	// 		break
	// 	}

	// 	switch currentToken.(type) {

	// 	case xml.StartElement:
	// 		element := currentToken.(xml.StartElement)

	// 		fmt.Printf("StartElement: %s\n", element.Name.Local)

	// 	case xml.CharData:
	// 		value := bytes.TrimSpace([]byte(currentToken.(xml.CharData)))

	// 		if len(value) == 0 {
	// 			continue
	// 		}

	// 		fmt.Printf("CharData: %s\n", value)
	// 	}
	// }
}
