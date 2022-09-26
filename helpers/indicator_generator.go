package helpers

import "log"

func GenerateIndicator(valueType string, statusNum int) string {
	switch valueType {
	case "water":
		return waterIndicator(statusNum)
	case "wind":
		return windIndicator(statusNum)
	default:
		log.Fatal("Type doesn't exist")
		return ""
	}
}

func waterIndicator(statusNum int) string {
	switch {
	case statusNum < 5:
		return "aman"
	case statusNum > 5 && statusNum < 9:
		return "siaga"
	case statusNum >= 9:
		return "bahaya"
	default:
		log.Fatal("status out of bounds")
		return ""
	}
}

func windIndicator(statusNum int) string {
	switch {
	case statusNum < 6:
		return "aman"
	case statusNum > 6 && statusNum < 16:
		return "siaga"
	case statusNum >= 16:
		return "bahaya"
	default:
		log.Fatal("status out of bounds")
		return ""
	}
}
