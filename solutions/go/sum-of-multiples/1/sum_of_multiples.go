package summultiples


func SumMultiples(limit int, divisors ...int) int {
    sumList := make(map[int]bool)
	for i := range divisors {
		if divisors[i] > limit || divisors[i] == 0 {
            continue
        }
        
        for j := 1; j < limit; j++ {
			if _, exist := sumList[j]; exist {
                continue
            }
            
            if j % divisors[i] == 0 {
                sumList[j] = true
            }
        }
    }

    ans := 0
    for k := range sumList {
        ans = ans + k
    }
    return ans
}
