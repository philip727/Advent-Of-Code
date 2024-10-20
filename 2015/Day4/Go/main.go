package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func FindLowestIterOfKey(secretKey string) int {
    i := 0;
    for {
        key := fmt.Sprintf("%s%d", secretKey, i);

        sum := md5.Sum([]byte(key));
        hex := hex.EncodeToString(sum[:]);
        check := hex[:5];
        if check == "00000" {
            break;
        }


        i++;
    }

    return i;
}

func main() {
    var secretKey string;

    fmt.Print("Please input a secret key: ");
    _, err := fmt.Scan(&secretKey);
    if err != nil {
        fmt.Println(err);
        return;
    }
    fmt.Print("\n");

    i := FindLowestIterOfKey(secretKey);

    fmt.Printf("Lowest num: %d\n", i);
}
