package workshops

import (
	"encoding/json"
	"fmt"
)

type movie4 struct {
	Title       string
	Year        int
	Rating      float32
	Genres      []string
	IsSuperHero bool
}

func Json() {

	body := `[
  {
    "imdbID": "tt4154796",
    "title": "Avengers: Endgame",
    "year": 2019,
    "rating": 8.4,
    "genres": ["Action", "Drama"],
    "isSuperHero": true
  },
  {
    "imdbID": "tt4154756",
    "title": "Avengers: Infinity War",
    "year": 2018,
    "rating": 8.4,
    "genres": ["Action", "Sci-Fi"],
    "isSuperHero": true
  }
]`

	ms := []movie4{}

	err := json.Unmarshal([]byte(body), &ms)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", ms)
}
