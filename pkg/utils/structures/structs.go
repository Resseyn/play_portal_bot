package structures

type MessageData struct {
	ChatID      int64
	MessageID   int
	Command     string
	PrevCommand string
}

type Command struct {
	Text    string
	Command string
}
