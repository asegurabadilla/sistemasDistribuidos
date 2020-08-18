package main

import (
    "encoding/json"
    "net/http"
    "path"
  //  "fmt"
  //  "reflect"
    "strconv" // convert from string to int
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := find(id)
    if i == -1 {
      getAllData, e := json.Marshal(books[1:])
      w.Header().Set("Content-Type", "application/json")
      w.Write(getAllData)
      return e
    }
    dataJson, err := json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json")
    w.Write(dataJson)
    return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    bookId,_ := strconv.Atoi(id)//id del book a modificar en books
    checkError("Parse error", err)
    for _, bookToModify := range books {
        if id == bookToModify.Id {
          err := json.NewDecoder(r.Body).Decode(&bookToModify)
          bookModified, err := json.Marshal(bookToModify)
          w.Header().Set("Content-Type", "application/json")
          w.Write(bookModified)
          if err != nil{
            panic(err)
          }
          books[bookId] = bookToModify
        }
    }
    return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := find(id)
    copy(books[i:], books[i+1:])
    books[len(books)-1] = Book{}
    books = books[:len(books)-1]
    listAll, err := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json")
    w.Write(listAll)
    return
}
