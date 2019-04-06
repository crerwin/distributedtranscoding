package dtc

type Job struct {
	Item       *Item
	InboxPath  string
	OutboxPath string
}
