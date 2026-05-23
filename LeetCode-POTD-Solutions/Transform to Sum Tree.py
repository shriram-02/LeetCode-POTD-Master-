class Solution:
    def toSumTree(self, root):
        
        def solve(node):
            if not node:
                return 0
            
            old_val = node.data
            
            left_sum = solve(node.left)
            right_sum = solve(node.right)
            
            node.data = left_sum + right_sum
            
            return old_val + node.data
        
        solve(root)