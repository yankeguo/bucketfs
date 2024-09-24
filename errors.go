package bucketfs

// AdapterNotFound is an error that is returned when an adapter is not found.
type AdapterNotFound struct {
	AdapterName string
}

func (e *AdapterNotFound) Error() string {
	return "adapter not found: " + e.AdapterName
}

// AdapterOptionMissing is an error that is returned when an adapter option is missing.
type AdapterOptionMissing struct {
	OptionName string
}

func (e *AdapterOptionMissing) Error() string {
	return "adapter option missing: " + e.OptionName
}

// AdapterOptionInvalid is an error that is returned when an adapter option is invalid.
type AdapterOptionInvalid struct {
	OptionName string
}

func (e *AdapterOptionInvalid) Error() string {
	return "adapter option invalid: " + e.OptionName
}
