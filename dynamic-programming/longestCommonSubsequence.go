//The function takes two parameters, both strings. It then finds the longest common subsequence 
//of both the strings and then outputs the length of common subsequence.




package longestCommonSubsequence

import "fmt"

func Max(x,y int64) int64 {
	if x>=y {
		return x
	}
	return y
}

func LCS(a string,b string) int64 {
	var n int = len(a)
	var m int = len(b)
	LCS := make([][]int64,n+1)
	for i := range LCS {
		LCS[i] = make([]int64,m+1)
	}
	for i := 0;i<=n;i++ {
		for j:=0;j<=m;j++ {
			if i==0||j==0 {
				LCS[i][j]=0;
			} else {
				if a[i-1]==b[j-1] {
					LCS[i][j]=1+LCS[i-1][j-1];
				} else {
					LCS[i][j]=Max(LCS[i-1][j],LCS[i][j-1])
				}
			}
		}
	}
	return LCS[n][m]
}
