package reactions

import (
	"github.com/lovelaced/glitzz/config"
	"github.com/lovelaced/glitzz/modules"
	"testing"
)

func TestCuteNoParameters(t *testing.T) {
	p := New(nil, config.Default())
	output, err := p.RunCommand(modules.Command{Text: ".cute", Nick: "nick"})
	if err != nil {
		t.Errorf("error was not nil %s", err)
	}
	if len(output) != 1 {
		t.Errorf("invalid output length %d", len(output))
	}
}

func TestCute(t *testing.T) {
	p := New(nil, config.Default())
	output, err := p.RunCommand(modules.Command{Text: ".cute param1 param2", Nick: "nick"})
	if err != nil {
		t.Errorf("error was not nil %s", err)
	}
	if len(output) != 1 {
		t.Errorf("invalid output length %d", len(output))
	}
}

func TestMagicNoParameters(t *testing.T) {
	p := New(nil, config.Default())
	_, err := p.RunCommand(modules.Command{Text: ".magic", Nick: "nick"})
	if err == nil {
		t.Error("error was nil")
	}
}

func TestMagic(t *testing.T) {
	p := New(nil, config.Default())
	output, err := p.RunCommand(modules.Command{Text: ".magic param1 param2", Nick: "nick"})
	if err != nil {
		t.Errorf("error was not nil %s", err)
	}
	if len(output) != 1 {
		t.Errorf("invalid output length %d", len(output))
	}
}

func TestStumpNoParameters(t *testing.T) {
	p := New(nil, config.Default())
	_, err := p.RunCommand(modules.Command{Text: ".stump", Nick: "nick"})
	if err == nil {
		t.Error("error was nil")
	}
}

func TestStump(t *testing.T) {
	p := New(nil, config.Default())
	output, err := p.RunCommand(modules.Command{Text: ".stump param1 param2", Nick: "nick"})
	if err != nil {
		t.Errorf("error was not nil %s", err)
	}
	if len(output) != 1 {
		t.Errorf("invalid output length %d", len(output))
	}
}