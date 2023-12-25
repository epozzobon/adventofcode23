#include <algorithm>
#include <assert.h>
#include <fstream>
#include <iostream>
#include <string>
#include <tuple>
#include <vector>
using namespace std;

vector<string> schematic;

vector<string> readCharMatrix(ifstream& myfile)
{
    vector<string> output;
    string line;
    int columns = 0;
    while (getline(myfile, line)) {
        if (line == "") {
            break;
        }

        if (columns == 0) {
            columns = line.length();
        } else if (columns != line.length()) {
            throw runtime_error("Inconsistent column size");
        }
        output.push_back(line);
    }
    return output;
}

bool isDigit(char c)
{
    return string("0123456789").find(c) != string::npos;
}

char fetch(int row, int column)
{
    if (row < 0 || column < 0 || row >= schematic.size() || column >= schematic[row].length()) {
        return '.';
    } else {
        return schematic[row][column];
    }
};

char findComponentAround(int r, int c, int l, int& r0, int& c0)
{
    char component = '.';
    for (r0 = r - 1; r0 <= r + 1; r0++) {
        for (c0 = c - 1; c0 <= c + l; c0++) {
            char h = fetch(r0, c0);
            if ('.' != h && !isDigit(h)) {
                component = h;
                return h;
            }
        }
    }
    return '.';
}

int main()
{
    string line;
    ifstream myfile("input.txt");
    if (!myfile.is_open()) {
        cout << "Unable to open file" << endl;
        return 1;
    }

    vector<tuple<int, int, int>> potentialGears = {};
    unsigned partNumberSum = 0;
    unsigned gearsRatioSum = 0;

    schematic = readCharMatrix(myfile);
    myfile.close();

    unsigned V = schematic.size();
    unsigned H = schematic[0].length();
    for (int r = 0; r < V; r++) {
        for (int c = 0; c < H; c++) {
            int l = 0;
            while (isDigit(fetch(r, c + l))) {
                l++;
            }
            if (l == 0) {
                continue;
            }
            string numStr = schematic[r].substr(c, l);
            int num = stoi(numStr);

            int r0, c0;
            char component = findComponentAround(r, c, l, r0, c0);

            if (component != '.') {
                cout << num << endl;
                partNumberSum += num;
                if ('*' == component) {
                    for (auto& otherCoords : potentialGears) {
                        int r1, c1, num1;
                        tie(r1, c1, num1) = otherCoords;
                        if (r1 == r0 && c1 == c0) {
                            gearsRatioSum += num * num1;
                            break;
                        }
                    }
                    potentialGears.push_back(tuple<int, int, int>(r0, c0, num));
                }
            }

            c += l - 1;
        }
    }

    cout << partNumberSum << endl;
    cout << gearsRatioSum << endl;

    return 0;
}