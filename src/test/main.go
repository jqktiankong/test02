package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type WorldHandler struct {
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", p.ByName("name"))
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(writer, request)
	}
}

func log2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	r.ParseMultipartForm(1024)
	//fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	//fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	//fmt.Fprintln(w, "(3)", r.PostForm)
	//fmt.Fprintln(w, "(4)", r.MultipartForm)

	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	//if err == nil {
	//	data, err := ioutil.ReadAll(file)
	//	if err == nil {
	//		fmt.Fprintln(w, string(data))
	//	}
	//}

	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := "<html> <head><title>Go web Programming</title></head><body><h1>" +
		"hello world</h1></body></html>"
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302)
}

//type Post struct {
//	User    string
//	Threads []string
//}
//
//func jsonExample(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	post := &Post{
//		User:    "Sau Sheong",
//		Threads: []string{"first", "second", "third"},
//	}
//	json, _ := json.Marshal(post)
//	w.Write(json)
//}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:       "first_cookie",
		Value:      "Go Web Programming",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	c2 := http.Cookie{
		Name:       "second_cookie",
		Value:      "Manning Publications Co",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)

	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:       "flash",
		Value:      base64.URLEncoding.EncodeToString(msg),
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:       "flash",
			Value:      "",
			Path:       "",
			Domain:     "",
			Expires:    time.Unix(1, 0),
			RawExpires: "",
			MaxAge:     -1,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

var mPath = "D:/MyProgram/Go/github/"

//var mPath = "E:/go/project/"

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl.html")
	t.Execute(w, "Hello Workd!")

}

func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl2.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func process4(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl4.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	//daysOfWeek := []string{}
	t.Execute(w, daysOfWeek)
}

func process5(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl5.html")
	t.Execute(w, "hello")
}

func process6(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath+"test02/src/test/t1.html", mPath+"test02/src/test/t2.html")
	t.Execute(w, "Hello World!")
}

func formatDate(t time.Time) string {
	layout := "2020-6-21"
	return t.Format(layout)
}

func process7(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl7.html").Funcs(funcMap)
	t, _ = t.ParseFiles(mPath + "test02/src/test/tmpl7.html")
	t.Execute(w, time.Now())
}

func process8(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl8.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func process9(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/tmpl9.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(mPath + "test02/src/test/form.html")
	t.Execute(w, nil)
}

func process10(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles(mPath+"test02/src/test/layout.html", mPath+"test02/src/test/red_hello.html")
	} else {
		//t, _ = template.ParseFiles(mPath + "test02/src/test/layout.html", mPath + "test02/src/test/blue_hello.html")
		t, _ = template.ParseFiles(mPath + "test02/src/test/layout.html")
	}

	t.ExecuteTemplate(w, "layout", "")
}

//func main() {
//	//hello := HelloHandler{}
//	//world := WorldHandler{}
//
//	//mux:= httprouter.New()
//	//mux.GET("/hello/:name", hello)
//
//	server := http.Server{
//		Addr: "127.0.0.1:8080",
//	}
//
//	//http.Handle("/hello", &hello)
//	//http.Handle("/world", &world)
//
//	//http.HandleFunc("/hello", hello)
//	//http.HandleFunc("/world", world)
//
//	//http.HandleFunc("/hello", log(hello))
//	//http.Handle("/hello", protect(log2(&hello)))
//
//	//http.HandleFunc("/headers", headers)
//	//http.HandleFunc("/body", body)
//	//http.HandleFunc("/process", process)
//	//http.HandleFunc("/write", writeExample)
//	//http.HandleFunc("/writeheader", writeHeaderExample)
//	//http.HandleFunc("/redirect", headerExample)
//	//http.HandleFunc("/json", jsonExample)
//	//http.HandleFunc("/set_cookie", setCookie)
//	//http.HandleFunc("/get_cookie", getCookie)
//	//http.HandleFunc("/set_message", setMessage)
//	//http.HandleFunc("/show_message", showMessage)
//	http.HandleFunc("/process2", process2)
//	http.HandleFunc("/process3", process3)
//	http.HandleFunc("/process4", process4)
//	http.HandleFunc("/process5", process5)
//	http.HandleFunc("/process6", process6)
//	http.HandleFunc("/process7", process7)
//	http.HandleFunc("/process8", process8)
//	http.HandleFunc("/process9", process9)
//	http.HandleFunc("/form", form)
//	http.HandleFunc("/process10", process10)
//
//	server.ListenAndServe()
//}

