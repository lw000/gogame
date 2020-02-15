package platform

type Room struct {
	Rid  int
	Name string
}

func NewRoom(rid int, name string) *Room {
	return &Room{
		Rid:  rid,
		Name: name,
	}
}
