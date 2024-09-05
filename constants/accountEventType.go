package constants
const (
	// Event types
	EVENT_TYPE_DEPOSIT  = "deposit"
	EVENT_TYPE_WITHDRAW = "withdraw"
	EVENT_TYPE_TRANSFER = "transfer"
)
type EventType int

const  (
	UNKNOWN EventType = iota
	DEPOSIT
	WITHDRAW
	TRANSFER

)

func (e EventType) String() string {
	switch e {
	case DEPOSIT:
		return EVENT_TYPE_DEPOSIT
	case WITHDRAW:
		return EVENT_TYPE_WITHDRAW
	case TRANSFER:
		return EVENT_TYPE_TRANSFER
	default:
		return ""		
	}
}

func (e EventType) Index() int {
	return int(e)
}

func ToEventTypeEnum(s string) EventType {
	switch s {
	case EVENT_TYPE_DEPOSIT:
		return DEPOSIT
	case EVENT_TYPE_WITHDRAW:
		return WITHDRAW
	case EVENT_TYPE_TRANSFER:
		return TRANSFER
	default:
		return UNKNOWN
	}
}