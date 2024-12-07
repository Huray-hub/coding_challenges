#include <bits/stdc++.h>
#include <string>
#include <unordered_map>
#include <vector>

#define fi first
#define se second

using namespace std;

typedef vector<int> vi;
typedef vector<vi> vvi;
typedef pair<int, int> ii;
typedef vector<ii> vii;
typedef tuple<int, int, int> iii;
typedef vector<iii> viii;
typedef vector<string> vs;
typedef vector<vs> vvs;
typedef unordered_map<int, int> mii;

const ii ud = {-1, 0};
const ii dd = {1, 0};
const ii rd = {0, 1};
const ii ld = {0, -1};

int ans = 0;

vs g;

ii sp;
ii sd;
ii cp; // current position
ii cd; // current direction

vvi visited;

bool isOutOfBounds(int r, int c) {
    return r < 0 || c < 0 || r >= g.size() || c >= g[0].size();
}

ii rotate(ii d) {
    if (d == ud) {
        return rd;
    } else if (d == rd) {
        return dd;
    } else if (d == dd) {
        return ld;
    } else {
        return ud;
    }
}

set<pair<ii, ii>> turningPoints;

bool moveFirst() {
    turningPoints.insert({cp, cd});
    cp = {cp.fi + cd.fi, cp.se + cd.se};
    while (!isOutOfBounds(cp.fi, cp.se)) {
        if (g[cp.fi][cp.se] == '#') {
            cp = {cp.fi - cd.fi, cp.se - cd.se};
            cd = rotate(cd);
            turningPoints.insert({cp, cd});
            return true;
        }
        if (!visited[cp.fi][cp.se]) {
            visited[cp.fi][cp.se] = true;
            ans++;
        }
        cp = {cp.fi + cd.fi, cp.se + cd.se};
    }
    return false;
}

int main(int argc, char *argv[]) {
    string line;
    while (getline(cin, line)) {
        g.push_back(line);
        for (int i = 0; i < line.size(); i++) {
            auto t = line[i];
            if (t == '^' || t == '>' || t == 'v' || t == '<') {
                cp.fi = g.size() - 1;
                cp.se = i;
                switch (t) {
                case '^':
                    cd = ud;
                    break;
                case '>':
                    cd = rd;
                    break;
                case 'v':
                    cd = dd;
                    break;
                case '<':
                    cd = ld;
                    break;
                }
            }
        }
    }

    sp = cp;
    sd = cd;

    visited = vvi(g.size(), vi(g[0].size(), false));
    while (moveFirst())
        ;

    for (auto [p, d] : turningPoints) {
        if (d == ld) {
            int r = p.fi;
            int c = p.se - 1;
            ii nd = rotate(d);
            while (!isOutOfBounds(r, c) && g[r][c] != '#') {
                if (turningPoints.count({{r, c}, nd})) {
                    ans++;
                    break;
                }
                c--;
            }
        }
        if (d == rd) {
            int r = p.fi;
            int c = p.se + 1;
            ii nd = rotate(d);
            while (!isOutOfBounds(r, c) && g[r][c] != '#') {
                if (turningPoints.count({{r, c}, nd})) {
                    ans++;
                    break;
                }
                c++;
            }
        }
        if (d == ud) {
            int r = p.fi - 1;
            int c = p.se;
            ii nd = rotate(d);
            while (!isOutOfBounds(r, c) && g[r][c] != '#') {
                if (turningPoints.count({{r, c}, nd})) {
                    ans++;
                    break;
                }
                r--;
            }
        }
        if (d == dd) {
            int r = p.fi + 1;
            int c = p.se;
            ii nd = rotate(d);
            while (!isOutOfBounds(r, c) && g[r][c] != '#') {
                if (turningPoints.count({{r, c}, nd})) {
                    ans++;
                    break;
                }
                r++;
            }
        }
    }

    cout << ans << '\n';
    return 0;
}
