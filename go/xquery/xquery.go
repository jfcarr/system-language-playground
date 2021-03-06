package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
<channel>
  <title>W3Schools Home Page</title>
  <link>https://www.w3schools.com</link>
  <description>Free web building tutorials</description>
  <item>
    <title>RSS Tutorial</title>
    <link>https://www.w3schools.com/xml/xml_rss.asp</link>
    <description>New RSS tutorial on W3Schools</description>
  </item>
  <item>
    <title>XML Tutorial</title>
    <link>https://www.w3schools.com/xml</link>
    <description>New XML tutorial on W3Schools</description>
  </item>
</channel>
</rss>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	channel := xmlquery.FindOne(doc, "//channel")
	fmt.Printf("title: %s\n", channel.SelectElement("title").InnerText())
	fmt.Printf("link: %s\n", channel.SelectElement("link").InnerText())
	for i, n := range xmlquery.Find(doc, "//item") {
		fmt.Printf("%d %s -> %s\n", i, n.SelectElement("title").InnerText(), n.SelectElement("link").InnerText())
	}
}
