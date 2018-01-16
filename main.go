package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//const xmlFilePath = "xmlSamples/usersSocial.xml"
const xmlFilePath = "xmlSamples/jampp_ztt_br.xml"

func main() {

	fmt.Println("Opening file...")

	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	fmt.Println("Reading all...")
	byteValue, _ := ioutil.ReadAll(xmlFile)

	fmt.Println("Unmarshal file...")
	var products Products
	xml.Unmarshal(byteValue, &products)

	productsLen := len(products.Products)
	fmt.Printf("productsLen %d\n", productsLen)

	w := csv.NewWriter(os.Stdout)
	records := [][]string{
		{"id", "title", "description"},
	}

	fmt.Println("Printing...")
	for i := 0; i < productsLen; i++ {
		records = append(records, []string{
			products.Products[i].Id, products.Products[i].Title,
			products.Products[i].Description})
	}

	w.WriteAll(records)
}
