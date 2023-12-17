package workshops

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func Struct() {

	ae := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	iw := movie{
		title:       "Avengers: Infinity War",
		year:        2018,
		rating:      8.4,
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}

	var mvs []movie = []movie{ae, iw}

	for _, mv := range mvs {
		fmt.Printf("%#v\n", mv)
	}

}
