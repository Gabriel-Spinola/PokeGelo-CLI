package lib

import "sync"

// Generic utility function for processing files concurrently
// and communicating the results through a channel.
//
// **Make sure to have your channel the same type as your process function return type**
func ConcurrentFileProcessor[T any](filePaths []string, channel chan<- T, process func(string) T) {
	var waitGroup sync.WaitGroup

	for _, filePath := range filePaths {
		waitGroup.Add(1)

		go func(fp string) {
			defer waitGroup.Done()

			response := process(fp)
			channel <- response
		}(filePath)
	}

	go func() {
		waitGroup.Wait()

		close(channel)
	}()
}
