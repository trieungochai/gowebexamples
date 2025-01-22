This Go code sets up a simple HTTP server with two endpoints: `/decode` and `/encode`. The code demonstrates how to handle JSON data, both decoding incoming JSON to a Go struct and encoding a Go struct into JSON to send as a response.

### `User` Struct:

```go
type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}
```

- `User` is a struct that represents a user, with three fields: `Firstname`, `Lastname`, and `Age`.
- The struct tags (`json:"firstname"`, `json:"lastname"`, `json:"age"`) are used to specify how the fields should be encoded/decoded to/from JSON. These tags ensure the JSON keys match the field names in the struct.

### `/decode` Endpoint:

```go
http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
})
```

- This endpoint expects an incoming POST request with JSON data in the request body.
- The `json.NewDecoder(r.Body).Decode(&user)` reads and decodes the JSON from the request body into the `user` variable, which is of type `User`.
- After decoding, it sends a formatted response, displaying the `Firstname`, `Lastname`, and `Age` fields of the `user` struct.

### `/encode` Endpoint:

```go
http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
    peter := User{
        Firstname: "John",
        Lastname:  "Doe",
        Age:       25,
    }

    json.NewEncoder(w).Encode(peter)
})
```

- This endpoint creates a `User` struct named `peter` with hardcoded values: `"John"` for `Firstname`, `"Doe"` for `Lastname`, and `25` for `Age`.
- The `json.NewEncoder(w).Encode(peter)` encodes the `peter` struct into JSON and sends it as the HTTP response.

### Main Function:

```go
http.ListenAndServe(":8080", nil)
```

- Starts the HTTP server on port `8080`, handling requests for the defined endpoints (`/decode` and `/encode`).

### Summary of Endpoints:

1. **`/decode`**:

   - Expects a `POST` request with JSON data in the body.
   - Decodes the JSON into a `User` struct and responds with a message like: `"John Doe is 25 years old!"`.

2. **`/encode`**:
   - Responds with a JSON representation of a predefined `User` struct (`{"firstname": "John", "lastname": "Doe", "age": 25}`).

### Example Usage:

- **POST to `/decode`**:

  ```json
  {
    "firstname": "Alice",
    "lastname": "Smith",
    "age": 30
  }
  ```

  The response will be:  
  `"Alice Smith is 30 years old!"`

- **GET from `/encode`**:  
  The response will be:
  ```json
  {
    "firstname": "John",
    "lastname": "Doe",
    "age": 25
  }
  ```
