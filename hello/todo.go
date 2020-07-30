package main 
 
import(
	"net/http"
	"fmt"
     
)
//ここからしたのコメント　いらぬ

//    func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		fmt.Fprint(w, "GET hello!\n")
// 	case "POST":
// 		fmt.Fprint(w, "POST hello!\n")
// 	// ...省略
// 	default:
// 		fmt.Fprint(w, "Method not allowed.\n")
// 	}
// }

// 	http.HandleFunc("/todo",func(w http.ResponseWriter,r*http.Request){
// 		fmt.Fprint(w,"majimanji!!")
// 	})

// 	http.HandleFunc("/",func(w http.ResponseWriter,r*http.Request){
// 		fmt.Fprint(w,"Hello,World")
// 	})
// http.HandleFunc("/hoge",func(w http.ResponseWriter,r*http.Request){
// 	fmt.Fprint(w,"hoge")// })
// http.ListenAndServe(":8080",nil) 
　// ここから↑　いらぬ
func main (){ 
	func main() {
		http.HandleFunc("/todo", todo)
	
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
     var todolist []string;
	 func  todo(w http.ResponseWriter,r http.Request,t todolist){
	switch r.Method {
	case "GET":
	println(w,"majimanji")
	case "POST":
		var td string;
	  todolist=td;
	}
	}