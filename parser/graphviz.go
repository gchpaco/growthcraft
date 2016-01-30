package parser

import "io"
import "fmt"

type Transition interface {
	Visit(io.Writer) error
}

func escape(in string) string {
	return "\"" + in + "\""
}

func (fluid *Fluid) ID() string {
	return escape(fluid.Name)
}

func (item *Object) Render() string {
	if item.Amount == 1 {
		return item.FullName()
	}
	return fmt.Sprintf("%s x %d", item.FullName(), item.Amount)
}

func (residue *Residue) Render() string {
	if residue.Amount == 1 {
		return fmt.Sprintf("%s (%2.0f%%)", residue.FullName(), residue.Pomace*100.0)
	}
	return fmt.Sprintf("%s x %d (%2.0f%%)", residue.FullName(),
		residue.Amount, residue.Pomace*100.0)
}

func (tr *BrewTransition) Visit(w io.Writer) error {
	label := tr.Item.Render()
	if tr.Residue != nil {
		label += "\\nOut: " + tr.Residue.Render()
	}
	fmt.Fprintf(w, "\t%s->%s[arrowhead=diamond, label=\"%s\", style=dashed];\n",
		tr.InputFluid.ID(), tr.OutputFluid.ID(), label)
	return nil
}

// Super dirty hacks
var bannedList = map[string]bool{
	"grc.grapewine2":  true,
	"grc.grapewine3":  true,
	"grc.applecider2": true,
	"grc.applecider3": true,
	"grc.hopale2":     true,
	"grc.hopale3":     true,
	"grc.lager2":      true,
	"grc.lager3":      true,
	"grc.ricesake2":   true,
	"grc.ricesake3":   true,
	"grc.honeymead2":  true,
	"grc.honeymead3":  true,
}

func (tr *FermentTransition) Visit(w io.Writer) error {
	if bannedList[tr.InputFluid.Name] && bannedList[tr.OutputFluid.Name] {
		return nil
	}
	label := tr.Item.Render()
	fmt.Fprintf(w, "\t%s->%s[arrowhead=normal, label=\"%s\", style=solid];\n",
		tr.InputFluid.ID(), tr.OutputFluid.ID(), label)
	return nil
}

func (tr *PressingTransition) Visit(w io.Writer) error {
	label := tr.Item.Render()
	if tr.Residue != nil {
		label += "\\nOut: " + tr.Residue.Render()
	}
	fmt.Fprintf(w, "\tnull->%s[arrowhead=dot, label=\"%s\", style=dotted];\n",
		tr.Fluid.ID(), label)
	return nil
}
