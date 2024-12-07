#include <bits/stdc++.h>

using namespace std;

typedef unordered_map<string, int> msi;
typedef vector<string> vs;

int ans = 0;
unordered_map<string, vs> rules;

bool isCorrect(msi &tokenPositions, string &midToken) {
    int midIdx = tokenPositions.size() / 2;
    for (auto [token, idx] : tokenPositions) {
        string x = token;
        if (!rules.count(x)) {
            continue;
        }
        for (auto &y : rules[x]) {
            if (!tokenPositions.count(y)) {
                continue;
            }
            int px = idx;               // position of X
            int py = tokenPositions[y]; // position of Y
            if (px > py) {
                return false;
            }
        }
        if (idx == midIdx) {
            midToken = token;
        }
    }
    return true;
}

int main(int argc, char *argv[]) {
    string line;
    while (getline(cin, line) && !line.empty()) {
        auto pos = line.find('|');
        string x = line.substr(0, pos);
        string y = line.substr(pos + 1);
        rules[x].push_back(y);
    }
    while (getline(cin, line)) {
        istringstream stream(line);
        msi tokenPositions; // using map<token, position> to get tokens' idx in O(1)
        string token;
        int i = 0;
        while (getline(stream, token, ',')) {
            tokenPositions[token] = i;
            i++;
        }
        string midToken; // token in the middle
        if (isCorrect(tokenPositions, midToken)) {
            ans += stoi(midToken);
        }
    }
    cout << ans << '\n';
    return 0;
}
