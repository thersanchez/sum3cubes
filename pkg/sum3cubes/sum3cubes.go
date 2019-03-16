package sum3cubes

import "fmt"

type Result struct {
	Impossible bool
	Known      bool
	Addends    [3]int
}

func (r Result) String() string {
	if r.Impossible {
		return "Impossible"
	}
	if !r.Known {
		return "Unknown"
	}
	return fmt.Sprintf("(%d, %d, %d)", r.Addends[0], r.Addends[1], r.Addends[2])
}

func Do(bound int) ([]Result, error) {
	if bound <= 0 {
		return nil, fmt.Errorf("bound must be >0, was %d", bound)
	}

	ret := make([]Result, numberOfResults(bound))
	setAllImpossibles(ret)
	setAllKnowns(ret)

	return ret, nil
}

func numberOfResults(b int) int {
	return (3 * b * b * b) + 1
}

func setAllImpossibles(r []Result) {
	for i := range r {
		if m := i % 9; m == 4 || m == 5 {
			r[i] = Result{Impossible: true}
		}
	}
}

func setAllKnowns(r []Result) {
	r[0].Known = true
	r[1].Known, r[1].Addends = true, [3]int{1, 0, 0}
	// TODO more knowns
}
