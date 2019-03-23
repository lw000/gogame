package tymail

import "encoding/xml"

type XmlMailConfig struct {
	XMLName xml.Name `xml:"mail"`
	Dest    string   `xml:"dest,attr"`
	Host    string   `xml:"host"`
	Port    string   `xml:"port"`
	Form    string   `xml:"form"`
	Pass    string   `xml:"pass"`
	To      string   `xml:"to"`
}
