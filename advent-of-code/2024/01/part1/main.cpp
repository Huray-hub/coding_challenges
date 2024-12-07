#include <bits/stdc++.h>

using namespace std;

typedef vector<int> vi;

int main(int argc, char *argv[]) {
    vi lv, rv;     // left side, right side
    int l, r;
    while (scanf("%d %d", &l, &r) == 2) {
        lv.push_back(l);
        rv.push_back(r);
    }

    sort(lv.begin(), lv.end());
    sort(rv.begin(), rv.end());

    int td = 0; // total distance
    for (int i = 0; i < lv.size(); i++) {
        td += abs(lv[i] - rv[i]);
    }

    printf("%d\n", td);
    return 0;
}
