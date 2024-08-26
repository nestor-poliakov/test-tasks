package app

// 30x faster then strconv.ParseUint
func ParseByte(b []byte) uint8 {
	switch len(b) {
	case 1:
		return b[0] - '0'
	case 2:
		return (b[0]-'0')*10 + b[1] - '0'
	case 3:
		return (b[0]-'0')*100 + (b[1]-'0')*10 + b[2] - '0'
	default:
		return 0
	}
}
