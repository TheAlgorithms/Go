package main


func NaiveApproach(n int) bool {
	if n==1{
		return false
	}
	for i:=2; i < n ; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func PairApproach(n int) bool {
	if n==1{
		return false
	}
	for i:=2;i*i<=n;i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//goos: windows
//goarch: amd64
//BenchmarkNaiveApproach-12       660980034                1.81 ns/op
//BenchmarkPairApproach-12        664641356                1.82 ns/op









