class Solution {
    struct Node {
        long long cnt;
        long long sum;
    };

    string s;
    Node dp[17][11][11][2][2];
    bool vis[17][11][11][2][2];

    Node dfs(int pos, int p2, int p1, bool started, bool tight) {
        if (pos == (int)s.size()) {
            return {1, 0};
        }

        if (!tight && vis[pos][p2 + 1][p1 + 1][started][0]) {
            return dp[pos][p2 + 1][p1 + 1][started][0];
        }

        int lim = tight ? (s[pos] - '0') : 9;
        Node res{0, 0};

        for (int d = 0; d <= lim; d++) {
            bool ntight = tight && (d == lim);

            if (!started && d == 0) {
                Node nxt = dfs(pos + 1, -1, -1, false, ntight);
                res.cnt += nxt.cnt;
                res.sum += nxt.sum;
            } else {
                bool add = false;

                if (started && p2 != -1) {
                    add = ((p1 > p2 && p1 > d) || (p1 < p2 && p1 < d));
                }

                int np2, np1;
                if (!started) {
                    np2 = -1;
                    np1 = d;
                } else {
                    np2 = p1;
                    np1 = d;
                }

                Node nxt = dfs(pos + 1, np2, np1, true, ntight);

                res.cnt += nxt.cnt;
                res.sum += nxt.sum + (add ? nxt.cnt : 0);
            }
        }

        if (!tight) {
            vis[pos][p2 + 1][p1 + 1][started][0] = true;
            dp[pos][p2 + 1][p1 + 1][started][0] = res;
        }

        return res;
    }

    long long solve(long long n) {
        if (n <= 0) return 0;

        s = to_string(n);
        memset(vis, 0, sizeof(vis));

        Node ans = dfs(0, -1, -1, false, true);
        return ans.sum;
    }

public:
    long long totalWaviness(long long num1, long long num2) {
        return solve(num2) - solve(num1 - 1);
    }
};