class Solution {
    public int[] queryResults(int limit, int[][] queries) {
        Map<Integer, Integer> mapColor = new HashMap<>();
        Map<Integer, Integer> countColor = new HashMap<>();
        int n = queries.length();
        int[] result = new int[n];
        int i = 0;
        while (i < n){
            int[] query = queries[i];
            int newColor = query[1];
            int currColor = mapColor.getOrDefault(query[0], -1);
            if (currColor != newColor){
                mapColor.put(countColor.getOrDefault(newColor, 0) + 1);
                if (currColor > 0){
                    int currCount = countColor.getOrDefault(currColor, 0);
                    if (currCount <= 1){
                        countColor.remove(currColor);
                    } else {
                        countColor.put(currColor, currCount - 1);
                    }
                }
            }

            result[i] = countColor.size();
        }

        return result;
    }
}