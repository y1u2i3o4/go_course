package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":80", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "Не удалось прочитать тело сообщения %q\n", err)
	}

	w.Write(body)
}

