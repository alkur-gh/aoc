package main

import (
    "fmt"
    "os"
    "math"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

const (
    INP string = "inp"
    ADD string = "add"
    MUL string = "mul"
    DIV string = "div"
    MOD string = "mod"
    EQL string = "eql"
)

type CPU struct {
    registers map[string]int
    input, inputRequest chan int
}

type Instruction struct {
    itype string
    args []string
}

func MakeCPU() *CPU {
    registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
    input := make(chan int)
    inputRequest := make(chan int)
    return &CPU{registers, input, inputRequest}
}

func (cpu *CPU) Execute(inst *Instruction) {
    if inst.itype == INP {
        cpu.inputRequest <- 1
        cpu.registers[inst.args[0]] = <-cpu.input
        return
    }
    a := inst.args[0]
    b, isRegister := cpu.registers[inst.args[1]]
    if !isRegister {
        b, _ = strconv.Atoi(inst.args[1])
    }
    switch inst.itype {
    case ADD:
        cpu.registers[a] += b
    case MUL:
        cpu.registers[a] *= b
    case DIV:
        if b == 0 {
            panic("division by 0")
        }
        cpu.registers[a] /= b
    case MOD:
        if b == 0 {
            panic("division by 0")
        }
        cpu.registers[a] %= b
    case EQL:
        if cpu.registers[a] == b {
            cpu.registers[a] = 1
        } else {
            cpu.registers[a] = 0
        }
    default:
        panic("unexpected type")
    }
}

func ReadInstructions(path string) []*Instruction {
    f, err := os.Open(path)
    if err != nil { panic(err) }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    insts := []*Instruction{}
    for scanner.Scan() {
        fields := strings.Fields(scanner.Text())
        insts = append(insts, &Instruction{fields[0], fields[1:]})
    }
    return insts
}

func GenerateInputs(inputs []int, ch, request, stop chan int) {
    i := 0
    for {
        select {
        case <- stop:
            return
        case <- request:
            if i >= len(inputs) {
                panic("no more input")
            }
            ch <- inputs[i]
            i++
        }
    }
}

func TryInputs(insts []*Instruction, inputs []int) {
    cpu := MakeCPU()
    stop := make(chan int)
    go GenerateInputs(inputs, cpu.input, cpu.inputRequest, stop)
    for _, inst := range insts {
        cpu.Execute(inst)
    }
    fmt.Println(inputs)
    fmt.Println(cpu.registers)
    stop <- 1
}

type Node struct {
    isValue bool
    id int
    value string
    op string
    args []*Node
    possibleValues []int
}

func MakeValueNode(id int, value string) *Node {
    return &Node{true, id, value, "", []*Node{}, []int{}}
}

func MakeOperationNode(id int, op string, args []*Node) *Node {
    return &Node{false, id, "", op, args, []int{}}
}

func sortedValues(m map[int]bool) []int {
    values := []int{}
    for k, p := range m {
        if p {
            values = append(values, k)
        }
    }
    sort.Ints(values)
    if len(values) > 5 {
        return []int{values[0], values[len(values) - 1]}
    } else {
        return values
    }
}

func (node *Node) String() string {
    if node.isValue {
        return fmt.Sprintf("\"%d:%s\"", node.id, node.value)
    } else if len(node.possibleValues) == 2 {
        //pstrs := []string{}
        //for _, p := range sortedValues(node.possibleValues) {
        //    pstrs = append(pstrs, strconv.Itoa(p))
        //}
        //return fmt.Sprintf("\"%d:%s[%d]\"", node.id, node.op, len(node.possibleValues))
        return fmt.Sprintf("\"%d:%s[%d,%d]\"", node.id, node.op, node.possibleValues[0], node.possibleValues[1])
    } else {
        return fmt.Sprintf("\"%d:%s\"", node.id, node.op)
    }
}

func (node *Node) IsInput() bool {
    return node.isValue && strings.HasPrefix(node.value, "INPUT")
}

func (node *Node) FindById(id int) *Node {
    if node.id == id {
        return node
    }
    if node.isValue {
        return nil
    }
    left := node.args[0].FindById(id)
    if left != nil {
        return left
    }
    right := node.args[1].FindById(id)
    if right != nil {
        return right
    }
    return nil
}

func rec(node *Node, visited map[string]int, ch chan string) {
    if visited[node.String()] > 0 {
        return
    }
    visited[node.String()]++
    src := node.String()
    for _, child := range node.args {
        dest := child.String()
        ch <- fmt.Sprintf("%s -> %s", src, dest)
        rec(child, visited, ch)
    }
}

func TraverseNodes(root *Node, ch chan string) {
    visited := make(map[string]int)
    rec(root, visited, ch)
    close(ch)
}

func DumpGraph(root *Node) {
    f, err := os.OpenFile("graph.gv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil { panic(err) }
    defer f.Close()
    writer := bufio.NewWriter(f)
    ch := make(chan string)
    go TraverseNodes(root, ch)

    writer.WriteString("digraph G {\n")
    writer.WriteString(fmt.Sprintf("root -> %s\n", root))

    for line := range ch {
        writer.WriteString(fmt.Sprintf("%s\n", line))
    }

    writer.WriteString("}\n")
    writer.Flush()
}

func simp_rec(node *Node, simplified map[string]*Node) *Node {
    if node.isValue {
        return node
    }
    repr := node.String()
    simp, prs := simplified[repr]
    if prs {
        return simp
    }
    left, right := simp_rec(node.args[0], simplified), simp_rec(node.args[1], simplified)
    node.args = []*Node{left, right}
    result := node
    switch node.op {
    case MUL:
        if left.isValue && right.isValue && !left.IsInput() && !right.IsInput() {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            result = MakeValueNode(node.id, strconv.Itoa(a * b))
        } else if (left.isValue && (left.value == "0")) ||
                  (right.isValue && (right.value == "0")) {
            result = MakeValueNode(node.id, "0")
        } else if left.isValue && (left.value == "1") {
            result = right
        } else if right.isValue && (right.value == "1") {
            result = left
        }
    case ADD:
        if left.isValue && right.isValue && !left.IsInput() && !right.IsInput() {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            result = MakeValueNode(node.id, strconv.Itoa(a + b))
        } else if left.isValue && (left.value == "0") {
            result = right
        } else if right.isValue && (right.value == "0") {
            result = left
        }
    case DIV:
        if left.isValue && right.isValue && !left.IsInput() && !right.IsInput() {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            result = MakeValueNode(node.id, strconv.Itoa(a / b))
        } else if right.isValue && (right.value == "1") {
            result = left
        }
    case MOD:
        if left.isValue && right.isValue && !left.IsInput() && !right.IsInput() {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            result = MakeValueNode(node.id, strconv.Itoa(a % b))
        }
    case EQL:
        if left.isValue && right.isValue && !left.IsInput() && !right.IsInput() {
            value := 0
            if left.value == right.value {
                value = 1
            }
            result = MakeValueNode(node.id, strconv.Itoa(value))
        } else if !((left.IsInput() && right.isValue) || (right.IsInput() && left.isValue)) {
            result = node
        } else {
            req := 0
            if left.IsInput() {
                req, _ = strconv.Atoi(right.value)
            } else {
                req, _ = strconv.Atoi(left.value)
            }
            if (req <= 0) || (req >= 10) {
                result = MakeValueNode(node.id, "0")
            }
        }
    }
    result.possibleValues = node.possibleValues
    simplified[repr] = result
    return result
}

func Simplify(node *Node) *Node {
    return simp_rec(node, make(map[string]*Node))
}

func PopulatePossibleValues(node *Node) *Node {
    if len(node.possibleValues) != 0 {
        return node
    }
    if node.IsInput() {
        node.possibleValues = []int{1, 9}
        return node
    }
    if node.isValue {
        v, _ := strconv.Atoi(node.value)
        node.possibleValues = []int{v, v}
        return node
    }
    left := PopulatePossibleValues(node.args[0])
    right := PopulatePossibleValues(node.args[1])
    node.args[0] = left
    node.args[1] = right
    switch node.op {
    case ADD:
        node.possibleValues = []int{
            left.possibleValues[0] + right.possibleValues[0],
            left.possibleValues[1] + right.possibleValues[1],
        }
    case MOD:
        min, max := math.MaxInt64, math.MinInt64
        for a := left.possibleValues[0]; a <= left.possibleValues[1]; a++ {
            for b := right.possibleValues[0]; b <= right.possibleValues[1]; b++ {
                v := a % b
                if v < min {
                    min = v
                }
                if v > max {
                    max = v
                }
            }
        }
        node.possibleValues = []int{min, max}
    case MUL:
        min, max := math.MaxInt64, math.MinInt64
        for a := left.possibleValues[0]; a <= left.possibleValues[1]; a++ {
            for b := right.possibleValues[0]; b <= right.possibleValues[1]; b++ {
                v := a * b
                if v < min {
                    min = v
                }
                if v > max {
                    max = v
                }
            }
        }
        node.possibleValues = []int{min, max}
    case DIV:
        min, max := math.MaxInt64, math.MinInt64
        for a := left.possibleValues[0]; a <= left.possibleValues[1]; a++ {
            for b := right.possibleValues[0]; b <= right.possibleValues[1]; b++ {
                v := a / b
                if v < min {
                    min = v
                }
                if v > max {
                    max = v
                }
            }
        }
        node.possibleValues = []int{min, max}
    case EQL:
        one, zero := false, false
        for a := left.possibleValues[0]; a <= left.possibleValues[1]; a++ {
            for b := right.possibleValues[0]; b <= right.possibleValues[1]; b++ {
                if a == b {
                    one = true
                } else {
                    zero = false
                }
            }
            if one && zero {
                break
            }
        }
        min, max := 0, 0
        if one && zero {
            min, max = 0, 1
        } else if zero {
            min, max = 0, 0
        } else {
            min, max = 1, 1
        }
        node.possibleValues = []int{min, max}
    }
    if (len(node.possibleValues) == 2) && (node.possibleValues[0] == node.possibleValues[1]) {
        v := node.possibleValues[0]
        result := MakeValueNode(node.id, strconv.Itoa(v))
        return PopulatePossibleValues(result)
    }
    return node
}

func demo() {
    path := "./files/input.txt"
    insts := ReadInstructions(path)
    registers := map[string]*Node{
        "w": MakeValueNode(0, "0"),
        "x": MakeValueNode(1, "0"),
        "y": MakeValueNode(2, "0"),
        "z": MakeValueNode(3, "0"),
    }
    id := 4
    input_count := 0
    for _, inst := range insts {
        if inst.itype == INP {
            input_count++
            value := fmt.Sprintf("INPUT-%d", input_count)
            id++
            registers[inst.args[0]] = MakeValueNode(id, value)
            continue
        }
        a := inst.args[0]
        var b *Node
        if strings.Contains("wxyz", inst.args[1]) {
            b = registers[inst.args[1]]
        } else {
            id++
            b = MakeValueNode(id, inst.args[1])
        }
        id++
        registers[a] = MakeOperationNode(id, inst.itype, []*Node{
            registers[a], b,
        })
    }

    root := Simplify(registers["z"])
    root = PopulatePossibleValues(root)
    fmt.Println("POPULATED")
    DumpGraph(root)
    //DumpGraph(Simplify(registers["z"]))
    //DumpGraph(registers["z"])
    //traverse(registers["z"])
}

func traverse(node *Node) {
    fmt.Println(node)
    if node.isValue {
        return
    }
    traverse(node.args[0])
    traverse(node.args[1])
}

func primary() {
    path := "./files/input.txt"
    insts := ReadInstructions(path)
    TryInputs(insts, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
}

func main() {
    demo()
    //primary()
}
