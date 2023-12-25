#include <assert.h>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
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

vector<int> splitInts(string s)
{
    vector<int> output = {};
    vector<string> pieces = split(s, " ");
    for (unsigned i = 0; i < pieces.size(); i++) {
        if (pieces[i].length() > 0) {
            output.push_back(stoi(pieces[i]));
        }
    }
    return output;
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
    vector<int> bonus = {};

    long long unsigned sumScore = 0;
    long long unsigned totalCards = 0;
    while (getline(myfile, line)) {
        bonus.push_back(0);

        vector<string> pieces = split(line, ":");
        assert(pieces.size() == 2);
        int cardID = stoi(pieces[0].substr(5));
        pieces = split(pieces[1], "|");
        vector<int> winNums = splitInts(pieces[0]);
        vector<int> cardNums = splitInts(pieces[1]);

        cout << cardID << ":";
        for (unsigned i = 0; i < winNums.size(); i++) {
            cout << winNums[i] << " ";
        }
        cout << "|";
        unsigned matches = 0;
        for (unsigned i = 0; i < cardNums.size(); i++) {
            cout << cardNums[i] << " ";
            if (winNums.end() != find(winNums.begin(), winNums.end(), cardNums[i])) {
                matches++;
            }
        }

        // Score only for first part
        unsigned score = 0;
        for (int i = 0; i < matches; i++) {
            if (score == 0) {
                score++;
            } else {
                score *= 2;
            }
        }

        unsigned numCards = 1 + bonus[cardID - 1];
        totalCards += numCards;
        for (int i = 0; i < matches; i++) {
            while (cardID + i >= bonus.size()) {
                bonus.push_back(0);
            }
            bonus[cardID + i] += numCards;
        }
        cout << "|" << score << endl;
        sumScore += score;
    }
    myfile.close();
    cout << sumScore << endl;
    cout << totalCards << endl;
    return 0;
}