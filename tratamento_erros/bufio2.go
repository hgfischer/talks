scanner := bufio.NewScanner(input)
for {
	token, err := scanner.Scan()
	if err != nil {
		return err
	}
	// processa token
}
