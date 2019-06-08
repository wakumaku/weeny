package hasher

type Hash interface {
	Encode(string) (string, error)
}
