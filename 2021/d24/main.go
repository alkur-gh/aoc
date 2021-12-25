package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
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
}

func MakeValueNode(id int, value string) *Node {
    return &Node{true, id, value, "", []*Node{}}
}

func MakeOperationNode(id int, op string, args []*Node) *Node {
    return &Node{false, id, "", op, args}
}

func (node *Node) String() string {
    if node.isValue {
        return fmt.Sprintf("\"%d:%s\"", node.id, node.value)
    } else {
        return fmt.Sprintf("\"%d:%s\"", node.id, node.op)
    }
}

func (node *Node) IsInput() bool {
    return node.isValue && strings.HasPrefix(node.value, "INPUT")
}

func (node *Node) Copy() *Node {
    return &Node{node.isValue, node.id, node.value, node.op, node.args}
}


func rec(node *Node, visited map[string]int, ch chan string) {
    if visited[node.String()] == 2 {
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

func simp_rec(node *Node, visited map[string]bool) *Node {
    if visited[node.String()] {
        fmt.Printf("HIT VISITED ON: %s\n", node)
        return node
    }

    if node.isValue {
        return node
    }
    visited[node.String()] = true
    left, right := simp_rec(node.args[0], visited), simp_rec(node.args[1], visited)
    node.args = []*Node{left, right}
    switch node.op {
    case MUL:
        if left.isValue && right.isValue {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            return MakeValueNode(node.id, strconv.Itoa(a * b))
        }
        if (left.isValue && (left.value == "0")) ||
           (right.isValue && (right.value == "0")) {
            node.isValue = true
            node.value = "0"
            node.args = []*Node{}
        }
        if left.isValue && (left.value == "1") {
            return right
        }
        if right.isValue && (right.value == "1") {
            return left
        }
        return node
    case ADD:
        if left.isValue && right.isValue {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            return MakeValueNode(node.id, strconv.Itoa(a + b))
        }
        if left.isValue && (left.value == "0") {
            return right
        }
        if right.isValue && (right.value == "0") {
            return left
        }
        return node
    case DIV:
        if left.isValue && right.isValue {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            return MakeValueNode(node.id, strconv.Itoa(a / b))
        }
        if right.isValue && (right.value == "1") {
            return left
        }
        return node
    case MOD:
        if left.isValue && right.isValue {
            a, _ := strconv.Atoi(left.value)
            b, _ := strconv.Atoi(right.value)
            return MakeValueNode(node.id, strconv.Itoa(a % b))
        }
        return node
    case EQL:
        //if left.isValue && right.isValue {
        //    value := 0
        //    if left.value == right.value {
        //        value = 1
        //    }
        //    return MakeValueNode(node.id, strconv.Itoa(value))
        //}
        if !((left.IsInput() && right.isValue) || (right.IsInput() && left.isValue)) {
            return node
        }

        req := 0
        if left.IsInput() {
            req, _ = strconv.Atoi(right.value)
        } else {
            req, _ = strconv.Atoi(left.value)
        }
        if (req <= 0) || (req >= 10) {
            return MakeValueNode(node.id, "0")
        }
        return node
    default:
        return node
    }
}

func Simplify(node *Node) *Node {
    return simp_rec(node, make(map[string]bool))
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
    for _, inst := range insts[:18] {
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

//    DumpGraph(Simplify(registers["z"]))
    DumpGraph(registers["z"])
}

func primary() {
    path := "./files/input.txt"
    insts := ReadInstructions(path)
    TryInputs(insts, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
}

func main() {
    demo()
    //primary()
    //for i0 := 9; i0 >= 1; i0-- {
    //for i1 := 9; i1 >= 1; i1-- {
    //for i2 := 9; i2 >= 1; i2-- {
    //for i3 := 9; i3 >= 1; i3-- {
    //for i4 := 9; i4 >= 1; i4-- {
    //for i5 := 9; i5 >= 1; i5-- {
    //for i6 := 9; i6 >= 1; i6-- {
    //for i7 := 9; i7 >= 1; i7-- {
    //for i8 := 9; i8 >= 1; i8-- {
    //for i9 := 9; i9 >= 1; i9-- {
    //for i10 := 9; i10 >= 1; i10-- {
    //for i11 := 9; i11 >= 1; i11-- {
    //for i12 := 9; i12 >= 1; i12-- {
    //for i13 := 9; i13 >= 1; i13-- {
    //    go TryInputs(insts, []int{i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13})
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
    //}
}
