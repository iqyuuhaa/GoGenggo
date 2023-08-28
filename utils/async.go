package utils

var AsyncFunc = func(fn func()) {
	go fn()
}

func AsyncToSync() {
	AsyncFunc = func(fn func()) {
		fn()
	}
}

func SyncToAsync() {
	AsyncFunc = func(fn func()) {
		go fn()
	}
}
