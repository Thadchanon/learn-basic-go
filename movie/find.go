package movie

func FindMovieName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found"
}
