#include <assert.h>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

vector<string> split(string s, string delimiter)
{
    size_t pos_start = 0, pos_end, delim_len = delimiter.length();
    string token;
    vector<string> res;

    while ((pos_end = s.find(delimiter, pos_start)) != string::npos) {
        token = s.substr(pos_start, pos_end - pos_start);
        pos_start = pos_end + delim_len;
        res.push_back(token);
    }

    res.push_back(s.substr(pos_start));
    return res;
}

int main()
{
    string line;
    ifstream myfile("input.txt");
    if (!myfile.is_open()) {
        cout << "Unable to open file" << endl;
        return 1;
    }

    vector<string> colors = { "red", "green", "blue" };

    int sumGG = 0;
    long long unsigned sumPower = 0;
    while (getline(myfile, line)) {
        vector<string> pieces = split(line, ":");
        assert(pieces.size() == 2);
        vector<string> pgame = split(pieces[0], " ");
        assert(pgame.size() == 2);
        unsigned gameID = stoi(pgame[1]);
        vector<string> subgame = split(pieces[1], ";");
        unsigned max[3] = { 0, 0, 0 };
        for (unsigned i = 0; i < subgame.size(); i++) {
            vector<string> draws = split(subgame[i], ",");
            for (unsigned j = 0; j < draws.size(); j++) {
                vector<string> drawPieces = split(draws[j].substr(1), " ");
                string color = drawPieces[1];
                unsigned count = stoi(drawPieces[0]);
                for (unsigned k = 0; k < 3; k++) {
                    if (color == colors[k]) {
                        if (count > max[k]) {
                            max[k] = count;
                        }
                    }
                }
            }
        }
        if (max[0] > 12 || max[1] > 13 || max[2] > 14) {
            // bad game
        } else {
            sumGG += gameID;
        }
        sumPower += max[0] * max[1] * max[2];
    }
    myfile.close();
    cout << sumGG << endl;
    cout << sumPower << endl;
    return 0;
}