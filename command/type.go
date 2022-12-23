package command

type Type string

type Wrapper struct {
	EventType Type `json:"kind"`
}

type TopicHolder struct {
	NotificationTopics []string
	AlertTopics        []string
}

type TemperatureReading struct {
	TopicHolder
	Reading float32 `json:"value"`
}

type HumidityReading struct {
	TopicHolder
	Reading float32 `json:"value"`
}

type WindReading struct {
	TopicHolder
	Gust      int32   `json:"gust"`
	Lull      int32   `json:"lull"`
	Direction float32 `json:"direction"`
}

type LightningReading struct {
	TopicHolder
	power int32 `json:"power"`
}
