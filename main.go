package main

func main() {
	startRepl()

}

//type LocationAreasResponse struct {
//Count    int            `json:"count"`
//Next     *string        `json:"next"`
//Previous *string        `json:"previous"`
//Results  []LocationArea `json:"results"`
//}
//type LocationArea struct {
//Name string `json:"name"`
//URL  string `json:"url"`
//}

//func testAPI() {
// Make a simple GET request to the PokeAPI
//baseURL := "https://pokeapi.co/api/v2/location-area"
//res, err := http.Get(baseURL)

//if err != nil {
//fmt.Println("error:", err)
//}

//defer res.Body.Close()

//body, err := io.ReadAll(res.Body)
//if err != nil {
//fmt.Println("error:", err)
//}

//var LocationAreasResponse LocationAreasResponse
//if err := json.Unmarshal(body, &LocationAreasResponse); err != nil {
//fmt.Println("Error Unmarshal:", err)
//}

//for _, location := range LocationAreasResponse.Results {
//fmt.Println(location.Name)
//}
//}
