package Player

type Player interface {
	PlayCard([]string) string
}

type PLayerInfo struct  {
	name string;
	deck []string;
}

