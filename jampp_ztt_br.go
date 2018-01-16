package main

import "encoding/xml"

type Products struct {
	XMLName  xml.Name  `xml:"PRODUCTS"`
	Products []Product `xml:"PRODUCT"`
}

type Product struct {
	XMLName         xml.Name `xml:"PRODUCT"`
	Id              string   `xml:"ID"`
	Title           string   `xml:"TITLE"`
	Description     string   `xml:"DESCRIPTION"`
	Price           string   `xml:"PRICE"`
	ImageLink       string   `xml:"IMAGE_LINK"`
	Link            string   `xml:"LINK"`
	Condition       string   `xml:"CONDITION"`
	IosUrl          string   `xml:"IOS_URL"`
	IphoneUrl       string   `xml:"IPHONE_URL"`
	IpadUrl         string   `xml:"IPAD_URL"`
	AndroidUrl      string   `xml:"ANDROID_URL"`
	WindowsPhoneUrl string   `xml:"WINDOWS_PHONE_URL"`
	Price2          string   `xml:"PRICE_"`
}
