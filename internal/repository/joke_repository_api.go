package repository

type Joke struct {
	UUID string `json:"uuid,omitempty"`
	Joke string `json:"joke,omitempty"`
}

type JokeRepository interface {
	Get(id string) (*Joke, error)
	Random() (*Joke, error)
	Add(joke string) (*Joke, error)
}
