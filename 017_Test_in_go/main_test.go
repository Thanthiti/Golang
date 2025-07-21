package main

import "testing"

func TestAdd(t *testing.T) {
    // go test 
    // go test -v to show verbose where tasecase is fails
    testCase :=  []struct{
        name string
        a,b int
        expected int
    }{
        {"Add positive numbers",2,3,5},
        {"Add nagative numbers",-1,-2,-5},
        {"Add zero",0,0,0},
    }

    for _,tc := range testCase{
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a,tc.b)
            expectResult := tc.expected
            if result != expectResult{
                t.Errorf("Add(%d , %d) = %d is wrong , correct is %d",
                tc.a,tc.b,result,expectResult )
            }
        })
    }
}