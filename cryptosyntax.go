// ham bam ma hoa 1 chuoi string
package main
 
import "fmt"
import "crypto/sha1"
 
func main() {
  
    str1 := []byte("Song la de yeu thuong, song la de hoan thanh giac mo su nghiep") 
    crypto := sha1.New()
    crypto.Write(str1)
    fmt.Println(crypto.Sum([]byte{}))
    fmt.Printf("%x", crypto.Sum([]byte{}))
}
