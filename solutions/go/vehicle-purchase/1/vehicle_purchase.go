package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	if kind == "truck" {
		return true
	}

	if kind == "car" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var choice string

	if option1 < option2 {
		choice = option1
	}

	if option2 < option1 {
		choice = option2
	}

	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.80
	}

	if age >= 10 {
		return originalPrice * 0.50
	}

	if age >= 3 && age < 10 {
		return originalPrice * 0.70
	}
	return 0.0
}
