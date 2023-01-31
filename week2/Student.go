package week2

type Student struct {
	name string
	id   int
	gpa  float32
}

func (s *Student) Getname() string {
	return s.name
}

func (s *Student) Setname(name string) {
	s.name = name
}

func (s *Student) Getid() int {
	return s.id
}

func (s *Student) Setid(id int) {
	s.id = id
}

func (s *Student) Getgpa() float32 {
	return s.gpa
}

func (s *Student) Setgpa(gpa float32) {
	s.gpa = gpa
}
