package handle

import "fmt"

func SuccessHandle(success bool) {
	if !success {
		fmt.Println("Didnt succeed")
	}
	fmt.Println("Successful")
}
