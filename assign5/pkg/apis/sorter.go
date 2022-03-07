package apis

type Sorter interface {
	Sort(*MutexFatRateRankMachine)
}

type BubbleSort struct {
}

func (b BubbleSort) Sort(machine *MutexFatRateRankMachine) {
	p := machine.Scoreboard
	m := machine.Registrants
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			if p[j].CurrentFatRate < p[i].CurrentFatRate {
				(m[p[i].Name]).rank, (m[p[j].Name]).rank = j, i
				p[i], p[j] = p[j], p[i]
			}
		}
	}

}

type QuickSort struct {
}

func (q QuickSort) Sort(machine *MutexFatRateRankMachine) {
	quickSort(machine, 0, len(machine.Scoreboard)-1)
}

func quickSort(machine *MutexFatRateRankMachine, start, end int) {
	if start < end {
		var p int
		p = partition(machine, start, end)
		quickSort(machine, start, p-1)
		quickSort(machine, p+1, end)
	}
}

func partition(machine *MutexFatRateRankMachine, start, end int) int {
	p := machine.Scoreboard
	m := machine.Registrants
	pivot := p[end]
	i := start
	for j := start; j < end; j++ {
		if p[j].CurrentFatRate < pivot.CurrentFatRate {
			(m[p[i].Name]).rank, (m[p[j].Name]).rank = j, i
			p[i], p[j] = p[j], p[i]
			i++
		}
	}
	(m[p[i].Name]).rank, (m[p[end].Name]).rank = end, i
	p[i], p[end] = p[end], p[i]

	return i
}
