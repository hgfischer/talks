// repeated1.go
func doAllThatThing() error {
	err := doSomething()
	if err != nil {
		return err
	}
	err = doSomethingElse()
	if err != nil {
		return err
	}
	err = doMore()
	if err != nil {
		return err
	}
	return nil
}