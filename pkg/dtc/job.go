package dtc

type Job struct {
	Item        *Item
	ItemSubPath string
	InboxPath   string
	OutboxPath  string
}
