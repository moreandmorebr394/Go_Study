package main

import "fmt"

func boxDetails(length, width, height float64) [2]float64 {
    surfaceArea := 2 * (length*width + length*height + width*height)
    volume := length * width * height

    return [2]float64{surfaceArea, volume}
}

func main() {
    details := boxDetails(3.0, 4.0, 5.0)
    fmt.Printf("Surface Area: %.2f, Volume: %.2f\n", details[0], details[1])
}
