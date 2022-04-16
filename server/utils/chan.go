package utils

import "fmt"

func IsChanClosed(ch <-chan string) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func OutPutChan(inCh chan string, quitCh chan bool) {
	for {
		select {
		case <-inCh:
			if log, ok := <-inCh; ok {
				fmt.Println("ch", log)
			}
		case <-quitCh:
			break
		}
	}
}
