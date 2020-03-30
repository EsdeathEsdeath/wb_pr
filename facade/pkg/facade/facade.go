//Package facade implement method to create a movie
package facade

type filmScreenwriter interface {
	WriteScript(name string)
}

type filmProducers interface {
	MakeMovie(name string)
}

type filmActors interface {
	PlayRole(name string)
}

type filmStuntman interface {
	DoTricks(name string)
}

type Film interface {
	MakeGreatFilm()
}

type film struct {
	screenwriter   filmScreenwriter
	producer   filmProducers
	actors filmActors
	stuntman  filmStuntman
}

//Set group to create a great film
func (f *film) MakeGreatFilm(){
	f.screenwriter.WriteScript("Billy Wilder")
	f.producer.MakeMovie("Michael Mann")
	f.actors.PlayRole("Stan Smith")
	f.stuntman.DoTricks("Dar Robinson")
}

// Film squad creating
func FilmSquad(
	screenwriter filmScreenwriter,
	producer filmProducers,
	actors filmActors,
	stuntman filmStuntman) Film {
	return &film{
		screenwriter:   screenwriter,
		producer:   producer,
		actors: actors,
		stuntman:  stuntman,
	}
}

