package ai

type AI interface {
	Text(prompt string) (string, error)
	Audio(prompt string) ([]byte, error)
	Video(prompt string) ([]byte, error)
}
