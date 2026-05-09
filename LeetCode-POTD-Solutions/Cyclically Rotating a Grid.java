class Solution {
    public int[][] rotateGrid(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;

        int layers = Math.min(m, n) / 2;

        for (int layer = 0; layer < layers; layer++) {

            // boundaries of current layer
            int top = layer;
            int left = layer;
            int bottom = m - layer - 1;
            int right = n - layer - 1;

            // extract layer elements in clockwise order
            java.util.List<Integer> list = new java.util.ArrayList<>();

            // top row
            for (int j = left; j <= right; j++) {
                list.add(grid[top][j]);
            }

            // right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                list.add(grid[i][right]);
            }

            // bottom row
            for (int j = right; j >= left; j--) {
                list.add(grid[bottom][j]);
            }

            // left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                list.add(grid[i][left]);
            }

            int len = list.size();

            // counter-clockwise rotation by k
            int rot = k % len;

            int idx = 0;

            // place back rotated values
            // top row
            for (int j = left; j <= right; j++) {
                grid[top][j] = list.get((idx + rot) % len);
                idx++;
            }

            // right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                grid[i][right] = list.get((idx + rot) % len);
                idx++;
            }

            // bottom row
            for (int j = right; j >= left; j--) {
                grid[bottom][j] = list.get((idx + rot) % len);
                idx++;
            }

            // left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                grid[i][left] = list.get((idx + rot) % len);
                idx++;
            }
        }

        return grid;
    }
}