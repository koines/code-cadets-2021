package taxCalculation

type taxLevel struct {
	UpperBound float64
	Percentage float64
}

func checkLevels(taxLevels []taxLevel) bool {
	for idx, val := range taxLevels {
		if idx != len(taxLevels)-1 {
			if val.UpperBound >= taxLevels[idx+1].UpperBound || val.Percentage >= taxLevels[idx+1].Percentage {
				return false
			}
		}
	}

	return true
}