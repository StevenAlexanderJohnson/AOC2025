package main

func findBatteryValue(batteryBank []uint8) uint64 {
	tensIndex := 0
	var tensValue uint64 = 0
	for i := range len(batteryBank) - 1 {
		if uint64(batteryBank[i]) > tensValue {
			tensValue = uint64(batteryBank[i])
			tensIndex = i
		}
	}

	var onesValue uint64 = 0
	for i := tensIndex + 1; i < len(batteryBank); i++ {
		if uint64(batteryBank[i]) > onesValue {
			onesValue = uint64(batteryBank[i])
		}
	}

	return (tensValue * 10) + onesValue
}

func part1(input [][]uint8) (uint64, error) {
	var output uint64

	for _, batteryBank := range input {
		output += findBatteryValue(batteryBank)
	}

	return output, nil
}
