package main

import (
    //"encoding/json"
    "context"
    "github.com/go-kit/kit/log"
    "fmt"
)
// Book struct-model
type Book struct {
    BookId    string `json:"bookId,omitempty"`
    Title     string `json:"title,omitempty"`
    Edition   string `json:"edition,omitempty"`
    Copyright string `json:"copyright,omitempty"`
    Language  string `json:"language,omitempty"`
    Pages     string `json:"pages,omitempty"`
    Author    string `json:"author,omitempty"`
    PublisherArray []Publisher `json:"publisherArray,omitempty"`
}
// Author struct-model
type Author struct {
  	AuthorId string `json:"authorId,omitempty"`
  	Name string `json:"name,omitempty"`
  	Nationality string `json:"nationality,omitempty"`
  	Birth string `json:"birth,omitempty"`
  	Genere string `json:"genere,omitempty"`
  	BookArray []Book `json:"bookArray,omitempty"`
}
// Publisher struct-model
type Publisher struct {
  	PublisherId string `json:"publisherId,omitempty"`
  	Name string `json:"name,omitempty"`
  	Country string `json:"country,omitempty"`
  	Founded string `json:"founded,omitempty"`
  	Genere string `json:"genere,omitempty"`
}


type bookservice struct {
    logger log.Logger
}
type publisherservice struct {
    logger log.Logger
}
type authorservice struct {
    logger log.Logger
}

// Service describes the Book service.
type BookService interface {
    CreateBook(ctx context.Context, book Book) (string, error)
    GetBookById(ctx context.Context, id string) (interface{}, error)
    UpdateBook(ctx context.Context, book Book) (string, error)
    DeleteBook(ctx context.Context, id string) (string, error)
    GetAuthorsFromBookById(ctx context.Context, id string) (interface{}, error)
    GetPublishersFromBookById(ctx context.Context, id string) (interface{}, error)
}
// Service describes the Author service.
type AuthorService interface {
    CreateAuthor(ctx context.Context, author Author) (string, error)
    GetAuthorById(ctx context.Context, id string) (interface{}, error)
    UpdateAuthor(ctx context.Context, author Author) (string, error)
    DeleteAuthor(ctx context.Context, id string) (string, error)
    GetBooksFromAuthorById(ctx context.Context, id string) (interface{}, error)
}
// Service describes the Publisher service.
type PublisherService interface {
    CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
    GetPublisherById(ctx context.Context, id string) (interface{}, error)
    UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
    DeletePublisher(ctx context.Context, id string) (string, error)
    GetBooksFromPublisherById(ctx context.Context, id string) (interface{}, error)
}

//author storage structure
var authors = []Author{
	Author{AuthorId:"Author1",Name: "Allen",Nationality: "Costarricense",Birth: "2020-08-28",Genere: "M",
		BookArray:[]Book{
				Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
					Copyright: "2012", Language: "ENGLISH", Pages: "976",
					PublisherArray:[]Publisher{
                                    Publisher{PublisherId:"Publisher1",Name:"Roberto Siluosa",Country:"Francia",Founded:"1999",Genere:"M"}}},
				Book{BookId: "Book2", Title: "Computer Networks", Edition: "5th",
					Copyright: "2010", Language: "ENGLISH", Pages: "960",
					PublisherArray:[]Publisher{Publisher{PublisherId:"Publisher2",Name:"Allen Fabioni",Country:"Italia",Founded:"2009",Genere:"M"},
																		 Publisher{PublisherId:"Publisher3",Name:"Juanito",Country:"Italia",Founded:"10",Genere:"M"}}},
		}},
}

//publisher storage structure
var publishers = []Publisher{
	Publisher{PublisherId:"Publisher1",Name:"Roberto Siluosa",Country:"Francia",Founded:"1999",Genere:"M"},
	Publisher{PublisherId:"Publisher2",Name:"Allen Fabioni",Country:"Italia",Founded:"2009",Genere:"M"},
	Publisher{PublisherId:"Publisher3",Name:"Juanito",Country:"Italia",Founded:"10",Genere:"M"},
}
//book storage structure
var books = []Book{
  Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
    Copyright: "2012", Language: "ENGLISH", Pages: "976",
    PublisherArray:[]Publisher{
                              Publisher{PublisherId:"Publisher1",Name:"Roberto Siluosa",Country:"Francia",Founded:"1999",Genere:"M"}}},
  Book{BookId: "Book2", Title: "Computer Networks", Edition: "5th",
    Copyright: "2010", Language: "ENGLISH", Pages: "960",
    PublisherArray:[]Publisher{Publisher{PublisherId:"Publisher2",Name:"Allen Fabioni",Country:"Italia",Founded:"2009",Genere:"M"},
                               Publisher{PublisherId:"Publisher3",Name:"Juanito",Country:"Italia",Founded:"10",Genere:"M"}}},
}
//find books
func find(x string) int {
    for i, book := range books {
        if x == book.BookId {
            return i
        }
    }
    return -1
}
//find authors
func findAuthor(x string) int {
	for i, author := range authors {
		if x == author.AuthorId {
			return i
		}
	}
	return -1
}
//find publishers
func findPublisher(x string) int {
	for i, publisher := range publishers {
		if x == publisher.PublisherId {
			return i
		}
	}
	return -1
}

