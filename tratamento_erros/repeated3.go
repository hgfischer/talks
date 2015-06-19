// repeated3.go
func doAllThatThing() error {
	return callFuncErr(doSomething, doSomethingElse, doMore)
}

func callFuncErr(functions ...func() error) error {
	for _, f := range functions {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
