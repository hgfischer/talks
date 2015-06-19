// repeated2.go
func doAllThatThing() error {
	if err := doSomething(); err != nil {
		return err
	}
	if err = doSomethingElse(); err != nil {
		return err
	}
	if err = doMore(); err != nil {
		return err
	}
	return nil
}