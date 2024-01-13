package pets

type Pet struct {
	Type  string
	Name  string
	Age   int16
	Color string
	Sex   string
}

type Pets struct {
	Pets []Pet
}

func AddPet() Pets {
	var pets []Pet

	pets = append(pets, Pet{Type: "rooster", Name: "Wilson", Age: 3, Color: "brown and red", Sex: "Male"})

	return Pets{Pets: pets}
}