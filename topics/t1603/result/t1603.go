package result

import "sync"

var mutex1 sync.Mutex
var mutex2 sync.Mutex
var mutex3 sync.Mutex

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	result := false
	switch carType {
	case 1:
		{
			mutex1.Lock()
			defer mutex1.Unlock()
			if this.Big > 0 {
				this.Big--
				result = true
			}
		}
	case 2:
		{
			mutex2.Lock()
			defer mutex2.Unlock()
			if this.Medium > 0 {
				this.Medium--
				result = true
			}
		}
	case 3:
		{
			mutex3.Lock()
			defer mutex3.Unlock()
			if this.Small > 0 {
				this.Small--
				result = true
			}
		}
	}

	return result
}
