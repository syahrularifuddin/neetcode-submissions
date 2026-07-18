class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key = lambda interval: interval[0])
        output = []
        for i in range(len(intervals)):
            if i != 0 and intervals[i][0] <= output[len(output)-1][1]:
                if intervals[i][1] > output[len(output)-1][1]:
                    output[len(output)-1][1] = intervals[i][1]
            else:
                output.append(intervals[i])
        return output
