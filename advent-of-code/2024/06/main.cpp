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

ii sp; // starting position
ii sd; // starting direction
ii cp; // current position
ii cd; // current direction

vvi visited;

bool isOutOfBounds(int r, int c) {
    return r < 0 || c < 0 || r >= g.size() || c >= g[0].size();
}

void rotate() {
    if (cd == ud) {
        cd = rd;
        return;
    } else if (cd == rd) {
        cd = dd;
        return;
    } else if (cd == dd) {
        cd = ld;
        return;
    } else if (cd == ld) {
        cd = ud;
        return;
    }
}

vector<pair<ii, ii>> turningPoints;

bool contains() {
    for (auto [p, d] : turningPoints) {
        /* if (p == cp) { */
        /*     return true; */
        /* } */
        if (p == cp && d == cd) {
            return true;
        }
    }
    return false;
}

bool movee(bool &foundCircle) {
    turningPoints.emplace_back(sp, sd);
    cp = {sp.fi + sd.fi, sp.se + sd.se};
    while (!isOutOfBounds(cp.fi, cp.se)) {
        if (g[cp.fi][cp.se] == '#') {
            cp = {cp.fi - cd.fi, cp.se - cd.se};
            rotate();
            if (contains()) {
                foundCircle = true;
                return false;
            }
            turningPoints.emplace_back(cp, cd);
            return true;
        }
        if (!visited[cp.fi][cp.se]) {
            visited[cp.fi][cp.se] = true;
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

    for (int i = 0; i < g.size(); i++) {
        for (int j = 0; j < g[0].size(); j++) {
            if (g[i][j] == '#' || (i == sp.fi && j == sp.se) || g[i][j] != '.') {
                continue;
            }
            g[i][j] = '#';
            visited = vvi(g.size(), vi(g[0].size(), false));
            turningPoints = vector<pair<ii, ii>>();
            bool foundCircle = false;
            while (movee(foundCircle))
                ;
            if (foundCircle) {
                ans++;
            }
            g[i][j] = '.';
            cp = sp;
            cd = sd;
        }
    }

    cout << ans << '\n';
    return 0;
}
