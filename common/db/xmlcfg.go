package ggdb

import "encoding/xml"

type XmlMysqlConf struct {
	XMLName      xml.Name `xml:"mysql"`
	Dest         string   `xml:"dest,attr"`
	Username     string   `xml:"username"`
	Password     string   `xml:"password"`
	Host         string   `xml:"host"`
	Database     string   `xml:"database"`
	MaxOpenConns int      `xml:"MaxOpenConns"`
	MaxOdleConns int      `xml:"MaxOdleConns"`
}
