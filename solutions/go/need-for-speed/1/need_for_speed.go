package speed

import "fmt"

type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

type Track struct{
    distance int
}

func NewTrack(distance int) Track {
	return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

    fmt.Println("car debug before:", car)
    
    if car.battery >= car.batteryDrain {
        car.distance += car.speed
   		car.battery -= car.batteryDrain
    }
    fmt.Println("car debug after:", car)
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    
    fmt.Println("car - track:", car, track)
    
    batteryAvailability := float64(car.battery) / float64(car.batteryDrain)
    fmt.Println("batteryAvailability:", batteryAvailability)
    
    metersCarCanRun := batteryAvailability * float64(car.speed)
    fmt.Println("metersCarCanRun:", metersCarCanRun)
    
    return metersCarCanRun >= float64(track.distance)
}
