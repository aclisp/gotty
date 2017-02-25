package webtty

const (
	// Unknown message type, maybe sent by a bug
	InputUnknown = '0'
	// User input typically from a keyboard
	Input = '1'
	// Ping to the server
	Ping = '2'
	// Notify that the browser size has been changed
	ResizeTerminal = '3'
)

const (
	// Unknown message type, maybe set by a bug
	Unknown = '0'
	// Normal output to the terminal
	Output = '1'
	// Pong to the browser
	Pong = '2'
	// Set window title of the terminal
	SetWindowTitle = '3'
	// Set preference to the terminal
	SetPreferences = '4'
	// Set reconnect configuration to the terminal
	SetReconnect = '5'
)