package src

type GasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo Owner
}

type Owner struct {
	name string
	age  int8
}

// This means Drive and milesLeft methods must be implemented by any type that wants to be considered an Engine.

// New Constructor
func New(i int, i2 int) *GasEngine {
	return &GasEngine{mpg: uint8(i), gallons: uint8(i2)}
}

// Methods

func (g *GasEngine) Drive() {
	g.gallons--
}

func (g *GasEngine) MilesLeft() uint16 {
	return uint16(g.GetGallons()) * uint16(g.GetMpg())
}

func CanMakeIt(g *GasEngine) bool {
	return g.MilesLeft() > 0
}

// Getter Methods

func (g *GasEngine) GetMpg() uint8 {
	return g.mpg
}

func (g *GasEngine) GetGallons() uint8 {
	return g.gallons
}

func (g *GasEngine) GetOwnerInfo() Owner {
	return g.ownerInfo
}

func (g *GasEngine) GetOwnerName() string {
	return g.ownerInfo.name
}

func (g *GasEngine) GetOwnerAge() int8 {
	return g.ownerInfo.age
}

// Setter methods

func (g *GasEngine) SetOwnerInfo(name string, age int8) {
	g.ownerInfo = Owner{name: name, age: age}
}

func (g *GasEngine) SetMpg(mpg uint8) {
	g.mpg = mpg
}

func (g *GasEngine) SetGallons(gallons uint8) {
	g.gallons = gallons
}
