package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/patdhlk/fury/evaluator"
	"github.com/patdhlk/fury/lexer"
	"github.com/patdhlk/fury/object"
	"github.com/patdhlk/fury/parser"
)

const FURY_PROMPT = "fury ==> "

func Start(in io.Reader, out io.Writer) {
	io.WriteString(out, FURY_LOGO)

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(FURY_PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const FURY_LOGO = `_________________________________________
        __    ___    _     ____   __   __
\    ___) |  |   |  | |    \  (  (  )  ) 
 |  (__   |  |   |  | | ()  )  \  \/  /  
 |   __)  |  |   |  | |    /    \    /   
 |  (     |   \_/   | | |\ \     )  /    
/    \_____\       /__| |_\ \___/  (_____
_________________________________________

`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Fatal fury error!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
