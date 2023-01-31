package week2

type Student struct {
	name string
	id   int
	gpa  float32
}

func (s *Student) getname() string {
	return s.name
}

func (s *Student) setname(name string) {
	s.name = name
}

func (s *Student) getid() int {
	return s.id
}

func (s *Student) setid(id int) {
	s.id = id
}

func (s *Student) getgpa() float32 {
	return s.gpa
}

func (s *Student) setgpa(gpa float32) {
	s.gpa = gpa
}
