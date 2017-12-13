package main

import (
    . "github.com/lsegal/gucumber"
    "github.com/stretchr/testify/assert"
)

func init() {
    var initVal int

    // Examples - from https://gnuu.org/2015/03/30/gucumber-is-cucumber-for-go/

    // Given(`^I have an initial value of (d+)$`, func(val int) {
    //     initVal = val
    // })
    //
    // And(`^I add (d+) to the value$`, func(val int) {
    //     initVal += val
    // })
    //
    // Then(`^I should have (d+)$`, func(val int) {
    //     assert.Equal(T, val, initVal)
    // })
}
