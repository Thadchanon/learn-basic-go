package workshops

import "fmt"

type movie2 struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie2) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

func Pointer() {
	eg := &movie2{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	eg.addVote(8)
	fmt.Println("votes:", eg.votes)

}
