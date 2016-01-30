package parser

import "encoding/json"
import "io"
import "fmt"

type Fluid struct {
	Comment string
	Name    string
	Amount  int
}

type Object struct {
	Comment string
	ModID   string `json:"mod_id"`
	Name    string
	Amount  int
	Meta    int
}

func (o *Object) FullName() string {
	lookup := fmt.Sprintf("%s:%s:%d", o.ModID, o.Name, o.Meta)
	if name, ok := Items[lookup]; ok {
		return name
	}
	return lookup
}

type Residue struct {
	*Object
	Pomace float64
}

type BrewTransition struct {
	Comment     string
	Item        *Object
	InputFluid  *Fluid `json:"input_fluid"`
	OutputFluid *Fluid `json:"output_fluid"`
	Residue     *Residue
	Time        int
}

type FermentTransition struct {
	Comment     string
	Item        *Object
	InputFluid  *Fluid `json:"input_fluid"`
	OutputFluid *Fluid `json:"output_fluid"`
	Time        int
}

type PressingTransition struct {
	Comment string
	Item    *Object
	Fluid   *Fluid
	Residue *Residue
	Time    int
}

func DecodeBrewing(r io.Reader) ([]*BrewTransition, error) {
	dec := json.NewDecoder(r)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return nil, err
	}

	results := []*BrewTransition{}
	for dec.More() {
		bt := new(BrewTransition)
		err := dec.Decode(bt)
		if err != nil {
			return results, err
		}

		results = append(results, bt)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return results, err
	}

	return results, nil
}

func DecodeFerment(r io.Reader) ([]*FermentTransition, error) {
	dec := json.NewDecoder(r)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return nil, err
	}

	results := []*FermentTransition{}
	for dec.More() {
		ft := new(FermentTransition)
		err := dec.Decode(ft)
		if err != nil {
			return results, err
		}

		results = append(results, ft)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return results, err
	}

	return results, nil
}

func DecodePressing(r io.Reader) ([]*PressingTransition, error) {
	dec := json.NewDecoder(r)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return nil, err
	}

	results := []*PressingTransition{}
	for dec.More() {
		pt := new(PressingTransition)
		err := dec.Decode(pt)
		if err != nil {
			return results, err
		}

		results = append(results, pt)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return results, err
	}

	return results, nil
}
