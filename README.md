---

# A CRUD API With Golang

This tutorial demonstrates how to build a CRUD (Create, Read, Update, Delete) API using Golang. The API allows users to perform basic operations on a collection of movies without using a database. We'll be using structs and slices to store data in memory.

## Prerequisites

Before you begin, make sure you have the following installed:

- Go programming language
- Postman (for testing the API)

## Setup

1. Clone the repository to your local machine:

```bash
git clone <repository-url>
```

2. Navigate to the project directory:

```bash
cd A-CRUD-API-With-Golang
```

3. Build and run the application:

```bash
go build
go run main.go
```

The server will start running on port 8000.

## Libraries Used

- `github.com/gorilla/mux`: A powerful HTTP router and URL matcher for building Go web servers.

## Project Structure

- `main.go`: Contains the main code for the CRUD API.
- `README.md`: The documentation you are currently reading.

## Understanding the Code

### Package Imports

```go
import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
)
```

Explanation:
- `fmt`: Provides formatted I/O functions for printing messages.
- `log`: Logging package to log errors and other messages.
- `math/rand`: Used for generating random numbers.
- `net/http`: Used to create an HTTP server.
- `strconv`: Package for string conversions.
- `encoding/json`: Used for encoding and decoding JSON data.
- `github.com/gorilla/mux`: HTTP router for handling routes and requests.

### Structs

#### Movie Struct

```go
type Movie struct {
    ID       string   `json:"id"`
    ISBN     int      `json:"isbn"`
    Title    string   `json:"title"`
    Director *Director `json:"director"`
}
```

Explanation:
- Defines the structure of a movie.
- Contains fields for ID, ISBN, Title, and Director.
- Utilizes JSON tags for marshaling and unmarshaling JSON data.

#### Director Struct

```go
type Director struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}
```

Explanation:
- Defines the structure of a movie director.
- Contains fields for first name and last name.

### Route Handlers

#### `getMovies`

```go
func getMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}
```

Explanation:
- This function handles a GET request to fetch all movies.
- `w http.ResponseWriter` is used to write the response back to the client.
- `r *http.Request` is the request sent by the client.
- `w.Header().Set("Content-Type", "application/json")` sets the response header to indicate that the response will be in JSON format.
- `json.NewEncoder(w).Encode(movies)` encodes the `movies` slice into JSON format and writes it to the response.

#### `deleteMovie`

```go
func deleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range movies {
        if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(movies)
}
```

Explanation:
- This function handles a DELETE request to delete a movie by its ID.
- `w.Header().Set("Content-Type", "application/json")` sets the response header to indicate that the response will be in JSON format.
- `params := mux.Vars(r)` retrieves the parameters from the request URL, specifically the movie ID.
- It then iterates over the `movies` slice to find the movie with the specified ID.
- Once found, it uses slice manipulation to remove the movie from the `movies` slice.
- Finally, it encodes the updated `movies` slice into JSON format and writes it to the response.

#### `getMovie`

```go
func getMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range movies {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Movie{})
}
```

Explanation:
- This function handles a GET request to fetch a single movie by its ID.
- `w.Header().Set("Content-Type", "application/json")` sets the response header to indicate that the response will be in JSON format.
- `params := mux.Vars(r)` retrieves the parameters from the request URL, specifically the movie ID.
- It then iterates over the `movies` slice to find the movie with the specified ID.
- If the movie is found, it encodes the movie into JSON format and writes it to the response.
- If the movie is not found, it encodes an empty `Movie` struct into JSON format and writes it to the response.

#### `createMovie`

```go
func createMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var movie Movie
    _ = json.NewDecoder(r.Body).Decode(&movie)
    movie.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe for production
    movies = append(movies, movie)
    json.NewEncoder(w).Encode(movie)
}
```

Explanation:
- This function handles a POST request to create a new movie.
- `w.Header().Set("Content-Type", "application/json")` sets the response header to indicate that the response will be in JSON format.
- It creates a new `Movie` struct to hold the data sent in the request body.
- `json.NewDecoder(r.Body).Decode(&movie)` decodes the JSON data from the request body into the `movie` struct.
- It generates a mock ID for the new movie using `strconv.Itoa(rand.Intn(1000000))`.
- The new movie is then appended to the `movies` slice.
- Finally, it encodes the newly created movie into JSON format and writes it to the response.

#### `updateMovie`

```go
func updateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range movies {
        if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            var movie Movie
            _ = json.NewDecoder(r.Body).Decode(&movie)
            movie.ID = params["id"]
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
    json.NewEncoder(w).Encode(movies)
}
```

Explanation:


- This function handles a PUT request to update an existing movie by its ID.
- `w.Header().Set("Content-Type", "application/json")` sets the response header to indicate that the response will be in JSON format.
- `params := mux.Vars(r)` retrieves the parameters from the request URL, specifically the movie ID.
- It then iterates over the `movies` slice to find the movie with the specified ID.
- Once found, it uses slice manipulation to remove the existing movie from the `movies` slice.
- It then decodes the JSON data from the request body into a new `movie` struct.
- It sets the ID of the new `movie` struct to the ID specified in the request URL.
- The new `movie` struct is then appended to the `movies` slice.
- Finally, it encodes the updated `movie` into JSON format and writes it to the response.

### Route Definitions

```go
r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```

Explanation:
- Defines routes for various CRUD operations.
- Specifies the corresponding HTTP methods for each route.
- Utilizes path variables to handle dynamic IDs.

## Testing the API

1. Open Postman and create requests for each CRUD operation.
2. Send requests to the appropriate endpoints to test the API functionality.

## Conclusion

Congratulations! You have successfully built a CRUD API with Golang using structs and slices. Feel free to explore further and enhance the functionality of the API as per your requirements.

---

