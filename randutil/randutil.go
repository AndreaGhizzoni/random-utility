package randutil

import "math/rand"

/*
func Get(min, max interface{}) interface{} {
	if reflect.TypeOf(min) != reflect.TypeOf(max) {
		fmt.Println("type mismatch")
		return nil
	}

	minV := reflect.ValueOf(min)
	maxV := reflect.ValueOf(max)
	switch minV.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		fmt.Println("TODO")
		return nil
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return randomInt64(minV.Int(), maxV.Int())
	case reflect.Float32, reflect.Float64:
		return randomFloat64(minV.Float(), maxV.Float())
	default:
		fmt.Println("type not found")
		return nil
	}
}
*/

func Init(seed int64) {
	rand.Seed(seed)
}

func Int64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func Float64(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
