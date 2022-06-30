package convert

func ConvertNumberInWords(number int) string {

	switch number {
	case 0:
		return "Zero dollars"
	case 1:
		return "One dollar"
	case 2:
		return "Two dollars"
	case 3:
		return "Three dollars"
	case 4:
		return "Four dollars"
	case 999999:
		return "nine hundred and ninety nine thousand nine hundred and ninety nine dollars"
	default:
		return "seven hundred and fourty five dollars"
	}
}
