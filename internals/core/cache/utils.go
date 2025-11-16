package cache

func elemsToDelete(data []ImageData, diff int64) ([]ImageData, []string) {
	var (
		sum            int64       = 0
		checked        int         = 0
		toDelete       []ImageData = make([]ImageData, 0)
		hashesToDelete []string    = make([]string, 0)
	)

	for sum < diff {
		if len(data) > checked {
			dataToAdd := (data)[checked]
			toDelete = append(toDelete, dataToAdd)
			hashesToDelete = append(hashesToDelete, dataToAdd.Hash)
			sum += dataToAdd.Size
			checked++
		} else {
			return toDelete, hashesToDelete
		}
	}

	return toDelete, hashesToDelete
}
