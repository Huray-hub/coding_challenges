#include <bits/stdc++.h>

using namespace std;

typedef long long ll;
typedef vector<ll> vll;

ll ans = 0;

vll vec;

bool backtrack(ll r, ll a, int i) {
    if (i == vec.size()) {
        return a == r;
    }
    if (a > r) {
        return false;
    }

    ll b = vec[i];

    ll next = a + b;
    if (backtrack(r, next, i + 1)) {
        return true;
    }

    next = a * b;
    if (backtrack(r, next, i + 1)) {
        return true;
    }

    string nextStr = to_string(a) + to_string(b);
    if (nextStr.size() > to_string(r).size()) {
        return false;
    }
    next = stoll(nextStr);
    if (backtrack(r, next, i + 1)) {
        return true;
    }

    return false;
}

int main(int argc, char *argv[]) {
    string line;
    while (getline(cin, line)) {
        istringstream iss(line);
        ll r, tmp;
        iss >> r;
        iss.ignore(1); // ignoring ':'
        while (iss >> tmp) {
            vec.push_back(tmp);
        }
        bool found = backtrack(r, vec[0], 1);
        if (found) {
            ans += r;
        }
        vec.clear();
    }
    cout << ans << '\n';
    return 0;
}
