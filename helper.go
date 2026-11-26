package main

func isReserved(code string) bool {
	switch code {
	case "healthz", "metrics":
		return true
	default:
		return false
	}
}
