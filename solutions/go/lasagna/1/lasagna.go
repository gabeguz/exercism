package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40  // How much time the dish should be in the oven

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(t int) int {
    return OvenTime - t
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    return PreparationTime(numberOfLayers) + actualMinutesInOven
}
