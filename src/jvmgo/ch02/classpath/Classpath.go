package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	return nil
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}
func (self *Classpath) String() (string) {
	return ""
}
