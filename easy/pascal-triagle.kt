class Solution {
    fun generate(numRows: Int): List<List<Int>> {
        val result: MutableList<List<Int>> = mutableListOf()
        
        for (i in 0 until numRows) {
            val tmpRow = mutableListOf<Int>()

            for (j in 0..i) {
                if (j==0 || j==i) {
                    tmpRow.add(1)
                } else {
                    val tmpNum = result[i-1][j-1] + result[i-1][j]
                    tmpRow.add(tmpNum)
                }
            }
            result.add(tmpRow)
        }
        
        return result
    }
}
