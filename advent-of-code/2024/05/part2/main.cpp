#include <bits/stdc++.h>

#define fi first
#define se second

using namespace std;

typedef vector<int> vi;
typedef pair<int, int> ii;
typedef vector<ii> vii;

int ans = 0;
vii rules;

int findPos(vi &vec, int val) {
    auto it = find(vec.begin(), vec.end(), val);
    if (it != vec.end()) {
        int index = distance(vec.begin(), it);
        return index;
    }
    return -1; // not found
}

bool isSorted(vi &update) {
    for (auto rule : rules) {
        int fpos = findPos(update, rule.fi);
        if (fpos == -1) {
            continue;
        }
        int spos = findPos(update, rule.se);
        if (spos == -1) {
            continue;
        }
        if (fpos > spos) {
            return false;
        }
    }
    return true;
}

int main(int argc, char *argv[]) {
    string line;
    while (getline(cin, line) && !line.empty()) {
        size_t pos = line.find('|');
        ii r; // rule
        r.fi = stoi(line.substr(0, pos));
        r.se = stoi(line.substr(pos + 1));
        rules.push_back(r);
    }
    while (getline(cin, line)) {
        vi update;
        istringstream stream(line);
        string token;
        while (getline(stream, token, ',')) {
            update.push_back(stoi(token));
        }
        if (!isSorted(update)) {
            do {
                for (auto rule : rules) {
                    int fpos = findPos(update, rule.fi);
                    if (fpos == -1) {
                        continue;
                    }
                    int spos = findPos(update, rule.se);
                    if (spos == -1) {
                        continue;
                    }
                    if (fpos > spos) {
                        int tmp = update[fpos];
                        update[fpos] = update[spos];
                        update[spos] = tmp;
                    }
                }
            } while (!isSorted(update));
            ans += update[update.size() / 2];
        }
    }
    cout << ans << '\n';
    return 0;
}
