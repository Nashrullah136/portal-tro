package audit

type Actor interface {
	Identity() string
}
