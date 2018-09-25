package misc

func SecondsToHMS(inSeconds int) string {
	hours := inSeconds / (60 * 60)
	seconds := inSeconds % (60 * 60)
	minutes := seconds / 60
	seconds = inSeconds % 60
	str := fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)
	return str
}
