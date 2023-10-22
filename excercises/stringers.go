package excercises

import "fmt"

type IPAddr [4]byte

// TODO add a 'String' string method to IPAddr

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDns": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Println(name, ip)
	}

}
