package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		c, buf, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer c.Close()

		buf.WriteString("Hello from tcp!")
		buf.Flush()

		str, err := buf.ReadString('\n')
		if err != nil {
			log.Printf("err when reading string %v", err)
			return
		}

		fmt.Fprintf(buf, "You said %q\nByte\n", str)
		buf.Flush()
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
