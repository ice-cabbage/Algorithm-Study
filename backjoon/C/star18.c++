#include <iostream>
#include <vector>
#include <algorithm>
#include <cstring>
using namespace std;
 
char map[1023][2045];
void drawStar(int n, int x, int y, bool dir) {
    map[x][y] = '*';
    if (n == 1) return;
 
    int height = (1 << n) - 1;
    int left = y, right = y;
    for (int i = 1; i < height; i++) {
        if (dir) x--; 
        else x++; 
 
        left--; right++;
        map[x][left] = '*'; map[x][right] = '*';
    }
 
    for (int j = left + 1; j < right; j++) {
        map[x][j] = '*';
    }
 
    if (dir) x++;
    else x--;
 
    drawStar(n - 1, x, y, !dir);
}
 
int main(void) {
    int n;
    cin >> n;
 
    int height = (1 << n) - 1;
    int width = 2 * height - 1;
    for (int i = 0; i < height; i++) {
        for (int j = 0; j < width; j++) {
            map[i][j] = ' ';
        }
    }
 
    bool dir = false; 
    if (n % 2 == 0) dir = true; 
 
    int x = 0;
    if (dir) x = (1 << n) - 2;
    int y = (1 << n) - 2;
 
    drawStar(n, x, y, dir);
 
    if (dir) {
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                printf("%c", map[i][j]);
            }
            width--;
            cout << "\n";
        }
    }
    else {
        for (int i = 0; i < height; i++) {
            for (int j = 0; j <= y; j++) {
                printf("%c", map[i][j]);
            }
            y++;
            cout << "\n";
        }
    }
 
    return 0;
}