package alert

type Priority string

const (
	Low      Priority = "Low"
	Med               = "Med"
	High              = "High"
	Critical          = "Critical"
)

type Alert struct {
	Priority Priority
	Reading  interface{}
}
