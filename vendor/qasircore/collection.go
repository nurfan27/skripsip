package qasircore

// Collection is a helper for easier to manipulate data query
type Collection struct {
	rows []string
	data []map[string]string
}

/**
 * Setup data for collection
 */
func (this *Collection) AppendData(initdata map[string]string) {
	this.data = append(this.data, initdata)
}

/**
 * Filtering data based on function filter that been set
 * @param filterFunction func(value, key string) bool
 * @return collection *Collection
 */
func (this *Collection) Filter(filterFunction func(value map[string]string) bool) *Collection {
	var newFilterData []map[string]string

	for _, value := range this.data {
		if filterFunction(value) {
			newFilterData = append(newFilterData, value)
		}
	}

	return this
}

/**
 * This will reject all data from data based on validation that been design closure function.
 * @param rejectFunction func(value, key string) bool
 * @return collection *Collection
 */
func (this *Collection) Reject(rejectFunction func(value map[string]string) bool) *Collection {
	var newRejectData []map[string]string

	for _, value := range this.data {
		if !rejectFunction(value) {
			newRejectData = append(newRejectData, value)
		}
	}

	this.data = newRejectData
	return this
}

/**
 * Get lists key from data in collection
 * @param key string
 * @return listsData []string
 */
func (this *Collection) Lists(key string) []string {
	var listsData []string

	for _, val := range this.data {
		for index, data := range val {
			if index == key {
				listsData = append(listsData, data)
			}
		}
	}

	return listsData
}

/**
 * Transform data based on closure function we defined
 * @param closureFunction func(data map[string]string) map[string]string
 * @return *Collection
 */
func (this *Collection) Transform(closureFunction func(data map[string]string) map[string]string) *Collection {
	for key, data := range this.data {
		this.data[key] = closureFunction(data)
	}
	return this
}

/**
 * Where function for key and value
 * @param key string
 * @param value string
 * @return *Collection
 */
func (this *Collection) Where(key, value string) *Collection {
	var newlistdata []map[string]string
	for _, arraydata := range this.data {
		for index, data := range arraydata {
			if index == key && data == value {
				newlistdata = append(newlistdata, arraydata)
			}
		}
	}

	this.data = newlistdata

	return this
}

/**
 * Return all of data array
 * @return data []map[string]string
 */
func (this *Collection) All() []map[string]string {
	return this.data
}

/**
 * Return First of object on array
 * @return data map[string]string
 */
func (this *Collection) First() map[string]string {
	return this.data[0]
}

/**
 * Return object last of data
 * @return map[string]string
 */
func (this *Collection) Last() map[string]string {
	length := len(this.data)

	return this.data[length-1]
}
