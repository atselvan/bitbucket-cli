package utils

func EntryExists(slice []string, entry string) bool{
	for i:=0; i<len(slice);i++{
		if slice[i] == entry {
			return true
		}
	}
	return false
}