func BookNewService(logger log.Logger) BookService {
    return &bookservice{
        logger: logger,
    }
}
func PublisherNewService(logger log.Logger) PublisherService {
    return &publisherservice{
        logger: logger,
    }
}
func AuthorNewService(logger log.Logger) AuthorService {
    return &authorservice{
        logger: logger,
    }
}

// Books CRUD
func (s bookservice) CreateBook(ctx context.Context, book Book) (string, error) {
  var msg = "success"
  publisherArray := book.PublisherArray
  for _,publisher := range publisherArray {
    publishers = append(publishers, publisher) // The Publisher is added by referential integrity (simulation)
  }
  books = append(books, book)
  fmt.Println("SI")
  return msg, nil
}
func (s bookservice) GetPublishersFromBookById(ctx context.Context, id string) (interface{}, error) {
  var err error
  var empty interface{}
  i := find(id)
  if i == -1 {
      return empty, err
  }else{
  	for _, book := range books { // loop over authors to find books to return
  		if id == book.BookId {
  			return book.PublisherArray,nil
  		}
  	}
  }
  return empty, err
}
func (s bookservice) GetAuthorsFromBookById(ctx context.Context, id string) (interface{}, error) {
  var err error
  var empty interface{}
  i := find(id)
  if i == -1 {
      return empty, err
  }else{
    for _, author := range authors { // loop over authors to find books to return
        for _, book := range author.BookArray {
          if id == book.BookId {
            authorUpdated := Author{
              AuthorId:   author.AuthorId,
              Name:        author.Name,
              Nationality: author.Nationality,
              Birth:       author.Birth,
              Genere:      author.Genere,
              BookArray:   []Book{},
            }
            return authorUpdated,nil
          }
        }
      }
  }
  return empty, err
}
func (s bookservice) GetBookById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var book interface{}
    var empty interface{}
    i := find(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
  for _, book := range books { // loop over books to find BookId to delete
    if id == book.BookId {
      for _,publisher := range book.PublisherArray { // we have to delete publishers then by referencial integrity(simulation)
        indexPublisher := findPublisher(publisher.PublisherId)
        newListPublishers := append(publishers[:indexPublisher], publishers[indexPublisher+1:]...)
        publishers = newListPublishers
      }
      indexBook := find(book.BookId) // we have to delete books then by referencial integrity(simulation)
      newListBooks := append(books[:indexBook], books[indexBook+1:]...)
      books = newListBooks
    }
  }
  for i, author := range authors { // loop to find theirs referential integrity (simulation) Authors
    for j, book := range author.BookArray {
      if book.BookId == id {
        author.BookArray = append(author.BookArray[:j],author.BookArray[j+1:]...)
        authors[i].BookArray = author.BookArray
      }
    }
  }
    msg := ""
    return msg, nil
}
func (s bookservice) UpdateBook(ctx context.Context, book Book) (string, error) {
  var empty = ""
  var err error
  bookId := find(book.BookId)
  if bookId == -1 {
      return empty, err
  }
  for _,publisherParam := range book.PublisherArray {
    for k,publisherActual := range publishers { // we have to update publishers then by referencial integrity(simulation)
      if publisherParam.PublisherId == publisherActual.PublisherId{
        publishers[k] = publisherParam
      }
    }
  }
  for _, author := range authors { // loop to find theirs referential integrity (simulation) Authors
    for g, bookActual := range author.BookArray {
      if book.BookId == bookActual.BookId{
        author.BookArray [g] = book
      }
    }
  }
  books[bookId] = book
  var msg = "success"
  return msg, nil
}

// Publisher CRUD
func (s publisherservice) CreatePublisher(ctx context.Context,  publisher Publisher) (string, error) {
  var msg = "success"
  publishers = append(publishers, publisher)
  return msg, nil
}

