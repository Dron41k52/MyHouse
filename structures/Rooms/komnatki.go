package rooms

type Room struct {
	Name   string
	Length float32
	Width  float32
	Square float32
	Style  string
}

type Rooms struct {
	kitchen  Room
	bathroom Room
	bedroom  Room
	hall     Room
}

func AddRoom() Rooms {
	return Rooms{
		kitchen:  Room{Name: "kuhnya", Length: 5.0, Width: 4.0, Square: 20.0, Style: "opium"},
		bathroom: Room{Name: "vannaya", Length: 4.0, Width: 4.0, Square: 16.0, Style: "opium"},
		bedroom:  Room{Name: "spalnya", Length: 6.0, Width: 4.0, Square: 24.0, Style: "opium"},
		hall:     Room{Name: "gostinaya", Length: 8.0, Width: 4.0, Square: 32.0, Style: "opium"},
	}
}