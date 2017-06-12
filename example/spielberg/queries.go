package main


var moviesSchema = `
mutation {
  schema {
    movie.name: string @index .
    person.name: string @index .
    movie.release_date: date @index .
    movie.revenue: float .
    movie.running_time: int .
  }
}

`

var allMovies = `
query {
	director(func:allofterms(name, "michael bay")) {
		name@en
		director.film (orderdesc: initial_release_date) {
			name@en
			initial_release_date
		}
	}
}

`
