package singleton 
  
type singleton struct { 
    count int 
}

type Singleton interface { 
    AddOne() int 
} 

var instance *singleton 

func GetInstance() Singleton { 
	if instance == nil {
		instance = new(singleton)
	}
    return instance 
} 

func (s *singleton) AddOne() int { 
	s.count++
	return s.count
} 