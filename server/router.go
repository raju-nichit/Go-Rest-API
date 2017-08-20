package server

//Init -- server setting
func Init() {
	r := NewRouter()
	r.Run(":8090")
}
