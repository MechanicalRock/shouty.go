package main

import (
    "fmt"
    "testing"
    "strconv"
)

func main3() {
    fmt.Printf("hello, world\n")
}

// Testcoordinate works out the distance from itself
func TestCoordinate_should_calculate_the_distance_from_itself(t *testing.T) {
    // it "should calculate the distance from itself"
    coord := coordinate(0,0)
    result := distanceFrom(coord, coord)

    if result != 0 {
        t.Error(strconv.Itoa(result))
    }
}

// Testcoordinate should calculate the distance from another coordinate along X axis
func TestCoordinate_should_calculate_the_distance_along_the_x_axis(t *testing.T) {
    // it "should calculate the distance from itself"
    coord1 := coordinate(0,0)
    coord2 := coordinate(400,0)
    result := distanceFrom(coord1, coord2)

    if result != 400 {
        t.Error(strconv.Itoa(result))
    }

    result = distanceFrom(coord2, coord1)
    
        if result != 400 {
            t.Error(strconv.Itoa(result))
        }
}