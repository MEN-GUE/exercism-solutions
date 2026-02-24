package jedlik
import "fmt"

//  Drive update the given car battery and distance
// based on its batteryDrain and speed.
// If the battery would be fully drain, nothing happens.
func (car *Car) Drive() Car {
    if car.battery - car.batteryDrain > 0 {
        car.distance += car.speed
        car.battery -= car.batteryDrain
        fmt.Println(car)
        return *car
    }
    return *car
}

// DisplayDistance return a string showing the given Car
// distance driven.
func (car Car) DisplayDistance() string {
    return fmt.Sprint("Driven ", car.distance," meters")
}

// DisplayBattery return a string showing the given Car
// battery.
func (car Car) DisplayBattery() string {
    return fmt.Sprint("Battery at ", car.battery, "%")
}

// CanFinish return a bool that checks if the Car can finish
// the given trackDistance without depleting fully the battery.
func (car Car) CanFinish(trackDistance int) bool {
    if (car.battery/car.batteryDrain) * car.speed < trackDistance {
        return false
    }
    return true
}


