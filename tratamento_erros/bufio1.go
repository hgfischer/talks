scanner := bufio.NewScanner(input)
for scanner.Scan() {
	token := scanner.Text()
	// processa token
}
if err := scanner.Err(); err != nil {
	// processa erro
}
