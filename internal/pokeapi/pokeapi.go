package pokeapi

import (
	// "encoding/json"
	"io"
	"net/http"
	"strings"
	"fmt"
)


func (c *Client) GetResponseById(baseURL *string, id string) (jsonBlob []byte, http_err error) {
	// Using "non-simple" HTTP GET method
	fullURL := *baseURL
	if len(id) > 0 {
		if strings.HasSuffix(fullURL, "/") {
			fullURL = fullURL + id
		} else {
			fullURL = fullURL + "/" + id
		}
	}
	// fmt.Printf("reading %s\n", fullURL)

	jsonBlob, ok := c.cache.Get(fullURL)
	if ok {
		return jsonBlob, nil
	}

    // create a new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []byte{}, err
	}

    // create a new client and make the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		fmt.Printf("%s not found\n", fullURL)
		return []byte{}, nil
	} else if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("HTTP failure (%d) for %s\n", fullURL)
	}

	jsonBlob, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	c.cache.Add(fullURL, jsonBlob)
	return jsonBlob, nil
}


// func getUserFromResponse(data []byte]) (User, error) {
// 	// decode the json data from the response
// 	// into a new User struct
// 	var user User
// 	err := json.Unmarshal(data, &user)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func getUsersFromResponse(data []byte]) ([]User, error) {
// 	// decode the json data from the response
// 	// into a new []User slice
// 	var users []User
// 	err := json.Unmarshal(data, &users)
// 	if err != nil {
// 		return []User{}, err
// 	}

// 	return users, nil
// }



// func getUsers(url string) ([]User, error) {
// 	// Using "simple" HTTP GET method, typically the server responds with an HTTP response of 200

// 	// By convention:
// 	// HTTP   CRUD
// 	// ====== ======
// 	// GET    Read
// 	// POST   Create
// 	// PUT    Update
// 	// PATCH  "Partially" Update
// 	// DELETE Delete

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
	
// 	users := []User{}
// 	decoder := json.NewDecoder(resp.Body)
// 	if err := decoder.Decode(&users); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func createUser(url, apiKey string, data User) (User, error) {
// 	// Using HTTP POST method, typically the server responds with an HTTP response of 201
// 	// (See https://pkg.go.dev/net/http#Response)

//     // encode our user as json
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return User{}, err
// 	}

//     // create a new request
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return User{}, err
// 	}

//     // set request headers
// 	req.Header.Set("Content-Type", "application/json")
//     req.Header.Set("X-API-Key", apiKey)

//     // create a new client and make the request
// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	defer res.Body.Close()

//     // decode the json data from the response
// 	// into a new User struct
// 	var user User
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&user)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func updateUser(baseURL, id, apiKey string, data User) (User, error) {
// 	// Using HTTP PUT method, typically the server responds with an HTTP response of 200
// 	fullURL := baseURL + "/" + id

//     // encode our user as json
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return User{}, err
// 	}

//     // create a new request
// 	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return User{}, err
// 	}

//     // set request headers
// 	req.Header.Set("Content-Type", "application/json")
//     req.Header.Set("X-API-Key", apiKey)

//     // create a new client and make the request
// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	defer res.Body.Close()

//     // decode the json data from the response
// 	// into a new User struct
// 	var user User
// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&user)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func deleteUser(baseURL, id, apiKey string) error {
// 	fullURL := baseURL + "/" + id

//     // create a new request
// 	req, err := http.NewRequest("DELETE", fullURL, nil)
// 	if err != nil {
// 		return err
// 	}

//     // set request headers
//     req.Header.Set("X-API-Key", apiKey)

//     // create a new client and make the request
// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode > 299 {
// 		return fmt.Errorf("request to delete location unsuccessful")
// 	}
// 	return nil
// }

