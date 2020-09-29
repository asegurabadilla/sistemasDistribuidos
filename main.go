package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()
//Book services
    var svc BookService
    svc = BookNewService(logger)
//Publisher services
    var svcP PublisherService
    svcP = PublisherNewService(logger)
//Publisher services
    var svcA AuthorService
    svcA = AuthorNewService(logger)

//Book microservice
    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(svc),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(svc),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(svc),
        decodeUpdateBookRequest,
        encodeResponse,
    )
    GetAuthorsFromBookHandler := httptransport.NewServer(
        makeAuthorsFromBookEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    GetPublishersFromBookHandler := httptransport.NewServer(
        makePublishersFromBookEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    http.Handle("/", r)
    http.Handle("/book", CreateBookHandler)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")
    r.Handle("/books/{bookid}/authors", GetAuthorsFromBookHandler).Methods("GET")
    r.Handle("/books/{bookid}/publishers", GetPublishersFromBookHandler).Methods("GET")
//Publisher microservice
    CreatePublisherHandler := httptransport.NewServer(
        makeCreatePublisherEndpoint(svcP),
        decodeCreatePublisherRequest,
        encodeResponse,
    )
    GetByPublisherIdHandler := httptransport.NewServer(
        makeGetPublisherByIdEndpoint(svcP),
        decodeGetPublisherByIdRequest,
        encodeResponse,
    )
    DeletePublisherHandler := httptransport.NewServer(
        makeDeletePublisherEndpoint(svcP),
        decodeDeletePublisherRequest,
        encodeResponse,
    )
    UpdatePublisherHandler := httptransport.NewServer(
        makeUpdatePublisherendpoint(svcP),
        decodeUpdatePublisherRequest,
        encodeResponse,
    )
    GetBooksFromPublisherHandler := httptransport.NewServer(
        makeGetBooksFromPublishersEndpoint(svcP),
        decodeGetPublisherByIdRequest,
        encodeResponse,
    )
    http.Handle("/publisher", CreatePublisherHandler)
    http.Handle("/publisher/update", UpdatePublisherHandler)
    r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")
    r.Handle("/publishers/{publisherid}/books", GetBooksFromPublisherHandler).Methods("GET")
//Author microservice
    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(svcA),
        decodeCreateAuthorRequest,
        encodeResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(svcA),
        decodeGetAuthorByIdRequest,
        encodeResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(svcA),
        decodeDeleteAuthorRequest,
        encodeResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(svcA),
        decodeUpdateAuthorRequest,
        encodeResponse,
    )
    GetBooksfromAuthorHandler := httptransport.NewServer(
        makeGetBooksFromAuthorendpoint(svcA),
        decodeGetAuthorByIdRequest,
        encodeResponse,
    )
    http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")
    r.Handle("/authors/{authorid}/books", GetBooksfromAuthorHandler).Methods("GET")

    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
