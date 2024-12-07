#include <bits/stdc++.h>

using namespace std;

int ans = 0;

int main(int argc, char *argv[]) {
    string line;
    bool enabled = true;
    while (cin >> line) {
        regex pattern(R"(mul\((\d+),(\d+)\)|do\(\)|don't\(\))");

        auto begin = sregex_iterator(line.begin(), line.end(), pattern);
        auto end = sregex_iterator();

        for (auto it = begin; it != end; ++it) {
            smatch match = *it;
            if (match.str().find("mul(") == 0) {
                int first = stoi(match[1].str());
                int second = stoi(match[2].str());
                ans += first * second * enabled;
            } else if (match.str() == "don't()") {
                enabled = false;
            } else if (match.str() == "do()") {
                enabled = true;
            }
        }
    }
    cout << ans << '\n';
    return 0;
}
