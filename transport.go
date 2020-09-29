package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)
// Books
func makeCreateBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateBookRequest)
        msg, err := s.CreateBook(ctx, req.book)
        return CreateBookResponse{Msg: msg, Err: err}, nil
    }
}
func makePublishersFromBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookByIdRequest)
        bookDetails, err := s.GetPublishersFromBookById(ctx, req.Id)
        if err != nil {
            return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
    }
}
func makeAuthorsFromBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookByIdRequest)
        bookDetails, err := s.GetAuthorsFromBookById(ctx, req.Id)
        if err != nil {
            return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
    }
}
func makeGetBookByIdEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookByIdRequest)
        bookDetails, err := s.GetBookById(ctx, req.Id)
        if err != nil {
            return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
    }
}
func makeDeleteBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteBookRequest)
        msg, err := s.DeleteBook(ctx, req.Bookid)
        if err != nil {
            return DeleteBookResponse{Msg: msg, Err: err}, nil
        }
        return DeleteBookResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateBookendpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateBookRequest)
        msg, err := s.UpdateBook(ctx, req.book)
        return msg, err
    }
}

func decodeCreateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateBookRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetBookByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetBookByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetBookByIdRequest{
        Id: vars["bookid"],
    }
    return req, nil
}
func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteBookRequest
    vars := mux.Vars(r)
    req = DeleteBookRequest{
        Bookid: vars["bookid"],
    }
    return req, nil
}
func decodeUpdateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateBookRequest
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreateBookRequest struct {
        book Book
    }
    CreateBookResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetBookByIdRequest struct {
        Id string `json:"bookid"`
    }
    GetBookByIdResponse struct {
        Book interface{} `json:"book,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteBookRequest struct {
        Bookid string `json:"bookid"`
    }

    DeleteBookResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateBookRequest struct {
        book Book
    }
    UpdateBookResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)
// Publisher
func makeCreatePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreatePublisherRequest)
        msg, err := s.CreatePublisher(ctx, req.publisher)
        return CreatePublisherResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetBooksFromPublishersEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetPublisherByIdRequest)
        publisherDetails, err := s.GetBooksFromPublisherById(ctx, req.Id)
        if err != nil {
            return GetPublisherByIdResponse{Publisher: publisherDetails, Err: "Id not found"}, nil
        }
        return GetPublisherByIdResponse{Publisher: publisherDetails, Err: ""}, nil
    }
}
func makeGetPublisherByIdEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetPublisherByIdRequest)
        publisherDetails, err := s.GetPublisherById(ctx, req.Id)
        if err != nil {
            return GetPublisherByIdResponse{Publisher: publisherDetails, Err: "Id not found"}, nil
        }
        return GetPublisherByIdResponse{Publisher: publisherDetails, Err: ""}, nil
    }
}
func makeDeletePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeletePublisherRequest)
        msg, err := s.DeletePublisher(ctx, req.Publisherid)
        if err != nil {
            return DeletePublisherResponse{Msg: msg, Err: err}, nil
        }
        return DeletePublisherResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdatePublisherendpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdatePublisherRequest)
        msg, err := s.UpdatePublisher(ctx, req.publisher)
        return msg, err
    }
}

func decodeCreatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreatePublisherRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetPublisherByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetPublisherByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetPublisherByIdRequest{
        Id: vars["publisherid"],
    }
    return req, nil
}
func decodeDeletePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeletePublisherRequest
    vars := mux.Vars(r)
    req = DeletePublisherRequest{
        Publisherid: vars["publisherid"],
    }
    return req, nil
}
func decodeUpdatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdatePublisherRequest
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}


type (
    CreatePublisherRequest struct {
        publisher Publisher
    }
    CreatePublisherResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetPublisherByIdRequest struct {
        Id string `json:"publisherid"`
    }
    GetPublisherByIdResponse struct {
        Publisher interface{} `json:",omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeletePublisherRequest struct {
        Publisherid string `json:"publisherid"`
    }

    DeletePublisherResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdatePublisherRequest struct {
        publisher Publisher
    }
    UpdatePublisherResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)
// Author
func makeCreateAuthorEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateAuthorRequest)
        msg, err := s.CreateAuthor(ctx, req.author)
        return CreateAuthorResponse{Msg: msg, Err: err}, nil
    }
}

func makeGetBooksFromAuthorendpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetAuthorByIdRequest)
        authorDetails, err := s.GetBooksFromAuthorById(ctx, req.Id)
        if err != nil {
            return GetAuthorByIdResponse{Author: authorDetails, Err: "Id not found"}, nil
        }
        return GetAuthorByIdResponse{Author: authorDetails, Err: ""}, nil
    }
}

func makeGetAuthorByIdEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetAuthorByIdRequest)
        authorDetails, err := s.GetAuthorById(ctx, req.Id)
        if err != nil {
            return GetAuthorByIdResponse{Author: authorDetails, Err: "Id not found"}, nil
        }
        return GetAuthorByIdResponse{Author: authorDetails, Err: ""}, nil
    }
}
func makeDeleteAuthorEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteAuthorRequest)
        msg, err := s.DeleteAuthor(ctx, req.Authorid)
        if err != nil {
            return DeleteAuthorResponse{Msg: msg, Err: err}, nil
        }
        return DeleteAuthorResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateAuthorendpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateAuthorRequest)
        msg, err := s.UpdateAuthor(ctx, req.author)
        return msg, err
    }
}

func decodeCreateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateAuthorRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetAuthorByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetAuthorByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetAuthorByIdRequest{
        Id: vars["authorid"],
    }
    return req, nil
}
func decodeDeleteAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteAuthorRequest
    vars := mux.Vars(r)
    req = DeleteAuthorRequest{
        Authorid: vars["authorid"],
    }
    return req, nil
}
func decodeUpdateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateAuthorRequest
    if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
        return nil, err
    }
    return req, nil
}


type (
    CreateAuthorRequest struct {
        author Author
    }
    CreateAuthorResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetAuthorByIdRequest struct {
        Id string `json:"authorid"`
    }
    GetAuthorByIdResponse struct {
        Author interface{} `json:",omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteAuthorRequest struct {
        Authorid string `json:"authorid"`
    }

    DeleteAuthorResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateAuthorRequest struct {
        author Author
    }
    UpdateAuthorResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)
