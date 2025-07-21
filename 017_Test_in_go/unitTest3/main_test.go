package main

import "testing"

func TestAdd(t *testing.T) {
    testCase :=  []struct{
        name string
        num int
        expected int
    }{
        {"case 5",5,120},
        {"case 2",2,2},
        {"case -1",-1,0},
    }

    for _,tc := range testCase{
        t.Run(tc.name, func(t *testing.T) {
            result := Factorial(tc.num)
            expectResult := tc.expected
            if result != expectResult{
                t.Errorf("Factorial(%d) = %d is wrong , correct is %d",
                tc.num,result,expectResult )
            }
        })
    }
}