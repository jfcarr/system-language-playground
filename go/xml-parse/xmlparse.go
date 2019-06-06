package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	var inputXML = []byte(`
	<?xml version="1.0"?>
	<dwml version="1.0">
	  <head>
		  <product srsName="WGS 1984" concise-name="time-series" operational-mode="official">
			  <title>NOAA's National Weather Service Forecast Data</title>
				<category>forecast</category>
			</product>
		</head>
	  <data>
	    <location lat="35" lon="-84">
	      <location-key>point1</location-key>
				<point latitude="39.44" longitude="-84.30"/>
	    </location>
			<parameters applicable-location="point1">
			  <temperature type="maximum" units="Fahrenheit" time-layout="k-p24h-n1-1">
				  <name>Daily Maximum Temperature</name>
				  <value>87</value>
				</temperature>
			  <temperature type="minimum" units="Fahrenheit" time-layout="k-p24h-n2-2">
				  <name>Daily Minimum Temperature</name>
				  <value>66</value>
        </temperature>
			</parameters>
	  </data>
	</dwml>`)

	type DWML struct {
		Head struct {
			Product struct {
				OperationalMode string `xml:"operational-mode,attr"`
				Title           string `xml:"title"`
				Category        string `xml:"category"`
			} `xml:"product"`
		} `xml:"head"`
		Data struct {
			Location struct {
				LocationKey string `xml:"location-key"`
				Point       struct {
					Latitude  string `xml:"latitude,attr"`
					Longitude string `xml:"longitude,attr"`
				} `xml:"point"`
			} `xml:"location"`
			Parameters struct {
				ApplicableLocation string `xml:"applicable-location,attr"`
				Temperatures       []struct {
					Name  string `xml:"name"`
					Value string `xml:"value"`
				} `xml:"temperature"`
			} `xml:"parameters"`
		} `xml:"data"`
	}

	var result DWML
	xml.Unmarshal(inputXML, &result)

	fmt.Printf("Operational Mode:    %s\n", result.Head.Product.OperationalMode)
	fmt.Printf("Title:               %s\n", result.Head.Product.Title)
	fmt.Printf("Category:            %s\n", result.Head.Product.Category)
	fmt.Printf("Location Key:        %s\n", result.Data.Location.LocationKey)
	fmt.Printf("Latitude:            %s\n", result.Data.Location.Point.Latitude)
	fmt.Printf("Longitude:           %s\n", result.Data.Location.Point.Longitude)
	fmt.Printf("Applicable Location: %s\n", result.Data.Parameters.ApplicableLocation)
	fmt.Printf("Temperature 1:       %s\n", result.Data.Parameters.Temperatures[0].Name)
	fmt.Printf("Temperature 1:       %s\n", result.Data.Parameters.Temperatures[0].Value)
	fmt.Printf("Temperature 2:       %s\n", result.Data.Parameters.Temperatures[1].Name)
	fmt.Printf("Temperature 2:       %s\n", result.Data.Parameters.Temperatures[1].Value)
}
