package fa

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func (d *DFA) Save() error {
	file, err := os.Create("temp/dfa.json")
	if err != nil {
		return err
	}
	defer file.Close()

	marsh, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = file.Write(marsh)
	if err != nil {
		return err
	}

	return nil
}

func Read() (*DFA, error) {
	marsh, err := os.ReadFile("temp/dfa.json")
	if err != nil {
		return nil, err
	}

	dfa := &DFA{}
	err = json.Unmarshal(marsh, dfa)
	if err != nil {
		return nil, err
	}

	return dfa, nil
}

func (d *DFA) Show(filename string) error {
	dot := filename + ".dot"
	png := filename + ".png"

	file, err := os.Create(dot)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("digraph DFA {\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("rankdir=LR;\n")
	if err != nil {
		return err
	}

	if len(d.States) > 0 {
		_, err = fmt.Fprintf(file, "start [shape=point];\n")
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(file, "start -> \"%s\" [label=\"\"];\n", d.States[0])
		if err != nil {
			return err
		}
	}

	for _, s := range d.States {
		if s.Last {
			_, err := fmt.Fprintf(file, "\"%s\" [peripheries=2 label=\"%s\"];\n", s, s.String())
			if err != nil {
				return err
			}
		} else {
			_, err := fmt.Fprintf(file, "\"%s\" [label=\"%s\"];\n", s, s.String())
			if err != nil {
				return err
			}
		}

	}

	for sStr, transitions := range d.Tran {
		for a, uPtr := range transitions {
			_, err := fmt.Fprintf(file, "\"%s\" -> \"%s\" [label=\"%c\"];\n", sStr, uPtr, a)
			if err != nil {
				return err
			}
		}
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tpng", dot, "-o", png)
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("open", png)
	return cmd.Run()
}

func (d *NFA) Show(filename string) error {
	dot := filename + ".dot"
	png := filename + ".png"

	file, err := os.Create(dot)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("digraph DFA {\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("rankdir=LR;\n")
	if err != nil {
		return err
	}

	for _, s := range d.States {
		if s.Start {
			_, err = fmt.Fprintf(file, "start [shape=point];\n")
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(file, "start -> \"%s\" [label=\"\"];\n", s)
			if err != nil {
				return err
			}
		}
	}

	for _, s := range d.States {
		if s.Last {
			_, err := fmt.Fprintf(file, "\"%s\" [peripheries=2 label=\"%s\"];\n", s, s.String())
			if err != nil {
				return err
			}
		} else {
			_, err := fmt.Fprintf(file, "\"%s\" [label=\"%s\"];\n", s, s.String())
			if err != nil {
				return err
			}
		}

	}

	for sStr, transitions := range d.Tran {
		for a, uPtrs := range transitions {
			for _, uPtr := range uPtrs {
				_, err := fmt.Fprintf(file, "\"%s\" -> \"%s\" [label=\"%c\"];\n", sStr, uPtr, a)
				if err != nil {
					return err
				}
			}
		}
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tpng", dot, "-o", png)
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("open", png)
	return cmd.Run()
}
