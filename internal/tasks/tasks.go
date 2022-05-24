package tasks

type CNameTask struct {
	Name string
}

func (t *CNameTask) Check() bool {
	return true
}

type ATask struct {
	Records []string
}

func (t *ATask) Check() bool {
	return true
}
