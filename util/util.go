package util

func IsError(err error) bool{
	if err!=nil{
		return true
	}

	return false
}
