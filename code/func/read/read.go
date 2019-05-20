package read

func GetContent() func(string) string{
	return func(s string) string{
		return "hello func3333!!!" + s
	}
}

