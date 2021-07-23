package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFiltro(t *testing.T) {
	RunTest(t, []string{"a1bG93foEìr92L"}, "", "b\nf\nL\n")
	RunTest(t, []string{"È3e34$33òñe3"}, "", "e\n")
	RunTest(t, []string{"3eÉ311d$4"}, "", "e\nd\n")
	RunTest(t, []string{"k⒑e2uWe²23d"}, "", "u\nd\n")
}

func RunTest(t *testing.T, args []string, stdin string, stdout string) { //si possono fare slice per stdin e stdout ma attento ai \n finali, se si vuole testare anche il tempo bisognerebbe fare controlli aggiuntivi
	
	os.Args = append([]string{""},args...)
	
	readStdin, writeStdin, _ := os.Pipe() //pipe written to writeStdin can be read from readStdin
	readStdin, os.Stdin = os.Stdin, readStdin //now we can write program input to writeStdin
	
	readStdout, writeStdout, _ := os.Pipe() //pipe written to writeStdout can be read from readStdout
	writeStdout, os.Stdout = os.Stdout, writeStdout //now we can read program output from readStdout

	fmt.Fprint(writeStdin, stdin)

	main()

	//revert correct stdin and stdout files
	readStdin, os.Stdin = os.Stdin, readStdin
	writeStdout, os.Stdout = os.Stdout, writeStdout
	readStdin.Close()
	writeStdin.Close()
	writeStdout.Close()

	result, _ := io.ReadAll(readStdout)

	readStdout.Close()
	
	if stdout != string(result) {
		t.Error(fmt.Sprintf("\nArgs: %s\nInput: %s\nExpected result:\n-------\n%s\n-------\nGiven result:\n-------\n%s\n-------\n", strings.Join(args, " "), stdin, stdout, string(result)))
	}
}