package graphics

import "fmt"

func DemoRectangle() {
    b := MakeBorders(-15, 15, -15, 15, -20, 20)
    r1 := MakeRectangle(-10, 10, -5, 5, 'x')
//    r2 := MakeRectangle(5, 13, -10, -7, 'o')
    r2 := MakeRectangle(-15, 0, -4, 14, 'o')
//    r2 := MakeRectangle(-5, 5, -2, 2, 'o')
    //cvs1 := MakeCanvas(b)
    //cvs1.DrawRectangle(r1)
    //cvs1.DrawRectangle(r2)
    //cvs1.DrawRectangle(r1.Intersection(r2))
    //cvs1.Plot()

    r := r1.Intersection(r2)
    ds := r1.Difference(r2)
    cvs := MakeCanvas(b)
    cvs.DrawRectangle(r1)
    cvs.DrawRectangle(r2)
    cvs.DrawRectangle(r)
    for _, d := range ds {
        fmt.Println(d.Area())
        cvs.DrawRectangle(d)
    }
    cvs.Plot()
}

func DemoLine() {
    cvs := MakeCanvas(MakeBorders(-20, 20, -20, 20, -50, 50))

    l1 := MakeLine(-10, 0, '+')
    l2 := MakeLine(-12, -5, '-')

    cvs.DrawLine(l1)
    cvs.DrawLine(l2)
    cvs.DrawLine(l1.Difference(l2)...)

    cvs.Plot()
}
