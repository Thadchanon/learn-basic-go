package workshops

import "fmt"

type voter interface {
	addVote(rating float64) // or addVote(float64)
}

func vote(v voter, rating float64) {
	v.addVote(rating)
}

type movie3 struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie3) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

func Interface() {

	eg := &movie3{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	vote(eg, 8)
	fmt.Println("votes:", eg.votes)
}
