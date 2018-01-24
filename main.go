package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

// const xmlFilePath = "xmlSamples/jampp_ztt_br.xml"
const xmlFilePath = "xmlSamples/usersSocial.xml"

func main() {

	fmt.Println("Opening file...")

	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	xmlDecoder := xml.NewDecoder(xmlFile)

	for {
		currentToken, _ := xmlDecoder.Token()
		if currentToken == nil {
			break
		}

		switch currentToken.(type) {

		case xml.StartElement:
			element := currentToken.(xml.StartElement)

			fmt.Printf("StartElement: %s\n", element.Name.Local)

		case xml.CharData:
			value := bytes.TrimSpace([]byte(currentToken.(xml.CharData)))

			if len(value) == 0 {
				continue
			}

			fmt.Printf("CharData: %s\n", value)
		}
	}
}
