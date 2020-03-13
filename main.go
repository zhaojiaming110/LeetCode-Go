package main
import (
    "fmt"
    "github.com/go-redis/redis/v7"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)  // handler 为向url发送请求时，调用的函数
    log.Fatal(http.ListenAndServe("localhost:1314",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    err = client.Set("key", "value1", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := client.Get("key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "Url.path = %q\n", r.URL.Path)
    fmt.Fprintf(w, "Hello, Docker!\n")
    fmt.Fprintf(w, "Hello, Rediss!\n")
    fmt.Fprintf(w, "val[%s]\n", val)
}
