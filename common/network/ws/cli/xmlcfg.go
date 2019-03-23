package tyws

import "encoding/xml"

type XmlWsConf struct {
	XMLName xml.Name `xml:"ws"`
	Dest    string   `xml:"dest,attr"`
	Host    string   `xml:"host"`
	Path    string   `xml:"path"`
}
