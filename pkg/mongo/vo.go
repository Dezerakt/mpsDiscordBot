package mongoPkg

type Database string

const (
	Events Database = "events"
)

func (obj Database) Str() string {
	return string(obj)
}

type Collection string

const (
	Channels Collection = "channels"
)

func (obj Collection) Str() string {
	return string(obj)
}
