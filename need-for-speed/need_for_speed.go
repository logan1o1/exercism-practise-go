package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var remBatr int
	var carDist int
	if car.battery >= car.batteryDrain {
		remBatr = car.battery - car.batteryDrain
		carDist = car.distance + car.speed
	} else {
		remBatr = car.battery
		carDist = car.distance
	}
	return Car{
		battery:      remBatr,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     carDist,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	turns := car.battery / car.batteryDrain
	if car.speed*turns >= track.distance {
		return true
	}
	return false
}
