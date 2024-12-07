#include <bits/stdc++.h>

using namespace std;

typedef vector<string> vs;

int ans = 0;
vs grid;

bool isOutOfBounds(int r, int c) {
    return r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size();
}

void search(int r, int c) {
    // down
    if (!isOutOfBounds(r + 3, c)) {
        string word = "";
        for (int i = r; i <= r + 3; i++) {
            word += grid[i][c];
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // up
    if (!isOutOfBounds(r - 3, c)) {
        string word = "";
        for (int i = r; i >= r - 3; i--) {
            word += grid[i][c];
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // right
    if (!isOutOfBounds(r, c + 3)) {
        string word = "";
        for (int i = c; i <= c + 3; i++) {
            word += grid[r][i];
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // left
    if (!isOutOfBounds(r, c - 3)) {
        string word = "";
        for (int i = c; i >= c - 3; i--) {
            word += grid[r][i];
        }
        if (word == "XMAS") {
            ans++;
        }
    }

    // up-left
    if (!isOutOfBounds(r - 3, c - 3)) {
        string word = "";
        int i = r, j = c;
        while (i >= r - 3 && j >= c - 3) {
            word += grid[i][j];
            i--;
            j--;
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // up-right
    if (!isOutOfBounds(r - 3, c + 3)) {
        string word = "";
        int i = r, j = c;
        while (i >= r - 3 && j <= c + 3) {
            word += grid[i][j];
            i--;
            j++;
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // down-right
    if (!isOutOfBounds(r + 3, c + 3)) {
        string word = "";
        int i = r, j = c;
        while (i <= r + 3 && j <= c + 3) {
            word += grid[i][j];
            i++;
            j++;
        }
        if (word == "XMAS") {
            ans++;
        }
    }
    // down-left
    if (!isOutOfBounds(r + 3, c - 3)) {
        string word = "";
        int i = r, j = c;
        while (i <= r + 3 && j >= c - 3) {
            word += grid[i][j];
            i++;
            j--;
        }
        if (word == "XMAS") {
            ans++;
        }
    }
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
