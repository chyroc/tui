package internal

func IsChanIntClosed(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func CloseChanInt(ch chan int) {
	// select {
	// case <-ch:
	// 	return
	// default:
	// }

	close(ch)
}

func CloseChanStruct(ch chan struct{}) {
	select {
	case <-ch:
		return
	default:
	}

	close(ch)
}
