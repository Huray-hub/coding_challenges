#include <bits/stdc++.h>

using namespace std;

typedef vector<string> vs;

int ans = 0;
vs grid;

bool isOutOfBounds(int r, int c) {
    return r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size();
}

void search(int r, int c) {
    if (grid[r][c] != 'A') {
        return;
    }

    if (isOutOfBounds(r - 1, c - 1) || isOutOfBounds(r + 1, c + 1) || isOutOfBounds(r - 1, c + 1)
        || isOutOfBounds(r + 1, c - 1)) {
        return;
    }

    if (!(grid[r - 1][c - 1] == 'M' && grid[r + 1][c + 1] == 'S')
        && !(grid[r - 1][c - 1] == 'S' && grid[r + 1][c + 1] == 'M')) {
        return;
    }
    if (!(grid[r + 1][c - 1] == 'M' && grid[r - 1][c + 1] == 'S')
        && !(grid[r + 1][c - 1] == 'S' && grid[r - 1][c + 1] == 'M')) {
        return;
    }

    ans++;
}

int main(int argc, char *argv[]) {
    string line;
    while (cin >> line) {
        grid.push_back(line);
    }
    for (int i = 0; i < grid.size(); i++) {
        for (int j = 0; j < grid[0].size(); j++) {
            search(i, j);
        }
    }
    cout << ans << '\n';
    return 0;
}
