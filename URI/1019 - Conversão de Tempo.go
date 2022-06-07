package main
 
import (
    "fmt"
)
 
func main() {
    segundos := 0
    minutos := 0
    horas := 0
    
    fmt.Scanln(&segundos)
    for segundos >= 60{
        segundos -= 60
        minutos++
    }
    
    for minutos >= 60{
        minutos -= 60
        horas++
    }
    
    
    fmt.Printf("%v:%v:%v\n", horas, minutos, segundos)
}
