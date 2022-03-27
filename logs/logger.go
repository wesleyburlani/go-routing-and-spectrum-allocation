package logs

type Logger interface {
	Log(logs ...interface{})
}
