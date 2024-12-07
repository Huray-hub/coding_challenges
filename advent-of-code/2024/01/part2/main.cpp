#include <bits/stdc++.h>

using namespace std;

typedef vector<int> vi;
typedef unordered_map<int, int> mii;

int main(int argc, char *argv[]) {
    vi lv;  // left side stored in a vector
    mii rm; // right side stored in a map
    int l, r;
    while (scanf("%d %d", &l, &r) == 2) {
        lv.push_back(l);
        rm[r]++;
    }

    int ss = 0; // similarity score
    for (auto n : lv) {
        ss += n * rm[n];
    }
    printf("%d\n", ss);
    return 0;
}
