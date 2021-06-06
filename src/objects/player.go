package objects

type Player struct {
	Name string
}

func (p Player) GetPlayerName() string {
	return p.Name
}