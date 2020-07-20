package main

//type Post struct {
//	Id       int       `json:"id"`
//	Content  string    `json:"content"`
//	Author   Author    `json:"author"`
//	Comments []Comment `json:"comments"`
//}
//
//type Author struct {
//	Id   int    `json:"id"`
//	Name string `json:"name"`
//}
//
//type Comment struct {
//	Id      int    `json:"id"`
//	Content string `json:"content"`
//	Author  string `json:"author"`
//}
//
//var mPath = "D:/MyProgram/Go/github/"

func main() {
	//jsonFile, err := os.Open(mPath + "test02/src/post.json")
	//if err != nil {
	//	fmt.Println("Error opening JSON file:", err)
	//	return
	//}
	//defer jsonFile.Close()
	//jsonData, err := ioutil.ReadAll(jsonFile)
	//if err != nil {
	//	fmt.Println("Error reading JSON data:", err)
	//	return
	//}
	//
	//var post Post
	//json.Unmarshal(jsonData, &post)
	//fmt.Println(post)

	//decoder := json.NewDecoder(jsonFile)
	//for {
	//	var post Post
	//	err := decoder.Decode(&post)
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error decoding JSON:", err)
	//		return
	//	}
	//	fmt.Println(post)
	//}

	//post := Post{
	//	Id:      1,
	//	Content: "Hello World!",
	//	Author: Author{
	//		Id:   2,
	//		Name: "Sau Sheong",
	//	},
	//	Comments: []Comment{
	//		{Id: 3,
	//			Content: "Have a great day!",
	//			Author:  "Adam",
	//		},
	//		{Id: 4,
	//			Content: "How are you today?",
	//			Author:  "Betty",
	//		},
	//	},
	//}

	//output, err := json.MarshalIndent(&post, "", "\t")
	//if err != nil {
	//	fmt.Println("Error marshallint to JSON:", err)
	//	return
	//}
	//err = ioutil.WriteFile("post.json", output, 0644)
	//if err != nil {
	//	fmt.Println("Error writing JSON to file:", err)
	//	return
	//}

	//jsonFile, err := os.Create("post.json")
	//if err != nil {
	//	fmt.Println("Error creating JSON file:", err)
	//	return
	//}
	//encoder := json.NewEncoder(jsonFile)
	//err = encoder.Encode(&post)
	//if err != nil {
	//	fmt.Println("Error encoding JSON to file:", err)
	//	return
	//}
}
