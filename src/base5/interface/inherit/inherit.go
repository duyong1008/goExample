package inherit

import "fmt"

type Humaner interface { //子集
	Sayhi()
}

type Personer interface { //超集
	Humaner //匿名字段，继承了sayhi()
	Sing(lrc string)
}

type Student struct {
	Name string
	Id   int
}

//Student实现了sayhi()
func (this *Student) Sayhi() {
	fmt.Printf("Student[%s,%d] sayhi\n", this.Name, this.Id)
}

func (this *Student) Sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}
