package types

type BroadcastMessage struct {
	Name string
	id   string
}

func (b *BroadcastMessage) ScratchID() string {
	return b.id
}

func NewScratchBroadcastMessage(name, id string) *BroadcastMessage {
	return &BroadcastMessage{
		Name: name,
		id:   id,
	}
}
