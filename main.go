package przepisnik

import "fmt"

func main()  {
    Hello("stefan")	
}
func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

