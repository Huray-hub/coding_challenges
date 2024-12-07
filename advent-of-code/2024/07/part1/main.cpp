#include <bits/stdc++.h>
#include <string>

using namespace std;

typedef vector<long> vl;

long ans = 0;

int main(int argc, char *argv[]) {
    string line;
    while (getline(cin, line)) {
        istringstream iss(line);
        long r, tmp;
        iss >> r;
        vl vec;
        iss.ignore(1); // ignoring ':'
        while (iss >> tmp) {
            vec.push_back(tmp);
        }

        int n = vec.size() - 1;
        ulong totalCombs = 1 << n;
        for (size_t i = 0; i < totalCombs; ++i) {
            long res = vec[0];
            for (size_t j = 0; j < n; ++j) {
                bool add = (i & (1 << j)) != 0;
                if (add) {
                    res += vec[j + 1];
                } else {
                    res *= vec[j + 1];
                }
            }
            if (res == r) {
                ans += r;
                break;
            }
        }
    }
    cout << ans << '\n';
    return 0;
}
