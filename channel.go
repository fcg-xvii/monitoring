package monitoring

type Channel interface {
	Log(title, message string)
}
