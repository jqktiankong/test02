package main

//type Post struct {
//	XMLName  xml.Name  `xml:"post"`
//	Id       string    `xml:"id,attr"`
//	Content  string    `xml:"content"`
//	Author   Author    `xml:"author"`
//	Xml      string    `xml:",innerxml"`
//	Comments []Comment `xml:"comments>comment"`
//}
//
//type Post02 struct {
//	XMLName xml.Name `xml:"post"`
//	Id      string   `xml:"id,attr"`
//	Content string   `xml:"content"`
//	Author  Author   `xml:"author"`
//	Xml     string   `xml:",innerxml"`
//}
//
//type Comment struct {
//	Id      string `xml:"id,attr"`
//	Content string `xml:"content"`
//	Author  Author `xml:"author"`
//}
//
//type Author struct {
//	Id   string `xml:"id,attr"`
//	Name string `xml:",chardata"`
//}
//
//var mPath = "D:/MyProgram/Go/github/"

func main() {
	//xmlFile, err := os.Open(mPath + "test02/src/post.xml")
	//if err != nil {
	//	fmt.Println("Error opening XML file:", err)
	//	return
	//}
	//defer xmlFile.Close();
	//xmlData, err := ioutil.ReadAll(xmlFile)
	//if err != nil {
	//	fmt.Println("Error reading XML data:", err)
	//	return
	//}
	//
	//var post Post
	//xml.Unmarshal(xmlData, &post)
	//fmt.Println(post)

	//xmlFile, err := os.Open(mPath + "test02/src/post.xml")
	//if err != nil {
	//	fmt.Println("Error opening XML file:", err)
	//	return
	//}
	//defer xmlFile.Close()
	//
	//decoder := xml.NewDecoder(xmlFile)
	//for {
	//	t, err := decoder.Token()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error decoding XML into tokens:", err)
	//		return
	//	}
	//
	//	switch se := t.(type) {
	//	case xml.StartElement:
	//		if se.Name.Local == "comment" {
	//			var comment Comment
	//			decoder.DecodeElement(&comment, &se)
	//			fmt.Println(comment)
	//		}
	//	}
	//}

	//post := Post02{
	//	Id:      "1",
	//	Content: "Hello World!",
	//	Author: Author{
	//		Id:   "2",
	//		Name: "Sau Sheong",
	//	},
	//}

	//output, err := xml.MarshalIndent(&post, "", "\t")
	//if err != nil {
	//	fmt.Println("Error marshalling to XML:", err)
	//	return
	//}
	//err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	//if err != nil {
	//	fmt.Println("Error writing XML to file:", err)
	//	return
	//}

	//xmlFile, err := os.Create("post.xml")
	//if err != nil {
	//	fmt.Println("Error creating XML file:", err)
	//	return
	//}
	//encoder := xml.NewEncoder(xmlFile)
	//encoder.Indent("", "\t")
	//err = encoder.Encode(&post)
	//if err != nil {
	//	fmt.Println("Error encoding XML to file:", err)
	//	return
	//}
}
