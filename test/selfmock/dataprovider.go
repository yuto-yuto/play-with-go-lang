package selfmock

import "fmt"

type Provider struct {
	reader Reader
}

func NewProvider(reader Reader) *Provider{
	instance := new(Provider)
	instance.reader = reader

	return instance
}

func (p *Provider) ProvideData() (string, error) {
	err := p.reader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open: %w", err)
	}
	defer p.reader.Close()

	data, err := p.reader.Read(99)
	if err != nil {
		return "", fmt.Errorf("failed to read: %w", err)
	}
	return data, nil
}
