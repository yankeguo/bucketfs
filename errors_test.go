package bucketfs

var (
	_ error = &AdapterNotFound{}
	_ error = &AdapterOptionMissing{}
	_ error = &AdapterOptionInvalid{}
)