//type Post struct {
//	Id      int
//	Content string
//	Author  string
//}
//
//var PostById map[int]*Post
//var PostsByAuthor map[string][]*Post
//
//func store(post Post) {
//	PostById[post.Id] = &post
//	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
//}
//
//func main() {
//	PostById = make(map[int]*Post)
//	PostsByAuthor = make(map[string][]*Post)
//
//	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
//	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
//	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
//	post4 := Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}
//
//	store(post1)
//	store(post2)
//	store(post3)
//	store(post4)
//
//	fmt.Println(PostById[1])
//	fmt.Println(PostById[2])
//
//	for _, post := range PostsByAuthor["Sau Sheong"] {
//		fmt.Println(post)
//	}
//
//	for _, post := range PostsByAuthor["Pedro"] {
//		fmt.Println(post)
//	}
//}

//func main()  {
//	data := []byte("Hello World!\n")
//	err := ioutil.WriteFile("data1", data, 0644)
//	if err != nil {
//		panic(err)
//	}
//
//	read1, _ := ioutil.ReadFile("data1")
//	fmt.Print(string(read1))
//
//	file1, _ := os.Create("data2")
//	defer file1.Close()
//
//	bytes, _ := file1.Write(data)
//	fmt.Printf("Wrote %d bytes to file\n", bytes)
//
//	file2, _ := os.Open("data2")
//	defer file2.Close()
//
//	read2 := make([]byte, len(data))
//	bytes, _ = file2.Read(read2)
//	fmt.Printf("Read %d bytes from file\n", bytes)
//	fmt.Println(string(read2))
//}

//type Post struct {
//	Id      int
//	Content string
//	Author  string
//}

//func store(data interface{}, filename string) {
//	buffer := new(bytes.Buffer)
//	encoder := gob.NewEncoder(buffer)
//	err := encoder.Encode(data)
//	if err != nil {
//		panic(err)
//	}
//	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func load(data interface{}, filename string) {
//	raw, err := ioutil.ReadFile(filename)
//	if err != nil {
//		panic(err)
//	}
//	buffer := bytes.NewBuffer(raw)
//	dec := gob.NewDecoder(buffer)
//	err = dec.Decode(data)
//	if err != nil {
//		panic(err)
//	}
//}

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "myuser:mypass@(127.0.0.1:3306)/gwp?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}

	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id ,content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2);"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("err = " + err.Error())
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	//csvFile, err := os.Create("posts.csv")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer csvFile.Close()
	//
	//allPosts := []Post{
	//	Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
	//	Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
	//	Post{Id: 3, Content: "Hola Mundo!", Author: "pedro"},
	//	Post{Id: 4, Content: "Greetings Earthings!", Author: "Sau Sheong"},
	//}
	//
	//writer := csv.NewWriter(csvFile)
	//for _, post := range allPosts {
	//	line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
	//	err := writer.Write(line)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//writer.Flush()
	//
	//file, err := os.Open("posts.csv")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//
	//reader := csv.NewReader(file)
	//reader.FieldsPerRecord = -1
	//record, err := reader.ReadAll()
	//if err != nil {
	//	panic(err)
	//}
	//
	//var posts []Post
	//for _, item := range record {
	//	id, _ := strconv.ParseInt(item[0], 0,0)
	//	post := Post{Id: int(id), Content: item[1], Author: item[2]}
	//	posts = append(posts, post)
	//}
	//fmt.Println(posts[0].Id)
	//fmt.Println(posts[0].Content)
	//fmt.Println(posts[0].Author)

	//post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	//store(post, "post1")
	//var postRead Post
	//load(&postRead, "post1")
	//fmt.Println(postRead)

	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	//readPost, _ := GetPost(post.Id)
	//fmt.Println(readPost)
	//
	//readPost.Content = "Bonjour Monde!"
	//readPost.Author = "Pierre"
	//readPost.Update()
	//
	//posts, _ := Posts(0)
	//fmt.Println(posts)
	//
	//readPost.Delete()
}
