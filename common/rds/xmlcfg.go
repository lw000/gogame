package ggrds

import "encoding/xml"

//<?xml version="1.0" encoding="utf-8"?>
//
//<config>
//	<redis dest="redis配置项">
//		<host>192.168.1.201:6379</host>
//		<psd>123456</psd>
//		<db>0</db>
//		<poolsize>20</poolsize>
//		<minIdConns>5</minIdConns>
//	</redis>
//</config>

type XmlRedisConf struct {
	XMLName    xml.Name `xml:"redis"`
	Dest       string   `xml:"dest,attr"`
	Host       string   `xml:"host"`
	Psd        string   `xml:"psd"`
	Db         int      `xml:"db"`
	Poolsize   int      `xml:"poolsize"`
	MinIdConns int      `xml:"minIdConns"`
}
