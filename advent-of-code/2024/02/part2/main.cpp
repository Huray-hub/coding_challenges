#include <bits/stdc++.h>

using namespace std;

typedef vector<int> vi;

/* bool isSafe(vi &v) { */
/*     int prevDiff = 0; */
/*     for (int i = 1; i < v.size(); i++) { */
/*         int diff = v[i] - v[i - 1]; */
/*         if ((diff < 0 && prevDiff > 0) || (diff > 0 && prevDiff < 0)) { */
/*             return false; */
/*         } */
/*         if (abs(diff) < 1 || abs(diff) > 3) { */
/*             return false; */
/*         } */
/*         prevDiff = diff; */
/*     } */
/*     return true; */
/* } */
/**/
/* int main(int argc, char *argv[]) { */
/*     string line, tmp; */
/*     int ans = 0; */
/*     while (getline(cin, line)) { */
/*         istringstream ss(line); */
/*         vi v; */
/*         while (ss >> tmp) { */
/*             v.push_back(stoi(tmp)); */
/*         } */
/*         if (isSafe(v)) { */
/*             ans++; */
/*             continue; */
/*         } */
/*         for (int i = 0; i < v.size(); i++) { */
/*             vi sv = v; */
/*             sv.erase(sv.begin() + i); */
/*             if (isSafe(sv)) { */
/*                 ans++; */
/*                 break; */
/*             } */
/*         } */
/*     } */
/*     cout << ans << '\n'; */
/*     return 0; */
/* } */

bool isSafe(vi &v) {
    int prevIdx = 0;
    int prevDiff = 0;
    bool unsafeOnce = false;
    for (int currIdx = 1; currIdx < v.size(); currIdx++) {
        int diff = v[currIdx] - v[prevIdx];
        if ((diff < 0 && prevDiff > 0) || (diff > 0 && prevDiff < 0)) {
            if (unsafeOnce) {
                return false;
            }
            unsafeOnce = true;
            continue;
        }
        if (abs(diff) < 1 || abs(diff) > 3) {
            if (unsafeOnce) {
                return false;
            }
            unsafeOnce = true;
            continue;
        }
        prevIdx = currIdx;
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
            continue;
        }
    }
    cout << ans << '\n';
    return 0;
}
