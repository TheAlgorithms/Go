package disjointSetTree


const OFFLINEMINIMUM_EXTRACT = -1

//input sequence: -1 means extract, minimum insert value should be 1
func offLineMinimum(seq []int) []int {
	//init sets
	if len(seq) == 0 {
		panic("seq must not be empty!")
	}
	if seq[0] == OFFLINEMINIMUM_EXTRACT {
		panic("first item must be insert!")
	}
	//insert set, value is index by m
	insertSets := make([]*disjointSet, 0, 0)
	//value set, value is insert value
	values := make([]*disjointSet, len(seq), cap(seq))
	n, m := 1, 0
	insertSets = append(insertSets, MakeSet(0))
	for i := range seq {
		if seq[i] == OFFLINEMINIMUM_EXTRACT {
			m ++
			insertSets = append(insertSets, MakeSet(m))
		} else {
			if seq[i] < 1 {
				panic("minimum insert value must >= 1!")
			}
			values[seq[i]] = MakeSet(seq[i])
			insertSets[m] = Union(insertSets[m], values[seq[i]])
			n++
		}
	}

	//get minimum sequence
	extractSeq := make([]int, m, m)
	for i := 1; i < n; i ++ {
		j := FindSet(values[i]).Value.(int)
		if j != m {
			extractSeq[j] = i
			for l := j + 1; l <= m; l ++ {
				if insertSets[l] != nil {
					insertSets[l] = Union(insertSets[l],insertSets[j])
					insertSets[l].Value = l
					insertSets[j] = nil
					break
				}
			}
		}
	}
	return extractSeq
}
