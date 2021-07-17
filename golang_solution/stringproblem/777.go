package main

func canTransform(start string, end string) bool {
    a := getExtract(start)
    b := getExtract(end)
    if a != b {
        return false 
    }
    
    t := 0
    for i := 0; i < len(start); i++ {
        if start[i] == 'L' {
            for end[t] != 'L' {
                t++
            }
            if (i < t) {
                return false
            }
            t++
        }
    }
    t = 0
    for i := 0; i < len(start); i++ {
        if start[i] == 'R' {
            for end[t] != 'R' {
                t++
            }
            if (i > t) {
                return false
            }
            t++
        }
    }
    return true
}

func getExtract(str string) string { 
    res := []byte{}
    for i := range str {
        if str[i] != 'X' {
            res = append(res, str[i])
        }
    }
    return string(res)
}