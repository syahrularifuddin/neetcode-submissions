# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
      path = []
      total = 0
      def traverse(root, path, total):
        if not root:
          return False
        path.append(root.val)
        total += root.val
        if not root.left and not root.right: # leaf
          if total == targetSum:
            return True
          return False
        if traverse(root.left, path, total):
          return True
        if traverse(root.right, path, total):
          return True
        path.pop()
        total -= root.val
        return False
      res = traverse(root, path, total)
      print(path, total)
      return res

        
        