#include <bits/stdc++.h>

using namespace std;

typedef vector<int> vi;

bool isSafe(vi &v) {
    int prevDiff = 0;
    for (int i = 1; i < v.size(); i++) {
        int diff = v[i] - v[i - 1];
        if ((diff < 0 && prevDiff > 0) || (diff > 0 && prevDiff < 0)) {
            return false;
        }
        if (abs(diff) < 1 || abs(diff) > 3) {
            return false;
        }
        prevDiff = diff;
    }
    return true;
}

int main(int argc, char *argv[]) {
    string line, tmp;
    int ans = 0;
    while (getline(cin, line)) {
        istringstream ss(line);
        vi v;
        while (ss >> tmp) {
            v.push_back(stoi(tmp));
        }
        if (isSafe(v)) {
            ans++;
        }
    }
    cout << ans << '\n';
    return 0;
}
