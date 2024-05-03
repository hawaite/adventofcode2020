package day14

func int64To36BitString(value int64) string {
	bit_string := ""

	for i := 0; i < 36; i++ {
		bit_value := value & (1 << i)
		if bit_value != 0 {
			bit_string = "1" + bit_string
		} else {
			bit_string = "0" + bit_string
		}
	}
	return bit_string
}
