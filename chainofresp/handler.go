package chainofresp

//handler

type Department interface{
     execute(*patient)
	 SetNext(Department)
}

func Executor(department Department,patient *patient){
    department.execute(patient)
}