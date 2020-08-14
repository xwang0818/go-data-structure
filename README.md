# go-data-structure

## Interview
### Flow
 1. Listen
 2. Ask Clarifications
 3. Brute Force
 4. Optimization
 5. Write outline comments //
 6. Implement
 7. Test

### Inspection
1. `package main`
2. `import "xxx"`
3. `func main () {}`
4. Return Value: `int` or `(int, int)`
5. All vars are initialized
6. array index out of bound
7. validate param, args
8. Function names match
9. node vs root
10. `list.Front().Value.(*ListNode)`

## Data structure types
```
// 1D Array
    res := make([]int, size)
    res = append(res, element1, element2)
    res = append(res, array…)

// 2D Array
    // make([]type, size, cap)
    var visited [][]bool
    visited = make([][]bool, size)
    for i := range visited {
        visited[i] = make([]bool, size)
    }

// Stack - LIFO (Last In First Out)
    arr = append(arr, num) // push
    num := arr[len(arr)-1] // pop
    arr = arr[:len(arr)-1]

// Hash Table/Map
    dict := make(map[int]int)
    dict[key] = value
    delete(dict, key)

// Function on struct
    /*
    type my_type struct { }

    func (m my_type) my_func() int {
        //code
    }
    */

    type Rectangle struct {
      length, width int
    }

    func (r Rectangle) Area() int {
        return r.length * r.width
    }

    func main() {
        r1 := Rectangle{4, 3}
        fmt.Println("Rectangle is: ", r1)
        fmt.Println("Rectangle area is: ", r1.Area())
    }


// Sort
    import "sort"

    // To sort chars in a string, the string needs to be converted to []bytes
    func main()
        str := "cdsafkleiowm"
        arr := []byte(str)
        sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
        str = string(arr)
    }


    // sort any struct array
    type Person struct {
    	Name string
    	Age  int
    }

    func main() {
    	people := []Person{
    		{"Bob", 31},
    		{"John", 42},
    		{"Michael", 17},
    		{"Jenny", 26},
    	}
        peter := Person{"Peter", 29}
        people = append(people, peter)

        sort.Slice(people, func (a, b int) bool { return people[a].Age < people[b].Age })
    	fmt.Println(people)
    }


    // implementing the sort interface
    type Person struct {
    	Name string
    	Age  int
    }

    func (p Person) String() string {
    	return fmt.Sprintf("%s: %d", p.Name, p.Age)
    }

    type ByAge []Person
    func (a ByAge) Len() int           { return len(a) }
    func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
    func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

    func main() {
    	people := []Person{
    		{"Bob", 31},
    		{"John", 42},
    		{"Michael", 17},
    		{"Jenny", 26},
    	}
    	fmt.Println(people)

    	sort.Sort(ByAge(people))
    	fmt.Println(people)

    	sort.Slice(people, func(i, j int) bool {
    		return people[i].Age < people[j].Age
    	})
    	fmt.Println(people)
    }
```
