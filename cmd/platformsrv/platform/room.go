package platform

type Room struct {
	Rid  int
	Name string
}

func NewRoom(rid int, rname string) *Room {
	return &Room{
		Rid:  rid,
		Name: rname,
	}
}
