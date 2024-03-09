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
git clone https://github.com/monster0freason/A-CRUD-API-With-Golang
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

var movies []Movie
```

Explanation:
- Defines the structure of a movie director.
- Contains fields for first name and last name.

### Route Handlers

#### `getMovies`

```go
// getMovies is a handler function for the GET request to fetch all movies.
func getMovies(w http.ResponseWriter, r *http.Request) {
    // Set the Content-Type header of the response to indicate JSON format.
    w.Header().Set("Content-Type", "application/json")
    
    // Encode the 'movies' slice into JSON format and write it to the response.
    // 'json.NewEncoder(w)' creates a new JSON encoder that writes to the ResponseWriter 'w'.
    // 'Encode(movies)' encodes the 'movies' slice and writes it to the ResponseWriter.
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
// deleteMovie handles the HTTP DELETE request to delete a movie by its ID.
func deleteMovie(w http.ResponseWriter, r *http.Request) {
    // Set the response header to indicate JSON content type
    w.Header().Set("Content-Type", "application/json")

    // Extract parameters from the request URL
    params := mux.Vars(r)

    // Iterate through the list of movies
    for index, item := range movies {
        // Check if the ID of the current movie matches the ID specified in the request
        if item.ID == params["id"] {
            // If found, remove the movie from the slice by slicing it
            movies = append(movies[:index], movies[index+1:]...)
            // Exit the loop since the movie is found and deleted
            break
        }
    }

    // Encode the updated list of movies into JSON format and send it in the response
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
// Define a function named getMovie with parameters w (http.ResponseWriter) and r (*http.Request)
func getMovie(w http.ResponseWriter, r *http.Request) {
    // Set the response header to indicate JSON content type
    w.Header().Set("Content-Type", "application/json")
    
    // Extract the parameters (including the movie ID) from the request URL
    params := mux.Vars(r)
    
    // Iterate through the list of movies
    for _, item := range movies {
        // Check if the ID of the current movie matches the ID provided in the request
        if item.ID == params["id"] {
            // If a match is found, encode the movie details into JSON format and write it to the response
            json.NewEncoder(w).Encode(item)
            return // Exit the function
        }
    }
    
    // If no matching movie is found, encode an empty Movie struct into JSON format and write it to the response
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
    // Set the content type of the response to JSON
    w.Header().Set("Content-Type", "application/json")

    // Declare a variable to hold the decoded JSON data
    var movie Movie
    
    // Decode the JSON data from the request body into the movie variable
    _ = json.NewDecoder(r.Body).Decode(&movie)
    
    // Generate a mock ID for the new movie (not safe for production)
    movie.ID = strconv.Itoa(rand.Intn(1000000))
    
    // Append the new movie to the movies slice
    movies = append(movies, movie)
    
    // Encode the newly created movie as JSON and write it to the response
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
    // Set the response header to indicate JSON content type
    w.Header().Set("Content-Type", "application/json")
    
    // Extract route parameters from the request
    params := mux.Vars(r)
    
    // Iterate through the movies slice to find the movie with the specified ID
    for index, item := range movies {
        if item.ID == params["id"] { // Check if the movie ID matches the requested ID
            // Remove the existing movie from the movies slice
            movies = append(movies[:index], movies[index+1:]...)
            
            // Decode the request body JSON into a new movie object
            var movie Movie
            _ = json.NewDecoder(r.Body).Decode(&movie)
            
            // Set the ID of the updated movie to the requested ID
            movie.ID = params["id"]
            
            // Append the updated movie to the movies slice
            movies = append(movies, movie)
            
            // Encode the updated movie as JSON and write it to the response
            json.NewEncoder(w).Encode(movie)
            
            // Return to exit the function after updating the movie
            return
        }
    }
    
    // If the movie with the specified ID is not found, encode the movies slice and write it to the response
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
// Define routes for various CRUD operations using HandleFunc.
// For each route, specify the corresponding HTTP method.
// Utilize path variables to handle dynamic IDs.

// Handle GET request to fetch all movies
r.HandleFunc("/movies", getMovies).Methods("GET")

// Handle GET request to fetch a single movie by ID
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

// Handle POST request to create a new movie
r.HandleFunc("/movies", createMovie).Methods("POST")

// Handle PUT request to update an existing movie by ID
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

// Handle DELETE request to delete a movie by ID
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

// Print a message indicating the server is starting at port 8000
fmt.Print("Starting server at port 8000\n")

// Start the HTTP server listening on port 8000
// Log any fatal errors that occur during server operation
log.Fatal(http.ListenAndServe(":8000", r))

```

Explanation:

The route definitions establish endpoints for CRUD operations, with each route specifying the corresponding HTTP method. Path variables are employed to handle dynamic IDs, enabling the API to interact with specific resources efficiently. This approach adheres to RESTful design principles, ensuring clarity, maintainability, and security in API routing.

- **`r.HandleFunc(path, handler).Methods(httpMethod)`:**
  - This statement is used to define routes for various CRUD operations in the API.
  - `r` is the instance of the Gorilla Mux router (`mux.Router`) that handles incoming HTTP requests.
  - `HandleFunc` is a method provided by the Gorilla Mux router to register a new route with the router.
  - `path` is the URL path pattern for the route. It specifies the endpoint where the HTTP request should be routed.
  - `handler` is the function that should be called to handle the HTTP request when it matches the specified path.
  - `httpMethod` is the HTTP method (e.g., GET, POST, PUT, DELETE) associated with the route.

- **Using Path Variables `{id}`:**
  - The routes `/movies/{id}` are examples of routes that use path variables. Path variables are specified by curly braces `{}` in the route definition.
  - The path variable `{id}` is used to dynamically capture the ID of a movie from the URL. This allows for fetching, updating, or deleting a specific movie based on its unique identifier.

- **Specifying HTTP Methods:**
  - The `.Methods(httpMethod)` part of each route definition specifies the HTTP method allowed for accessing that route.
  - For example, `Methods("GET")` means that the route only responds to HTTP GET requests, while `Methods("POST")` means it only responds to HTTP POST requests, and so on.
  - This ensures that the API adheres to the principles of RESTful design by mapping HTTP methods to CRUD operations: GET for reading, POST for creating, PUT for updating, and DELETE for deleting resources.

- **Why Using Methods:**
  - Specifying HTTP methods for each route helps in enforcing the proper usage of the API and ensures that each endpoint is only accessible via the intended HTTP method.
  - It enhances the clarity and maintainability of the code by clearly indicating the allowed actions for each route.
  - Additionally, it provides an additional layer of security by preventing unauthorized access to sensitive operations.


## Testing the API with Postman

1. **Setting Up Postman:**
   - Ensure you have Postman installed on your system.
   - Open Postman to start testing the API endpoints.

2. **Creating Requests:**
   - Create a new folder within Postman to organize your requests. For example, name it "Go Movies".

3. **Fetching All Movies (GET Request):**
   - Create a new request within the "Go Movies" folder to retrieve all movies from the API.
   - Set the request URL to `http://localhost:8000/movies`.
   - Send the request to fetch all movies in JSON format.

4. **Fetching a Movie by ID (GET Request):**
   - Create another request within the same folder to fetch a specific movie by its ID.
   - Set the request URL to `http://localhost:8000/movies/{id}`, replacing `{id}` with the ID of the desired movie.
   - Specify the ID of the movie to retrieve, for example, `/1` to fetch the first movie.
   - Send the request to retrieve the details of the specified movie.

5. **Creating a New Movie (POST Request):**
   - Create a new request within the "Go Movies" folder to add a new movie to the API.
   - Set the request URL to `http://localhost:8000/movies` and choose the HTTP method as POST.
   - In the request body, provide JSON data representing the new movie to be created, excluding the ID since it will be generated automatically by the API.
   - Send the request to create the new movie and receive a response containing the details of the newly created movie, including the automatically generated ID.

6. **Updating an Existing Movie (PUT Request):**
   - Create another request to update an existing movie in the API.
   - Set the request URL to `http://localhost:8000/movies/{id}`, replacing `{id}` with the ID of the movie to be updated.
   - Specify the ID of the movie to update and set the HTTP method to PUT.
   - In the request body, provide JSON data representing the updated details of the movie.
   - Send the request to update the movie and receive a response containing the updated details.

7. **Deleting a Movie (DELETE Request):**
   - Create a new request to delete a movie from the API.
   - Set the request URL to `http://localhost:8000/movies/{id}`, replacing `{id}` with the ID of the movie to be deleted.
   - Specify the ID of the movie to delete and set the HTTP method to DELETE.
   - Send the request to delete the movie and receive a confirmation response.

8. **Verifying Changes:**
   - Verify the changes made to the movies by examining the responses from each request.
   - Check the addition, retrieval, updating, and deletion of movies to ensure that the API functions correctly.



Postman serves as a versatile tool for testing the functionality of the CRUD API, allowing you to interact with the API endpoints and verify their behavior effectively.

## Testing the API

1. Open Postman and create requests for each CRUD operation.
2. Send requests to the appropriate endpoints to test the API functionality.

## Conclusion

Congratulations! You have successfully built a CRUD API with Golang using structs and slices. Feel free to explore further and enhance the functionality of the API as per your requirements.

---

