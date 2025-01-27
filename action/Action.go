package action

type Action int

const (
	NONE = -1
	SET  = 1
	// OPEN = 1
	// LIST = 2
	// NEW  = 3
	GET = 2 // 4
	// PUT  = 5
	// REM  = 6
)

func (me Action) ToString() string {
	switch me {
	case SET:
		return "set"
	/*
		case OPEN:
			return "open"
		case LIST:
			return "list"
		case NEW:
			return "new"
	*/
	case GET:
		return "get"
	/*
		case PUT:
			return "put"
		case REM:
			return "rem"
	*/
	default:
		return ""
	}
}
