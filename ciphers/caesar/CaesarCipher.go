package ceaser
 
import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)
 
func main() {
	cipherKey := flag.Int("c", 0, "Cipher shift amount (-26 - 26)")
	input := flag.String("i", "", "Input")
	flag.Parse()
 
	if *cipherKey > 26 || *cipherKey < -26 {
		flag.PrintDefaults()
	} else {
		fmt.Println(caesarCipher(*input, *cipherKey))
	}
 
}
 
func caesarCipher(input string, key int) string {
	var outputBuffer bytes.Buffer
	for _, r := range strings.ToLower(input) {
		newByte := byte(r)
 
		if newByte >= 'a' && newByte <= 'z' {
			newByte += byte(key)
 
			if newByte > 'z' {
				newByte -= 26
			} else if newByte < 'a' {
				newByte += 26
			}	
		}
 
		outputBuffer.WriteString(string(newByte))
	}
	return outputBuffer.String()
}