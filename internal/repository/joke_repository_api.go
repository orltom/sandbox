package repository

type Joke struct {
	UUID string `json:"uuid,omitempty"`
	Joke string `json:"joke,omitempty"`
}

type JokeRepository interface {
	FindByID(id string) (*Joke, error)
	Random() (*Joke, error)
	Create(uuid string, joke string) (*Joke, error)
}
