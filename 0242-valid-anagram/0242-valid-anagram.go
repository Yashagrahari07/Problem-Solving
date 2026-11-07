func isAnagram(s string, t string) bool {
    if (len(s) != len(t)){
        return false
    }

    arr := make([]int, 26)

    for i:=0; i<len(s); i++{
        arr[s[i]-'a']++
        arr[t[i]-'a']--
    }

    for _,char := range arr{
        if(char != 0){
            return false
        }
    }

    return true
}