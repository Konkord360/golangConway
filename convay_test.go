package main

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCountNeighboursSliceReturnsEightForAFieldWithAllNeighboursAlive(t *testing.T) {
    a := [5][5]Field{
        {Field{false, "0", 0}, Field{true,"1", 0}, Field{false,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"1", 0}, Field{true,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"4", 0}, Field{false,"5", 0}, Field{true,"6", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"7", 0}, Field{true,"8", 0}, Field{true,"9", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"1", 0}, Field{false,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
    }

    var s = make([][]Field, 3)
    for i := 1; i <= 3 ; i++  {
        s[i - 1] = a[i][1:4]
    }
    count := CountNeighboursSlice(s)
    expected := 8
    if count != expected {
        t.Errorf("Excpected count to be %d but was %d", expected, count)
    }
}

func TestCountNeighboursSliceCountsCorrectlyNeighboursFromTopAndBottom(t *testing.T) {
    a := [5][5]Field{
        {Field{false, "0", 0}, Field{true,"1", 0}, Field{false,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
        {Field{true, "0", 0}, Field{true,"1", 0}, Field{true,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"4", 0}, Field{false,"5", 0}, Field{true,"6", 0}, Field{false, "0", 0}},
        {Field{true, "0", 0}, Field{true,"7", 0}, Field{true,"8", 0}, Field{true,"9", 0}, Field{false, "0", 0}},
        {Field{false, "0", 0}, Field{true,"1", 0}, Field{false,"2", 0}, Field{true,"3", 0}, Field{false, "0", 0}},
    }
    //s := a[1:4]
    var s = make([][]Field, 3)
    for i := 1; i <= 3 ; i++  {
        s[i - 1] = a[i][0:3]
    }
    count := CountNeighboursSlice(s)
    expected := 6
    if count != expected {
        t.Errorf("Excpected count to be %d but was %d", expected, count)
    }

}