func (s publisherservice) GetBooksFromPublisherById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var empty interface{}
    i := findPublisher(id)
    if i == -1 {
        return empty, err
    }else{
      for _, book := range books {
    		for _, publisher := range book.PublisherArray {
    			if(publisher.PublisherId == id){
            bookUpdated := Book{
              BookId:   book.BookId,
              Title:    book.Title,
              Copyright:book.Copyright,
              Language: book.Language,
              Pages:    book.Pages,
              PublisherArray:   []Publisher{},
            }
            return bookUpdated,nil
    			}
    		}
      }
    }
    return empty, err
}
func (s publisherservice) GetPublisherById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var publisher interface{}
    var empty interface{}
    i := findPublisher(id)
    if i == -1 {
        return empty, err
    }
    publisher = publishers[i]
    return publisher, nil
}
func (s publisherservice) DeletePublisher(ctx context.Context, id string) (string, error) {
  for j, book := range books { // loop to find theirs referential integrity (simulation) Books
			for i, publisher := range book.PublisherArray {
				if id == publisher.PublisherId {
					book.PublisherArray = append(book.PublisherArray[:i], book.PublisherArray[i+1:]...)
					books[j].PublisherArray = book.PublisherArray
					break
				}
			}
		}
  for i, author := range authors { // loop to find theirs referential integrity (simulation) Authors
  		for j, book := range author.BookArray {
  			for k, publisher := range book.PublisherArray {
  				if id == publisher.PublisherId {
  					book.PublisherArray = append(book.PublisherArray[:k], book.PublisherArray[k+1:]...)
  					authors[i].BookArray[j].PublisherArray = book.PublisherArray
  					break
  				}
  			}
  		}
  }
  indexPublisher := findPublisher(id)
  newListPublishers := append(publishers[:indexPublisher], publishers[indexPublisher+1:]...)
  publishers = newListPublishers
  msg := ""
  return msg, nil
}
func (s publisherservice) UpdatePublisher(ctx context.Context, publisher Publisher) (string, error) {
  var empty = ""
  var err error
  publisherId := findPublisher(publisher.PublisherId)
  if publisherId == -1 {
      return empty, err
  }
  for _, book := range books { // loop to find theirs referential integrity (simulation) Books
      for i, publisherActual := range book.PublisherArray {
        if publisher.PublisherId == publisherActual.PublisherId {
          fmt.Println("UPDATED")
          book.PublisherArray[i] = publisher
        }
      }
  }
  for _, author := range authors { // loop to find theirs referential integrity (simulation) Authors
    for _, book := range author.BookArray {
      for i, publisherActual := range book.PublisherArray {
        if publisher.PublisherId == publisherActual.PublisherId {
          book.PublisherArray[i] = publisher
        }
      }
    }
  }
  publishers[publisherId] = publisher
  var msg = "success"
  return msg, nil
}
// Author CRUD
func (s authorservice) CreateAuthor(ctx context.Context,  author Author) (string, error) {
  var msg = "success"
	bookArray := author.BookArray
	for _,book := range bookArray {
			for _,publisher := range book.PublisherArray {
				publishers = append(publishers, publisher) // ingreso de las editoriales por llave foranea (simulación)
			}
			books = append(books, book) // ingreso de los libros por llave foranea (simulación)
	}
	authors = append(authors, author)
  return msg, nil
}

func (s authorservice) GetBooksFromAuthorById(ctx context.Context, id string) (interface{}, error) {
  var err error
  var empty interface{}
  i := findAuthor(id)
  if i == -1 {
      return empty, err
  }else{
    for _, author := range authors { // loop over authors to find books to return
      if id == author.AuthorId {
        return author.BookArray,nil
      }else{
        return empty, err
      }
    }
  }
  return empty, err
}

func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var author interface{}
    var empty interface{}
    i := findAuthor(id)
    if i == -1 {
        return empty, err
    }
    author = authors[i]
    return author, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
  for i, author := range authors { // loop over authors to find AuthordId to delete
    if id == author.AuthorId {
      bookArray := author.BookArray
      for _,book := range bookArray {
          for _,publisher := range book.PublisherArray { // we have to delete publishers then by referencial integrity(simulation)
            indexPublisher := findPublisher(publisher.PublisherId)
            newListPublishers := append(publishers[:indexPublisher], publishers[indexPublisher+1:]...)
            publishers = newListPublishers
          }
          indexBook := find(book.BookId) // we have to delete books then by referencial integrity(simulation)
          newListBooks := append(books[:indexBook], books[indexBook+1:]...)
          books = newListBooks
      }
      newListAuthors := append(authors[:i], authors[i+1:]...)
      authors = newListAuthors
    }
  }
  msg := ""
  return msg, nil
}
func (s authorservice) UpdateAuthor(ctx context.Context, author Author) (string, error) {
  var empty = ""
  var err error
  authorId := findAuthor(author.AuthorId)
  if authorId == -1 {
      return empty, err
  }
	bookArray := author.BookArray
	for _,bookParam := range bookArray {
		for _,publisherParam := range bookParam.PublisherArray {
			for k,publisherActual := range publishers { // we have to delete publishers then by referencial integrity(simulation)
				if publisherParam.PublisherId == publisherActual.PublisherId{
					publishers[k] = publisherParam
				}
			}
			for j,bookActual := range books { // we have to delete books then by referencial integrity(simulation)
				if bookParam.BookId == bookActual.BookId{
					books[j] = bookParam
				}
			}
		}
  }
	authors[authorId] = author
  var msg = "success"
  return msg, nil
}
