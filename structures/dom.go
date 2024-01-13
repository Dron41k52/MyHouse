package dom

import (
	"MyHouse/structures/Pets"
	"MyHouse/structures/Members"
	"MyHouse/structures/Furnitures"
	"MyHouse/structures/Gadgets"
	"MyHouse/structures/Rooms"
)

type Dom struct {
	Animals pets.Pets
	Family members.Fam
	Furnitures furnitures.FurnitureSet
	Gadgets gadgets.Gadget
	Rooms rooms.Rooms
}

func AddHouse() Dom {
	return Dom{
		Animals: pets.AddPet(),
		Family: members.AddFamily(),
		Furnitures: furnitures.AddFurnitureSet(),
		Gadgets: gadgets.AddGadgets(),
		Rooms: rooms.AddRoom(),
	}
}