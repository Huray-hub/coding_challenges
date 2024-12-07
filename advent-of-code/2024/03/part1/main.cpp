#include <bits/stdc++.h>

using namespace std;

int ans = 0;

int main(int argc, char *argv[]) {
    string line;
    while (cin >> line) {
        regex pattern(R"(mul\((\d+),(\d+)\))");

        auto begin = sregex_iterator(line.begin(), line.end(), pattern);
        auto end = sregex_iterator();

        for (auto it = begin; it != end; ++it) {
            smatch match = *it;
            int first = stoi(match[1].str());
            int second = stoi(match[2].str());
            ans += first * second;
        }
    }
    cout << ans << '\n';
    return 0;
}
