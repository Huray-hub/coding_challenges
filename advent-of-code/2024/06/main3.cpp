#include <bits/stdc++.h>

#define fi first
#define se second

using namespace std;

typedef vector<int> vi;
typedef vector<vi> vvi;
typedef pair<int, int> ii;
typedef vector<string> vs;

const ii ud = {-1, 0};
const ii dd = {1, 0};
const ii rd = {0, 1};
const ii ld = {0, -1};

vs g;
int ans = 0;

ii rotate(ii d) {
    if (d == ud) {
        return rd;
    }
    if (d == rd) {
        return dd;
    }
    if (d == dd) {
        return ld;
    }
    return ud;
}

bool isOutOfBounds(int r, int c) {
    return r < 0 || c < 0 || r >= g.size() || c >= g[0].size();
}

bool isLoop(ii p, ii d) {
    set<pair<ii, ii>> states;
    states.insert({p, d});

    while (true) {
        ii np = {p.fi + d.fi, p.se + d.se};
        if (isOutOfBounds(np.fi, np.se)) {
            return false;
        }
        if (g[np.fi][np.se] == '#') {
            d = rotate(d);
        } else {
            p = np; // Move forward
        }
        pair<ii, ii> state = {p, d};
        if (states.count(state)) {
            return true;
        }
        states.insert(state);
    }
};

int main() {
    string line;
    while (getline(cin, line)) {
        g.push_back(line);
    }

    ii sp; // starting position
    ii sd; // starting direction
    for (int i = 0; i < g.size(); i++) {
        for (int j = 0; j < g[i].size(); j++) {
            char t = g[i][j];
            if (t == '^') {
                sp = {i, j};
                sd = ud;
            } else if (t == '>') {
                sp = {i, j};
                sd = rd;
            } else if (t == 'v') {
                sp = {i, j};
                sd = dd;
            } else if (t == '<') {
                sp = {i, j};
                sd = ld;
            }
        }
    }

    for (int i = 0; i < g.size(); i++) {
        for (int j = 0; j < g[i].size(); j++) {
            if (g[i][j] == '#' || (i == sp.fi && j == sp.se)) {
                continue;
            }
            if (g[i][j] != '.') {
                continue;
            }
            g[i][j] = '#';
            if (isLoop(sp, sd)) {
                ans++;
            }
            g[i][j] = '.';
        }
    }
    cout << ans << "\n";
    return 0;
}
