package reg


type fn func ([]string)

var REGISTERED_COMMANDS = make(map[string]fn)

func RegisterNewCommand(name string, f fn) {
	REGISTERED_COMMANDS[name] = f
}