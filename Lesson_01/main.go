package main

import (
    "encoding/base64"
    "fmt"
    "bufio"
    "os"
)

func base64Encode(src []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
    return base64.StdEncoding.DecodeString(string(src))
}

func main() {

    //says := "dijibil.com!"

     // Klavyeden girilen veriyi elde et.
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("************\n")
    fmt.Print("Şifrelenecek Veri: ")
    says, _ := reader.ReadString('\n')
    fmt.Print("************\n\n")


    // Şifrele
    encoded := base64Encode([]byte(says))
    // Şifreli veriyi ekrana yazdır
    fmt.Println("Şifreli :", encoded)


    // Şifre Çöz
    decoded, err := base64Decode(encoded)

    if err != nil {
        fmt.Println(err.Error())
    }

    if says != string(decoded) {
        fmt.Println("'says' ile 'decoded' eşit değil.")
    }

    fmt.Println("Çözülen Veri :", string(decoded))
}
