package main
import "fmt"
func main(){
	var f,polegadas float64
	fmt.Scan(&f)
	fmt.Scan("polegadas")
	c:= (5*f - 160)/ 9
	mm:= polegadas*25.4
	fmt.Printf ( "o valor em cesius = %.2f n\,""c)
	fmt.Printf(" A quantidade de chuva e =%.2f n\,"mm)
}
