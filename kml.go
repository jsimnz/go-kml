package kml

import (
	//"fmt"
	"encoding/xml"
)

//Placemark template
type Placemark struct {
	Name		string	`xml:"name"`
	Description	string	`xml:"description"`
	Point		string	`xml:"Point>coordinates"`
}

//Kml template
type Kml struct {
	XMLName		xml.Name	`xml:"kml"`
	Namespace	string		`xml:xmlns,attr`
	Placemark	[]Placemark	
}

func NewKML(namespace string, numPlacemarks int) *Kml {
	//Initiate new kml layout
	kml := new(Kml)
	if namespace == "" {
		namespace = "http://www.opengis.net/kml/2.2"
	}
	kml.Namespace = namespace
	kml.Placemark = make([]Placemark, numPlacemarks)
	return kml
}

func (k *Kml) AddPlacemark(name string, desc string, point string) {
	placemark := Placemark{}
	placemark.Name 	= name
	placemark.Description = desc
	placemark.Point = point
	k.Placemark = append(k.Placemark, placemark)
}

func (k *Kml) Marshal() ([]byte, error){
	return xml.MarshalIndent(k, "", "   ")
}