package groupAnagrams

func groupAnagrams(strs []string) [][]string {
    mp:= make(map[string][]string)
    for _, val := range strs {
        key := sortString(val)
        mp[key] = append(mp[key], val)
    }
    
    var res [][]string
    for _, words := range mp {
        res = append(res, words)
    }
    return res
}

func sortString(str string) string {
    runes := []rune(str)
    sort.Slice(runes, func(a, b int) bool {
        return runes[a] < runes[b]
    })
    return string(runes)
}
