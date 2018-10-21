func minSwapsCouples(row []int) int {
    couples := len(row)
    count := 0
    for i:=0; i<couples; i++ {
        if row[i]%2==1 {
            row[i] -= 1
        }
    }

    for i:=0; i<couples; i+=2 {
        if row[i] == row[i+1] {
            continue
        }

        for j:=i+2; j<couples; j++ {
            if row[i] == row[j] {
                row[i+1], row[j] = row[j], row[i+1]
                count++
            }
        } 
    }
    return count
}
