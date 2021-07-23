package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

type Location struct{
    x,y int
}

func Newloc(x, y int) Location{
    var pos Location
    pos.x = x
    pos.y = y
    return pos
}

func Move(loc Location, direction rune) (newloc Location, err error){
    if direction == 'u' || direction == 'U'{
        loc.y++
    }else if direction == 'd' || direction == 'D'{
        loc.y--
    }else if direction == 'r' || direction == 'R'{
        loc.x++
    }else if direction == 'l' || direction == 'L'{
        loc.x--
    }else{
        loc.x = 0
        loc.y = 0
        err = fmt.Errorf("unknown direction")
    }
    newloc = loc
    return
}

func (loc Location) String() string {
    return fmt.Sprintf("(%d,%d)",loc.x,loc.y)
}

func RandomWalk(start Location, length uint) []Location{
    var random,l int
    var locations []Location
    locations = append(locations,start)
    l = int(length)
    rand.Seed(time.Now().UnixNano())
    for i:=0 ; i<l ; i++{
        random = rand.Intn(4)
        if random == 0{
            temp,_ := Move(start,'u')
            locations = append(locations,temp)
            start = temp
        }
        if random == 1{
            temp,_ := Move(start,'r')
            locations = append(locations,temp)
            start = temp
        }
        if random == 2{
            temp,_ := Move(start,'d')
            locations = append(locations,temp)
            start = temp
        }
        if random == 3{
            temp,_ := Move(start,'l')
            locations = append(locations,temp)
            start = temp
        }
    }
    return locations
}

func ContainsLoop(walk []Location) (found bool){

    for i:=0 ; i<len(walk)-1; i++{
        for j:= i+1 ; j<len(walk) ; j++{
            if walk[i].x == walk[j].x && walk[i].y == walk[j].y{
                return true
            }
        }
    }
    return false
}


func main(){
    var bot Location
    if len(os.Args) < 4{
        fmt.Errorf("too few arguments")
        return
    }
    x,_ := strconv.Atoi(os.Args[1])
    y,_ := strconv.Atoi(os.Args[2])
    temp,_ := strconv.Atoi(os.Args[3])
    lung := uint(temp)
    bot.x = x
    bot.y = y
    fmt.Println(lung)
    camminata := RandomWalk(bot,lung)
    if ContainsLoop(camminata){
        fmt.Println("loop")
    }else{
        fmt.Println("no loops")
    }
    if (len(os.Args) == 5 && os.Args[5] == "-v"){
        fmt.Println(camminata)
    }
}
