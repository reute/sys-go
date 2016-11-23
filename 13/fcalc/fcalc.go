package main

import ("fmt";"os"; "bufio"; "strings"; "strconv"; "math")

type stack []float64

func (s stack) Peek() float64   {
     return s[len(s)-1] 
}

func (s *stack) Put(f float64)  {
     *s = append(*s, f) 
}

func (s *stack) Pop() float64 {
	f := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return f
}

func main() {    
    fmt.Print("Function: ")
    function := inputFunction()  
    result := calculate(function)  
    fmt.Printf("Result : %f\n", result)
}

func calculate(function []string) float64 {
    var ops stack
    vars := make(map[string]float64)
    for _, e := range(function) {
        f, err := strconv.ParseFloat(e, 64)        
        if err == nil {
            ops.Put(f)
        } else {            
            switch e {
                case "+" : 
                    op1 := ops.Pop()
                    op2 := ops.Pop() 
                    ops.Put(op1 + op2)
                case "-" :
                    op1 := ops.Pop()
                    op2 := ops.Pop() 
                    ops.Put(op1 - op2)
                case "*" :
                    op1 := ops.Pop()
                    op2 := ops.Pop() 
                    ops.Put(op1 * op2)
                case "/" :
                    op1 := ops.Pop()
                    op2 := ops.Pop() 
                    ops.Put(op2 / op1)
                case "sin" :
                    op1 := ops.Pop()
                    ops.Put(math.Sin(op1))
                case "cos" :
                    op1 := ops.Pop()
                    ops.Put(math.Cos(op1))
                case "tan":
                    op1 := ops.Pop()
                    ops.Put(math.Tan(op1)) 
                default:
                    var num float64
                    val, err := vars[e]
                    if err == true {
                        ops.Put(val) 
                    } else {
                        fmt.Printf("Please enter value for variable %s : ", e)
                        fmt.Scanf("%f", &num)
                        vars[e] = num
                        ops.Put(num)  
                    }
           } 
        }
    }
    return ops.Pop()
}

func inputFunction() []string {
    in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
    return strings.Fields(str)
}