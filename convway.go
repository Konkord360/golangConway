package main

import (
    "fmt"
    "time"
)

type Field struct {
    alive bool
    display string
    neighbours int
}

const x = 30
const y = 100
var Alive = "O"
var Dead = " "
func main() {
    clearTerm()
    var a[x][y] Field

    var wholeAraySlice = make([][]Field, len(a))
    for i := range wholeAraySlice {
        wholeAraySlice[i] = a[i][:]
    }

    for i := range a {
        for j := range a[i] {
            wholeAraySlice[i][j] = Field{false, Dead, 0}
        }
    }

    setStartingPattern(wholeAraySlice)
    DisplayArray(wholeAraySlice)
    time.Sleep(200 * time.Millisecond)

    for t := 0; t < 500; t = t + 1 {
        for i := 1; i < x - 1; i++  {
            for j := 1; j < y - 2; j++ {
                var s = make([][]Field, 3)
                for k := 0; k < 3; k++ {
                    s[k] = a[i-1+k][j-1:j+2]
                }
                a[i][j].neighbours = CountNeighboursSlice(s)
            }
        }

        for i := 1; i < x - 1; i++  {
            for j := 1; j < y - 1; j++ {
                if a[i][j].neighbours < 2 {
                    a[i][j].alive = false
                    a[i][j].display = Dead
                }

                if a[i][j].neighbours > 3 {
                    a[i][j].alive = false
                    a[i][j].display = Dead
                }

                if a[i][j].neighbours == 3 {
                    a[i][j].alive = true
                    a[i][j].display = Alive
                }
            }
        }
        clearTerm()
        DisplayArray(wholeAraySlice)
        time.Sleep(50 * time.Millisecond)
    }
    fmt.Printf("\033[?25h");
}

func setStartingPattern(a[][] Field) {

    //---++
    //---++
    //-+++-
    //--+--
    a[15][50].alive = true
    a[15][50].display = Alive
    a[15][51].alive = true
    a[15][51].display = Alive

    a[16][50].alive = true
    a[16][50].display = Alive
    a[16][51].alive = true
    a[16][51].display = Alive

    a[17][48].alive = true
    a[17][48].display = Alive
    a[17][49].alive = true
    a[17][49].display = Alive
    a[17][50].alive = true
    a[17][50].display = Alive
    
    a[18][49].alive = true
    a[18][49].display = Alive
}

func CountNeighboursSlice(a[][] Field) int {
    count := 0
    for i, b := range a {
        for j, c := range b {
            if i == 1 && j == 1 {
                continue;
            }
            if c.alive {
                count++
            }
        }
    }
    return count
}


func DisplayArray(a[][] Field) {
    for i := range a {
        for j := range a[i] {
            fmt.Print(a[i][j].display)                
        }
        fmt.Println()
    }
}

func clearTerm() {
    fmt.Printf("\033c"); // clears the terminal
    fmt.Printf("\033[?25l");
}
