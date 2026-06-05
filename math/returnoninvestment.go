// returnoninvestment.go
// description: Return on Investment (ROI)
// details:
// Return on Investment (ROI) measures the profitability of an investment
// relative to its cost, expressed as a percentage.
// Formula: ROI = (Gain - Cost) / Cost * 100
// reference: https://www.investopedia.com/terms/r/returnoninvestment.asp
// author(s) [InukaWijerathna](https://github.com/InukaWijerathna)
// see returnoninvestment_test.go

package math

import (
	"errors"
)

var ErrNonPositiveCost = errors.New("costOfInvestment must be greater than 0")

// ReturnOnInvestment calculates the return on investment as a percentage.
func ReturnOnInvestment(gainFromInvestment, costOfInvestment float64) (float64, error) {
	if costOfInvestment <= 0 {
		return 0, ErrNonPositiveCost
	}
	return (gainFromInvestment - costOfInvestment) / costOfInvestment * 100, nil
}
