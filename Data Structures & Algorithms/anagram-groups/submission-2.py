class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        cmap = {}
        for str1 in strs:
          cmap.setdefault("".join(sorted(str1)), []).append(str1)
        return [v for v in cmap.values()]