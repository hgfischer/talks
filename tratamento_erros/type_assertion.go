if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
	time.Sleep(10 * time.Second)
	continue
}

if err != nil {
	log.Fatal(err)
}

//...

switch netErr := err.(type) {
net.Error:
	if netErr.Temporary() {
		//...
	} else if netErr.Timeout() {
		//...
	}
default:
	//...
}
