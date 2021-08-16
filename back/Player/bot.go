package Player

type BotInfo struct {
	Name string `json:"username" binding:"-"`;
	Deck []string `json:"deck" binding:"-"`;
}