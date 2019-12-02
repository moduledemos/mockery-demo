package demo

type Story interface {
	Open()
	ReadChapter(chapter int) ([]byte, error)
	HasChapter(chapter int) bool
	Close()
}
