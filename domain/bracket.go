package domain

type (
	bracketRange struct {
		from float64
		to   float64
	}
)

func init() {

	//Policies and regulations are subject to change
	minimumDownPaymentPercentage = make(map[bracketRange]float64)
	minimumDownPaymentPercentage[bracketRange{from: 0, to: 500000}] = 5
	minimumDownPaymentPercentage[bracketRange{from: 500001, to: 1000000}] = 10
	minimumDownPaymentPercentage[bracketRange{from: 1000001, to: -1}] = 20

	insuranceRateIndex = make(map[bracketRange]float64)
	insuranceRateIndex[bracketRange{from: 5, to: 9.99}] = 0.04
	insuranceRateIndex[bracketRange{from: 10, to: 14.99}] = 0.031
	insuranceRateIndex[bracketRange{from: 15, to: 19.99}] = 0.028
	insuranceRateIndex[bracketRange{from: 20, to: -1}] = 0
}

func (b bracketRange) withinRange(v float64) bool {
	return (b.isUpperLimit() && v >= b.from) || (b.from <= v && v <= b.to)
}

func (b bracketRange) isUpperLimit() bool {
	return b.to < 0
}
