package main

import (
	"fmt"
	"net/http"
	"log"
	// "encoding/json"
)

func main() {
	// 映射路由到函数，使用默认mux
	// http.HandleFunc("/welcome",welcome)
	// http.HandleFunc("/getProfile",getProfile)
	// 自定义mux方式
	mux := http.NewServeMux()

	v1Mux := http.NewServeMux()

	v1Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "v1 Profile")
	})

	v1Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "v1 Posts")
	})

	v2Mux := http.NewServeMux()

	v2Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "v2 Profile")
	})

	v2Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "v2 Posts")
	})

	// The http.StripPrefix function returns a handler that serves HTTP requests by removing the given prefix from the request URL's path and invoking the handler.
	mux.Handle("/v1/", http.StripPrefix("/v1", v1Mux))
	mux.Handle("/v2/", http.StripPrefix("/v2", v2Mux))

	// 映射静态资源路径
	// fs := http.FileServer(http.Dir("static"))
	// // http.Handle("/", fs)
	// mux.Handle("/",fs)

	// 增加日志中间件，和java的aop类似。
	loggedHandler := loggingMiddleware(mux)
	// 监听端口，提供自定义mux
	if err := http.ListenAndServe(":8080", loggedHandler); err != nil {
			log.Fatal(err)
	}
	// 监听端口，提供默认mux，nil
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 		log.Fatal(err)
	// }
}

func loggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					log.Printf("Got a %s request for: %v\n", r.Method, r.URL)
					handler.ServeHTTP(w, r)
					log.Printf("Handler finished processing request")
	})
}

// type profile struct {
// 	Name    string
// 	Hobbies []string
// }

// func getProfile(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 					http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
// 					return
// 	}

// 	profile := profile{
// 					Name:    "Yashish",
// 					Hobbies: []string{"sports", "programming"},
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(profile); err != nil {
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func welcome(w http.ResponseWriter,r *http.Request)  {
// 	fmt.Fprintln(w,"welcome my friends! from coral.")
// }